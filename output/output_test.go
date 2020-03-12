package output

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func init() {
	SetPrefixFormat("%s")
}

func TestHr(t *testing.T) {
	type (
		testin struct {
			buf bytes.Buffer
			fn  func()
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
				buf: bytes.Buffer{},
				fn:  StdoutHr,
			},
			out: testout{},
		},
		{
			in: testin{
				buf: bytes.Buffer{},
				fn:  StderrHr,
			},
			out: testout{},
		},
	}

	for _, c := range cases {
		func() {
			orig := Writer()
			SetWriter(&c.in.buf)
			defer SetWriter(orig)

			// Act
			c.in.fn()

			// Assert
			got := strings.TrimSpace(c.in.buf.String())
			allOk := true
			for _, v := range got {
				if v != '-' {
					allOk = false
					break
				}
			}

			if !allOk {
				t.Error("want: all '-'\tgot: not all '-'")
			}
		}()
	}
}

func TestStdoutl(t *testing.T) {
	type (
		testin struct {
			buf     bytes.Buffer
			prefix  string
			message string
		}
		testout struct {
			want string
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	// Arrange
	cases := []testcase{
		{
			in: testin{
				buf:     bytes.Buffer{},
				prefix:  "test",
				message: "hello",
			},
			out: testout{
				want: fmt.Sprintf("%s %s", "test", "hello"),
			},
		},
	}

	for _, c := range cases {
		func() {
			orig := Writer()
			SetWriter(&c.in.buf)
			defer SetWriter(orig)

			// Act
			Stdoutl(c.in.prefix, c.in.message)

			// Assert
			got := strings.TrimSpace(c.in.buf.String())
			if got != c.out.want {
				t.Errorf("want:%s\tgot:%s", c.out.want, got)
			}
		}()
	}
}

func TestStdoutf(t *testing.T) {
	type (
		testin struct {
			buf     bytes.Buffer
			prefix  string
			format  string
			message string
		}
		testout struct {
			want string
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	// Arrange
	cases := []testcase{
		{
			in: testin{
				buf:     bytes.Buffer{},
				prefix:  "test",
				format:  "%s world",
				message: "hello",
			},
			out: testout{
				want: fmt.Sprintf("%s %s world", "test", "hello"),
			},
		},
	}

	for _, c := range cases {
		func() {
			orig := Writer()
			SetWriter(&c.in.buf)
			defer SetWriter(orig)

			// Act
			Stdoutf(c.in.prefix, c.in.format, c.in.message)

			// Assert
			got := strings.TrimSpace(c.in.buf.String())
			if got != c.out.want {
				t.Errorf("want:%s\tgot:%s", c.out.want, got)
			}
		}()
	}
}

func TestStderrl(t *testing.T) {
	type (
		testin struct {
			buf     bytes.Buffer
			prefix  string
			message string
		}
		testout struct {
			want string
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	// Arrange
	cases := []testcase{
		{
			in: testin{
				buf:     bytes.Buffer{},
				prefix:  "test",
				message: "hello",
			},
			out: testout{
				want: fmt.Sprintf("%s %s", "test", "hello"),
			},
		},
	}

	for _, c := range cases {
		func() {
			orig := ErrWriter()
			SetErrWriter(&c.in.buf)
			defer SetErrWriter(orig)

			// Act
			Stderrl(c.in.prefix, c.in.message)

			// Assert
			got := strings.TrimSpace(c.in.buf.String())
			if got != c.out.want {
				t.Errorf("want:%s\tgot:%s", c.out.want, got)
			}
		}()
	}
}

func TestStderrf(t *testing.T) {
	type (
		testin struct {
			buf     bytes.Buffer
			prefix  string
			format  string
			message string
		}
		testout struct {
			want string
		}
		testcase struct {
			in  testin
			out testout
		}
	)

	// Arrange
	cases := []testcase{
		{
			in: testin{
				buf:     bytes.Buffer{},
				prefix:  "test",
				format:  "%s world",
				message: "hello",
			},
			out: testout{
				want: fmt.Sprintf("%s %s world", "test", "hello"),
			},
		},
	}

	for _, c := range cases {
		func() {
			orig := ErrWriter()
			SetErrWriter(&c.in.buf)
			defer SetErrWriter(orig)

			// Act
			Stderrf(c.in.prefix, c.in.format, c.in.message)

			// Assert
			got := strings.TrimSpace(c.in.buf.String())
			if got != c.out.want {
				t.Errorf("want:%s\tgot:%s", c.out.want, got)
			}
		}()
	}
}
