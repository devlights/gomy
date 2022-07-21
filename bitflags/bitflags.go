package bitflags

// Has -- 指定したビットが立ってるかどうかを返します。
func Has[T ~int](v T, flag T) bool {
	return v|flag == v
}

// Unset -- 指定したビットを落とします。
func Unset[T ~int](v *T, flag T) {
	*v = *v &^ flag
}

// Force -- 指定したビットで上書きします。
func Force[T ~int](v *T, flag T) {
	*v = flag
}

// Set -- 指定したビットを立てます。
func Set[T ~int](v *T, flags ...T) {
	if len(flags) == 0 {
		return
	}

	for _, flag := range flags {
		*v = *v | flag
	}
}
