package strlit

import (
	"github.com/reiver/go-buffers"
	"github.com/reiver/go-oi"
	"github.com/reiver/go-utf8s"
	"github.com/reiver/go-whitespace"

	"bytes"
	"fmt"
	"io"
	"unicode/utf8"
)

// Bare provides methods to decode, and encode Bare String Literals
type Bare struct {}

// Decode decodes a Bare String Literal.
//
// ‘dst’ can be a []byte, or an io.Writer.
//
// ‘src’ can be a []byte, or an io.ReaderAt, or an io.ReadSeeker.
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

	return receiver.decode(writer, readSeeker)
}

func (receiver Bare) decode(writer io.Writer, readSeeker io.ReadSeeker) (bytesWritten int, bytesRead int, err error) {

	if nil == writer {
		return 0, 0, errNilDestination
	}

	if nil == readSeeker {
		return 0, 0, errNilSource
	}

	Loop: for {
		r, size, err := utf8s.ReadRune(readSeeker)
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
			_, err := readSeeker.Seek(int64(-size), io.SeekCurrent)
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
