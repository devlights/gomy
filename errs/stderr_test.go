package errs

import (
	"bytes"
	"errors"
	"os"
	"testing"
)

func TestStderr(t *testing.T) {
	defer func() { stdErrWriter = os.Stderr }()
	buf := new(bytes.Buffer)
	stdErrWriter = buf

	fn := func() (string, error) {
		return "", errors.New("test")
	}

	Stderr(fn())

	if buf.String() != "test\n" {
		t.Errorf("[want] test\t[got] %s", buf.String())
	}
}
