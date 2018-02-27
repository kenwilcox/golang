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

//func (rot *rot13Reader) Read(p []byte) (n int, err error) {
//	n, err = rot.r.Read(p)
//	for i, c := range p {
//		switch {
//		case c >= 'A' && c <= 'M' || c >= 'a' && c <= 'm':
//			p[i] += 13
//		case c >= 'N' && c <= 'Z' || c >= 'n' && c <= 'z':
//			p[i] -= 13
//		}
//	}
//	return
//}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot.r.Read(p)
	if err != nil {
		return 0, err
	}
	var normal = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var rot13 = "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm"
	for i, c := range p {
		// get the index from normal
		var idx = strings.IndexByte(normal, c)
		if idx != -1 {
			// if there is a match then
			// read the index from rot13
			p[i] = rot13[idx]
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
