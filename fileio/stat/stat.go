package stat

// Stat は、ファイルの読み取り・書き込み・実行の権限状態を判定するためのインターフェースを持ちます。
type Stat interface {
	// 読み込み可能かどうか
	CanRead() bool
	// 書き込み可能かどうか
	CanWrite() bool
	// 実行可能かどうか
	CanExecute() bool
}
