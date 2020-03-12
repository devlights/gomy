package jp

type (
	// Encoding は、日本語圏内で利用するエンコーディングを表します。
	Encoding int
)

//noinspection GoUnusedConst
const (
	Utf8     Encoding = 0x01 // Utf8
	ShiftJis Encoding = 0x10 // Shift-JIS
	EucJp    Encoding = 0x20 // EUC-JP
	Ascii    Encoding = 0xFF // Ascii
)
