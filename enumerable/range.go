package enumerable

type (
	// Range は、範囲を表すインターフェースです。
	Range interface {
		// Start は、開始点の値を返します.
		Start() int
		// End は、終了点の値を返します.
		End() int
		// Next は、次の値に進みます. 進むことが出来ない場合は false を返します.
		Next() bool
		// Current は、現在の値を返します.
		Current() int
		// Reset は、現在の値をリセットして開始点の値に戻します. 戻り値は (リセット直前の値, 処理で発生したエラー) です.
		Reset() (int, error)
	}

	enumerableRange struct {
		start, end, current int
	}
)

// NewRange は、指定された値を元に Range を生成して返します.
func NewRange(start, end int) Range {
	return &enumerableRange{
		start:   start,
		end:     end,
		current: start,
	}
}

// Start は、開始店の値を返します。
func (e *enumerableRange) Start() int {
	return e.start
}

// End は、終了点の値を返します。
func (e *enumerableRange) End() int {
	return e.end
}

// Next は、次の値に進みます。進むことが出来ない場合は false を返します。
func (e *enumerableRange) Next() bool {
	if e.current == e.end {
		return false
	}

	e.current++
	return true
}

// Current は、現在の値を返します。
func (e *enumerableRange) Current() int {
	return e.current
}

// Reset は、現在の値をリセットして開始点の値に戻します. 戻り値は (リセット直前の値, 処理で発生したエラー) です。
func (e *enumerableRange) Reset() (int, error) {
	cur := e.current
	e.current = e.start
	return cur, nil
}
