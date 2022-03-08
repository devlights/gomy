package iter

type (
	ZipItem struct {
		Item1, Item2 interface{}
	}
)

// Zip は、指定された２つのスライスから項目を取り出し返します。
//
// python の zip 関数と同じです。
func Zip(a, b []interface{}) []ZipItem {
	var (
		result = make([]ZipItem, 0)
	)

	if a == nil || b == nil {
		return result
	}

	if len(a) == 0 || len(b) == 0 {
		return result
	}

	length := len(a)
	if len(b) < length {
		length = len(b)
	}

	for i := 0; i < length; i++ {
		item := ZipItem{a[i], b[i]}
		result = append(result, item)
	}

	return result
}
