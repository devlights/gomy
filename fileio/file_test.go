package fileio

import (
	"bufio"
	"io"
	"os"
	"testing"

	"github.com/devlights/gomy/fileio/jp"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

type (
	setupInfo struct {
		dir       string
		sjisFile  string
		eucjpFile string
	}
)

func TestOpenRead(t *testing.T) {
	type (
		testin struct {
			fpath string
			enc   jp.Encoding
		}
		testout struct {
			want string
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	info, err := setup()
	if err != nil {
		t.Error(err)
	}

	defer teardown(info)

	cases := []testcase{
		{
			in: testin{
				fpath: info.sjisFile,
				enc:   jp.ShiftJis,
			},
			out: testout{
				want: "こんにちわWorld",
			},
		},
		{
			in: testin{
				fpath: info.eucjpFile,
				enc:   jp.EucJp,
			},
			out: testout{
				want: "こんにちわWorld",
			},
		},
	}

	for _, c := range cases {
		func() {
			reader, releaseFn, ioErr := OpenRead(c.in.fpath, c.in.enc)
			if ioErr != nil {
				t.Error(ioErr)
			}

			//noinspection GoUnhandledErrorResult
			defer releaseFn()

			all, ioErr := io.ReadAll(reader)
			if ioErr != nil {
				t.Error(ioErr)
			}

			data := string(all)
			if data != c.out.want {
				t.Errorf("want: %s\tgot: %s", c.out.want, data)
			}
		}()
	}
}

func TestOpenWrite(t *testing.T) {
	type (
		testin struct {
			fpath string
			enc   jp.Encoding
			data  string
		}
		testout struct {
			want string
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	info, err := setup()
	if err != nil {
		t.Error(err)
	}

	defer teardown(info)

	cases := []testcase{
		{
			in: testin{
				fpath: info.sjisFile,
				enc:   jp.ShiftJis,
				data:  "Hello世界",
			},
			out: testout{
				want: "Hello世界",
			},
		},
		{
			in: testin{
				fpath: info.eucjpFile,
				enc:   jp.EucJp,
				data:  "Hello世界",
			},
			out: testout{
				want: "Hello世界",
			},
		},
	}

	for _, c := range cases {
		func() {
			writer, releaseFn, ioErr := OpenWrite(c.in.fpath, c.in.enc)
			if ioErr != nil {
				t.Error(ioErr)
			}

			//noinspection GoUnhandledErrorResult
			defer releaseFn()

			_, ioErr = writer.Write([]byte(c.in.data))
			if ioErr != nil {
				t.Error(ioErr)
			}
		}()

		func() {
			reader, f, _ := OpenRead(c.in.fpath, c.in.enc)

			//noinspection GoUnhandledErrorResult
			defer f()

			all, _ := io.ReadAll(reader)

			data := string(all)
			if data != c.out.want {
				t.Errorf("want: %s\tgot %s", c.out.want, data)
			}
		}()
	}
}

func TestOpenAppend(t *testing.T) {
	type (
		testin struct {
			fpath string
			enc   jp.Encoding
			data  string
		}
		testout struct {
			want string
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	info, err := setup()
	if err != nil {
		t.Error(err)
	}

	defer teardown(info)

	cases := []testcase{
		{
			in: testin{
				fpath: info.sjisFile,
				enc:   jp.ShiftJis,
				data:  "Hello世界",
			},
			out: testout{
				want: "こんにちわWorldHello世界",
			},
		},
		{
			in: testin{
				fpath: info.eucjpFile,
				enc:   jp.EucJp,
				data:  "Hello世界",
			},
			out: testout{
				want: "こんにちわWorldHello世界",
			},
		},
	}

	for _, c := range cases {
		func() {
			writer, releaseFn, ioErr := OpenAppend(c.in.fpath, c.in.enc)
			if ioErr != nil {
				t.Error(ioErr)
			}

			//noinspection GoUnhandledErrorResult
			defer releaseFn()

			_, ioErr = writer.Write([]byte(c.in.data))
			if ioErr != nil {
				t.Error(ioErr)
			}
		}()

		func() {
			reader, f, _ := OpenRead(c.in.fpath, c.in.enc)

			//noinspection GoUnhandledErrorResult
			defer f()

			all, _ := io.ReadAll(reader)

			data := string(all)
			if data != c.out.want {
				t.Errorf("want: %s\tgot %s", c.out.want, data)
			}
		}()
	}
}

func teardown(s *setupInfo) {
	_ = os.RemoveAll(s.dir)
}

func setup() (*setupInfo, error) {
	dir, err := os.MkdirTemp("", "gomy")
	if err != nil {
		return nil, err
	}

	sjisFile, err := writeSjisFile(dir)
	if err != nil {
		return nil, err
	}

	eucjpFile, err := writeEucJpFile(dir)
	if err != nil {
		return nil, err
	}

	r := &setupInfo{
		dir:       dir,
		sjisFile:  sjisFile,
		eucjpFile: eucjpFile,
	}

	return r, nil
}

func writeSjisFile(dir string) (string, error) {

	file, err := os.CreateTemp(dir, "gomy-sjis")
	if err != nil {
		return "", err
	}

	//noinspection GoUnhandledErrorResult
	defer file.Close()

	writer := bufio.NewWriter(transform.NewWriter(file, japanese.ShiftJIS.NewEncoder()))
	_, err = writer.WriteString("こんにちわWorld")
	if err != nil {
		return "", err
	}

	err = writer.Flush()
	if err != nil {
		return "", err
	}

	return file.Name(), nil
}

func writeEucJpFile(dir string) (string, error) {
	file, err := os.CreateTemp(dir, "gomy-eucjp")
	if err != nil {
		return "", err
	}

	//noinspection GoUnhandledErrorResult
	defer file.Close()

	writer := bufio.NewWriter(transform.NewWriter(file, japanese.EUCJP.NewEncoder()))
	_, err = writer.WriteString("こんにちわWorld")
	if err != nil {
		return "", err
	}

	err = writer.Flush()
	if err != nil {
		return "", err
	}

	return file.Name(), nil
}
