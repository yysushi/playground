package main_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestAny(t *testing.T) {
	now := timestamppb.New(time.Now())
	any, _ := anypb.New(now)
	assert.Equal(t, "type.googleapis.com/google.protobuf.Timestamp", any.TypeUrl)
	assert.Equal(t, "google.protobuf.Timestamp", string(any.MessageName()))
	assert.Equal(t, "google.protobuf.Timestamp", string(now.ProtoReflect().Descriptor().FullName()))
	assert.Equal(t, "Timestamp", string(now.ProtoReflect().Descriptor().Name()))
	assert.Equal(t, "google.protobuf.Any", string(any.ProtoReflect().Descriptor().FullName()))
	// unmarshal by using type url
	//  1 anypb.UnmarshalTo
	var someTimestamppb = new(timestamppb.Timestamp)
	assert.NoError(t, any.UnmarshalTo(someTimestamppb))
	//  2 anypb.UnmarshalNew
	pm, _ := any.UnmarshalNew()
	assert.IsType(t, pm, new(timestamppb.Timestamp))
	// type check without unmarshal
	assert.True(t, any.MessageIs(new(timestamppb.Timestamp)))
}
