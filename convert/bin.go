package convert

import (
	"fmt"
	"strconv"
	"strings"
)

// Bin2Dec -- 指定された2進数文字列を10進数文字列にします.
//
// - 空文字を指定した場合は空文字が返ります.
//
// - 変換に失敗した場合は err に値が設定されます.
//
// - prefixを指定した場合、変換後の文字列の先頭に付与します.
//
// - lengthを指定した場合、その長さに合うようにゼロパディングします. 0の場合はパディング無しで変換されます.
//   (e.g. length=4 で 2進数 1000 を指定した場合 0008 となります.)
func Bin2Dec(val string, prefix string, length int) (string, error) {
	if val == "" {
		return "", nil
	}

	val = strings.ToLower(val)
	if strings.HasPrefix(val, "0b") {
		val = strings.Replace(val, "0b", "", 1)
	}

	num, err := strconv.ParseInt(val, 2, 64)
	if err != nil {
		return "", err
	}

	format := "%s%d"
	result := fmt.Sprintf(format, prefix, num)

	if length > 0 {
		result = fmt.Sprintf("%s%0*d", prefix, length, num)
	}

	return result, nil
}

// Bin2Hex -- 指定された2進数文字列を16進数文字列にします.
//
// - 空文字を指定した場合は空文字が返ります.
//
// - 変換に失敗した場合は err に値が設定されます.
//
// - prefixを指定した場合、変換後の文字列の先頭に付与します.
//
// - lengthを指定した場合、その長さに合うようにゼロパディングします. 0の場合はパディング無しで変換されます.
//   (e.g. length=4 で 2進数 1111 を指定した場合 000F となります.)
func Bin2Hex(val string, prefix string, length int) (string, error) {
	if val == "" {
		return "", nil
	}

	val = strings.ToLower(val)
	if strings.HasPrefix(val, "0b") {
		val = strings.Replace(val, "0b", "", 1)
	}

	num, err := strconv.ParseInt(val, 2, 64)
	if err != nil {
		return "", err
	}

	format := "%s%X"
	result := fmt.Sprintf(format, prefix, num)

	if length > 0 {
		result = fmt.Sprintf("%s%0*X", prefix, length, num)
	}

	return result, nil
}
