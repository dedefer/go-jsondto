package jsondto

import (
	"reflect"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

type nillableTypesEmptyOnNilExt struct {
	jsoniter.DecoderExtension
}

func (ext nillableTypesEmptyOnNilExt) DecorateEncoder(t reflect2.Type, e jsoniter.ValEncoder) jsoniter.ValEncoder {
	if t, ok := t.(*reflect2.UnsafeSliceType); ok {
		return sliceEmptyOnNilEncoder{
			ValEncoder: e, type_: t,
			bytes: t.Elem().Kind() == reflect.Uint8,
		}
	}

	if t, ok := t.(*reflect2.UnsafeMapType); ok {
		return mapEmptyOnNilEncoder{ValEncoder: e, type_: t}
	}

	return e
}

type sliceEmptyOnNilEncoder struct {
	jsoniter.ValEncoder
	type_ *reflect2.UnsafeSliceType
	bytes bool
}

func (e sliceEmptyOnNilEncoder) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	if e.type_.UnsafeIsNil(ptr) {
		if e.bytes {
			stream.WriteString("")
			return
		}

		stream.WriteEmptyArray()
		return
	}

	e.ValEncoder.Encode(ptr, stream)
}

type mapEmptyOnNilEncoder struct {
	jsoniter.ValEncoder
	type_ *reflect2.UnsafeMapType
}

func (e mapEmptyOnNilEncoder) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	if e.type_.UnsafeIsNil(ptr) {
		stream.WriteEmptyObject()
		return
	}

	e.ValEncoder.Encode(ptr, stream)
}
