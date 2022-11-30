package stat

import "io/fs"

type groupStat struct {
	fi fs.FileInfo
}

// Group は、指定された fs.FileInfo からグループの 読み取り・書き込み・実行の権限状態を判定します。
func Group(fi fs.FileInfo) Stat {
	return groupStat{fi}
}

func (me groupStat) CanRead() bool {
	return me.fi.Mode().Perm()&0040 == 0040
}

func (me groupStat) CanWrite() bool {
	return me.fi.Mode().Perm()&0020 == 0020
}

func (me groupStat) CanExecute() bool {
	return me.fi.Mode().Perm()&0010 == 0010
}
