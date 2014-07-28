package main

import (
	"io"
	"os"
	"strings"
)

var rot13map = map[byte]byte{
	'A': 'N',
	'B': 'O',
	'C': 'P',
	'D': 'Q',
	'E': 'R',
	'F': 'S',
	'G': 'T',
	'H': 'U',
	'I': 'V',
	'J': 'W',
	'K': 'X',
	'L': 'Y',
	'M': 'Z',
	'N': 'a',
	'O': 'b',
	'P': 'c',
	'Q': 'd',
	'R': 'e',
	'S': 'f',
	'T': 'g',
	'U': 'h',
	'V': 'i',
	'W': 'j',
	'X': 'k',
	'Y': 'l',
	'Z': 'm',
	'a': 'n',
	'b': 'o',
	'c': 'p',
	'd': 'q',
	'e': 'r',
	'f': 's',
	'g': 't',
	'h': 'u',
	'i': 'v',
	'j': 'w',
	'k': 'x',
	'l': 'y',
	'm': 'z',
	'n': 'a',
	'o': 'b',
	'p': 'c',
	'q': 'd',
	'r': 'e',
	's': 'f',
	't': 'g',
	'u': 'h',
	'v': 'i',
	'w': 'j',
	'x': 'k',
	'y': 'l',
	'z': 'm',
}

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	for i := 0; i < n; i++ {
		rot, isRot := rot13map[p[i]]
		if isRot {
			p[i] = rot
		}
	}
	return
}

func main() {
	s := strings.NewReader(
		"Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
