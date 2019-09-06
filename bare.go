package strlit

import (
	"github.com/reiver/go-buffers"
	"github.com/reiver/go-whitespace"

	"fmt"
	"io"
	"unicode/utf8"
)

// Bare provides methods to decode, and encode Bare String Literals
type Bare struct {}

// Decode decodes a Bare String Literal.
func (receiver Bare) Decode(dst interface{}, src []byte) (bytesWritten int, bytesRead int, err error) {

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
	}

	return
}
