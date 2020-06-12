package fileio

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestReadDir(t *testing.T) {
	type (
		testin struct {
			dirpath string
		}
		testout struct {
			results []string
		}
		testcase struct {
			name string
			in   testin
			out  testout
		}
	)

	cwd, _ := os.Getwd()
	absDirPath, _ := filepath.Abs(cwd)

	cases := []testcase{
		{
			name: "current dir",
			in:   testin{dirpath: "."},
			out:  testout{getFileList(absDirPath)},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			files, err := ReadDir(c.in.dirpath)
			if err != nil {
				t.Error(err)
			}

			for _, fi := range files {
				found := false
				for _, r := range c.out.results {
					if fi.FullPath == r {
						found = true
						break
					}
				}

				if !found {
					t.Errorf("[want] contains\t[got] not contains (%v)", fi.FullPath)
				}
			}
		})
	}
}

func getFileList(p string) []string {
	r := make([]string, 0, 0)

	files, _ := ioutil.ReadDir(p)
	for _, fi := range files {
		r = append(r, filepath.Join(p, fi.Name()))
	}

	return r
}
