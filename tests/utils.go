package tests

import (
	"bytes"
	"io"
	"io/ioutil"
)

func ByteToReadCloser(data []byte) io.ReadCloser {
	reader := bytes.NewReader(data)
	return ioutil.NopCloser(reader)
}
