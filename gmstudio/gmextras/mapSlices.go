package gmextras

import (
	"reflect"
	"unsafe"
)

func MapSliceFloat32AsByte(src []float32) []byte {
	dst := (*((*[1]byte)(unsafe.Pointer(&src[0]))))[:]
	dstHeader := (*reflect.SliceHeader)(unsafe.Pointer(&dst))
	dstHeader.Len = len(src) * 4 // 1 float32 = 4 bytes
	dstHeader.Cap = cap(src) * 4 // 1 float32 = 4 bytes
	return dst
}
