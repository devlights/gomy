package exts

import "fmt"

type (
	// Num は、int のエイリアス型です。
	// int を扱う上で少し便利なメソッドを提供します。
	Num int
)

// ToStr は、自身の値を文字列で返します。
func (me Num) ToStr() string {
	return fmt.Sprintf("%d", int(me))
}

// Times は、自身の値と同じ回数ループし、ループの度に fn を実行します。
// fn の引数に指定されるのは index となります。
func (me Num) Times(fn func(i int)) {
	for i := 0; i < int(me); i++ {
		fn(i)
	}
}

// Upto は、自身の値から to の値までの回数ループし、ループの度に fn を実行します。
// fn の引数に指定されるのは index となります。
// Downto との違いは、Uptoはインデックスが正の方向に上がっていく(i++)のに対し、Downto は負の方向へ下がっていきます。(i--)
func (me Num) Upto(to int, fn func(i int)) {
	for i := int(me); i <= to; i++ {
		fn(i)
	}
}

// Downto は、自身の値から to 値までの回数ループし、ループの度に fn を実行します。
// fn の引数に指定されるのは index となります。
// Upto との違いは、Uptoはインデックスが正の方向に上がっていく(i++)のに対し、Downto は負の方向へ下がっていきます。(i--)
func (me Num) Downto(to int, fn func(i int)) {
	for i := int(me); i >= to; i-- {
		fn(i)
	}
}
