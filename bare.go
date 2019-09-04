package strlit

import (
	"github.com/reiver/go-whitespace"

	"fmt"
	"unicode/utf8"
)

// Bare provides methods to decode, and encode Bare String Literals
type Bare struct {}

// Decode decodes a Bare String Literal.
func (receiver Bare) Decode(dst []byte, src []byte) (bytesWritten int, bytesRead int, err error) {

	if nil == dst {
		return 0, 0, errNilDestination
	}
	if nil == src {
		return 0, 0, errNilSource
	}
	if 0 == len(src) {
		return 0, 0, nil
	}

	var pSrc []byte = src

	var pDst []byte = dst

	Loop: for {
		r, size := utf8.DecodeRune(pSrc)
		if utf8.RuneError == r && 0 == size {
			break Loop
		}
		if utf8.RuneError == r {
			return bytesWritten, bytesRead, errUTF8RuneError
		}

		switch {
		case whitespace.IsWhitespace(r):
			break Loop
		}

		n := copy(pDst, pSrc[:size])

		bytesWritten += size
		pDst = pDst[size:]

		bytesRead += size
		pSrc = pSrc[size:]

		if expected, actual := size, n; expected != actual {
			return bytesWritten, bytesRead, fmt.Errorf("strlit: Internal Error: wrong number of bytes copied; expected=%d actual=%d", expected, actual)
		}
	}

	return
}
