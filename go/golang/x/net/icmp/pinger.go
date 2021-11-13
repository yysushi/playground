package icmp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"runtime"
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

func (r *Request) embedMetaData(buf *bytes.Buffer) {
	switch runtime.GOOS {
	case "linux":
		// we can't use ID field in a header of ICMP packet. this is a constraint in linux environment.
		// int type, which is return value by os.Getpid(), is not proper to embed because the size depends on OS.
		// man 5 proc says that PID_MAX_LIMIT is 2^22 in 64 bit machine. it is smaller than 2^31.
		// hence, we convert a process ID to uint32 and embed it to the ICMP payload.
		embeddedPID := int32(r.pid)
		binary.Write(buf, binary.LittleEndian, embeddedPID)
	default:
	}
}

func (r *Request) paddingBodySize() int {
	switch runtime.GOOS {
	case "linux":
		return r.bodySize - 4 // 4 byte = 32 bit / 8
	default:
		return r.bodySize
	}
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

func (r *Request) decodePID(m *icmp.Echo) (int, error) {
	switch runtime.GOOS {
	case "linux":
	default:
		return m.ID, nil
	}
	// linux
	buf := bytes.NewBuffer(m.Data[:4])
	var pid int32
	err := binary.Read(buf, binary.LittleEndian, &pid)
	if err != nil {
		return 0, err
	}
	return int(pid), nil
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
func NewRequest(pid, seq, bodySize int) *Request {
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
