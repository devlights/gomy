package zeromemcpy

import (
	"reflect"
	"unsafe"
)

// S2B -- string から []byte へメモリコピー無しで変換します。
//
// REREFENCES
//   - https://stackoverflow.com/questions/59209493/how-to-use-unsafe-get-a-byte-slice-from-a-string-without-memory-copy
func S2B(val string) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&val)).Data)), len(val))
}
