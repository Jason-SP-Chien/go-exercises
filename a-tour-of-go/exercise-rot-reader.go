package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 *rot13Reader) Read(b []byte) (n int, e error) {
	n, err := r13.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13mapping(b[i])
	}
	return n, err
}

func rot13mapping(b byte) byte {
	if b >= 'a' && b <= 'z' {
		if b >= 'm' {
			return b - 13
		} else {
			return b + 13
		}
	} else if b >= 'A' && b <= 'Z' {
		if b >= 'M' {
			return b - 13
		} else {
			return b + 13
		}
	}

	return b
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
