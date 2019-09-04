package strlit_test

import (
	"github.com/reiver/go-strlit"

	"math/rand"
	"reflect"
	"strings"
	"time"

	"testing"
)

func TestBareDecode(t *testing.T) {

	tests := []struct {
		Source             []byte
		ExpectedBytesRead    int
		ExpectedBytesWritten int
		Expected           []byte
	}{}

	var breaks []rune = []rune{
		'\u0009', // horizontal tab
		'\u000A', // line feed
		'\u000B', // vertical tab
		'\u000C', // form feed
		'\u000D', // carriage return
		'\u0020', // space
		'\u0085', // next line
		'\u00A0', // no-break space
		'\u1680', // ogham space mark
		'\u180E', // mongolian vowel separator
		'\u2000', // en quad
		'\u2001', // em quad
		'\u2002', // en space
		'\u2003', // em space
		'\u2004', // three-per-em space
		'\u2005', // four-per-em space
		'\u2006', // six-per-em space
		'\u2007', // figure space
		'\u2008', // punctuation space
		'\u2009', // thin space
		'\u200A', // hair space
		'\u2028', // line separator
		'\u2029', // paragraph separator
		'\u202F', // narrow no-break space
		'\u205F', // medium mathematical space
		'\u3000', // ideographic space
	}

	{
		content := []string {
			"apple",
			"BANANA",
			"Cherry",
			"dATE",

			"Hello_world!",
			"Hello-world!", // hyphen-minus
			"Hello‐world!", // hyphen
		}


		for _, s := range content {

			test := struct{
				Source             []byte
				ExpectedBytesRead    int
				ExpectedBytesWritten int
				Expected           []byte
			}{}

			test.Source = []byte(s)
			test.ExpectedBytesRead = len(test.Source)
			test.ExpectedBytesWritten = len(s)
			test.Expected = []byte(s)

			tests = append(tests, test)


			for _, b := range breaks {
				anotherTest := test
				copy(anotherTest.Source, test.Source)
				copy(anotherTest.Expected, test.Expected)
				anotherTest.Source = append(anotherTest.Source, string(b)...)
				anotherTest.Source = append(anotherTest.Source, `"wow"`...)

				tests = append(tests, anotherTest)
			}


			for _, b := range breaks {
				anotherTest := test
				copy(anotherTest.Source, test.Source)
				copy(anotherTest.Expected, test.Expected)
				anotherTest.Source = append(anotherTest.Source, string(b)...)
				anotherTest.Source = append(anotherTest.Source, `wow`...)

				tests = append(tests, anotherTest)
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

		innerLength := 1 + randomness.Intn(30)
		for ii:=0; ii<innerLength; ii++ {

			r := variousRunes[randomness.Int() % len(variousRunes)]

			builder.WriteRune(r)
		}


		test := struct{
			Source             []byte
			ExpectedBytesRead    int
			ExpectedBytesWritten int
			Expected           []byte
		}{}

		test.Source = []byte(nil)
		test.Source = append(test.Source, builder.String()...)
		test.ExpectedBytesRead = len(test.Source)
		test.ExpectedBytesWritten = len(test.Source)

		test.Expected = []byte(test.Source)

		tests = append(tests, test)


		for _, b := range breaks {
			anotherTest := test
			copy(anotherTest.Source, test.Source)
			copy(anotherTest.Expected, test.Expected)
			anotherTest.Source = append(anotherTest.Source, string(b)...)
			anotherTest.Source = append(anotherTest.Source, `"wow"`...)

			tests = append(tests, anotherTest)
		}


		for _, b := range breaks {
			anotherTest := test
			copy(anotherTest.Source, test.Source)
			copy(anotherTest.Expected, test.Expected)
			anotherTest.Source = append(anotherTest.Source, string(b)...)
			anotherTest.Source = append(anotherTest.Source, `wow`...)

			tests = append(tests, anotherTest)
		}
	}

	for testNumber, test := range tests {

		var bare strlit.Bare

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

			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
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

			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
