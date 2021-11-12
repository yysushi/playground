package request

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"
)

type Request struct {
	pid    int
	seq    int
	sentAt time.Time
}

// EmbeddedMessage ...
type embeddedMessage struct {
	PID int64
}

// this is workaround for linux
func (r *Request) embeddedMessage() *embeddedMessage {
	return &embeddedMessage{
		PID: int64(r.pid), // int is not propert to embed because a fixed size is necessary.
	}
}

func (r *Request) Encode() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, r.embeddedMessage())
	return buf.Bytes()
}

func decodeEmbeddedMessage(s []byte) (*embeddedMessage, error) {
	buf := bytes.NewReader(s)
	var em embeddedMessage
	err := binary.Read(buf, binary.LittleEndian, &em)
	if err != nil {
		return nil, err
	}
	return &em, nil
}

type stat struct {
	Seq                 int
	ElapsedMicroseconds time.Duration
}

func (r *Request) CalcStat(s []byte, seq int, recvAt time.Time) (*stat, error) {
	em, err := decodeEmbeddedMessage(s)
	if err != nil {
		return nil, fmt.Errorf("decode failure: %s", err)
	}
	if int(em.PID) != r.pid {
		return nil, fmt.Errorf("others (pid:%d)", em.PID)
	}
	return &stat{
		Seq:                 seq,
		ElapsedMicroseconds: recvAt.Sub(r.sentAt),
	}, nil
}

func NewRequest(pid int, seq int) *Request {
	return &Request{
		pid:    pid,
		seq:    seq,
		sentAt: time.Now(),
	}
}
