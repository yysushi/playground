package main

import (
	"context"
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"

	"github.com/chromedp/cdproto/har"
	goproxy "gopkg.in/elazarl/goproxy.v1"
)

var recordableContentTypes = []string{
	"text",
	"application/json",
}

type Server struct {
	addr      string
	adminAddr string

	entries []*Entry
	mux     sync.RWMutex
}

func NewServer(addr, adminAddr string) *Server {
	return &Server{
		addr:      addr,
		adminAddr: adminAddr,
	}
}

func (s *Server) NewEntry() *Entry {
	var ent Entry
	s.mux.Lock()
	defer s.mux.Unlock()
	s.entries = append(s.entries, &ent)
	return &ent
}

func (s *Server) Run(parentCtx context.Context) {
	ctx, cancel := context.WithCancel(parentCtx)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		s.serveProxy(ctx)
		cancel()
	}()
	go func() {
		defer wg.Done()
		s.serveController(ctx)
		cancel()
	}()
	wg.Wait()
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	harBytes, _ := s.har().MarshalJSON()
	if _, reset := r.URL.Query()["reset"]; reset {
		s.resetEntries()
	}
	w.Write(harBytes)
}

func (s *Server) har() *har.HAR {
	s.mux.RLock()
	defer s.mux.RUnlock()
	harLog := har.HAR{
		Log: &har.Log{
			Entries: s.harEntries(),
		},
	}
	return &harLog
}

func (s *Server) resetEntries() {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.entries = nil
}

func (s *Server) harEntries() []*har.Entry {
	entries := make([]*har.Entry, 0)
	for _, entry := range s.entries {
		if !entry.Recorded() {
			continue
		}
		entries = append(entries, entry.harEntry())
	}
	return entries
}

type Entry struct {
	req              *http.Request
	reqBodyRecorder  *BodyRecorder
	resp             *http.Response
	respBodyRecorder *BodyRecorder
}

func (e *Entry) addReq(req *http.Request, origBody io.ReadCloser) (recordableBody io.ReadCloser) {
	e.req = req
	e.reqBodyRecorder, recordableBody = NewBodyRecorder(req.Header.Get("Content-Type"), origBody)
	return recordableBody
}

func (e *Entry) addResp(resp *http.Response, origBody io.ReadCloser) (recordableBody io.ReadCloser) {
	e.resp = resp
	e.respBodyRecorder, recordableBody = NewBodyRecorder(resp.Header.Get("Content-Type"), origBody)
	return recordableBody
}

func (e *Entry) Recorded() bool {
	return e.reqBodyRecorder.recorded && e.respBodyRecorder.recorded
}

func (e *Entry) harEntry() *har.Entry {
	var reqHeaders, respHeaders []*har.NameValuePair
	var queryString []*har.NameValuePair
	for k, vset := range e.req.Header {
		for _, v := range vset {
			reqHeaders = append(reqHeaders, &har.NameValuePair{
				Name:  k,
				Value: v,
			})
		}
	}
	for k, vset := range e.resp.Header {
		for _, v := range vset {
			respHeaders = append(respHeaders, &har.NameValuePair{
				Name:  k,
				Value: v,
			})
		}
	}
	for k, vset := range e.req.URL.Query() {
		for _, v := range vset {
			queryString = append(queryString, &har.NameValuePair{
				Name:  k,
				Value: v,
			})
		}
	}
	return &har.Entry{
		Request: &har.Request{
			Method:      e.req.Method,
			URL:         e.req.URL.String(),
			HTTPVersion: e.req.Proto,
			Headers:     reqHeaders,
			QueryString: queryString,
			PostData: &har.PostData{
				MimeType: e.req.Header.Get("Content-Type"),
				Text:     string(e.reqBodyRecorder.recordedBody),
			},
			BodySize: int64(e.reqBodyRecorder.size),
		},
		Response: &har.Response{
			Status:      int64(e.resp.StatusCode),
			StatusText:  e.resp.Status,
			HTTPVersion: e.resp.Proto,
			Headers:     respHeaders,
			Content: &har.Content{
				MimeType: e.resp.Header.Get("Content-Type"),
				Text:     base64.StdEncoding.EncodeToString(e.respBodyRecorder.recordedBody),
				Encoding: "base64",
			},
			BodySize: int64(e.respBodyRecorder.size),
		},
	}
}

func NewBodyRecorder(contentType string, body io.ReadCloser) (*BodyRecorder, io.ReadCloser) {
	var needRecording bool = false
	for _, recordableContentType := range recordableContentTypes {
		needRecording = needRecording || strings.HasPrefix(contentType, recordableContentType)
	}
	r := &BodyRecorder{
		origBody:      body,
		skipRecording: !needRecording,
	}
	return r, struct {
		io.Reader
		io.Closer
	}{
		io.TeeReader(body, r),
		r,
	}
}

type BodyRecorder struct {
	origBody      io.ReadCloser
	skipRecording bool
	size          int
	recordedBody  []byte
	recorded      bool
}

func (r *BodyRecorder) Close() (err error) {
	defer func() { r.recorded = true }()
	return r.origBody.Close()
}

func (r *BodyRecorder) Write(data []byte) (n int, err error) {
	if r.skipRecording {
		return
	}
	r.recordedBody = append(r.recordedBody, data...)
	r.size += len(data)
	return
}

func (srv *Server) serveProxy(parentCtx context.Context) {
	ctx, cancel := context.WithCancel(parentCtx)
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	s := &http.Server{
		Handler: proxy,
		Addr:    srv.addr,
	}
	proxy.OnRequest().DoFunc(
		func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			ent := srv.NewEntry()
			req.Body = ent.addReq(ctx.Req, req.Body)
			req.Header.Add("X-Forwarded-For", req.RemoteAddr)
			ctx.UserData = ent
			return req, nil
		})

	proxy.OnResponse().DoFunc(
		func(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
			ent := (ctx.UserData).(*Entry)
			resp.Body = ent.addResp(ctx.Resp, resp.Body)
			// TODO: get remote address by connection
			// then store as har entry's ServerIPAddress and set XFF
			// https://pkg.go.dev/net/http/httptrace#GotConnInfo
			return resp
		})
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		if err := s.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("failed to serve proxy: %s\n", err)
		}
		cancel()
	}()
	go func() {
		defer wg.Done()
		<-ctx.Done()
		if err := s.Shutdown(context.Background()); err != nil {
			log.Printf("failed to shutdown proxy server: %s\n", err)
		}
	}()
	wg.Wait()
}

func (srv *Server) serveController(parentCtx context.Context) {
	ctx, cancel := context.WithCancel(parentCtx)
	s := &http.Server{
		Handler: srv,
		Addr:    srv.adminAddr,
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		if err := s.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("failed to serve proxy: %s\n", err)
		}
		cancel()
	}()
	go func() {
		defer wg.Done()
		<-ctx.Done()
		if err := s.Shutdown(context.Background()); err != nil {
			log.Printf("failed to shutdown proxy server: %s\n", err)
		}
	}()
	wg.Wait()
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	server := NewServer(":8080", ":8081")
	server.Run(ctx)
}
