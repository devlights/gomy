package slices

// Filter -- 要素をフィルタリングします。（元のスライスはそのまま）
func Filter(src []interface{}, fn func(v interface{}) bool) []interface{} {
	var (
		dst = make([]interface{}, 0)
	)

	return _filter(src, dst, fn)
}

// FilterD -- 要素をフィルタリングします。（元のスライスを変更します。破壊的変更）
func FilterD(src []interface{}, fn(func(v interface{}) bool)) []interface{} {
	var (
		dst = src[:0]
	)

	return _filter(src, dst, fn)
}

func _filter(src, dst []interface{}, fn(func(v interface{}) bool)) []interface{} {
	for _, v := range src {
		if fn(v) {
			dst = append(dst, v)
		}
	}

	return dst
}