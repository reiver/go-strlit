package strlit

import (
	"github.com/reiver/go-buffers"
	"github.com/reiver/go-oi"
	"github.com/reiver/go-utf8s"

	"bytes"
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
//
// ‘src’ can be a []byte, or an io.ReaderAt, or an io.ReadSeeker.
func (receiver Paired) Decode(dst interface{}, src interface{}) (bytesWritten int, bytesRead int, err error) {

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

	var readSeeker io.ReadSeeker
	{
		switch casted := src.(type) {
		case io.ReadSeeker:
			readSeeker = casted
		case io.ReaderAt:
			readSeeker = oi.ReadSeeker(casted)
		case []byte:
			readSeeker = bytes.NewReader(casted)
		default:
			return 0, 0, fmt.Errorf("strlit: Unsupport Source Type: %T", src)
		}
	}

	var depth uint
	{
		var r rune
		var size int
		var err error

		r, size, err = utf8s.ReadRune(readSeeker)
		if nil != err && io.EOF == err {
			if 0 < size {
				readSeeker.Seek(int64(-size), io.SeekCurrent)
			}
			return bytesWritten, bytesRead, errSyntaxError("", "End Of File (io.EOF) received before getting to end of paired string literal")
		}
		if nil != err {
			if 0 < size {
				readSeeker.Seek(int64(-size), io.SeekCurrent)
			}
			return bytesWritten, bytesRead, err
		}
		if utf8.RuneError == r && 0 == size {
			if 0 < size {
				readSeeker.Seek(int64(-size), io.SeekCurrent)
			}
			return 0, 0, nil
		}
		if utf8.RuneError == r {
			if 0 < size {
				readSeeker.Seek(int64(-size), io.SeekCurrent)
			}
			return 0, 0, errUTF8RuneError
		}

		if beginSymbol != r {
			if 0 < size {
				readSeeker.Seek(int64(-size), io.SeekCurrent)
			}
			return 0, 0, errNotLiteral(r)
		}

		depth++
	}

	Loop: for {
		r, size, err := utf8s.ReadRune(readSeeker)
		bytesRead += size
		if nil != err && io.EOF == err {
			break Loop
		}
		if nil != err {
			return bytesWritten, bytesRead, err
		}
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

				break Loop
			}
		}


		n, err := utf8s.WriteRune(writer, r)

		bytesWritten += size
		if nil != err {
			return bytesWritten, bytesRead, err
		}

		if expected, actual := size, n; expected != actual {
			return bytesWritten, bytesRead, fmt.Errorf("strlit: Internal Error: wrong number of bytes copied; expected=%d actual=%d", expected, actual)
		}
	}

	return
}
