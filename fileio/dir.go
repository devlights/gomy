package fileio

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type (
	// FileInfoEx -- os.FileInfo に 追加 のプロパティを付属させている構造体です.
	FileInfoEx struct {
		os.FileInfo
		FullPath string // フルパス
	}
)

// ReadDir -- ioutil.ReadDir() の 結果に追加情報を付与したデータを返します.
//
// 動作仕様は ioutil.ReadDir() と同じです.
func ReadDir(dirPath string) ([]FileInfoEx, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	absPath, err := filepath.Abs(dirPath)
	if err != nil {
		return nil, err
	}

	result := make([]FileInfoEx, 0, len(files))
	for _, fi := range files {
		fiEx := FileInfoEx{
			FileInfo: fi,
			FullPath: filepath.Join(absPath, fi.Name()),
		}

		result = append(result, fiEx)
	}

	return result, nil
}
