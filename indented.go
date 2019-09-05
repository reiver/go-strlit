package strlit

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

// Bare provides methods to decode, and encode Indented String Literals
type Indented struct {}

// Decode decodes a Indented String Literal.
func (receiver Indented) Decode(dst []byte, src []byte) (bytesWritten int, bytesRead int, err error) {

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

	var indentation []byte
	IndentationLoop: for {
		r, size := utf8.DecodeRune(pSrc)
		if utf8.RuneError == r && 0 == size {
			break IndentationLoop
		}
		if utf8.RuneError == r {
			return bytesWritten, bytesRead, errUTF8RuneError
		}

		switch r {
		case '\t',' ':

		default:
			break IndentationLoop
		}

		indentation = append(indentation, string(r)...)
		pSrc = pSrc[size:]
	}
	pSrc = src

	var pDst []byte = dst

	Loop: for {
		if !bytes.HasPrefix(pSrc, indentation) {
			break Loop
		}
		{
			size := len(indentation)

			bytesRead += size
			pSrc = pSrc[size:]
		}

		Inner: for {
			r, size := utf8.DecodeRune(pSrc)
			if utf8.RuneError == r && 0 == size {
				break Loop
			}
			if utf8.RuneError == r {
				return bytesWritten, bytesRead, errUTF8RuneError
			}

			n := copy(pDst, pSrc[:size])

			bytesWritten += size
			pDst = pDst[size:]

			bytesRead += size
			pSrc = pSrc[size:]

			if expected, actual := size, n; expected != actual {
				return bytesWritten, bytesRead, fmt.Errorf("strlit: Internal Error: wrong number of bytes copied; expected=%d actual=%d", expected, actual)
			}

			switch r {
			case '\u2028':
				break Inner
			}
		}
	}

	return
}
