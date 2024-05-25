package main_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"

	_ "github.com/yysushi/playground/proto/protocolbuffers/go/internal/gen/example/v1"
)

func TestA(t *testing.T) {
	var enumFullName = protoreflect.FullName("example.v1.HogeEnum")
	var enumValueName = protoreflect.Name("HOGE_ENUM_DISPLAY")
	var extensionFullName = protoreflect.FullName("example.v1.string_name")
	var err error
	var enumType protoreflect.EnumType
	var enumValueDescriptor protoreflect.EnumValueDescriptor
	var extensionType protoreflect.ExtensionType

	enumType, err = protoregistry.GlobalTypes.FindEnumByName(enumFullName)
	require.NoError(t, err)
	extensionType, err = protoregistry.GlobalTypes.FindExtensionByName(extensionFullName)
	require.NoError(t, err)

	enumValueDescriptor = enumType.Descriptor().Values().ByName(enumValueName)
	fmt.Printf("%#v\n", proto.GetExtension(enumValueDescriptor.Options(), extensionType))
}
