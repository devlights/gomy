package stat

import "io/fs"

type otherStat struct {
	fi fs.FileInfo
}

// Other は、指定された fs.FileInfo から他者の 読み取り・書き込み・実行の権限状態を判定します。
func Other(fi fs.FileInfo) Stat {
	return otherStat{fi}
}

func (me otherStat) CanRead() bool {
	return me.fi.Mode().Perm()&0004 == 0004
}

func (me otherStat) CanWrite() bool {
	return me.fi.Mode().Perm()&0002 == 0002
}

func (me otherStat) CanExecute() bool {
	return me.fi.Mode().Perm()&0001 == 0001
}
