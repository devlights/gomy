package bcd

// ToBcd -- 指定された数値とバイト数でBCDを生成して返します.
func ToBcd(num uint64, byteCount int) []byte {
	var (
		bcd = make([]byte, byteCount)
	)

	for index := 1; index <= byteCount; index++ {
		mod := num % 100

		digit2 := mod % 10
		digit1 := (mod - digit2) / 10

		bcd[(byteCount - index)] = byte((digit1 * 16) + digit2)

		num = (num - mod) / 100
	}

	return bcd
}

// ToUInt64 -- bcd を uint64 に変換します.
func ToUInt64(bcd []byte) uint64 {
	var (
		result uint64 = 0
	)

	for _, b := range bcd {
		digit1 := b >> 4
		digit2 := b & 0x0f

		result = (result * 100) + uint64(digit1*10) + uint64(digit2)
	}

	return result
}
