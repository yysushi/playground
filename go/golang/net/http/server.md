# Note for http server

- net/http

```golang
func Handle(pattern string, handler Handler)
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
func ListenAndServe(addr string, handler Handler) error
func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error
func Serve(l net.Listener, handler Handler) error
type Handler
type HandlerFunc
    func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)
type ServeMux
    func NewServeMux() *ServeMux
    func (mux *ServeMux) Handle(pattern string, handler Handler)
    func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
    func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string)
    func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)
type Server
    func (srv *Server) Close() error
    func (srv *Server) ListenAndServe() error
    func (srv *Server) ListenAndServeTLS(certFile, keyFile string) error
    func (srv *Server) Serve(l net.Listener) error
    func (srv *Server) Shutdown(ctx context.Context) error
```

- notable types

  - Handler is a interface to serve http.

  ```golang
  type Handler interface {
  	ServeHTTP(ResponseWriter, *Request)
  }
  ```

  - HandlerFunc is a just function which is something serving http.
    it implements a Handler interface

  ```golang
  type HandlerFunc func(ResponseWriter, *Request)
  func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request){
  	f(w, r)
  }
  ```

    => this enables easier registeration of a function which meats Handler interface

  ```golang
  http.HandlerFunc(func(w http.ResponseWriter, *http.Request) {fmt.Fprintln(w, "hello")})
  ```

  - ServeMux is a http request multiplexer (= send one signal as a result of receving two signals).
    it implements a Handler interface

- Handle/HandleFunc vs ServeMux.Handle/ServeMux.HandleFunc 

  - Handle/HandleFunc is a link to default server mux

- how to register a handler

  1. pass handler

  ```golang
  http.Handle("/", handler)
  ```

  2. pass handler func

  ```golang
  http.HandleFunc("/", func(w ResponseWriter, *Request) {fmt.Fprintln(w, "hello")})
  ```

  - http.HandleFunc just wraps http.Handle

  ```golang
  func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
  	if handler == nil {
  		panic("http: nil handler")
  	}
  	mux.Handle(pattern, HandlerFunc(handler))
  }
  ```

- how to serve 

  - ListenAndServe is mapped to Serve and it is default mapped to serve mux's ServeHTTP
    <https://golang.org/src/net/http/server.go?s=77114:81434#L1895>

  - serve mux's ServeHTTP searchs preferable a Handler and call its ServeHTTP
    <https://golang.org/src/net/http/server.go?s=77114:81434#L2387>
