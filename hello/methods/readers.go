package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot.r.Read(p)
	for i, c := range p {
		switch {
		case c >= 'A' && c <= 'M' || c >= 'a' && c <= 'm':
			p[i] += 13
		case c >= 'N' && c <= 'Z' || c >= 'n' && c <= 'z':
			p[i] -= 13
		}
	}
	return
}

func main() {
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	rd := rot13Reader{s}
	io.Copy(os.Stdout, &rd)
}
