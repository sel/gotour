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
	z := byte('Z')
	if 'a' <= c && c <= 'z' {
		z = 'z'
	} else if c < 'A' || c > 'Z' {
		return c
	}
	c += 13
	if c > z {
		c -= 26
	}
	return c
}

func (t rot13Reader) Read(buf []byte) (int, error) {
	n, err := t.r.Read(buf)
	if err != nil {
		return n, err
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
