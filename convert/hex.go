package convert

import (
	"fmt"
	"strconv"
	"strings"
)

// Hex2Dec -- 指定された16進数文字列を10進数文字列にします.
//
// - 空文字を指定した場合は空文字が返ります.
//
// - 変換に失敗した場合は err に値が設定されます.
//
// - prefixを指定した場合、変換後の文字列の先頭に付与します.
//
// - lengthを指定した場合、その長さに合うようにゼロパディングします. 0の場合はパディング無しで変換されます.
//   (e.g. length=4 で 16進数 FF を指定した場合 0255 となります.)
func Hex2Dec(val string, prefix string, length int) (string, error) {
	if val == "" {
		return "", nil
	}

	val = strings.ToLower(val)
	if strings.HasPrefix(val, "0x") {
		val = strings.Replace(val, "0x", "", 1)
	}

	num, err := strconv.ParseInt(val, 16, 32)
	if err != nil {
		return "", err
	}

	format := "%s%d"
	result := fmt.Sprintf(format, prefix, num)

	if length > 0 {
		format = "%s" + "%0" + strconv.Itoa(length) + "d"
		result = fmt.Sprintf(format, prefix, num)
	}

	return result, nil
}

// Hex2Bin -- 指定された16進数文字列を2進数文字列にします.
//
// - 空文字を指定した場合は空文字が返ります.
//
// - 変換に失敗した場合は err に値が設定されます.
//
// - prefixを指定した場合、変換後の文字列の先頭に付与します.
//
// - lengthを指定した場合、その長さに合うようにゼロパディングします. 0の場合はパディング無しで変換されます.
//   (e.g. length=4 で 16進数 8 を指定した場合 0100 となります.)
func Hex2Bin(val string, prefix string, length int) (string, error) {
	if val == "" {
		return "", nil
	}

	val = strings.ToLower(val)
	if strings.HasPrefix(val, "0x") {
		val = strings.Replace(val, "0x", "", 1)
	}

	num, err := strconv.ParseInt(val, 16, 32)
	if err != nil {
		return "", err
	}

	format := "%s%b"
	result := fmt.Sprintf(format, prefix, num)

	if length > 0 {
		format = "%s" + "%0" + strconv.Itoa(length) + "b"
		result = fmt.Sprintf(format, prefix, num)
	}

	return result, nil
}
