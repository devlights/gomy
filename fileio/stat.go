package fileio

import "io/fs"

// Readable は、指定された fs.FileInfo のパーミッションが読み取り可能かどうかを返します.
func Readable(fi fs.FileInfo) bool {
	return fi.Mode().Perm()&0400 == 0400
}

// Writable は、指定された fs.FileInfo のパーミッションが書き込み可能かどうかを返します.
func Writable(fi fs.FileInfo) bool {
	return fi.Mode().Perm()&0200 == 0200
}

// Executable は、指定された fs.FileInfo のパーミッションが実行可能かどうかを返します.
func Executable(fi fs.FileInfo) bool {
	return fi.Mode().Perm()&0100 == 0100
}
