package stat

import "io/fs"

type userStat struct {
	fi fs.FileInfo
}

// User は、指定された fs.FileInfo からユーザの 読み取り・書き込み・実行の権限状態を判定します。
func User(fi fs.FileInfo) Stat {
	return userStat{fi}
}

func (me userStat) CanRead() bool {
	return me.fi.Mode().Perm()&0400 == 0400
}

func (me userStat) CanWrite() bool {
	return me.fi.Mode().Perm()&0200 == 0200
}

func (me userStat) CanExecute() bool {
	return me.fi.Mode().Perm()&0100 == 0100
}
