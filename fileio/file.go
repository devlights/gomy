package fileio

import (
	"bytes"
	"io"
	"os"

	"github.com/devlights/gomy/fileio/jp"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// OpenRead は、読み込みモードでファイルをオープンします。
func OpenRead(name string, encoding jp.Encoding) (io.Reader, func() error, error) {
	var (
		flag = os.O_RDONLY
	)

	return openReadMode(name, flag, encoding)
}

// OpenWrite は、書き込みモードでファイルをオープンします。
func OpenWrite(name string, encoding jp.Encoding) (io.Writer, func() error, error) {
	var (
		flag = os.O_RDWR | os.O_CREATE | os.O_TRUNC
	)

	return openWriteMode(name, flag, encoding)
}

// OpenAppend は、追記モードでファイルをオープンします。
func OpenAppend(name string, encoding jp.Encoding) (io.Writer, func() error, error) {
	var (
		flag = os.O_APPEND | os.O_CREATE | os.O_WRONLY
	)

	return openWriteMode(name, flag, encoding)
}

// BufRead は、指定されたエンコーディングで読み出しを行う io.Reader を返します。
func BufRead(buf *bytes.Buffer, enc jp.Encoding) (io.Reader, error) {
	decoder := getDecoder(enc)
	if decoder == nil {
		return buf, nil
	}

	return transform.NewReader(buf, decoder), nil
}

// BufWrite は、指定されたエンコーディングで書き出しを行う io.Writer を返します。
func BufWrite(buf *bytes.Buffer, enc jp.Encoding) (io.Writer, error) {
	encoder := getEncoder(enc)
	if encoder == nil {
		return buf, nil
	}

	return transform.NewWriter(buf, encoder), nil
}

func getDecoder(enc jp.Encoding) *encoding.Decoder {
	if enc == jp.ShiftJis {
		return japanese.ShiftJIS.NewDecoder()
	}

	if enc == jp.EucJp {
		return japanese.EUCJP.NewDecoder()
	}

	return nil
}

func getEncoder(enc jp.Encoding) *encoding.Encoder {
	if enc == jp.ShiftJis {
		return japanese.ShiftJIS.NewEncoder()
	}

	if enc == jp.EucJp {
		return japanese.EUCJP.NewEncoder()
	}

	return nil
}

func openReadMode(name string, flag int, encoding jp.Encoding) (io.Reader, func() error, error) {
	var (
		decoder = getDecoder(encoding)
		reader  io.Reader
	)

	fp, err := os.OpenFile(name, flag, 0666)
	if err != nil {
		return nil, nil, err
	}

	releaseFn := func() error {
		return fp.Close()
	}

	if decoder == nil {
		reader = fp
	} else {
		reader = transform.NewReader(fp, decoder)
	}

	return reader, releaseFn, nil
}

func openWriteMode(name string, flag int, encoding jp.Encoding) (io.Writer, func() error, error) {
	var (
		encoder = getEncoder(encoding)
		writer  io.Writer
	)

	fp, err := os.OpenFile(name, flag, 0666)
	if err != nil {
		return nil, nil, err
	}

	releaseFn := func() error {
		_ = fp.Sync()
		return fp.Close()
	}

	if encoder == nil {
		writer = fp
	} else {
		writer = transform.NewWriter(fp, encoder)
	}

	return writer, releaseFn, nil
}
