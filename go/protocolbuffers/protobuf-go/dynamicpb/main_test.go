package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestA(t *testing.T) {
	msgd := timestamppb.File_google_protobuf_timestamp_proto.Messages().Get(0)
	msg := dynamicpb.NewMessage(msgd)
	assert.Equal(t, protoreflect.FullName("google.protobuf.Timestamp"), msg.Descriptor().FullName())
	assert.Equal(t, protoreflect.FullName("google.protobuf.Timestamp"), msg.Interface().ProtoReflect().Descriptor().FullName())
	_, ok := msg.Interface().(*timestamppb.Timestamp)
	assert.False(t, ok)
}
