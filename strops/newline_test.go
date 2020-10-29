package strops

import "testing"

func TestChop(t *testing.T) {
	cases := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "LF",
			in:   "helloworld\n",
			out:  "helloworld",
		},
		{
			name: "CRLF",
			in:   "helloworld\r\n",
			out:  "helloworld",
		},
		{
			name: "No newline",
			in:   "helloworld",
			out:  "helloworld",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r := Chop(c.in)
			if r != c.out {
				t.Errorf("want: %s\tgot: %s", c.out, r)
			}
		})
	}
}
