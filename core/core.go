package core

import (
	"errors"
	"reflect"
	"unsafe"
)

type FormFile string

type FormMem []byte

type H map[string]interface{}

type A []interface{}

var ErrUnknownType = errors.New("unknown type")

func LoopElem(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return v
		}
		v = v.Elem()
	}

	return v
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func StringToBytes(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Len = sh.Len
	bh.Cap = sh.Len
	return b
}

func NewPtrVal(defValue interface{}) interface{} {
	p := reflect.New(reflect.TypeOf(defValue))
	p.Elem().Set(reflect.ValueOf(defValue))
	return p.Interface()
}
