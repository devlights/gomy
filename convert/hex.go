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
//   - lengthを指定した場合、その長さに合うようにゼロパディングします. 0の場合はパディング無しで変換されます.
//     (e.g. length=4 で 16進数 FF を指定した場合 0255 となります.)
func Hex2Dec(val string, prefix string, length int) (string, error) {
	if val == "" {
		return "", nil
	}

	val = strings.ToLower(val)
	if strings.HasPrefix(val, "0x") {
		val = strings.Replace(val, "0x", "", 1)
	}

	num, err := strconv.ParseInt(val, 16, 64)
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

// Hex2Bin -- 指定された16進数文字列を2進数文字列にします.
//
// - 空文字を指定した場合は空文字が返ります.
//
// - 変換に失敗した場合は err に値が設定されます.
//
// - prefixを指定した場合、変換後の文字列の先頭に付与します.
//
// - lengthを指定した場合、その長さに合うようにゼロパディングします.
//   - 0の場合はパディング無しで変換されます.
//   - -1の場合は8の倍数でゼロパディングします.
func Hex2Bin(val string, prefix string, length int) (string, error) {
	if val == "" {
		return "", nil
	}

	val = strings.ToLower(val)
	if strings.HasPrefix(val, "0x") {
		val = strings.Replace(val, "0x", "", 1)
	}

	num, err := strconv.ParseInt(val, 16, 64)
	if err != nil {
		return "", err
	}

	format := "%s%b"
	result := fmt.Sprintf(format, prefix, num)

	switch {
	case length > 0:
		result = fmt.Sprintf("%s%0*b", prefix, length, num)
	case length < 0:
		strBin := fmt.Sprintf("%b", num)
		strLen := len(strBin)

		actualLength := 0
		for i := 0; ; i++ {
			v := 8 * i
			if strLen <= v {
				actualLength = v
				break
			}
		}

		result = fmt.Sprintf("%s%0*b", prefix, actualLength, num)
	}

	return result, nil
}
