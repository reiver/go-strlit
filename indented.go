package strlit

import (
	"github.com/reiver/go-buffers"

	"bytes"
	"fmt"
	"io"
	"unicode/utf8"
)

// Bare provides methods to decode, and encode Indented String Literals
type Indented struct {}

// Decode decodes a Indented String Literal.
//
// ‘dst’ can be a []byte, or an io.Writer.
func (receiver Indented) Decode(dst interface{}, src []byte) (bytesWritten int, bytesRead int, err error) {

	if nil == dst {
		return 0, 0, errNilDestination
	}

	var writer io.Writer
	{
		switch casted := dst.(type) {
		case io.Writer:
			writer = casted
		case []byte:
			writer = buffers.NewWriter(casted)
		default:
			return 0, 0, fmt.Errorf("strlit: Unsupport Destination Type: %T", dst)
		}
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

			n, err := writer.Write(pSrc[:size])

			bytesWritten += size
			if nil != err {
				return bytesWritten, bytesRead, err
			}

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
