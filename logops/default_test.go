package logops

import (
	"bytes"
	"log"
	"testing"
)

func TestDefaultAppLogger(t *testing.T) {
	cases := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "helloworld",
			in:   "helloworld",
			out:  "helloworld\n",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var (
				buf = new(bytes.Buffer)
			)

			a, _, _ := Default.Logger(true, func(a, _, _ *log.Logger) {
				a.SetOutput(buf)
			})

			a.Print(c.in)

			if buf.String() != c.out {
				t.Errorf("\nwant: %v(%v)\ngot: %v(%v)", c.out, []byte(c.out), buf.String(), buf.Bytes())
			}
		})
	}
}

func TestDefaultErrLogger(t *testing.T) {
	cases := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "helloworld",
			in:   "helloworld",
			out:  "helloworld\n",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var (
				buf = new(bytes.Buffer)
			)

			_, e, _ := Default.Logger(true, func(_, e, _ *log.Logger) {
				e.SetOutput(buf)
			})

			e.Print(c.in)

			if buf.String() != c.out {
				t.Errorf("\nwant: %v(%v)\ngot: %v(%v)", c.out, []byte(c.out), buf.String(), buf.Bytes())
			}
		})
	}
}

func TestDefaultDbgLogger(t *testing.T) {
	cases := []struct {
		name  string
		in    string
		out   string
		debug bool
	}{
		{
			name:  "debug(true)",
			in:    "helloworld",
			out:   "helloworld\n",
			debug: true,
		},
		{
			name:  "debug(false)",
			in:    "helloworld",
			out:   "",
			debug: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var (
				buf = new(bytes.Buffer)
			)

			_, _, d := Default.Logger(c.debug, func(_, _, d *log.Logger) {
				d.SetOutput(buf)
			})

			d.Print(c.in)

			if buf.String() != c.out {
				t.Errorf("\nwant: %v(%v)\ngot: %v(%v)", c.out, []byte(c.out), buf.String(), buf.Bytes())
			}
		})
	}
}
