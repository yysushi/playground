package main_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestProto(t *testing.T) {
	// FullName
	now := timestamppb.New(time.Now())
	fmt.Printf("%#v\n", now.ProtoReflect().Descriptor())
	assert.IsType(t, now, new(timestamppb.Timestamp))
	assert.Equal(t, "google.protobuf.Timestamp", string(now.ProtoReflect().Descriptor().FullName()))

	any, _ := anypb.New(now)
	assert.IsType(t, any, new(anypb.Any))
	assert.Equal(t, "google.protobuf.Any", string(any.ProtoReflect().Descriptor().FullName()))
}
