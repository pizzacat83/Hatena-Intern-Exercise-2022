package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(c byte) byte {
	switch {
	case 'a' <= c && c <= 'z':
		return (c-'a'+13)%26 + 'a'
	case 'A' <= c && c <= 'Z':
		return (c-'A'+13)%26 + 'A'
	default:
		return c
	}
}

func (r rot13Reader) Read(buf []byte) (int, error) {
	n, err := r.r.Read(buf)
	if err != nil {
		return 0, err
	}
	for i := 0; i < n; i++ {
		buf[i] = rot13(buf[i])
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
