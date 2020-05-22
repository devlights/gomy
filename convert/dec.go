package convert

import (
	"fmt"
	"strconv"
)

// Dec2Hex -- 指定された10進数文字列を16進数文字列にします.
//
// - 空文字を指定した場合は空文字が返ります.
//
// - 変換に失敗した場合は err に値が設定されます.
//
// - prefixを指定した場合、変換後の文字列の先頭に付与します.
//
// - lengthを指定した場合、その長さに合うようにゼロパディングします. 0の場合はパディング無しで変換されます.
//   (e.g. length=4 で 10進数 777 を指定した場合 0309 となります.)
func Dec2Hex(val string, prefix string, length int) (string, error) {
	if val == "" {
		return "", nil
	}

	num, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return "", err
	}

	format := "%s%X"
	result := fmt.Sprintf(format, prefix, num)

	if length > 0 {
		format = "%s" + "%0" + strconv.Itoa(length) + "X"
		result = fmt.Sprintf(format, prefix, num)
	}

	return result, nil
}

// Dec2Bin -- 指定された10進数文字列を16進数文字列にします.
//
// - 空文字を指定した場合は空文字が返ります.
//
// - 変換に失敗した場合は err に値が設定されます.
//
// - prefixを指定した場合、変換後の文字列の先頭に付与します.
//
// - lengthを指定した場合、その長さに合うようにゼロパディングします. 0の場合はパディング無しで変換されます.
//   (e.g. length=4 で 10進数 4 を指定した場合 0100 となります.)
func Dec2Bin(val string, prefix string, length int) (string, error) {
	if val == "" {
		return "", nil
	}

	num, err := strconv.ParseInt(val, 10, 32)
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
