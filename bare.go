package strlit

import (
	"github.com/reiver/go-buffers"
	"github.com/reiver/go-oi"
	"github.com/reiver/go-utf8s"
	"github.com/reiver/go-whitespace"

	"bytes"
	"fmt"
	"io"
	"string"
	"unicode/utf8"
)

// Bare provides methods to decode, and encode Bare String Literals
type Bare struct {}

// Decode decodes a Bare String Literal.
//
// ‘dst’ can be a []byte, or an io.Writer.
//
// ‘src’ can be a string, or a []byte, or an io.ReaderAt, or an io.ReadSeeker, or an io.RuneScanner.
func (receiver Bare) Decode(dst interface{}, src interface{}) (bytesWritten int, bytesRead int, err error) {

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
			return 0, 0, fmt.Errorf("strlit: Unsupported Destination Type: %T", dst)
		}
	}

	if nil == src {
		return 0, 0, errNilSource
	}

	var runeScanner io.RuneScanner
	{
		switch casted := src.(type) {
		case io.RuneScanner:
			runeScanner = casted
		case io.ReadSeeker:
			runeScanner = utf8s.NewRuneScanner(casted)
		case io.ReaderAt:
			runeScanner = utf8s.NewRuneScanner(oi.ReadSeeker(casted))
		case []byte:
			runeScanner = bytes.NewReader(casted)
		case string:
			runeScanner = strings.NewReader(casted)
		default:
			return 0, 0, fmt.Errorf("strlit: Unsupported Source Type: %T", src)
		}
	}

	return receiver.decode(writer, runeScanner)
}

func (receiver Bare) decode(writer io.Writer, runeScanner io.RuneScanner) (bytesWritten int, bytesRead int, err error) {

	if nil == writer {
		return 0, 0, errNilDestination
	}

	if nil == runeScanner {
		return 0, 0, errNilSource
	}

	Loop: for {
		r, size, err := runeScanner.ReadRune()
		bytesRead += size
		if nil != err && io.EOF == err {
			if 0 == bytesRead {
				return bytesWritten, bytesRead, errEmptySource
			}
			break Loop
		}
		if nil != err {
			return bytesWritten, bytesRead, err
		}
		if utf8.RuneError == r && 0 == size {
			break Loop
		}
		if utf8.RuneError == r {
			return bytesWritten, bytesRead, errUTF8RuneError
		}

		switch {
		case whitespace.IsWhitespace(r):
			err := runeScanner.UnreadRune()
			if nil != err {
				return bytesWritten, bytesRead, err
			}
			bytesRead -= size

			if 0 == bytesWritten {
				return bytesWritten, bytesRead, errNotBareLiteral(r)
			}

			break Loop
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
