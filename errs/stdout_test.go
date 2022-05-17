package errs

import (
	"bytes"
	"errors"
	"os"
	"testing"
)

func TestStdout(t *testing.T) {
	defer func() { stdOutWriter = os.Stdout }()
	buf := new(bytes.Buffer)
	stdOutWriter = buf

	fn := func() (string, error) {
		return "", errors.New("test")
	}

	Stdout(fn())

	if buf.String() != "test\n" {
		t.Errorf("[want] test\t[got] %s", buf.String())
	}
}
