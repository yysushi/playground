package main_test

import (
	"net/netip"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefix(t *testing.T) {
	prefix := netip.MustParsePrefix("192.168.100.0/24")
	assert.True(t, prefix.Contains(netip.MustParseAddr("192.168.100.100")))
}

func TestPrefixDump(t *testing.T) {
	prefix := netip.MustParsePrefix("192.168.100.0/24")
	text, _ := prefix.MarshalText()
	assert.Equal(t, "192.168.100.0/24", string(text))

	b, _ := prefix.MarshalBinary()
	assert.Len(t, b, 5)
	assert.Equal(t, "\xc0\xa8d\x00\x18", string(b))
	assert.Equal(t, []byte{0xc0 /* 192 */, 0xa8 /* 168 */, 0x64 /* 100 */, 0x00 /* 0 */, 0x18 /* 24 */}, b)
}
