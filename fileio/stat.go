package fileio

import (
	"io/fs"

	"github.com/devlights/gomy/fileio/stat"
)

// Readable は、指定された fs.FileInfo のパーミッションが読み取り可能かどうかを返します.
func Readable(fi fs.FileInfo) bool {
	return stat.User(fi).CanRead()
}

// Writable は、指定された fs.FileInfo のパーミッションが書き込み可能かどうかを返します.
func Writable(fi fs.FileInfo) bool {
	return stat.User(fi).CanWrite()
}

// Executable は、指定された fs.FileInfo のパーミッションが実行可能かどうかを返します.
func Executable(fi fs.FileInfo) bool {
	return stat.User(fi).CanExecute()
}
