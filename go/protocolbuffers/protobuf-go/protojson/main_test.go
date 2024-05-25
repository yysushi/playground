package main_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestMarshal(t *testing.T) {
	var tm time.Time
	tm, _ = time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
	tmpb := timestamppb.New(tm)
	var v []byte

	v, _ = json.Marshal(tmpb)
	assert.Equal(t, `{"seconds":1136189045}`, string(v))

	v, _ = protojson.Marshal(tmpb)
	assert.Equal(t, `"2006-01-02T08:04:05Z"`, string(v))

	v, _ = prototext.Marshal(tmpb)
	assert.Equal(t, `seconds:1136189045`, string(v))

	v, _ = proto.Marshal(tmpb)
	assert.Equal(t, "\b\xf5\xbc\xe3\x9d\x04", string(v))
}

func TestProtoJson(t *testing.T) {
	d := durationpb.New(90 * time.Second)
	any, _ := anypb.New(d)

	cases := []struct {
		m        proto.Message
		expected string
	}{
		{
			m:        d,
			expected: `"90s"`,
		},
		{
			m:        any,
			expected: `{"@type":"type.googleapis.com/google.protobuf.Duration","value":"90s"}`,
		},
		{
			m:        &emptypb.Empty{},
			expected: `{}`,
		},
	}

	for _, tt := range cases {
		v, err := protojson.Marshal(tt.m)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, tt.expected, string(v))
	}
}
