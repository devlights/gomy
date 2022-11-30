package stat_test

import (
	"io/fs"
	"os"
)

func tmpfile() (*os.File, error) {
	return os.CreateTemp(os.TempDir(), "")
}

func tmpAndChmod(perm fs.FileMode) (*os.File, error) {
	f, err := tmpfile()
	if err != nil {
		return nil, err
	}

	if err := f.Chmod(perm); err != nil {
		return nil, err
	}

	return f, nil
}

func tmpAndChmodAndStat(perm fs.FileMode) (*os.File, fs.FileInfo, error) {
	f, err := tmpAndChmod(perm)
	if err != nil {
		return nil, nil, err
	}

	fi, err := f.Stat()
	if err != nil {
		return nil, nil, err
	}

	return f, fi, nil
}
