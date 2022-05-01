package gormfilter

import (
	"fmt"

	"github.com/robertsong9972/gorm-gen/util"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

func GoType(t descriptor.FieldDescriptorProto_Type) string {
	return util.GoCamelCase(goType(t))
}

func goType(t descriptor.FieldDescriptorProto_Type) string {
	switch t {
	case descriptor.FieldDescriptorProto_TYPE_INT64:
		return "int64"
	case descriptor.FieldDescriptorProto_TYPE_UINT64:
		return "uint64"
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		return "string"
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		return "bytes"
	case descriptor.FieldDescriptorProto_TYPE_UINT32:
		return "uint32"
	case descriptor.FieldDescriptorProto_TYPE_INT32:
		return "int32"

	default:
		panic(fmt.Sprintf("invalid type|type=%d", t))
	}
}

func isNumber(t descriptor.FieldDescriptorProto_Type) bool {
	if t == descriptor.FieldDescriptorProto_TYPE_INT32 ||
		t == descriptor.FieldDescriptorProto_TYPE_INT64 ||
		t == descriptor.FieldDescriptorProto_TYPE_UINT64 ||
		t == descriptor.FieldDescriptorProto_TYPE_UINT32 {
		return true
	}
	return false
}

func isString(t descriptor.FieldDescriptorProto_Type) bool {
	return t == descriptor.FieldDescriptorProto_TYPE_STRING
}

func notLastIndex(i int, length int) bool {
	return i < length-1
}
