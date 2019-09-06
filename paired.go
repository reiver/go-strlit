package strlit

import (
	"github.com/reiver/go-buffers"

	"fmt"
	"io"
	"unicode/utf8"
)

// Paired provides methods to decode, and encode Paired String Literals.
type Paired struct {
	BeginSymbol rune
	EndSymbol   rune
}

// Decode decodes a Paired String Literal.
//
// ‘dst’ can be a []byte, or an io.Writer.
func (receiver Paired) Decode(dst interface{}, src []byte) (bytesWritten int, bytesRead int, err error) {

	beginSymbol := receiver.BeginSymbol
	endSymbol   := receiver.EndSymbol
	if beginSymbol == endSymbol {
		return 0, 0, fmt.Errorf("strlit: a paired literal cannot have the beginning symbol, and the ending symbol be the same: begin-symbol=%q begin-symbol=%q", beginSymbol, endSymbol)
	}

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
	var depth uint
	{
		var r rune
		var size int

		r, size = utf8.DecodeRune(pSrc)
		if utf8.RuneError == r && 0 == size {
			return 0, 0, nil
		}
		if utf8.RuneError == r {
			return 0, 0, errUTF8RuneError
		}

		if beginSymbol != r {
			return 0, 0, errNotLiteral(r)
		}

		bytesRead += size
		pSrc = pSrc[size:]

		depth++
	}

	Loop: for {
		r, size := utf8.DecodeRune(pSrc)
		if utf8.RuneError == r && 0 == size {
			return bytesWritten, bytesRead, fmt.Errorf("strlit: source ended before seeing end symbol %q.", endSymbol)
		}
		if utf8.RuneError == r {
			return bytesWritten, bytesRead, errUTF8RuneError
		}

		switch r {
		case beginSymbol:
			depth++
		case endSymbol:
			depth--
			if 0 == depth {
				bytesRead += size
				pSrc = pSrc[size:]

				break Loop
			}
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
