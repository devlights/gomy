package fileio

import (
	"os"
	"testing"
)

func TestTempDir(t *testing.T) {
	type (
		testin struct {
			pattern string
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
				pattern: "TestTempDir",
			},
			out: testout{},
		},
	}

	var targetDir string
	for _, c := range cases {
		func() {
			dir, removeFn, err := TempDir(c.in.pattern)
			if err != nil {
				t.Error(err)
			}

			defer removeFn()

			if _, statErr := os.Stat(dir); os.IsNotExist(statErr) {
				t.Errorf("dir is not exists [%s]", dir)
			}

			targetDir = dir
		}()
	}

	if targetDir == "" {
		t.Errorf("targetDir is empty.")
	}

	if _, statErr := os.Stat(targetDir); !os.IsNotExist(statErr) {
		t.Errorf("dir is exists [%s]", targetDir)
	}
}
