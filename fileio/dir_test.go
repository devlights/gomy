package fileio

import (
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
			out:  testout{files(absDirPath)},
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

func files(p string) []string {
	r := make([]string, 0)

	entries, _ := os.ReadDir(p)
	for _, e := range entries {
		r = append(r, filepath.Join(p, e.Name()))
	}

	return r
}
