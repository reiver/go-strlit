package strlit

import (
	"github.com/reiver/go-buffers"
	"github.com/reiver/go-indent"
	"github.com/reiver/go-oi"
	"github.com/reiver/go-utf8s"

	"bytes"
	"fmt"
	"io"
	"strings"
	"unicode/utf8"
)

// Bare provides methods to decode, and encode Indented String Literals
type Indented struct {}

// Decode decodes a Indented String Literal.
//
// ‘dst’ can be a []byte, or an io.Writer.
//
// ‘src’ can be a []byte, or an io.ReaderAt, or an io.ReadSeeker.
func (receiver Indented) Decode(dst interface{}, src interface{}) (bytesWritten int, bytesRead int, err error) {

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

func (receiver Indented) decode(writer io.Writer, readSeeker io.ReadSeeker) (bytesWritten int, bytesRead int, err error) {

	if nil == writer {
		return 0, 0, errNilDestination
	}

	if nil == readSeeker {
		return 0, 0, errNilSource
	}

	var indentation []byte
	{
		var dst strings.Builder

		if err := indent.Detect(&dst, readSeeker); nil != err {
			return 0, 0, err
		}

		indentation = append([]byte(nil), dst.String()...)
	}

	var indentationBuffer []byte = make([]byte, len(indentation))

	Loop: for {
		{
			n, err := readSeeker.Read(indentationBuffer)
			if nil != err && io.EOF == err {
				break Loop
			}
			if nil != err {
				return bytesWritten, bytesRead, err
			}
			if expected, actual := len(indentationBuffer), n; expected != actual {
				if expected < actual {
					return bytesWritten, bytesRead, fmt.Errorf("strlit: read from io.SeekReader was expected to be %d bytes but was actually %d bytes", expected, actual)
				}

				m, err := readSeeker.Read(indentationBuffer[n:])
				if nil != err && io.EOF == err {
					break Loop
				}

				if expected, actual := len(indentationBuffer), n+m; expected != actual {
					return bytesWritten, bytesRead, fmt.Errorf("strlit: read from io.SeekReader was expected to be %d bytes but was actually %d bytes", expected, actual)
				}
			}

			if expected, actual := indentation, indentationBuffer; !bytes.Equal(expected, actual) {
				if _, err = readSeeker.Seek(int64(-1 * n), io.SeekCurrent); nil != err {
					return bytesWritten, bytesRead, err
				}
				break Loop
			}
			bytesRead += n
		}

		Inner: for {
			r, size, err := utf8s.ReadRune(readSeeker)
			bytesRead += size
			if nil != err {
				return bytesWritten, bytesRead, err
			}
			if utf8.RuneError == r && 0 == size {
				break Loop
			}
			if utf8.RuneError == r {
				return bytesWritten, bytesRead, errUTF8RuneError
			}

			written, err := utf8s.WriteRune(writer, r)

			bytesWritten += written
			if nil != err {
				return bytesWritten, bytesRead, err
			}

			if expected, actual := size, written; expected != actual {
				return bytesWritten, bytesRead, fmt.Errorf("strlit: Internal Error: wrong number of bytes copied; expected=%d actual=%d", expected, actual)
			}

			switch r {
			case '\n',
			     '\u0085', // Next Line
			     '\u2028': // Line Separator
				break Inner
			}
		}
	}

	return
}
