package fileio

import (
	"os"
	"testing"
)

func TestTempFileName(t *testing.T) {
	type (
		testin struct {
			prefix string
		}
		testout struct {
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	cases := []testcase{
		{
			in: testin{
				prefix: "TestTempFileName",
			},
			out: testout{},
		},
	}

	for _, c := range cases {
		func() {
			dir, fn, err := TempDir(c.in.prefix)
			if err != nil {
				t.Error(err)
			}

			defer fn()

			fpath := TempFileName(dir, c.in.prefix)
			if fpath == "" {
				t.Errorf("name is empty")
			}

			if _, statErr := os.Stat(fpath); !os.IsNotExist(statErr) {
				t.Errorf("file is exists [%s]", fpath)
			}
		}()
	}
}
