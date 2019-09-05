package strlit_test

import (
	"github.com/reiver/go-strlit"

	"math/rand"
	"reflect"
	"strings"
	"time"

	"testing"
)

func TestIndentedDecode(t *testing.T) {

	const LS = '\u2028' // line separator

	tests := []struct {
		Source             []byte
		ExpectedBytesRead    int
		ExpectedBytesWritten int
		Expected           []byte
	}{}

	{
		test := struct{
			Source             []byte
			ExpectedBytesRead    int
			ExpectedBytesWritten int
			Expected           []byte
		}{}

		test.Expected = []byte(
			      "one two three"        +string(LS)+
			      "four five"            +string(LS)+
			      "six seven eight nine" +string(LS))
		test.ExpectedBytesWritten = len(test.Expected)

		test.Source = []byte(
			"\t"+ "one two three"        +string(LS)+
			"\t"+ "four five"            +string(LS)+
			"\t"+ "six seven eight nine" +string(LS))
		test.ExpectedBytesRead = len(test.Source)


		tests = append(tests, test)
	}

	{
		test := struct{
			Source             []byte
			ExpectedBytesRead    int
			ExpectedBytesWritten int
			Expected           []byte
		}{}

		test.Expected = []byte(
			      "one two three"        +string(LS)+
			      "four five"            +string(LS)+
			      "six seven eight nine" +string(LS))
		test.ExpectedBytesWritten = len(test.Expected)

		test.Source = []byte(
			"\t"+ "one two three"        +string(LS)+
			"\t"+ "four five"            +string(LS)+
			"\t"+ "six seven eight nine" +string(LS)+
			"catch"                      +string(LS))
		test.ExpectedBytesRead = len(test.Source) - len("catch"+string(LS))


		tests = append(tests, test)
	}


	indentations := []string {
		"\t",
		"\t\t",
		"\t\t\t",
		"\t\t\t\t",
		"\t\t\t\t\t",

		" ",
		"  ",
		"   ",
		"    ",
		"     ",

		" \t",

		"\t ",

		"\t\t  ",

		"  \t\t",

		"  \t   \t\t  ",
	}

	{
		content := []string {
			"apple",

			"apple"  +string(LS)+
			"BANANA",

			"apple"  +string(LS)+
			"BANANA" +string(LS)+
			"Cherry",

			"apple"  +string(LS)+
			"BANANA" +string(LS)+
			"Cherry" +string(LS)+
			"dATE",


			"Hello world!",

			"Hello world!" +string(LS)+
			"apple BANANA Cherry dATE",
		}

		for _, s := range content {

			for _, indentation := range indentations {
				test := struct{
					Source             []byte
					ExpectedBytesRead    int
					ExpectedBytesWritten int
					Expected           []byte
				}{}

				{
					ss := strings.ReplaceAll(s, string(LS), string(LS)+indentation)
					ss = indentation + ss + string(LS)

					test.Source = []byte(ss)

					test.ExpectedBytesRead = len(test.Source)
					test.ExpectedBytesWritten = len(s) + len(string(LS))
					test.Expected = []byte(nil)
					test.Expected = append(test.Expected, s...)
					test.Expected = append(test.Expected, string(LS)...)

					tests = append(tests, test)
				}

				{
					ss := strings.ReplaceAll(s, string(LS), string(LS)+indentation)
					ss = indentation + ss + string(LS) + "wow"

					test.Source = []byte(ss)

					test.ExpectedBytesRead = len(test.Source) - len("wow")
					test.ExpectedBytesWritten = len(s) + len(string(LS))
					test.Expected = []byte(nil)
					test.Expected = append(test.Expected, s...)
					test.Expected = append(test.Expected, string(LS)...)

					tests = append(tests, test)
				}
			}
		}
	}

	var variousRunes []rune = []rune{
		'0',
		'1',
		'2',
		'3',
		'4',
		'5',
		'6',
		'7',
		'8',
		'9',

		'۰',
		'۱',
		'۲',
		'۳',
		'۴',
		'۵',
		'۶',
		'۷',
		'۸',
		'۹',

		'A',
		'B',
		'C',
		'D',
		'E',
		'F',
		'G',
		'H',
		'I',
		'J',
		'K',
		'L',
		'M',
		'N',
		'O',
		'P',
		'Q',
		'R',
		'S',
		'T',
		'U',
		'V',
		'W',
		'X',
		'Y',
		'Z',

		'a',
		'b',
		'c',
		'd',
		'e',
		'f',
		'g',
		'h',
		'i',
		'j',
		'k',
		'l',
		'm',
		'n',
		'o',
		'p',
		'q',
		'r',
		's',
		't',
		'u',
		'v',
		'w',
		'x',
		'y',
		'z',

		'ا',
		'ب',
		'پ',
		'ت',
		'ث',
		'ج',
		'چ',
		'ح',
		'خ',
		'د',
		'ذ',
		'ر',
		'ز',
		'ژ',
		'س',
		'ش',
		'ص',
		'ض',
		'ط',
		'ظ',
		'ع',
		'غ',
		'ف',
		'ق',
		'ک',
		'گ',
		'ل',
		'م',
		'ن',
		'و',
		'ه',
		'ی',

		'ㄱ',
		'ㄲ',
		'ㄴ',
		'ㄷ',
		'ㄸ',
		'ㄹ',
		'ㅁ',
		'ㅂ',
		'ㅃ',
		'ㅅ',
		'ㅆ',
		'ㅇ',
		'ㅈ',
		'ㅉ',
		'ㅊ',
		'ㅋ',
		'ㅌ',
		'ㅍ',
		'ㅎ',
		'ㅏ',
		'ㅐ',
		'ㅑ',
		'ㅒ',
		'ㅓ',
		'ㅔ',
		'ㅕ',
		'ㅖ',
		'ㅗ',
		'ㅘ',
		'ㅙ',
		'ㅚ',
		'ㅛ',
		'ㅜ',
		'ㅝ',
		'ㅞ',
		'ㅟ',
		'ㅠ',
		'ㅡ',
		'ㅢ',
		'ㅣ',

		'!',
		'#',
		'$',
		'%',
		'&',
		'*',
		'+',
		',',
		'-',
		'.',
		'/',

		':',
		';',
		'<',
		'=',
		'>',
		'?',
		'@',

		'^',
		'_',
		'`',

		'|',
		'~',

		'😀',
		'😁',
		'😂',
		'🤣',
		'😃',
		'😄',
		'😅',
		'😆',
		'😉',
		'😊',
		'😋',
		'😎',
		'😍',
		'😘',

		'😈',
		'👿',
		'👹',
		'👺',
		'💀',
		'👻',
		'👽',
		'🤖',

		'👾',

		'👍',
		'👎',
		'👊',
		'✊',
		'🤛',
		'🤜',
		'🤞',
		'✌',
		'🤟',
		'🤘',
		'👌',
		'👈',
		'👉',
		'👆',
		'👇',
		'☝',
		'✋',
		'🤚',
		'🖐',
		'🖖',
		'👋',
		'🤙',
		'🖕',
		'✍',
	}

	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

	for i:=0; i<50; i++ {

		var builder strings.Builder
		{
			var numLines int = 1 + randomness.Intn(20)
			for lineNumber:=0; lineNumber<numLines; lineNumber++ {

				if 0 < lineNumber {
					builder.WriteRune('\u2028')
				}

				innerLength := 1 + randomness.Intn(30)
				for ii:=0; ii<innerLength; ii++ {

					r := variousRunes[randomness.Int() % len(variousRunes)]

					builder.WriteRune(r)
				}
			}
		}

		for _, indentation := range indentations {

			test := struct{
				Source             []byte
				ExpectedBytesRead    int
				ExpectedBytesWritten int
				Expected           []byte
			}{}

			s := builder.String()

			{
				ss := strings.ReplaceAll(s, string(LS), string(LS)+indentation)
				ss = indentation + ss + string(LS)

				test.Source = []byte(ss)

				test.ExpectedBytesRead = len(test.Source)
				test.ExpectedBytesWritten = len(s) + len(string(LS))

				test.Expected = []byte(nil)
				test.Expected = append(test.Expected, s...)
				test.Expected = append(test.Expected, string(LS)...)

				tests = append(tests, test)
			}

			{
				ss := strings.ReplaceAll(s, string(LS), string(LS)+indentation)
				ss = indentation + ss + string(LS) + "wow"

				test.Source = []byte(ss)

				test.ExpectedBytesRead = len(test.Source) - len("wow")
				test.ExpectedBytesWritten = len(s) + len(string(LS))

				test.Expected = []byte(nil)
				test.Expected = append(test.Expected, s...)
				test.Expected = append(test.Expected, string(LS)...)

				tests = append(tests, test)
			}
		}
	}

	for testNumber, test := range tests {

		var bare strlit.Indented

		var src []byte = test.Source
		lenSrc := len(src)

		var dst []byte = make([]byte, lenSrc)

		byteWritten, bytesRead, err := bare.Decode(dst, src)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one.", testNumber)

			t.Logf("SRC (string): [%s] {%d}", src, len(src))
			t.Logf("SRC (string): (%q) {%d}", src, len(src))
			t.Logf("DST (string): [%s] {%d}", dst[:byteWritten], len(dst[:byteWritten]))
			t.Logf("DST (string): (%q) {%d}", dst[:byteWritten], len(dst[:byteWritten]))
			t.Logf("DST (string): [%s] (full) {%d}", dst, len(dst))
			t.Logf("DST (string): (%q) (full) {%d}", dst, len(dst))

			t.Logf("ERROR TYPE: %T", err)
			t.Logf("ERROR: %q", err)
			continue
		}


		if expected, actual := test.ExpectedBytesRead, bytesRead; expected != actual {
			t.Errorf("For test #%d, actual number of bytes read is not what was expected.", testNumber)

			t.Logf("SRC (string): [%s] {%d}", src, len(src))
			t.Logf("SRC (string): (%q) {%d}", src, len(src))
			t.Logf("DST (string): [%s] {%d}", dst[:byteWritten], len(dst[:byteWritten]))
			t.Logf("DST (string): (%q) {%d}", dst[:byteWritten], len(dst[:byteWritten]))
			t.Logf("DST (string): [%s] (full) {%d}", dst, len(dst))
			t.Logf("DST (string): (%q) (full) {%d}", dst, len(dst))


			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			continue
		}

		if expected, actual := test.ExpectedBytesWritten, byteWritten; expected != actual {
			t.Errorf("For test #%d, actual number of bytes written is not what was expected.", testNumber)

			t.Logf("SRC (string): [%s] {%d}", src, len(src))
			t.Logf("SRC (string): (%q) {%d}", src, len(src))
			t.Logf("DST (string): [%s] {%d}", dst[:byteWritten], len(dst[:byteWritten]))
			t.Logf("DST (string): (%q) {%d}", dst[:byteWritten], len(dst[:byteWritten]))
			t.Logf("DST (string): [%s] (full) {%d}", dst, len(dst))
			t.Logf("DST (string): (%q) (full) {%d}", dst, len(dst))

			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}

		if expected, actual := test.Expected, dst[:byteWritten]; !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, what was actually written is not what was expected.", testNumber)

			t.Logf("SRC (string): [%s] {%d}", src, len(src))
			t.Logf("SRC (string): (%q) {%d}", src, len(src))
			t.Logf("DST (string): [%s] {%d}", dst[:byteWritten], len(dst[:byteWritten]))
			t.Logf("DST (string): (%q) {%d}", dst[:byteWritten], len(dst[:byteWritten]))
			t.Logf("DST (string): [%s] (full) {%d}", dst, len(dst))
			t.Logf("DST (string): (%q) (full) {%d}", dst, len(dst))

			t.Logf("EXPECTED: %q {%d}", expected, len(expected))
			t.Logf("ACTUAL:   %q {%d}", actual, len(actual))
			continue
		}
	}
}
