package zeromemcpy

import (
	"unsafe"
)

// B2S -- []byte] から string へメモリコピー無しで変換します。
//
// REREFENCES
//   - https://cs.opensource.google/go/go/+/refs/tags/go1.18.4:src/strings/builder.go;l=47
func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
