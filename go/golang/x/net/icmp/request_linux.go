package icmp

import (
	"bytes"
	"encoding/binary"

	"golang.org/x/net/icmp"
)

func (r *Request) embedMetaData(buf *bytes.Buffer) {
	// we can't use ID field in a header of ICMP packet. this is a constraint in linux environment.
	// int type, which is return value by os.Getpid(), is not proper to embed because the size depends on OS.
	// man 5 proc says that PID_MAX_LIMIT is 2^22 in 64 bit machine. it is smaller than 2^31.
	// hence, we convert a process ID to uint32 and embed it to the ICMP payload.
	embeddedPID := int32(r.pid)
	binary.Write(buf, binary.LittleEndian, embeddedPID)
}

func (r *Request) paddingBodySize() int {
	return r.bodySize - 4 // 4 byte = 32 bit / 8
}

func (r *Request) decodePID(m *icmp.Echo) (int, error) {
	buf := bytes.NewBuffer(m.Data[:4])
	var pid int32
	err := binary.Read(buf, binary.LittleEndian, &pid)
	if err != nil {
		return 0, err
	}
	return int(pid), nil
}
