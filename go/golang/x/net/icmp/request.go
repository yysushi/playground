package icmp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/net/icmp"
)

// Request ...
type Request struct {
	pid      int
	seq      int
	sentAt   time.Time
	bodySize int
}

func (r *Request) fillInPaddingData(buf *bytes.Buffer) {
	paddingMesage := make([]byte, r.paddingBodySize())
	rand.Read(paddingMesage)
	binary.Write(buf, binary.LittleEndian, paddingMesage)
}

// Encode ...
func (r *Request) Encode() []byte {
	buf := new(bytes.Buffer)
	r.embedMetaData(buf)
	r.fillInPaddingData(buf)
	return buf.Bytes()
}

// Stat ...
type Stat struct {
	Seq                 int
	ElapsedMicroseconds time.Duration
}

// CalcStat ...
func (r *Request) CalcStat(m *icmp.Echo, recvAt time.Time) (*Stat, error) {
	pid, err := r.decodePID(m)
	if err != nil {
		return nil, fmt.Errorf("decode failure: %s", err)
	}
	if pid != r.pid {
		return nil, fmt.Errorf("others (pid:%d)", pid)
	}
	return &Stat{
		Seq:                 m.Seq,
		ElapsedMicroseconds: recvAt.Sub(r.sentAt),
	}, nil
}

// NewRequest ...
func NewRequest(pid int, seq, bodySize int) *Request {
	return &Request{
		pid:      pid,
		seq:      seq,
		bodySize: bodySize,
	}
}

// MarkSentAt ...
func (r *Request) MarkSentAt() *Request {
	r.sentAt = time.Now()
	return r
}
