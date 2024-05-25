package main_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	// _ "google.golang.org/protobuf/types/known/emptypb"
)

func TestProto(t *testing.T) {
	// FullName
	now := timestamppb.New(time.Now())
	fmt.Printf("%#v\n", now.ProtoReflect().Descriptor())
	assert.IsType(t, now, new(timestamppb.Timestamp))
	assert.Equal(t, "google.protobuf.Timestamp", string(now.ProtoReflect().Descriptor().FullName()))

	any, err := anypb.New(now)
	require.NoError(t, err)
	assert.IsType(t, any, new(anypb.Any))
	assert.Equal(t, "google.protobuf.Any", string(any.ProtoReflect().Descriptor().FullName()))
}

func TestRegistryNotFound(t *testing.T) {
	_, err := protoregistry.GlobalTypes.FindMessageByURL("googleapis.com/google.protobuf.Empty")
	assert.ErrorIs(t, protoregistry.NotFound, err)
}

func TestRegistry(t *testing.T) {
	msgType, err := protoregistry.GlobalTypes.FindMessageByURL("googleapis.com/google.protobuf.Timestamp")
	require.NoError(t, err)
	// pmsg := msgType.New().Interface()
	pmsg := msgType.Zero().Interface()
	assert.Equal(t, protoreflect.FullName("google.protobuf.Timestamp"), msgType.Descriptor().FullName())
	assert.Equal(t, protoreflect.FullName("google.protobuf.Timestamp"), pmsg.ProtoReflect().Descriptor().FullName())
	timestamp := pmsg.(*timestamppb.Timestamp)
	assert.Equal(t, timestamp.AsTime(), time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC))
}
