package strlit_test

import (
	"github.com/reiver/go-strlit"

	"github.com/reiver/go-tmpl"

	"math/rand"
	"reflect"
	"strings"
	"time"

	"testing"
)

func TestPairedDecode(t *testing.T) {

	tests := []struct {
		BeginSymbol          rune
		EndSymbol            rune
		Source             []byte
		ExpectedBytesRead    int
		ExpectedBytesWritten int
		Expected           []byte
	}{}

	pairs := []struct{
		BeginSymbol rune
		EndSymbol   rune
	}{
		{
			BeginSymbol: '‘',
			EndSymbol:   '’',
		},
		{
			BeginSymbol: '“',
			EndSymbol:   '”',
		},


		{
			BeginSymbol: '‹',
			EndSymbol:   '›',
		},
		{
			BeginSymbol: '«',
			EndSymbol:   '»',
		},


		{
			BeginSymbol: '「',
			EndSymbol:   '」',
		},
		{
			BeginSymbol: '『',
			EndSymbol:   '』',
		},


		{
			BeginSymbol: '[', // Left Square Bracket
			EndSymbol:   ']', // Right Square Bracket
		},
		{
			BeginSymbol: '⁅', // Left Square Bracket with Quill
			EndSymbol:   '⁆', // Right Square Bracket with Quill
		},
		{
			BeginSymbol: '⟦', // Mathematical Left White Square Bracket
			EndSymbol:   '⟧', // Mathematical Right White Square Bracket
		},


		{
			BeginSymbol: '(', // Left Parenthesis
			EndSymbol:   ')', // Right Parenthesis
		},
		{
			BeginSymbol: '❨', // Medium Left Parenthesis Ornament
			EndSymbol:   '❩', // Medium Right Parenthesis Ornament
		},
		{
			BeginSymbol: '❪', // Medium Flattened Left Parenthesis Ornament
			EndSymbol:   '❫', // Medium Flattened Right Parenthesis Ornament
		},
		{
			BeginSymbol: '⦅', // Left White Parenthesis
			EndSymbol:   '⦆', // Right White Parenthesis
		},
		{
			BeginSymbol: '⸨', // Left Double Parenthesis
			EndSymbol:   '⸩', // Right Double Parenthesis
		},
		{
			BeginSymbol: '﴾', // Ornate Left Parenthesis
			EndSymbol:   '﴿', // Ornate Right Parenthesis
		},
		{
			BeginSymbol: '﹙', // Small Left Parenthesis
			EndSymbol:   '﹚', // Small Right Parenthesis
		},
		{
			BeginSymbol: '（', // Fullwidth Left Parenthesis
			EndSymbol:   '）', // Fullwidth Right Parenthesis
		},
		{
			BeginSymbol: '｟', // Fullwidth Left White Parenthesis
			EndSymbol:   '｠', // Fullwidth Right White Parenthesis
		},


		{
			BeginSymbol: '{', // Left Curly Bracket
			EndSymbol:   '}', // Right Curly Bracket
		},
		{
			BeginSymbol: '❴', // Medium Left Curly Bracket Ornament
			EndSymbol:   '❵', // Medium Right Curly Bracket Ornament
		},
		{
			BeginSymbol: '⦃', // Left White Curly Bracket
			EndSymbol:   '⦄', // Right White Curly Bracket
		},
		{
			BeginSymbol: '﹛', // Small Left Curly Bracket
			EndSymbol:   '﹜', // Small Right Curly Bracket
		},


		{
			BeginSymbol: '〈', // Left-Pointing Angle Bracket
			EndSymbol:   '〉', // Right-Pointing Angle Bracket
		},
		{
			BeginSymbol: '⟨', // Mathematical Left Angle Bracket
			EndSymbol:   '⟩', // Mathematical Right Angle Bracket
		},


		{
			BeginSymbol: '〔', // Left Tortoise Shell Bracket
			EndSymbol:   '〕', // Right Tortoise Shell Bracket
		},
		{
			BeginSymbol: '〘', // Left White Tortoise Shell Bracket
			EndSymbol:   '〙', // Right White Tortoise Shell Bracket
		},
		{
			BeginSymbol: '﹝', // Small Left Tortoise Shell Bracket
			EndSymbol:   '﹞', // Small Right Tortoise Shell Bracket
		},









		{
			BeginSymbol: '༺', // Tibetan Mark Gug Rtags Gyon
			EndSymbol:   '༻', // Tibetan Mark Gug Rtags Gyas
		},

		{
			BeginSymbol: '༼', // Tibetan Mark Ang Khang Gyon
			EndSymbol:   '༽', // Tibetan Mark Ang Khang Gyas
		},
	}

	{
		for _, pair := range pairs {
			test := struct{
				BeginSymbol          rune
				EndSymbol            rune
				Source             []byte
				ExpectedBytesRead    int
				ExpectedBytesWritten int
				Expected           []byte
			}{}

			test.BeginSymbol = pair.BeginSymbol
			test.EndSymbol = pair.EndSymbol
			test.Source = append(append([]byte(nil), string(pair.BeginSymbol)...), string(pair.EndSymbol)...)
			test.ExpectedBytesRead = len(string(pair.BeginSymbol)) + len(string(pair.EndSymbol))
			test.ExpectedBytesWritten = 0
			test.Expected = []byte{}

			tests = append(tests, test)



			test2 := test
			copy(test2.Source, test.Source)
			copy(test2.Expected, test.Expected)
			test2.Source = append(test2.Source, ` "wow"`...)

			tests = append(tests, test2)
		}
	}

	{
		content := []string {
			"apple",
			"BANANA",
			"Cherry",
			"dATE",

			"Hello world!",
		}


		for _, s := range content {
			for _, pair := range pairs {

				test := struct{
					BeginSymbol          rune
					EndSymbol            rune
					Source             []byte
					ExpectedBytesRead    int
					ExpectedBytesWritten int
					Expected           []byte
				}{}

				test.BeginSymbol = pair.BeginSymbol
				test.EndSymbol   = pair.EndSymbol
				test.Source =
					append(
						append(
							append(
								[]byte(nil),
								string(pair.BeginSymbol)...
							),
							s...
						),
						string(pair.EndSymbol)...
					)
				test.ExpectedBytesRead = len(test.Source)
				test.ExpectedBytesWritten = len(s)
				test.Expected = []byte(s)

				tests = append(tests, test)



				test2 := test
				copy(test2.Source, test.Source)
				copy(test2.Expected, test.Expected)
				test2.Source = append(test2.Source, ` "wow"`...)

				tests = append(tests, test2)
			}
		}
	}

	{
		templates := []string {
			"{{.BeginSymbol}}x{{.EndSymbol}}",
			"{{.BeginSymbol}}And she said, {{.BeginSymbol}}do you know him?{{.EndSymbol}}.{{.EndSymbol}}",
			"{{.BeginSymbol}}1 + {{.BeginSymbol}}2 + {{.BeginSymbol}}3 + ...{{.EndSymbol}}{{.EndSymbol}}{{.EndSymbol}}",
			"{{.BeginSymbol}}He said, {{.BeginSymbol}}howdy{{.EndSymbol}}. And she said, {{.BeginSymbol}}hi!{{.EndSymbol}}.{{.EndSymbol}}",
		}

		for _, t := range templates {
			for _, pair := range pairs {

				test := struct{
					BeginSymbol          rune
					EndSymbol            rune
					Source             []byte
					ExpectedBytesRead    int
					ExpectedBytesWritten int
					Expected           []byte
				}{}

				test.BeginSymbol = pair.BeginSymbol
				test.EndSymbol   = pair.EndSymbol
				test.Source = []byte(tmpl.Sprintt(t, struct{
					BeginSymbol string
					EndSymbol string
				}{
					BeginSymbol: string(pair.BeginSymbol),
					EndSymbol:   string(pair.EndSymbol),
				}))
				test.ExpectedBytesRead = len(test.Source)
				test.ExpectedBytesWritten = len(test.Source) - len(string(pair.BeginSymbol)) - len(string(pair.EndSymbol))

				test.Expected = []byte(test.Source)
				test.Expected = test.Expected[len(string(pair.BeginSymbol)):]
				test.Expected = test.Expected[:len(test.Expected)-len(string(pair.EndSymbol))]

				tests = append(tests, test)


				test2 := test
				copy(test2.Source, test.Source)
				copy(test2.Expected, test.Expected)
				test2.Source = append(test2.Source, ` "wow"`...)

				tests = append(tests, test2)
			}
		}
	}

	var variousRunes []rune = []rune{
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

		for _, pair := range pairs {

			test := struct{
				BeginSymbol          rune
				EndSymbol            rune
				Source             []byte
				ExpectedBytesRead    int
				ExpectedBytesWritten int
				Expected           []byte
			}{}

			test.BeginSymbol = pair.BeginSymbol
			test.EndSymbol   = pair.EndSymbol
			test.Source = []byte(nil)
			test.Source = append(test.Source, string(pair.BeginSymbol)...)
			test.Source = append(test.Source, builder.String()...)
			test.Source = append(test.Source, string(pair.EndSymbol)...)
			test.ExpectedBytesRead = len(test.Source)
			test.ExpectedBytesWritten = len(test.Source) - len(string(pair.BeginSymbol)) - len(string(pair.EndSymbol))

			test.Expected = []byte(test.Source)
			test.Expected = test.Expected[len(string(pair.BeginSymbol)):]
			test.Expected = test.Expected[:len(test.Expected)-len(string(pair.EndSymbol))]

			tests = append(tests, test)


			test2 := test
			copy(test2.Source, test.Source)
			copy(test2.Expected, test.Expected)
			test2.Source = append(test2.Source, ` "wow"`...)

			tests = append(tests, test2)
		}
	}

	for testNumber, test := range tests {

		var paired strlit.Paired
		paired.BeginSymbol = test.BeginSymbol
		paired.EndSymbol   = test.EndSymbol

		var src []byte = test.Source
		lenSrc := len(src)

		var dst []byte = make([]byte, lenSrc)

		byteWritten, bytesRead, err := paired.Decode(dst, src)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one.", testNumber)
			t.Logf("PAIRED: \x1b[93;41m%s...%s\x1b[0m", string(paired.BeginSymbol), string(paired.EndSymbol))

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
			t.Logf("PAIRED: \x1b[93;41m%s...%s\x1b[0m", string(paired.BeginSymbol), string(paired.EndSymbol))

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
			t.Logf("PAIRED: \x1b[93;41m%s...%s\x1b[0m", string(paired.BeginSymbol), string(paired.EndSymbol))

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
			t.Logf("PAIRED: \x1b[93;41m%s...%s\x1b[0m", string(paired.BeginSymbol), string(paired.EndSymbol))

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
