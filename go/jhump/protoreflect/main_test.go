package main_test

import (
	"testing"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoprint"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
	// "google.golang.org/protobuf/reflect/protodesc"
)

const expected string = `syntax = "proto3";

package google.protobuf;

option cc_enable_arenas = true;

option csharp_namespace = "Google.Protobuf.WellKnownTypes";

option go_package = "google.golang.org/protobuf/types/known/timestamppb";

option java_multiple_files = true;

option java_outer_classname = "TimestampProto";

option java_package = "com.google.protobuf";

option objc_class_prefix = "GPB";

message Timestamp {
  int64 seconds = 1;

  int32 nanos = 2;
}
`

func TestProtoFile(t *testing.T) {
	var fd *desc.FileDescriptor
	var err error
	fd, err = desc.WrapFile(timestamppb.File_google_protobuf_timestamp_proto)
	if err != nil {
		t.Fatal(err)
	}
	var p protoprint.Printer
	var protostring string
	protostring, err = p.PrintProtoToString(fd)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expected, protostring)
}
