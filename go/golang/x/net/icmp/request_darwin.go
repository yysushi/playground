package icmp

import (
	"bytes"

	"golang.org/x/net/icmp"
)

func (r *Request) embedMetaData(buf *bytes.Buffer) {
}

func (r *Request) paddingBodySize() int {
	return r.bodySize
}

func (r *Request) decodePID(m *icmp.Echo) (int, error) {
	return m.ID, nil
}
