package strlit


import (
	"strings"

	"testing"
)


func TestCompile(t *testing.T) {

	tests := []struct{
		Code     string
		Expected string
	}{
		{
			Code:     `''`,
			Expected:  ``,
		},
		{
			Code:     `""`,
			Expected:  ``,
		},
		{
			Code:     `‘’`,
			Expected:  ``,
		},
		{
			Code:     `“”`,
			Expected:  ``,
		},
		{
			Code:     `«»`,
			Expected:  ``,
		},



		{
			Code:     `' '`,
			Expected:  ` `,
		},
		{
			Code:     `" "`,
			Expected:  ` `,
		},
		{
			Code:     `‘ ’`,
			Expected:  ` `,
		},
		{
			Code:     `“ ”`,
			Expected:  ` `,
		},
		{
			Code:     `« »`,
			Expected:  ` `,
		},



		{
			Code:     `'a'`,
			Expected:  `a`,
		},
		{
			Code:     `"a"`,
			Expected:  `a`,
		},
		{
			Code:     `‘a’`,
			Expected:  `a`,
		},
		{
			Code:     `“a”`,
			Expected:  `a`,
		},
		{
			Code:     `«a»`,
			Expected:  `a`,
		},



		{
			Code:     `'A'`,
			Expected:  `A`,
		},
		{
			Code:     `"A"`,
			Expected:  `A`,
		},
		{
			Code:     `‘A’`,
			Expected:  `A`,
		},
		{
			Code:     `“A”`,
			Expected:  `A`,
		},
		{
			Code:     `«A»`,
			Expected:  `A`,
		},



		{
			Code:     `'b'`,
			Expected:  `b`,
		},
		{
			Code:     `"b"`,
			Expected:  `b`,
		},
		{
			Code:     `‘b’`,
			Expected:  `b`,
		},
		{
			Code:     `“b”`,
			Expected:  `b`,
		},
		{
			Code:     `«b»`,
			Expected:  `b`,
		},



		{
			Code:     `'B'`,
			Expected:  `B`,
		},
		{
			Code:     `"B"`,
			Expected:  `B`,
		},
		{
			Code:     `‘B’`,
			Expected:  `B`,
		},
		{
			Code:     `“B”`,
			Expected:  `B`,
		},
		{
			Code:     `«B»`,
			Expected:  `B`,
		},



		{
			Code:     `'c'`,
			Expected:  `c`,
		},
		{
			Code:     `"c"`,
			Expected:  `c`,
		},
		{
			Code:     `‘c’`,
			Expected:  `c`,
		},
		{
			Code:     `“c”`,
			Expected:  `c`,
		},
		{
			Code:     `«c»`,
			Expected:  `c`,
		},



		{
			Code:     `'C'`,
			Expected:  `C`,
		},
		{
			Code:     `"C"`,
			Expected:  `C`,
		},
		{
			Code:     `‘C’`,
			Expected:  `C`,
		},
		{
			Code:     `“C”`,
			Expected:  `C`,
		},
		{
			Code:     `«C»`,
			Expected:  `C`,
		},



		{
			Code:     `'d'`,
			Expected:  `d`,
		},
		{
			Code:     `"d"`,
			Expected:  `d`,
		},
		{
			Code:     `‘d’`,
			Expected:  `d`,
		},
		{
			Code:     `“d”`,
			Expected:  `d`,
		},
		{
			Code:     `«d»`,
			Expected:  `d`,
		},



		{
			Code:     `'D'`,
			Expected:  `D`,
		},
		{
			Code:     `"D"`,
			Expected:  `D`,
		},
		{
			Code:     `‘D’`,
			Expected:  `D`,
		},
		{
			Code:     `“D”`,
			Expected:  `D`,
		},
		{
			Code:     `«D»`,
			Expected:  `D`,
		},



		{
			Code:     `'e'`,
			Expected:  `e`,
		},
		{
			Code:     `"e"`,
			Expected:  `e`,
		},
		{
			Code:     `‘e’`,
			Expected:  `e`,
		},
		{
			Code:     `“e”`,
			Expected:  `e`,
		},
		{
			Code:     `«e»`,
			Expected:  `e`,
		},



		{
			Code:     `'E'`,
			Expected:  `E`,
		},
		{
			Code:     `"E"`,
			Expected:  `E`,
		},
		{
			Code:     `‘E’`,
			Expected:  `E`,
		},
		{
			Code:     `“E”`,
			Expected:  `E`,
		},
		{
			Code:     `«E»`,
			Expected:  `E`,
		},



		{
			Code:     `'0'`,
			Expected:  `0`,
		},
		{
			Code:     `"0"`,
			Expected:  `0`,
		},
		{
			Code:     `‘0’`,
			Expected:  `0`,
		},
		{
			Code:     `“0”`,
			Expected:  `0`,
		},
		{
			Code:     `«0»`,
			Expected:  `0`,
		},



		{
			Code:     `'1'`,
			Expected:  `1`,
		},
		{
			Code:     `"1"`,
			Expected:  `1`,
		},
		{
			Code:     `‘1’`,
			Expected:  `1`,
		},
		{
			Code:     `“1”`,
			Expected:  `1`,
		},
		{
			Code:     `«1»`,
			Expected:  `1`,
		},



		{
			Code:     `'2'`,
			Expected:  `2`,
		},
		{
			Code:     `"2"`,
			Expected:  `2`,
		},
		{
			Code:     `‘2’`,
			Expected:  `2`,
		},
		{
			Code:     `“2”`,
			Expected:  `2`,
		},
		{
			Code:     `«2»`,
			Expected:  `2`,
		},



		{
			Code:     `'3'`,
			Expected:  `3`,
		},
		{
			Code:     `"3"`,
			Expected:  `3`,
		},
		{
			Code:     `‘3’`,
			Expected:  `3`,
		},
		{
			Code:     `“3”`,
			Expected:  `3`,
		},
		{
			Code:     `«3»`,
			Expected:  `3`,
		},



		{
			Code:     `'4'`,
			Expected:  `4`,
		},
		{
			Code:     `"4"`,
			Expected:  `4`,
		},
		{
			Code:     `‘4’`,
			Expected:  `4`,
		},
		{
			Code:     `“4”`,
			Expected:  `4`,
		},
		{
			Code:     `«4»`,
			Expected:  `4`,
		},



		{
			Code:     `'5'`,
			Expected:  `5`,
		},
		{
			Code:     `"5"`,
			Expected:  `5`,
		},
		{
			Code:     `‘5’`,
			Expected:  `5`,
		},
		{
			Code:     `“5”`,
			Expected:  `5`,
		},
		{
			Code:     `«5»`,
			Expected:  `5`,
		},



		{
			Code:     `'6'`,
			Expected:  `6`,
		},
		{
			Code:     `"6"`,
			Expected:  `6`,
		},
		{
			Code:     `‘6’`,
			Expected:  `6`,
		},
		{
			Code:     `“6”`,
			Expected:  `6`,
		},
		{
			Code:     `«6»`,
			Expected:  `6`,
		},



		{
			Code:     `'7'`,
			Expected:  `7`,
		},
		{
			Code:     `"7"`,
			Expected:  `7`,
		},
		{
			Code:     `‘7’`,
			Expected:  `7`,
		},
		{
			Code:     `“7”`,
			Expected:  `7`,
		},
		{
			Code:     `«7»`,
			Expected:  `7`,
		},



		{
			Code:     `'8'`,
			Expected:  `8`,
		},
		{
			Code:     `"8"`,
			Expected:  `8`,
		},
		{
			Code:     `‘8’`,
			Expected:  `8`,
		},
		{
			Code:     `“8”`,
			Expected:  `8`,
		},
		{
			Code:     `«8»`,
			Expected:  `8`,
		},



		{
			Code:     `'9'`,
			Expected:  `9`,
		},
		{
			Code:     `"9"`,
			Expected:  `9`,
		},
		{
			Code:     `‘9’`,
			Expected:  `9`,
		},
		{
			Code:     `“9”`,
			Expected:  `9`,
		},
		{
			Code:     `«9»`,
			Expected:  `9`,
		},



		{
			Code:     `'123456789'`,
			Expected:  `123456789`,
		},
		{
			Code:     `"123456789"`,
			Expected:  `123456789`,
		},
		{
			Code:     `‘123456789’`,
			Expected:  `123456789`,
		},
		{
			Code:     `“123456789”`,
			Expected:  `123456789`,
		},
		{
			Code:     `«123456789»`,
			Expected:  `123456789`,
		},



		{
			Code:     `'a1b2c3d4e5f6g7h8i9'`,
			Expected:  `a1b2c3d4e5f6g7h8i9`,
		},
		{
			Code:     `"a1b2c3d4e5f6g7h8i9"`,
			Expected:  `a1b2c3d4e5f6g7h8i9`,
		},
		{
			Code:     `‘a1b2c3d4e5f6g7h8i9’`,
			Expected:  `a1b2c3d4e5f6g7h8i9`,
		},
		{
			Code:     `“a1b2c3d4e5f6g7h8i9”`,
			Expected:  `a1b2c3d4e5f6g7h8i9`,
		},
		{
			Code:     `«a1b2c3d4e5f6g7h8i9»`,
			Expected:  `a1b2c3d4e5f6g7h8i9`,
		},



		{
			Code:     `'!'`,
			Expected:  `!`,
		},
		{
			Code:     `"!"`,
			Expected:  `!`,
		},
		{
			Code:     `‘!’`,
			Expected:  `!`,
		},
		{
			Code:     `“!”`,
			Expected:  `!`,
		},
		{
			Code:     `«!»`,
			Expected:  `!`,
		},



		{
			Code:     `'@'`,
			Expected:  `@`,
		},
		{
			Code:     `"@"`,
			Expected:  `@`,
		},
		{
			Code:     `‘@’`,
			Expected:  `@`,
		},
		{
			Code:     `“@”`,
			Expected:  `@`,
		},
		{
			Code:     `«@»`,
			Expected:  `@`,
		},



		{
			Code:     `'#'`,
			Expected:  `#`,
		},
		{
			Code:     `"#"`,
			Expected:  `#`,
		},
		{
			Code:     `‘#’`,
			Expected:  `#`,
		},
		{
			Code:     `“#”`,
			Expected:  `#`,
		},
		{
			Code:     `«#»`,
			Expected:  `#`,
		},



		{
			Code:     `'$'`,
			Expected:  `$`,
		},
		{
			Code:     `"$"`,
			Expected:  `$`,
		},
		{
			Code:     `‘$’`,
			Expected:  `$`,
		},
		{
			Code:     `“$”`,
			Expected:  `$`,
		},
		{
			Code:     `«$»`,
			Expected:  `$`,
		},



		{
			Code:     `'%'`,
			Expected:  `%`,
		},
		{
			Code:     `"%"`,
			Expected:  `%`,
		},
		{
			Code:     `‘%’`,
			Expected:  `%`,
		},
		{
			Code:     `“%”`,
			Expected:  `%`,
		},
		{
			Code:     `«%»`,
			Expected:  `%`,
		},



		{
			Code:     `'^'`,
			Expected:  `^`,
		},
		{
			Code:     `"^"`,
			Expected:  `^`,
		},
		{
			Code:     `‘^’`,
			Expected:  `^`,
		},
		{
			Code:     `“^”`,
			Expected:  `^`,
		},
		{
			Code:     `«^»`,
			Expected:  `^`,
		},



		{
			Code:     `'&'`,
			Expected:  `&`,
		},
		{
			Code:     `"&"`,
			Expected:  `&`,
		},
		{
			Code:     `‘&’`,
			Expected:  `&`,
		},
		{
			Code:     `“&”`,
			Expected:  `&`,
		},
		{
			Code:     `«&»`,
			Expected:  `&`,
		},



		{
			Code:     `'*'`,
			Expected:  `*`,
		},
		{
			Code:     `"*"`,
			Expected:  `*`,
		},
		{
			Code:     `‘*’`,
			Expected:  `*`,
		},
		{
			Code:     `“*”`,
			Expected:  `*`,
		},
		{
			Code:     `«*»`,
			Expected:  `*`,
		},



		{
			Code:     `'('`,
			Expected:  `(`,
		},
		{
			Code:     `"("`,
			Expected:  `(`,
		},
		{
			Code:     `‘(’`,
			Expected:  `(`,
		},
		{
			Code:     `“(”`,
			Expected:  `(`,
		},
		{
			Code:     `«(»`,
			Expected:  `(`,
		},



		{
			Code:     `')'`,
			Expected:  `)`,
		},
		{
			Code:     `")"`,
			Expected:  `)`,
		},
		{
			Code:     `‘)’`,
			Expected:  `)`,
		},
		{
			Code:     `“)”`,
			Expected:  `)`,
		},
		{
			Code:     `«)»`,
			Expected:  `)`,
		},



		{
			Code:     `'apple banana cherry'`,
			Expected:  `apple banana cherry`,
		},
		{
			Code:     `"apple banana cherry"`,
			Expected:  `apple banana cherry`,
		},
		{
			Code:     `‘apple banana cherry’`,
			Expected:  `apple banana cherry`,
		},
		{
			Code:     `“apple banana cherry”`,
			Expected:  `apple banana cherry`,
		},
		{
			Code:     `«apple banana cherry»`,
			Expected:  `apple banana cherry`,
		},



//@TODO: ###############################################
		{
			Code:     `''''`,
			Expected: `'`,
		},
		{
			Code:     `""""`,
			Expected: `"`,
		},



		{
			Code:     `' '''`,
			Expected: ` '`,
		},
		{
			Code:     `" """`,
			Expected: ` "`,
		},



		{
			Code:     `''' '`,
			Expected: `' `,
		},
		{
			Code:     `""" "`,
			Expected: `" `,
		},



		{
			Code:     `' '' '`,
			Expected: ` ' `,
		},
		{
			Code:     `" "" "`,
			Expected: ` " `,
		},



		{
			Code:     `''''''`,
			Expected: `''`,
		},
		{
			Code:     `""""""`,
			Expected: `""`,
		},



		{
			Code:     `''''''''`,
			Expected: `'''`,
		},
		{
			Code:     `""""""""`,
			Expected: `"""`,
		},



		{
			Code:     `''''''''''`,
			Expected: `''''`,
		},
		{
			Code:     `""""""""""`,
			Expected: `""""`,
		},



		{
			Code:     `''''''''''''`,
			Expected: `'''''`,
		},
		{
			Code:     `""""""""""""`,
			Expected: `"""""`,
		},



		{
			Code:     `'''apple'' ''banana'' ''cherry'''`,
			Expected:     `'apple' 'banana' 'cherry'`,
		},
		{
			Code:     `"""apple"" ""banana"" ""cherry"""`,
			Expected:     `"apple" "banana" "cherry"`,
		},
//@TODO: ###############################################



		{
			Code:     `'\0'`,
			Expected:  "\u0000",
		},
		{
			Code:     `"\0"`,
			Expected:  "\u0000",
		},
		{
			Code:     `‘\0’`,
			Expected:  "\u0000",
		},
		{
			Code:     `“\0”`,
			Expected:  "\u0000",
		},
		{
			Code:     `«\0»`,
			Expected:  "\u0000",
		},



		{
			Code:     `'\a'`,
			Expected:  "\a",
		},
		{
			Code:     `"\a"`,
			Expected:  "\a",
		},
		{
			Code:     `‘\a’`,
			Expected:  "\a",
		},
		{
			Code:     `“\a”`,
			Expected:  "\a",
		},
		{
			Code:     `«\a»`,
			Expected:  "\a",
		},



		{
			Code:     `'\b'`,
			Expected:  "\b",
		},
		{
			Code:     `"\b"`,
			Expected:  "\b",
		},
		{
			Code:     `‘\b’`,
			Expected:  "\b",
		},
		{
			Code:     `“\b”`,
			Expected:  "\b",
		},
		{
			Code:     `«\b»`,
			Expected:  "\b",
		},



		{
			Code:     `'\e'`,
			Expected:  "\x1B",
		},
		{
			Code:     `"\e"`,
			Expected:  "\x1B",
		},
		{
			Code:     `‘\e’`,
			Expected:  "\x1B",
		},
		{
			Code:     `“\e”`,
			Expected:  "\x1B",
		},
		{
			Code:     `«\e»`,
			Expected:  "\x1B",
		},



		{
			Code:     `'\f'`,
			Expected:  "\f",
		},
		{
			Code:     `"\f"`,
			Expected:  "\f",
		},
		{
			Code:     `‘\f’`,
			Expected:  "\f",
		},
		{
			Code:     `“\f”`,
			Expected:  "\f",
		},
		{
			Code:     `«\f»`,
			Expected:  "\f",
		},



		{
			Code:     `'\n'`,
			Expected:  "\n",
		},
		{
			Code:     `"\n"`,
			Expected:  "\n",
		},
		{
			Code:     `‘\n’`,
			Expected:  "\n",
		},
		{
			Code:     `“\n”`,
			Expected:  "\n",
		},
		{
			Code:     `«\n»`,
			Expected:  "\n",
		},



		{
			Code:     `'\r'`,
			Expected:  "\r",
		},
		{
			Code:     `"\r"`,
			Expected:  "\r",
		},
		{
			Code:     `‘\r’`,
			Expected:  "\r",
		},
		{
			Code:     `“\r”`,
			Expected:  "\r",
		},
		{
			Code:     `«\r»`,
			Expected:  "\r",
		},



		{
			Code:     `'\t'`,
			Expected:  "\t",
		},
		{
			Code:     `"\t"`,
			Expected:  "\t",
		},
		{
			Code:     `‘\t’`,
			Expected:  "\t",
		},
		{
			Code:     `“\t”`,
			Expected:  "\t",
		},
		{
			Code:     `«\t»`,
			Expected:  "\t",
		},



		{
			Code:     `'\v'`,
			Expected:  "\v",
		},
		{
			Code:     `"\v"`,
			Expected:  "\v",
		},
		{
			Code:     `‘\v’`,
			Expected:  "\v",
		},
		{
			Code:     `“\v”`,
			Expected:  "\v",
		},
		{
			Code:     `«\v»`,
			Expected:  "\v",
		},



		{
			Code:     `'\\'`,
			Expected:   `\`,
		},
		{
			Code:     `"\\"`,
			Expected:   `\`,
		},
		{
			Code:     `‘\\’`,
			Expected:   `\`,
		},
		{
			Code:     `“\\”`,
			Expected:   `\`,
		},
		{
			Code:     `«\\»`,
			Expected:   `\`,
		},



		{
			Code:     `'\''`,
			Expected:   `'`,
		},
		{
			Code:     `"\'"`,
			Expected:   `'`,
		},
		{
			Code:     `‘\'’`,
			Expected:   `'`,
		},
		{
			Code:     `“\'”`,
			Expected:   `'`,
		},
		{
			Code:     `«\'»`,
			Expected:   `'`,
		},



		{
			Code:     `'\"'`,
			Expected:   `"`,
		},
		{
			Code:     `"\""`,
			Expected:   `"`,
		},
		{
			Code:     `‘\"’`,
			Expected:   `"`,
		},
		{
			Code:     `“\"”`,
			Expected:   `"`,
		},
		{
			Code:     `«\"»`,
			Expected:   `"`,
		},



		{
			Code:     `'\?'`,
			Expected:   `?`,
		},
		{
			Code:     `"\?"`,
			Expected:   `?`,
		},
		{
			Code:     `‘\?’`,
			Expected:   `?`,
		},
		{
			Code:     `“\?”`,
			Expected:   `?`,
		},
		{
			Code:     `«\?»`,
			Expected:   `?`,
		},



		{
			Code:     `'\‘'`,
			Expected:   `‘`,
		},
		{
			Code:     `"\‘"`,
			Expected:   `‘`,
		},
		{
			Code:     `‘\‘’`,
			Expected:   `‘`,
		},
		{
			Code:     `“\‘”`,
			Expected:   `‘`,
		},
		{
			Code:     `«\‘»`,
			Expected:   `‘`,
		},



		{
			Code:     `'\’'`,
			Expected:   `’`,
		},
		{
			Code:     `"\’"`,
			Expected:   `’`,
		},
		{
			Code:     `‘\’’`,
			Expected:   `’`,
		},
		{
			Code:     `“\’”`,
			Expected:   `’`,
		},
		{
			Code:     `«\’»`,
			Expected:   `’`,
		},



		{
			Code:     `'\“'`,
			Expected:   `“`,
		},
		{
			Code:     `"\“"`,
			Expected:   `“`,
		},
		{
			Code:     `‘\“’`,
			Expected:   `“`,
		},
		{
			Code:     `“\“”`,
			Expected:   `“`,
		},
		{
			Code:     `«\“»`,
			Expected:   `“`,
		},



		{
			Code:     `'\”'`,
			Expected:   `”`,
		},
		{
			Code:     `"\”"`,
			Expected:   `”`,
		},
		{
			Code:     `‘\”’`,
			Expected:   `”`,
		},
		{
			Code:     `“\””`,
			Expected:   `”`,
		},
		{
			Code:     `«\”»`,
			Expected:   `”`,
		},



		{
			Code:     `'\x00'`,
			Expected:  "\x00",
		},
		{
			Code:     `"\x00"`,
			Expected:  "\x00",
		},
		{
			Code:     `‘\x00’`,
			Expected:  "\x00",
		},
		{
			Code:     `“\x00”`,
			Expected:  "\x00",
		},
		{
			Code:     `«\x00»`,
			Expected:  "\x00",
		},



		{
			Code:     `'\x01'`,
			Expected:  "\x01",
		},
		{
			Code:     `"\x01"`,
			Expected:  "\x01",
		},
		{
			Code:     `‘\x01’`,
			Expected:  "\x01",
		},
		{
			Code:     `“\x01”`,
			Expected:  "\x01",
		},
		{
			Code:     `«\x01»`,
			Expected:  "\x01",
		},



		{
			Code:     `'\x02'`,
			Expected:  "\x02",
		},
		{
			Code:     `"\x02"`,
			Expected:  "\x02",
		},
		{
			Code:     `‘\x02’`,
			Expected:  "\x02",
		},
		{
			Code:     `“\x02”`,
			Expected:  "\x02",
		},
		{
			Code:     `«\x02»`,
			Expected:  "\x02",
		},



		{
			Code:     `'\x03'`,
			Expected:  "\x03",
		},
		{
			Code:     `"\x03"`,
			Expected:  "\x03",
		},
		{
			Code:     `‘\x03’`,
			Expected:  "\x03",
		},
		{
			Code:     `“\x03”`,
			Expected:  "\x03",
		},
		{
			Code:     `«\x03»`,
			Expected:  "\x03",
		},



		{
			Code:     `'\x04'`,
			Expected:  "\x04",
		},
		{
			Code:     `"\x04"`,
			Expected:  "\x04",
		},
		{
			Code:     `‘\x04’`,
			Expected:  "\x04",
		},
		{
			Code:     `“\x04”`,
			Expected:  "\x04",
		},
		{
			Code:     `«\x04»`,
			Expected:  "\x04",
		},



		{
			Code:     `'\x05'`,
			Expected:  "\x05",
		},
		{
			Code:     `"\x05"`,
			Expected:  "\x05",
		},
		{
			Code:     `‘\x05’`,
			Expected:  "\x05",
		},
		{
			Code:     `“\x05”`,
			Expected:  "\x05",
		},
		{
			Code:     `«\x05»`,
			Expected:  "\x05",
		},



		{
			Code:     `'\x06'`,
			Expected:  "\x06",
		},
		{
			Code:     `"\x06"`,
			Expected:  "\x06",
		},
		{
			Code:     `‘\x06’`,
			Expected:  "\x06",
		},
		{
			Code:     `“\x06”`,
			Expected:  "\x06",
		},
		{
			Code:     `«\x06»`,
			Expected:  "\x06",
		},



		{
			Code:     `'\x07'`,
			Expected:  "\x07",
		},
		{
			Code:     `"\x07"`,
			Expected:  "\x07",
		},
		{
			Code:     `‘\x07’`,
			Expected:  "\x07",
		},
		{
			Code:     `“\x07”`,
			Expected:  "\x07",
		},
		{
			Code:     `«\x07»`,
			Expected:  "\x07",
		},



		{
			Code:     `'\x08'`,
			Expected:  "\x08",
		},
		{
			Code:     `"\x08"`,
			Expected:  "\x08",
		},
		{
			Code:     `‘\x08’`,
			Expected:  "\x08",
		},
		{
			Code:     `“\x08”`,
			Expected:  "\x08",
		},
		{
			Code:     `«\x08»`,
			Expected:  "\x08",
		},



		{
			Code:     `'\x09'`,
			Expected:  "\x09",
		},
		{
			Code:     `"\x09"`,
			Expected:  "\x09",
		},
		{
			Code:     `‘\x09’`,
			Expected:  "\x09",
		},
		{
			Code:     `“\x09”`,
			Expected:  "\x09",
		},
		{
			Code:     `«\x09»`,
			Expected:  "\x09",
		},



		{
			Code:     `'\x0A'`,
			Expected:  "\x0A",
		},
		{
			Code:     `"\x0A"`,
			Expected:  "\x0A",
		},
		{
			Code:     `‘\x0A’`,
			Expected:  "\x0A",
		},
		{
			Code:     `“\x0A”`,
			Expected:  "\x0A",
		},
		{
			Code:     `«\x0A»`,
			Expected:  "\x0A",
		},



		{
			Code:     `'\x0a'`,
			Expected:  "\x0a",
		},
		{
			Code:     `"\x0a"`,
			Expected:  "\x0a",
		},
		{
			Code:     `‘\x0a’`,
			Expected:  "\x0a",
		},
		{
			Code:     `“\x0a”`,
			Expected:  "\x0a",
		},
		{
			Code:     `«\x0a»`,
			Expected:  "\x0a",
		},



		{
			Code:     `'\x0B'`,
			Expected:  "\x0B",
		},
		{
			Code:     `"\x0B"`,
			Expected:  "\x0B",
		},
		{
			Code:     `‘\x0B’`,
			Expected:  "\x0B",
		},
		{
			Code:     `“\x0B”`,
			Expected:  "\x0B",
		},
		{
			Code:     `«\x0B»`,
			Expected:  "\x0b",
		},



		{
			Code:     `'\x0b'`,
			Expected:  "\x0b",
		},
		{
			Code:     `"\x0b"`,
			Expected:  "\x0b",
		},
		{
			Code:     `‘\x0b’`,
			Expected:  "\x0b",
		},
		{
			Code:     `“\x0b”`,
			Expected:  "\x0b",
		},
		{
			Code:     `«\x0b»`,
			Expected:  "\x0b",
		},



		{
			Code:     `'\x0C'`,
			Expected:  "\x0C",
		},
		{
			Code:     `"\x0C"`,
			Expected:  "\x0C",
		},
		{
			Code:     `‘\x0C’`,
			Expected:  "\x0C",
		},
		{
			Code:     `“\x0C”`,
			Expected:  "\x0C",
		},
		{
			Code:     `«\x0C»`,
			Expected:  "\x0C",
		},



		{
			Code:     `'\x0c'`,
			Expected:  "\x0c",
		},
		{
			Code:     `"\x0c"`,
			Expected:  "\x0c",
		},
		{
			Code:     `‘\x0c’`,
			Expected:  "\x0c",
		},
		{
			Code:     `“\x0c”`,
			Expected:  "\x0c",
		},
		{
			Code:     `«\x0c»`,
			Expected:  "\x0c",
		},



		{
			Code:     `'\x0D'`,
			Expected:  "\x0D",
		},
		{
			Code:     `"\x0D"`,
			Expected:  "\x0D",
		},
		{
			Code:     `‘\x0D’`,
			Expected:  "\x0D",
		},
		{
			Code:     `“\x0D”`,
			Expected:  "\x0D",
		},
		{
			Code:     `«\x0D»`,
			Expected:  "\x0D",
		},



		{
			Code:     `'\x0d'`,
			Expected:  "\x0d",
		},
		{
			Code:     `"\x0d"`,
			Expected:  "\x0d",
		},
		{
			Code:     `‘\x0d’`,
			Expected:  "\x0d",
		},
		{
			Code:     `“\x0d”`,
			Expected:  "\x0d",
		},
		{
			Code:     `«\x0d»`,
			Expected:  "\x0d",
		},



		{
			Code:     `'\x0E'`,
			Expected:  "\x0e",
		},
		{
			Code:     `"\x0E"`,
			Expected:  "\x0e",
		},
		{
			Code:     `‘\x0E’`,
			Expected:  "\x0e",
		},
		{
			Code:     `“\x0E”`,
			Expected:  "\x0e",
		},
		{
			Code:     `«\x0E»`,
			Expected:  "\x0e",
		},



		{
			Code:     `'\x0e'`,
			Expected:  "\x0e",
		},
		{
			Code:     `"\x0e"`,
			Expected:  "\x0e",
		},
		{
			Code:     `‘\x0e’`,
			Expected:  "\x0e",
		},
		{
			Code:     `“\x0e”`,
			Expected:  "\x0e",
		},
		{
			Code:     `«\x0e»`,
			Expected:  "\x0e",
		},



		{
			Code:     `'\x0F'`,
			Expected:  "\x0f",
		},
		{
			Code:     `"\x0F"`,
			Expected:  "\x0f",
		},
		{
			Code:     `‘\x0F’`,
			Expected:  "\x0f",
		},
		{
			Code:     `“\x0F”`,
			Expected:  "\x0f",
		},
		{
			Code:     `«\x0F»`,
			Expected:  "\x0f",
		},



		{
			Code:     `'\x0f'`,
			Expected:  "\x0f",
		},
		{
			Code:     `"\x0f"`,
			Expected:  "\x0f",
		},
		{
			Code:     `‘\x0f’`,
			Expected:  "\x0f",
		},
		{
			Code:     `“\x0f”`,
			Expected:  "\x0f",
		},
		{
			Code:     `«\x0f»`,
			Expected:  "\x0f",
		},



		{
			Code:     `'\x10'`,
			Expected:  "\x10",
		},
		{
			Code:     `"\x10"`,
			Expected:  "\x10",
		},
		{
			Code:     `‘\x10’`,
			Expected:  "\x10",
		},
		{
			Code:     `“\x10”`,
			Expected:  "\x10",
		},
		{
			Code:     `«\x10»`,
			Expected:  "\x10",
		},



		{
			Code:     `'\x11'`,
			Expected:  "\x11",
		},
		{
			Code:     `"\x11"`,
			Expected:  "\x11",
		},
		{
			Code:     `‘\x11’`,
			Expected:  "\x11",
		},
		{
			Code:     `“\x11”`,
			Expected:  "\x11",
		},
		{
			Code:     `«\x11»`,
			Expected:  "\x11",
		},



		{
			Code:     `'\x12'`,
			Expected:  "\x12",
		},
		{
			Code:     `"\x12"`,
			Expected:  "\x12",
		},
		{
			Code:     `‘\x12’`,
			Expected:  "\x12",
		},
		{
			Code:     `“\x12”`,
			Expected:  "\x12",
		},
		{
			Code:     `«\x12»`,
			Expected:  "\x12",
		},



		{
			Code:     `'\x13'`,
			Expected:  "\x13",
		},
		{
			Code:     `"\x13"`,
			Expected:  "\x13",
		},
		{
			Code:     `‘\x13’`,
			Expected:  "\x13",
		},
		{
			Code:     `“\x13”`,
			Expected:  "\x13",
		},
		{
			Code:     `«\x13»`,
			Expected:  "\x13",
		},



		{
			Code:     `'\x14'`,
			Expected:  "\x14",
		},
		{
			Code:     `"\x14"`,
			Expected:  "\x14",
		},
		{
			Code:     `‘\x14’`,
			Expected:  "\x14",
		},
		{
			Code:     `“\x14”`,
			Expected:  "\x14",
		},
		{
			Code:     `«\x14»`,
			Expected:  "\x14",
		},



		{
			Code:     `'\x15'`,
			Expected:  "\x15",
		},
		{
			Code:     `"\x15"`,
			Expected:  "\x15",
		},
		{
			Code:     `‘\x15’`,
			Expected:  "\x15",
		},
		{
			Code:     `“\x15”`,
			Expected:  "\x15",
		},
		{
			Code:     `«\x15»`,
			Expected:  "\x15",
		},



		{
			Code:     `'\x16'`,
			Expected:  "\x16",
		},
		{
			Code:     `"\x16"`,
			Expected:  "\x16",
		},
		{
			Code:     `‘\x16’`,
			Expected:  "\x16",
		},
		{
			Code:     `“\x16”`,
			Expected:  "\x16",
		},
		{
			Code:     `«\x16»`,
			Expected:  "\x16",
		},



		{
			Code:     `'\x17'`,
			Expected:  "\x17",
		},
		{
			Code:     `"\x17"`,
			Expected:  "\x17",
		},
		{
			Code:     `‘\x17’`,
			Expected:  "\x17",
		},
		{
			Code:     `“\x17”`,
			Expected:  "\x17",
		},
		{
			Code:     `«\x17»`,
			Expected:  "\x17",
		},



		{
			Code:     `'\x18'`,
			Expected:  "\x18",
		},
		{
			Code:     `"\x18"`,
			Expected:  "\x18",
		},
		{
			Code:     `‘\x18’`,
			Expected:  "\x18",
		},
		{
			Code:     `“\x18”`,
			Expected:  "\x18",
		},
		{
			Code:     `«\x18»`,
			Expected:  "\x18",
		},



		{
			Code:     `'\x19'`,
			Expected:  "\x19",
		},
		{
			Code:     `"\x19"`,
			Expected:  "\x19",
		},
		{
			Code:     `‘\x19’`,
			Expected:  "\x19",
		},
		{
			Code:     `“\x19”`,
			Expected:  "\x19",
		},
		{
			Code:     `«\x19»`,
			Expected:  "\x19",
		},



		{
			Code:     `'\x1A'`,
			Expected:  "\x1A",
		},
		{
			Code:     `"\x1A"`,
			Expected:  "\x1A",
		},
		{
			Code:     `‘\x1A’`,
			Expected:  "\x1A",
		},
		{
			Code:     `“\x1A”`,
			Expected:  "\x1A",
		},
		{
			Code:     `«\x1A»`,
			Expected:  "\x1A",
		},



		{
			Code:     `'\x1a'`,
			Expected:  "\x1a",
		},
		{
			Code:     `"\x1a"`,
			Expected:  "\x1a",
		},
		{
			Code:     `‘\x1a’`,
			Expected:  "\x1a",
		},
		{
			Code:     `“\x1a”`,
			Expected:  "\x1a",
		},
		{
			Code:     `«\x1a»`,
			Expected:  "\x1a",
		},



		{
			Code:     `'\x1B'`,
			Expected:  "\x1B",
		},
		{
			Code:     `"\x1B"`,
			Expected:  "\x1B",
		},
		{
			Code:     `‘\x1B’`,
			Expected:  "\x1B",
		},
		{
			Code:     `“\x1B”`,
			Expected:  "\x1B",
		},
		{
			Code:     `«\x1B»`,
			Expected:  "\x1B",
		},



		{
			Code:     `'\x1b'`,
			Expected:  "\x1b",
		},
		{
			Code:     `"\x1b"`,
			Expected:  "\x1b",
		},
		{
			Code:     `‘\x1b’`,
			Expected:  "\x1b",
		},
		{
			Code:     `“\x1b”`,
			Expected:  "\x1b",
		},
		{
			Code:     `«\x1b»`,
			Expected:  "\x1b",
		},



		{
			Code:     `'\x1C'`,
			Expected:  "\x1C",
		},
		{
			Code:     `"\x1C"`,
			Expected:  "\x1C",
		},
		{
			Code:     `‘\x1C’`,
			Expected:  "\x1C",
		},
		{
			Code:     `“\x1C”`,
			Expected:  "\x1C",
		},
		{
			Code:     `«\x1C»`,
			Expected:  "\x1C",
		},



		{
			Code:     `'\x1c'`,
			Expected:  "\x1c",
		},
		{
			Code:     `"\x1c"`,
			Expected:  "\x1c",
		},
		{
			Code:     `‘\x1c’`,
			Expected:  "\x1c",
		},
		{
			Code:     `“\x1c”`,
			Expected:  "\x1c",
		},
		{
			Code:     `«\x1c»`,
			Expected:  "\x1c",
		},



		{
			Code:     `'\x1D'`,
			Expected:  "\x1D",
		},
		{
			Code:     `"\x1D"`,
			Expected:  "\x1D",
		},
		{
			Code:     `‘\x1D’`,
			Expected:  "\x1D",
		},
		{
			Code:     `“\x1D”`,
			Expected:  "\x1D",
		},
		{
			Code:     `«\x1D»`,
			Expected:  "\x1D",
		},



		{
			Code:     `'\x1d'`,
			Expected:  "\x1d",
		},
		{
			Code:     `"\x1d"`,
			Expected:  "\x1d",
		},
		{
			Code:     `‘\x1d’`,
			Expected:  "\x1d",
		},
		{
			Code:     `“\x1d”`,
			Expected:  "\x1d",
		},
		{
			Code:     `«\x1d»`,
			Expected:  "\x1d",
		},



		{
			Code:     `'\x1E'`,
			Expected:  "\x1E",
		},
		{
			Code:     `"\x1E"`,
			Expected:  "\x1E",
		},
		{
			Code:     `‘\x1E’`,
			Expected:  "\x1E",
		},
		{
			Code:     `“\x1E”`,
			Expected:  "\x1E",
		},
		{
			Code:     `«\x1E»`,
			Expected:  "\x1E",
		},



		{
			Code:     `'\x1e'`,
			Expected:  "\x1e",
		},
		{
			Code:     `"\x1e"`,
			Expected:  "\x1e",
		},
		{
			Code:     `‘\x1e’`,
			Expected:  "\x1e",
		},
		{
			Code:     `“\x1e”`,
			Expected:  "\x1e",
		},
		{
			Code:     `«\x1e»`,
			Expected:  "\x1e",
		},



		{
			Code:     `'\x1F'`,
			Expected:  "\x1F",
		},
		{
			Code:     `"\x1F"`,
			Expected:  "\x1F",
		},
		{
			Code:     `‘\x1F’`,
			Expected:  "\x1F",
		},
		{
			Code:     `“\x1F”`,
			Expected:  "\x1F",
		},
		{
			Code:     `«\x1F»`,
			Expected:  "\x1F",
		},



		{
			Code:     `'\x1f'`,
			Expected:  "\x1f",
		},
		{
			Code:     `"\x1f"`,
			Expected:  "\x1f",
		},
		{
			Code:     `‘\x1f’`,
			Expected:  "\x1f",
		},
		{
			Code:     `“\x1f”`,
			Expected:  "\x1f",
		},
		{
			Code:     `«\x1f»`,
			Expected:  "\x1f",
		},



//@TODO: ###############################################



		{
			Code:     `'\xAA'`,
			Expected:  "ª",
		},
		{
			Code:     `"\xAA"`,
			Expected:  "ª",
		},
		{
			Code:     `‘\xAA’`,
			Expected:  "ª",
		},
		{
			Code:     `“\xAA”`,
			Expected:  "ª",
		},
		{
			Code:     `«\xAA»`,
			Expected:  "ª",
		},



		{
			Code:     `'\xAa'`,
			Expected:  "ª",
		},
		{
			Code:     `"\xAa"`,
			Expected:  "ª",
		},
		{
			Code:     `‘\xAa’`,
			Expected:  "ª",
		},
		{
			Code:     `“\xAa”`,
			Expected:  "ª",
		},
		{
			Code:     `«\xAa»`,
			Expected:  "ª",
		},



		{
			Code:     `'\xaA'`,
			Expected:  "ª",
		},
		{
			Code:     `"\xaA"`,
			Expected:  "ª",
		},
		{
			Code:     `‘\xaA’`,
			Expected:  "ª",
		},
		{
			Code:     `“\xaA”`,
			Expected:  "ª",
		},
		{
			Code:     `«\xaA»`,
			Expected:  "ª",
		},



		{
			Code:     `'\xaa'`,
			Expected:  "ª",
		},
		{
			Code:     `"\xaa"`,
			Expected:  "ª",
		},
		{
			Code:     `‘\xaa’`,
			Expected:  "ª",
		},
		{
			Code:     `“\xaa”`,
			Expected:  "ª",
		},
		{
			Code:     `«\xaa»`,
			Expected:  "ª",
		},



//@TODO: ###############################################



		{
			Code:     `'\xFF'`,
			Expected:  "ÿ",
		},
		{
			Code:     `"\xFF"`,
			Expected:  "ÿ",
		},
		{
			Code:     `‘\xFF’`,
			Expected:  "ÿ",
		},
		{
			Code:     `“\xFF”`,
			Expected:  "ÿ",
		},
		{
			Code:     `«\xFF»`,
			Expected:  "ÿ",
		},



		{
			Code:     `'\xFf'`,
			Expected:  "ÿ",
		},
		{
			Code:     `"\xFf"`,
			Expected:  "ÿ",
		},
		{
			Code:     `‘\xFf’`,
			Expected:  "ÿ",
		},
		{
			Code:     `“\xFf”`,
			Expected:  "ÿ",
		},
		{
			Code:     `«\xFf»`,
			Expected:  "ÿ",
		},



		{
			Code:     `'\xfF'`,
			Expected:  "ÿ",
		},
		{
			Code:     `"\xfF"`,
			Expected:  "ÿ",
		},
		{
			Code:     `‘\xfF’`,
			Expected:  "ÿ",
		},
		{
			Code:     `“\xfF”`,
			Expected:  "ÿ",
		},
		{
			Code:     `«\xfF»`,
			Expected:  "ÿ",
		},



		{
			Code:     `'\xff'`,
			Expected:  "ÿ",
		},
		{
			Code:     `"\xff"`,
			Expected:  "ÿ",
		},
		{
			Code:     `‘\xff’`,
			Expected:  "ÿ",
		},
		{
			Code:     `“\xff”`,
			Expected:  "ÿ",
		},
		{
			Code:     `«\xff»`,
			Expected:  "ÿ",
		},



		{
			Code:     `'\u0000'`,
			Expected:  "\u0000",
		},
		{
			Code:     `"\u0000"`,
			Expected:  "\u0000",
		},
		{
			Code:     `‘\u0000’`,
			Expected:  "\u0000",
		},
		{
			Code:     `“\u0000”`,
			Expected:  "\u0000",
		},
		{
			Code:     `«\u0000»`,
			Expected:  "\u0000",
		},




		{
			Code:     `'\u0001'`,
			Expected:  "\u0001",
		},
		{
			Code:     `"\u0001"`,
			Expected:  "\u0001",
		},
		{
			Code:     `‘\u0001’`,
			Expected:  "\u0001",
		},
		{
			Code:     `“\u0001”`,
			Expected:  "\u0001",
		},
		{
			Code:     `«\u0001»`,
			Expected:  "\u0001",
		},




		{
			Code:     `'\u0002'`,
			Expected:  "\u0002",
		},
		{
			Code:     `"\u0002"`,
			Expected:  "\u0002",
		},
		{
			Code:     `‘\u0002’`,
			Expected:  "\u0002",
		},
		{
			Code:     `“\u0002”`,
			Expected:  "\u0002",
		},
		{
			Code:     `«\u0002»`,
			Expected:  "\u0002",
		},




		{
			Code:     `'\u0003'`,
			Expected:  "\u0003",
		},
		{
			Code:     `"\u0003"`,
			Expected:  "\u0003",
		},
		{
			Code:     `‘\u0003’`,
			Expected:  "\u0003",
		},
		{
			Code:     `“\u0003”`,
			Expected:  "\u0003",
		},
		{
			Code:     `«\u0003»`,
			Expected:  "\u0003",
		},




		{
			Code:     `'\u0004'`,
			Expected:  "\u0004",
		},
		{
			Code:     `"\u0004"`,
			Expected:  "\u0004",
		},
		{
			Code:     `‘\u0004’`,
			Expected:  "\u0004",
		},
		{
			Code:     `“\u0004”`,
			Expected:  "\u0004",
		},
		{
			Code:     `«\u0004»`,
			Expected:  "\u0004",
		},




		{
			Code:     `'\u0005'`,
			Expected:  "\u0005",
		},
		{
			Code:     `"\u0005"`,
			Expected:  "\u0005",
		},
		{
			Code:     `‘\u0005’`,
			Expected:  "\u0005",
		},
		{
			Code:     `“\u0005”`,
			Expected:  "\u0005",
		},
		{
			Code:     `«\u0005»`,
			Expected:  "\u0005",
		},



		{
			Code:     `'\u0006'`,
			Expected:  "\u0006",
		},
		{
			Code:     `"\u0006"`,
			Expected:  "\u0006",
		},
		{
			Code:     `‘\u0006’`,
			Expected:  "\u0006",
		},
		{
			Code:     `“\u0006”`,
			Expected:  "\u0006",
		},
		{
			Code:     `«\u0006»`,
			Expected:  "\u0006",
		},



		{
			Code:     `'\u0007'`,
			Expected:  "\u0007",
		},
		{
			Code:     `"\u0007"`,
			Expected:  "\u0007",
		},
		{
			Code:     `‘\u0007’`,
			Expected:  "\u0007",
		},
		{
			Code:     `“\u0007”`,
			Expected:  "\u0007",
		},
		{
			Code:     `«\u0007»`,
			Expected:  "\u0007",
		},



		{
			Code:     `'\u0008'`,
			Expected:  "\u0008",
		},
		{
			Code:     `"\u0008"`,
			Expected:  "\u0008",
		},
		{
			Code:     `‘\u0008’`,
			Expected:  "\u0008",
		},
		{
			Code:     `“\u0008”`,
			Expected:  "\u0008",
		},
		{
			Code:     `«\u0008»`,
			Expected:  "\u0008",
		},



		{
			Code:     `'\u0009'`,
			Expected:  "\u0009",
		},
		{
			Code:     `"\u0009"`,
			Expected:  "\u0009",
		},
		{
			Code:     `‘\u0009’`,
			Expected:  "\u0009",
		},
		{
			Code:     `“\u0009”`,
			Expected:  "\u0009",
		},
		{
			Code:     `«\u0009»`,
			Expected:  "\u0009",
		},



		{
			Code:     `'\u000A'`,
			Expected:  "\u000A",
		},
		{
			Code:     `"\u000A"`,
			Expected:  "\u000A",
		},
		{
			Code:     `‘\u000A’`,
			Expected:  "\u000A",
		},
		{
			Code:     `“\u000A”`,
			Expected:  "\u000A",
		},
		{
			Code:     `«\u000A»`,
			Expected:  "\u000A",
		},



		{
			Code:     `'\u000a'`,
			Expected:  "\u000a",
		},
		{
			Code:     `"\u000a"`,
			Expected:  "\u000a",
		},
		{
			Code:     `‘\u000a’`,
			Expected:  "\u000a",
		},
		{
			Code:     `“\u000a”`,
			Expected:  "\u000a",
		},
		{
			Code:     `«\u000a»`,
			Expected:  "\u000a",
		},



		{
			Code:     `'\u000B'`,
			Expected:  "\u000B",
		},
		{
			Code:     `"\u000B"`,
			Expected:  "\u000B",
		},
		{
			Code:     `‘\u000B’`,
			Expected:  "\u000B",
		},
		{
			Code:     `“\u000B”`,
			Expected:  "\u000B",
		},
		{
			Code:     `«\u000B»`,
			Expected:  "\u000B",
		},



		{
			Code:     `'\u000b'`,
			Expected:  "\u000b",
		},
		{
			Code:     `"\u000b"`,
			Expected:  "\u000b",
		},
		{
			Code:     `‘\u000b’`,
			Expected:  "\u000b",
		},
		{
			Code:     `“\u000b”`,
			Expected:  "\u000b",
		},
		{
			Code:     `«\u000b»`,
			Expected:  "\u000b",
		},



		{
			Code:     `'\u000C'`,
			Expected:  "\u000C",
		},
		{
			Code:     `"\u000C"`,
			Expected:  "\u000C",
		},
		{
			Code:     `‘\u000C’`,
			Expected:  "\u000C",
		},
		{
			Code:     `“\u000C”`,
			Expected:  "\u000C",
		},
		{
			Code:     `«\u000C»`,
			Expected:  "\u000C",
		},



		{
			Code:     `'\u000c'`,
			Expected:  "\u000c",
		},
		{
			Code:     `"\u000c"`,
			Expected:  "\u000c",
		},
		{
			Code:     `‘\u000c’`,
			Expected:  "\u000c",
		},
		{
			Code:     `“\u000c”`,
			Expected:  "\u000c",
		},
		{
			Code:     `«\u000c»`,
			Expected:  "\u000c",
		},



		{
			Code:     `'\u000D'`,
			Expected:  "\u000D",
		},
		{
			Code:     `"\u000D"`,
			Expected:  "\u000D",
		},
		{
			Code:     `‘\u000D’`,
			Expected:  "\u000D",
		},
		{
			Code:     `“\u000D”`,
			Expected:  "\u000D",
		},
		{
			Code:     `«\u000D»`,
			Expected:  "\u000D",
		},



		{
			Code:     `'\u000d'`,
			Expected:  "\u000d",
		},
		{
			Code:     `"\u000d"`,
			Expected:  "\u000d",
		},
		{
			Code:     `‘\u000d’`,
			Expected:  "\u000d",
		},
		{
			Code:     `“\u000d”`,
			Expected:  "\u000d",
		},
		{
			Code:     `«\u000d»`,
			Expected:  "\u000d",
		},



		{
			Code:     `'\u000E'`,
			Expected:  "\u000E",
		},
		{
			Code:     `"\u000E"`,
			Expected:  "\u000E",
		},
		{
			Code:     `‘\u000E’`,
			Expected:  "\u000E",
		},
		{
			Code:     `“\u000E”`,
			Expected:  "\u000E",
		},
		{
			Code:     `«\u000E»`,
			Expected:  "\u000E",
		},



		{
			Code:     `'\u000e'`,
			Expected:  "\u000e",
		},
		{
			Code:     `"\u000e"`,
			Expected:  "\u000e",
		},
		{
			Code:     `‘\u000e’`,
			Expected:  "\u000e",
		},
		{
			Code:     `“\u000e”`,
			Expected:  "\u000e",
		},
		{
			Code:     `«\u000e»`,
			Expected:  "\u000e",
		},



		{
			Code:     `'\u000F'`,
			Expected:  "\u000F",
		},
		{
			Code:     `"\u000F"`,
			Expected:  "\u000F",
		},
		{
			Code:     `‘\u000F’`,
			Expected:  "\u000F",
		},
		{
			Code:     `“\u000F”`,
			Expected:  "\u000F",
		},
		{
			Code:     `«\u000F»`,
			Expected:  "\u000F",
		},



		{
			Code:     `'\u000f'`,
			Expected:  "\u000f",
		},
		{
			Code:     `"\u000f"`,
			Expected:  "\u000f",
		},
		{
			Code:     `‘\u000f’`,
			Expected:  "\u000f",
		},
		{
			Code:     `“\u000f”`,
			Expected:  "\u000f",
		},
		{
			Code:     `«\u000f»`,
			Expected:  "\u000f",
		},



//@TODO: ###############################################



		{
			Code:     `'\u0010'`,
			Expected:  "\u0010",
		},
		{
			Code:     `"\u0010"`,
			Expected:  "\u0010",
		},
		{
			Code:     `‘\u0010’`,
			Expected:  "\u0010",
		},
		{
			Code:     `“\u0010”`,
			Expected:  "\u0010",
		},
		{
			Code:     `«\u0010»`,
			Expected:  "\u0010",
		},



		{
			Code:     `'\u0020'`,
			Expected:  "\u0020",
		},
		{
			Code:     `"\u0020"`,
			Expected:  "\u0020",
		},
		{
			Code:     `‘\u0020’`,
			Expected:  "\u0020",
		},
		{
			Code:     `“\u0020”`,
			Expected:  "\u0020",
		},
		{
			Code:     `«\u0020»`,
			Expected:  "\u0020",
		},



		{
			Code:     `'\u0030'`,
			Expected:  "\u0030",
		},
		{
			Code:     `"\u0030"`,
			Expected:  "\u0030",
		},
		{
			Code:     `‘\u0030’`,
			Expected:  "\u0030",
		},
		{
			Code:     `“\u0030”`,
			Expected:  "\u0030",
		},
		{
			Code:     `«\u0030»`,
			Expected:  "\u0030",
		},



		{
			Code:     `'\u0040'`,
			Expected:  "\u0040",
		},
		{
			Code:     `"\u0040"`,
			Expected:  "\u0040",
		},
		{
			Code:     `‘\u0040’`,
			Expected:  "\u0040",
		},
		{
			Code:     `“\u0040”`,
			Expected:  "\u0040",
		},
		{
			Code:     `«\u0040»`,
			Expected:  "\u0040",
		},



		{
			Code:     `'\u0050'`,
			Expected:  "\u0050",
		},
		{
			Code:     `"\u0050"`,
			Expected:  "\u0050",
		},
		{
			Code:     `‘\u0050’`,
			Expected:  "\u0050",
		},
		{
			Code:     `“\u0050”`,
			Expected:  "\u0050",
		},
		{
			Code:     `«\u0050»`,
			Expected:  "\u0050",
		},



		{
			Code:     `'\u0060'`,
			Expected:  "\u0060",
		},
		{
			Code:     `"\u0060"`,
			Expected:  "\u0060",
		},
		{
			Code:     `‘\u0060’`,
			Expected:  "\u0060",
		},
		{
			Code:     `“\u0060”`,
			Expected:  "\u0060",
		},
		{
			Code:     `«\u0060»`,
			Expected:  "\u0060",
		},



		{
			Code:     `'\u0070'`,
			Expected:  "\u0070",
		},
		{
			Code:     `"\u0070"`,
			Expected:  "\u0070",
		},
		{
			Code:     `‘\u0070’`,
			Expected:  "\u0070",
		},
		{
			Code:     `“\u0070”`,
			Expected:  "\u0070",
		},
		{
			Code:     `«\u0070»`,
			Expected:  "\u0070",
		},



		{
			Code:     `'\u0080'`,
			Expected:  "\u0080",
		},
		{
			Code:     `"\u0080"`,
			Expected:  "\u0080",
		},
		{
			Code:     `‘\u0080’`,
			Expected:  "\u0080",
		},
		{
			Code:     `“\u0080”`,
			Expected:  "\u0080",
		},
		{
			Code:     `«\u0080»`,
			Expected:  "\u0080",
		},



		{
			Code:     `'\u0090'`,
			Expected:  "\u0090",
		},
		{
			Code:     `"\u0090"`,
			Expected:  "\u0090",
		},
		{
			Code:     `‘\u0090’`,
			Expected:  "\u0090",
		},
		{
			Code:     `“\u0090”`,
			Expected:  "\u0090",
		},
		{
			Code:     `«\u0090»`,
			Expected:  "\u0090",
		},



		{
			Code:     `'\u00A0'`,
			Expected:  "\u00A0",
		},
		{
			Code:     `"\u00A0"`,
			Expected:  "\u00A0",
		},
		{
			Code:     `‘\u00A0’`,
			Expected:  "\u00A0",
		},
		{
			Code:     `“\u00A0”`,
			Expected:  "\u00A0",
		},
		{
			Code:     `«\u00A0»`,
			Expected:  "\u00A0",
		},



		{
			Code:     `'\u00a0'`,
			Expected:  "\u00a0",
		},
		{
			Code:     `"\u00a0"`,
			Expected:  "\u00a0",
		},
		{
			Code:     `‘\u00a0’`,
			Expected:  "\u00a0",
		},
		{
			Code:     `“\u00a0”`,
			Expected:  "\u00a0",
		},
		{
			Code:     `«\u00a0»`,
			Expected:  "\u00a0",
		},



		{
			Code:     `'\u00B0'`,
			Expected:  "\u00B0",
		},
		{
			Code:     `"\u00B0"`,
			Expected:  "\u00B0",
		},
		{
			Code:     `‘\u00B0’`,
			Expected:  "\u00B0",
		},
		{
			Code:     `“\u00B0”`,
			Expected:  "\u00B0",
		},
		{
			Code:     `«\u00B0»`,
			Expected:  "\u00B0",
		},



		{
			Code:     `'\u00b0'`,
			Expected:  "\u00b0",
		},
		{
			Code:     `"\u00b0"`,
			Expected:  "\u00b0",
		},
		{
			Code:     `‘\u00b0’`,
			Expected:  "\u00b0",
		},
		{
			Code:     `“\u00b0”`,
			Expected:  "\u00b0",
		},
		{
			Code:     `«\u00b0»`,
			Expected:  "\u00b0",
		},



		{
			Code:     `'\u00C0'`,
			Expected:  "\u00C0",
		},
		{
			Code:     `"\u00C0"`,
			Expected:  "\u00C0",
		},
		{
			Code:     `‘\u00C0’`,
			Expected:  "\u00C0",
		},
		{
			Code:     `“\u00C0”`,
			Expected:  "\u00C0",
		},
		{
			Code:     `«\u00C0»`,
			Expected:  "\u00C0",
		},



		{
			Code:     `'\u00c0'`,
			Expected:  "\u00c0",
		},
		{
			Code:     `"\u00c0"`,
			Expected:  "\u00c0",
		},
		{
			Code:     `‘\u00c0’`,
			Expected:  "\u00c0",
		},
		{
			Code:     `“\u00c0”`,
			Expected:  "\u00c0",
		},
		{
			Code:     `«\u00c0»`,
			Expected:  "\u00c0",
		},



		{
			Code:     `'\u00D0'`,
			Expected:  "\u00D0",
		},
		{
			Code:     `"\u00D0"`,
			Expected:  "\u00D0",
		},
		{
			Code:     `‘\u00D0’`,
			Expected:  "\u00D0",
		},
		{
			Code:     `“\u00D0”`,
			Expected:  "\u00D0",
		},
		{
			Code:     `«\u00D0»`,
			Expected:  "\u00D0",
		},



		{
			Code:     `'\u00d0'`,
			Expected:  "\u00d0",
		},
		{
			Code:     `"\u00d0"`,
			Expected:  "\u00d0",
		},
		{
			Code:     `‘\u00d0’`,
			Expected:  "\u00d0",
		},
		{
			Code:     `“\u00d0”`,
			Expected:  "\u00d0",
		},
		{
			Code:     `«\u00d0»`,
			Expected:  "\u00d0",
		},



		{
			Code:     `'\u00E0'`,
			Expected:  "\u00E0",
		},
		{
			Code:     `"\u00E0"`,
			Expected:  "\u00E0",
		},
		{
			Code:     `‘\u00E0’`,
			Expected:  "\u00E0",
		},
		{
			Code:     `“\u00E0”`,
			Expected:  "\u00E0",
		},
		{
			Code:     `«\u00E0»`,
			Expected:  "\u00E0",
		},



		{
			Code:     `'\u00e0'`,
			Expected:  "\u00e0",
		},
		{
			Code:     `"\u00e0"`,
			Expected:  "\u00e0",
		},
		{
			Code:     `‘\u00e0’`,
			Expected:  "\u00e0",
		},
		{
			Code:     `“\u00e0”`,
			Expected:  "\u00e0",
		},
		{
			Code:     `«\u00e0»`,
			Expected:  "\u00e0",
		},



		{
			Code:     `'\u00F0'`,
			Expected:  "\u00F0",
		},
		{
			Code:     `"\u00F0"`,
			Expected:  "\u00F0",
		},
		{
			Code:     `‘\u00F0’`,
			Expected:  "\u00F0",
		},
		{
			Code:     `“\u00F0”`,
			Expected:  "\u00F0",
		},
		{
			Code:     `«\u00F0»`,
			Expected:  "\u00F0",
		},



		{
			Code:     `'\u00f0'`,
			Expected:  "\u00f0",
		},
		{
			Code:     `"\u00f0"`,
			Expected:  "\u00f0",
		},
		{
			Code:     `‘\u00f0’`,
			Expected:  "\u00f0",
		},
		{
			Code:     `“\u00f0”`,
			Expected:  "\u00f0",
		},
		{
			Code:     `«\u00f0»`,
			Expected:  "\u00f0",
		},



//@TODO: ###############################################



		{
			Code:     `'\u0100'`,
			Expected:  "\u0100",
		},
		{
			Code:     `"\u0100"`,
			Expected:  "\u0100",
		},
		{
			Code:     `‘\u0100’`,
			Expected:  "\u0100",
		},
		{
			Code:     `“\u0100”`,
			Expected:  "\u0100",
		},
		{
			Code:     `«\u0100»`,
			Expected:  "\u0100",
		},



		{
			Code:     `'\u0200'`,
			Expected:  "\u0200",
		},
		{
			Code:     `"\u0200"`,
			Expected:  "\u0200",
		},
		{
			Code:     `‘\u0200’`,
			Expected:  "\u0200",
		},
		{
			Code:     `“\u0200”`,
			Expected:  "\u0200",
		},
		{
			Code:     `«\u0200»`,
			Expected:  "\u0200",
		},



		{
			Code:     `'\u0300'`,
			Expected:  "\u0300",
		},
		{
			Code:     `"\u0300"`,
			Expected:  "\u0300",
		},
		{
			Code:     `‘\u0300’`,
			Expected:  "\u0300",
		},
		{
			Code:     `“\u0300”`,
			Expected:  "\u0300",
		},
		{
			Code:     `«\u0300»`,
			Expected:  "\u0300",
		},



		{
			Code:     `'\u0400'`,
			Expected:  "\u0400",
		},
		{
			Code:     `"\u0400"`,
			Expected:  "\u0400",
		},
		{
			Code:     `‘\u0400’`,
			Expected:  "\u0400",
		},
		{
			Code:     `“\u0400”`,
			Expected:  "\u0400",
		},
		{
			Code:     `«\u0400»`,
			Expected:  "\u0400",
		},



		{
			Code:     `'\u0500'`,
			Expected:  "\u0500",
		},
		{
			Code:     `"\u0500"`,
			Expected:  "\u0500",
		},
		{
			Code:     `‘\u0500’`,
			Expected:  "\u0500",
		},
		{
			Code:     `“\u0500”`,
			Expected:  "\u0500",
		},
		{
			Code:     `«\u0500»`,
			Expected:  "\u0500",
		},



		{
			Code:     `'\u0600'`,
			Expected:  "\u0600",
		},
		{
			Code:     `"\u0600"`,
			Expected:  "\u0600",
		},
		{
			Code:     `‘\u0600’`,
			Expected:  "\u0600",
		},
		{
			Code:     `“\u0600”`,
			Expected:  "\u0600",
		},
		{
			Code:     `«\u0600»`,
			Expected:  "\u0600",
		},



		{
			Code:     `'\u0700'`,
			Expected:  "\u0700",
		},
		{
			Code:     `"\u0700"`,
			Expected:  "\u0700",
		},
		{
			Code:     `‘\u0700’`,
			Expected:  "\u0700",
		},
		{
			Code:     `“\u0700”`,
			Expected:  "\u0700",
		},
		{
			Code:     `«\u0700»`,
			Expected:  "\u0700",
		},



		{
			Code:     `'\u0800'`,
			Expected:  "\u0800",
		},
		{
			Code:     `"\u0800"`,
			Expected:  "\u0800",
		},
		{
			Code:     `‘\u0800’`,
			Expected:  "\u0800",
		},
		{
			Code:     `“\u0800”`,
			Expected:  "\u0800",
		},
		{
			Code:     `«\u0800»`,
			Expected:  "\u0800",
		},



		{
			Code:     `'\u0900'`,
			Expected:  "\u0900",
		},
		{
			Code:     `"\u0900"`,
			Expected:  "\u0900",
		},
		{
			Code:     `‘\u0900’`,
			Expected:  "\u0900",
		},
		{
			Code:     `“\u0900”`,
			Expected:  "\u0900",
		},
		{
			Code:     `«\u0900»`,
			Expected:  "\u0900",
		},



		{
			Code:     `'\u0A00'`,
			Expected:  "\u0A00",
		},
		{
			Code:     `"\u0A00"`,
			Expected:  "\u0A00",
		},
		{
			Code:     `‘\u0A00’`,
			Expected:  "\u0A00",
		},
		{
			Code:     `“\u0A00”`,
			Expected:  "\u0A00",
		},
		{
			Code:     `«\u0A00»`,
			Expected:  "\u0A00",
		},



		{
			Code:     `'\u0a00'`,
			Expected:  "\u0a00",
		},
		{
			Code:     `"\u0a00"`,
			Expected:  "\u0a00",
		},
		{
			Code:     `‘\u0a00’`,
			Expected:  "\u0a00",
		},
		{
			Code:     `“\u0a00”`,
			Expected:  "\u0a00",
		},
		{
			Code:     `«\u0a00»`,
			Expected:  "\u0a00",
		},



		{
			Code:     `'\u0B00'`,
			Expected:  "\u0B00",
		},
		{
			Code:     `"\u0B00"`,
			Expected:  "\u0B00",
		},
		{
			Code:     `‘\u0B00’`,
			Expected:  "\u0B00",
		},
		{
			Code:     `“\u0B00”`,
			Expected:  "\u0B00",
		},
		{
			Code:     `«\u0B00»`,
			Expected:  "\u0B00",
		},



		{
			Code:     `'\u0b00'`,
			Expected:  "\u0b00",
		},
		{
			Code:     `"\u0b00"`,
			Expected:  "\u0b00",
		},
		{
			Code:     `‘\u0b00’`,
			Expected:  "\u0b00",
		},
		{
			Code:     `“\u0b00”`,
			Expected:  "\u0b00",
		},
		{
			Code:     `«\u0b00»`,
			Expected:  "\u0b00",
		},



		{
			Code:     `'\u0C00'`,
			Expected:  "\u0C00",
		},
		{
			Code:     `"\u0C00"`,
			Expected:  "\u0C00",
		},
		{
			Code:     `‘\u0C00’`,
			Expected:  "\u0C00",
		},
		{
			Code:     `“\u0C00”`,
			Expected:  "\u0C00",
		},
		{
			Code:     `«\u0C00»`,
			Expected:  "\u0C00",
		},



		{
			Code:     `'\u0c00'`,
			Expected:  "\u0c00",
		},
		{
			Code:     `"\u0c00"`,
			Expected:  "\u0c00",
		},
		{
			Code:     `‘\u0c00’`,
			Expected:  "\u0c00",
		},
		{
			Code:     `“\u0c00”`,
			Expected:  "\u0c00",
		},
		{
			Code:     `«\u0c00»`,
			Expected:  "\u0c00",
		},



		{
			Code:     `'\u0D00'`,
			Expected:  "\u0D00",
		},
		{
			Code:     `"\u0D00"`,
			Expected:  "\u0D00",
		},
		{
			Code:     `‘\u0D00’`,
			Expected:  "\u0D00",
		},
		{
			Code:     `“\u0D00”`,
			Expected:  "\u0D00",
		},
		{
			Code:     `«\u0D00»`,
			Expected:  "\u0D00",
		},



		{
			Code:     `'\u0d00'`,
			Expected:  "\u0d00",
		},
		{
			Code:     `"\u0d00"`,
			Expected:  "\u0d00",
		},
		{
			Code:     `‘\u0d00’`,
			Expected:  "\u0d00",
		},
		{
			Code:     `“\u0d00”`,
			Expected:  "\u0d00",
		},
		{
			Code:     `«\u0d00»`,
			Expected:  "\u0d00",
		},



		{
			Code:     `'\u0E00'`,
			Expected:  "\u0E00",
		},
		{
			Code:     `"\u0E00"`,
			Expected:  "\u0E00",
		},
		{
			Code:     `‘\u0E00’`,
			Expected:  "\u0E00",
		},
		{
			Code:     `“\u0E00”`,
			Expected:  "\u0E00",
		},
		{
			Code:     `«\u0E00»`,
			Expected:  "\u0E00",
		},



		{
			Code:     `'\u0e00'`,
			Expected:  "\u0e00",
		},
		{
			Code:     `"\u0e00"`,
			Expected:  "\u0e00",
		},
		{
			Code:     `‘\u0e00’`,
			Expected:  "\u0e00",
		},
		{
			Code:     `“\u0e00”`,
			Expected:  "\u0e00",
		},
		{
			Code:     `«\u0e00»`,
			Expected:  "\u0e00",
		},



		{
			Code:     `'\u0F00'`,
			Expected:  "\u0F00",
		},
		{
			Code:     `"\u0F00"`,
			Expected:  "\u0F00",
		},
		{
			Code:     `‘\u0F00’`,
			Expected:  "\u0F00",
		},
		{
			Code:     `“\u0F00”`,
			Expected:  "\u0F00",
		},
		{
			Code:     `«\u0F00»`,
			Expected:  "\u0F00",
		},



		{
			Code:     `'\u0f00'`,
			Expected:  "\u0f00",
		},
		{
			Code:     `"\u0f00"`,
			Expected:  "\u0f00",
		},
		{
			Code:     `‘\u0f00’`,
			Expected:  "\u0f00",
		},
		{
			Code:     `“\u0f00”`,
			Expected:  "\u0f00",
		},
		{
			Code:     `«\u0f00»`,
			Expected:  "\u0f00",
		},



//@TODO: ###############################################



		{
			Code:     `'\u1000'`,
			Expected:  "\u1000",
		},
		{
			Code:     `"\u1000"`,
			Expected:  "\u1000",
		},
		{
			Code:     `‘\u1000’`,
			Expected:  "\u1000",
		},
		{
			Code:     `“\u1000”`,
			Expected:  "\u1000",
		},
		{
			Code:     `«\u1000»`,
			Expected:  "\u1000",
		},



		{
			Code:     `'\u2000'`,
			Expected:  "\u2000",
		},
		{
			Code:     `"\u2000"`,
			Expected:  "\u2000",
		},
		{
			Code:     `‘\u2000’`,
			Expected:  "\u2000",
		},
		{
			Code:     `“\u2000”`,
			Expected:  "\u2000",
		},
		{
			Code:     `«\u2000»`,
			Expected:  "\u2000",
		},



		{
			Code:     `'\u3000'`,
			Expected:  "\u3000",
		},
		{
			Code:     `"\u3000"`,
			Expected:  "\u3000",
		},
		{
			Code:     `‘\u3000’`,
			Expected:  "\u3000",
		},
		{
			Code:     `“\u3000”`,
			Expected:  "\u3000",
		},
		{
			Code:     `«\u3000»`,
			Expected:  "\u3000",
		},



		{
			Code:     `'\u4000'`,
			Expected:  "\u4000",
		},
		{
			Code:     `"\u4000"`,
			Expected:  "\u4000",
		},
		{
			Code:     `‘\u4000’`,
			Expected:  "\u4000",
		},
		{
			Code:     `“\u4000”`,
			Expected:  "\u4000",
		},
		{
			Code:     `«\u4000»`,
			Expected:  "\u4000",
		},



		{
			Code:     `'\u5000'`,
			Expected:  "\u5000",
		},
		{
			Code:     `"\u5000"`,
			Expected:  "\u5000",
		},
		{
			Code:     `‘\u5000’`,
			Expected:  "\u5000",
		},
		{
			Code:     `“\u5000”`,
			Expected:  "\u5000",
		},
		{
			Code:     `«\u5000»`,
			Expected:  "\u5000",
		},



		{
			Code:     `'\u6000'`,
			Expected:  "\u6000",
		},
		{
			Code:     `"\u6000"`,
			Expected:  "\u6000",
		},
		{
			Code:     `‘\u6000’`,
			Expected:  "\u6000",
		},
		{
			Code:     `“\u6000”`,
			Expected:  "\u6000",
		},
		{
			Code:     `«\u6000»`,
			Expected:  "\u6000",
		},



		{
			Code:     `'\u7000'`,
			Expected:  "\u7000",
		},
		{
			Code:     `"\u7000"`,
			Expected:  "\u7000",
		},
		{
			Code:     `‘\u7000’`,
			Expected:  "\u7000",
		},
		{
			Code:     `“\u7000”`,
			Expected:  "\u7000",
		},
		{
			Code:     `«\u7000»`,
			Expected:  "\u7000",
		},



		{
			Code:     `'\u8000'`,
			Expected:  "\u8000",
		},
		{
			Code:     `"\u8000"`,
			Expected:  "\u8000",
		},
		{
			Code:     `‘\u8000’`,
			Expected:  "\u8000",
		},
		{
			Code:     `“\u8000”`,
			Expected:  "\u8000",
		},
		{
			Code:     `«\u8000»`,
			Expected:  "\u8000",
		},



		{
			Code:     `'\u9000'`,
			Expected:  "\u9000",
		},
		{
			Code:     `"\u9000"`,
			Expected:  "\u9000",
		},
		{
			Code:     `‘\u9000’`,
			Expected:  "\u9000",
		},
		{
			Code:     `“\u9000”`,
			Expected:  "\u9000",
		},
		{
			Code:     `«\u9000»`,
			Expected:  "\u9000",
		},



		{
			Code:     `'\uA000'`,
			Expected:  "\uA000",
		},
		{
			Code:     `"\uA000"`,
			Expected:  "\uA000",
		},
		{
			Code:     `‘\uA000’`,
			Expected:  "\uA000",
		},
		{
			Code:     `“\uA000”`,
			Expected:  "\uA000",
		},
		{
			Code:     `«\uA000»`,
			Expected:  "\uA000",
		},



		{
			Code:     `'\ua000'`,
			Expected:  "\ua000",
		},
		{
			Code:     `"\ua000"`,
			Expected:  "\ua000",
		},
		{
			Code:     `‘\ua000’`,
			Expected:  "\ua000",
		},
		{
			Code:     `“\ua000”`,
			Expected:  "\ua000",
		},
		{
			Code:     `«\ua000»`,
			Expected:  "\ua000",
		},



		{
			Code:     `'\uB000'`,
			Expected:  "\uB000",
		},
		{
			Code:     `"\uB000"`,
			Expected:  "\uB000",
		},
		{
			Code:     `‘\uB000’`,
			Expected:  "\uB000",
		},
		{
			Code:     `“\uB000”`,
			Expected:  "\uB000",
		},
		{
			Code:     `«\uB000»`,
			Expected:  "\uB000",
		},



		{
			Code:     `'\ub000'`,
			Expected:  "\ub000",
		},
		{
			Code:     `"\ub000"`,
			Expected:  "\ub000",
		},
		{
			Code:     `‘\ub000’`,
			Expected:  "\ub000",
		},
		{
			Code:     `“\ub000”`,
			Expected:  "\ub000",
		},
		{
			Code:     `«\ub000»`,
			Expected:  "\ub000",
		},



		{
			Code:     `'\uC000'`,
			Expected:  "\uC000",
		},
		{
			Code:     `"\uC000"`,
			Expected:  "\uC000",
		},
		{
			Code:     `‘\uC000’`,
			Expected:  "\uC000",
		},
		{
			Code:     `“\uC000”`,
			Expected:  "\uC000",
		},
		{
			Code:     `«\uC000»`,
			Expected:  "\uC000",
		},



		{
			Code:     `'\uc000'`,
			Expected:  "\uc000",
		},
		{
			Code:     `"\uc000"`,
			Expected:  "\uc000",
		},
		{
			Code:     `‘\uc000’`,
			Expected:  "\uc000",
		},
		{
			Code:     `“\uc000”`,
			Expected:  "\uc000",
		},
		{
			Code:     `«\uc000»`,
			Expected:  "\uc000",
		},



		{
			Code:     `'\uD000'`,
			Expected:  "\uD000",
		},
		{
			Code:     `"\uD000"`,
			Expected:  "\uD000",
		},
		{
			Code:     `‘\uD000’`,
			Expected:  "\uD000",
		},
		{
			Code:     `“\uD000”`,
			Expected:  "\uD000",
		},
		{
			Code:     `«\uD000»`,
			Expected:  "\uD000",
		},



		{
			Code:     `'\ud000'`,
			Expected:  "\ud000",
		},
		{
			Code:     `"\ud000"`,
			Expected:  "\ud000",
		},
		{
			Code:     `‘\ud000’`,
			Expected:  "\ud000",
		},
		{
			Code:     `“\ud000”`,
			Expected:  "\ud000",
		},
		{
			Code:     `«\ud000»`,
			Expected:  "\ud000",
		},



		{
			Code:     `'\uE000'`,
			Expected:  "\uE000",
		},
		{
			Code:     `"\uE000"`,
			Expected:  "\uE000",
		},
		{
			Code:     `‘\uE000’`,
			Expected:  "\uE000",
		},
		{
			Code:     `“\uE000”`,
			Expected:  "\uE000",
		},
		{
			Code:     `«\uE000»`,
			Expected:  "\uE000",
		},



		{
			Code:     `'\ue000'`,
			Expected:  "\ue000",
		},
		{
			Code:     `"\ue000"`,
			Expected:  "\ue000",
		},
		{
			Code:     `‘\ue000’`,
			Expected:  "\ue000",
		},
		{
			Code:     `“\ue000”`,
			Expected:  "\ue000",
		},
		{
			Code:     `«\ue000»`,
			Expected:  "\ue000",
		},



		{
			Code:     `'\uF000'`,
			Expected:  "\uF000",
		},
		{
			Code:     `"\uF000"`,
			Expected:  "\uF000",
		},
		{
			Code:     `‘\uF000’`,
			Expected:  "\uF000",
		},
		{
			Code:     `“\uF000”`,
			Expected:  "\uF000",
		},
		{
			Code:     `«\uF000»`,
			Expected:  "\uF000",
		},



		{
			Code:     `'\uf000'`,
			Expected:  "\uf000",
		},
		{
			Code:     `"\uf000"`,
			Expected:  "\uf000",
		},
		{
			Code:     `‘\uf000’`,
			Expected:  "\uf000",
		},
		{
			Code:     `“\uf000”`,
			Expected:  "\uf000",
		},
		{
			Code:     `«\uf000»`,
			Expected:  "\uf000",
		},



//@TODO: ###############################################



		{
			Code:     `'\uf2A7'`,
			Expected:  "\uf2A7",
		},
		{
			Code:     `"\uf2A7"`,
			Expected:  "\uf2A7",
		},
		{
			Code:     `‘\uf2A7’`,
			Expected:  "\uf2A7",
		},
		{
			Code:     `“\uf2A7”`,
			Expected:  "\uf2A7",
		},
		{
			Code:     `«\uf2A7»`,
			Expected:  "\uf2A7",
		},



//@TODO: ###############################################



		{
			Code:     `'\U00000000'`,
			Expected:  "\U00000000",
		},
		{
			Code:     `"\U00000000"`,
			Expected:  "\U00000000",
		},
		{
			Code:     `‘\U00000000’`,
			Expected:  "\U00000000",
		},
		{
			Code:     `“\U00000000”`,
			Expected:  "\U00000000",
		},
		{
			Code:     `«\U00000000»`,
			Expected:  "\U00000000",
		},



		{
			Code:     `'\U00000001'`,
			Expected:  "\U00000001",
		},
		{
			Code:     `"\U00000001"`,
			Expected:  "\U00000001",
		},
		{
			Code:     `‘\U00000001’`,
			Expected:  "\U00000001",
		},
		{
			Code:     `“\U00000001”`,
			Expected:  "\U00000001",
		},
		{
			Code:     `«\U00000001»`,
			Expected:  "\U00000001",
		},



		{
			Code:     `'\U00000002'`,
			Expected:  "\U00000002",
		},
		{
			Code:     `"\U00000002"`,
			Expected:  "\U00000002",
		},
		{
			Code:     `‘\U00000002’`,
			Expected:  "\U00000002",
		},
		{
			Code:     `“\U00000002”`,
			Expected:  "\U00000002",
		},
		{
			Code:     `«\U00000002»`,
			Expected:  "\U00000002",
		},



		{
			Code:     `'\U00000003'`,
			Expected:  "\U00000003",
		},
		{
			Code:     `"\U00000003"`,
			Expected:  "\U00000003",
		},
		{
			Code:     `‘\U00000003’`,
			Expected:  "\U00000003",
		},
		{
			Code:     `“\U00000003”`,
			Expected:  "\U00000003",
		},
		{
			Code:     `«\U00000003»`,
			Expected:  "\U00000003",
		},



		{
			Code:     `'\U00000004'`,
			Expected:  "\U00000004",
		},
		{
			Code:     `"\U00000004"`,
			Expected:  "\U00000004",
		},
		{
			Code:     `‘\U00000004’`,
			Expected:  "\U00000004",
		},
		{
			Code:     `“\U00000004”`,
			Expected:  "\U00000004",
		},
		{
			Code:     `«\U00000004»`,
			Expected:  "\U00000004",
		},



		{
			Code:     `'\U00000005'`,
			Expected:  "\U00000005",
		},
		{
			Code:     `"\U00000005"`,
			Expected:  "\U00000005",
		},
		{
			Code:     `‘\U00000005’`,
			Expected:  "\U00000005",
		},
		{
			Code:     `“\U00000005”`,
			Expected:  "\U00000005",
		},
		{
			Code:     `«\U00000005»`,
			Expected:  "\U00000005",
		},



		{
			Code:     `'\U00000006'`,
			Expected:  "\U00000006",
		},
		{
			Code:     `"\U00000006"`,
			Expected:  "\U00000006",
		},
		{
			Code:     `‘\U00000006’`,
			Expected:  "\U00000006",
		},
		{
			Code:     `“\U00000006”`,
			Expected:  "\U00000006",
		},
		{
			Code:     `«\U00000006»`,
			Expected:  "\U00000006",
		},



		{
			Code:     `'\U00000007'`,
			Expected:  "\U00000007",
		},
		{
			Code:     `"\U00000007"`,
			Expected:  "\U00000007",
		},
		{
			Code:     `‘\U00000007’`,
			Expected:  "\U00000007",
		},
		{
			Code:     `“\U00000007”`,
			Expected:  "\U00000007",
		},
		{
			Code:     `«\U00000007»`,
			Expected:  "\U00000007",
		},



		{
			Code:     `'\U00000008'`,
			Expected:  "\U00000008",
		},
		{
			Code:     `"\U00000008"`,
			Expected:  "\U00000008",
		},
		{
			Code:     `‘\U00000008’`,
			Expected:  "\U00000008",
		},
		{
			Code:     `“\U00000008”`,
			Expected:  "\U00000008",
		},
		{
			Code:     `«\U00000008»`,
			Expected:  "\U00000008",
		},



		{
			Code:     `'\U00000009'`,
			Expected:  "\U00000009",
		},
		{
			Code:     `"\U00000009"`,
			Expected:  "\U00000009",
		},
		{
			Code:     `‘\U00000009’`,
			Expected:  "\U00000009",
		},
		{
			Code:     `“\U00000009”`,
			Expected:  "\U00000009",
		},
		{
			Code:     `«\U00000009»`,
			Expected:  "\U00000009",
		},



		{
			Code:     `'\U0000000A'`,
			Expected:  "\U0000000A",
		},
		{
			Code:     `"\U0000000A"`,
			Expected:  "\U0000000A",
		},
		{
			Code:     `‘\U0000000A’`,
			Expected:  "\U0000000A",
		},
		{
			Code:     `“\U0000000A”`,
			Expected:  "\U0000000A",
		},
		{
			Code:     `«\U0000000A»`,
			Expected:  "\U0000000A",
		},



		{
			Code:     `'\U0000000a'`,
			Expected:  "\U0000000a",
		},
		{
			Code:     `"\U0000000a"`,
			Expected:  "\U0000000a",
		},
		{
			Code:     `‘\U0000000a’`,
			Expected:  "\U0000000a",
		},
		{
			Code:     `“\U0000000a”`,
			Expected:  "\U0000000a",
		},
		{
			Code:     `«\U0000000a»`,
			Expected:  "\U0000000a",
		},



		{
			Code:     `'\U0000000B'`,
			Expected:  "\U0000000B",
		},
		{
			Code:     `"\U0000000B"`,
			Expected:  "\U0000000B",
		},
		{
			Code:     `‘\U0000000B’`,
			Expected:  "\U0000000B",
		},
		{
			Code:     `“\U0000000B”`,
			Expected:  "\U0000000B",
		},
		{
			Code:     `«\U0000000B»`,
			Expected:  "\U0000000B",
		},



		{
			Code:     `'\U0000000b'`,
			Expected:  "\U0000000b",
		},
		{
			Code:     `"\U0000000b"`,
			Expected:  "\U0000000b",
		},
		{
			Code:     `‘\U0000000b’`,
			Expected:  "\U0000000b",
		},
		{
			Code:     `“\U0000000b”`,
			Expected:  "\U0000000b",
		},
		{
			Code:     `«\U0000000b»`,
			Expected:  "\U0000000b",
		},



		{
			Code:     `'\U0000000C'`,
			Expected:  "\U0000000C",
		},
		{
			Code:     `"\U0000000C"`,
			Expected:  "\U0000000C",
		},
		{
			Code:     `‘\U0000000C’`,
			Expected:  "\U0000000C",
		},
		{
			Code:     `“\U0000000C”`,
			Expected:  "\U0000000C",
		},
		{
			Code:     `«\U0000000C»`,
			Expected:  "\U0000000C",
		},



		{
			Code:     `'\U0000000c'`,
			Expected:  "\U0000000c",
		},
		{
			Code:     `"\U0000000c"`,
			Expected:  "\U0000000c",
		},
		{
			Code:     `‘\U0000000c’`,
			Expected:  "\U0000000c",
		},
		{
			Code:     `“\U0000000c”`,
			Expected:  "\U0000000c",
		},
		{
			Code:     `«\U0000000c»`,
			Expected:  "\U0000000c",
		},



		{
			Code:     `'\U0000000D'`,
			Expected:  "\U0000000D",
		},
		{
			Code:     `"\U0000000D"`,
			Expected:  "\U0000000D",
		},
		{
			Code:     `‘\U0000000D’`,
			Expected:  "\U0000000D",
		},
		{
			Code:     `“\U0000000D”`,
			Expected:  "\U0000000D",
		},
		{
			Code:     `«\U0000000D»`,
			Expected:  "\U0000000D",
		},



		{
			Code:     `'\U0000000d'`,
			Expected:  "\U0000000d",
		},
		{
			Code:     `"\U0000000d"`,
			Expected:  "\U0000000d",
		},
		{
			Code:     `‘\U0000000d’`,
			Expected:  "\U0000000d",
		},
		{
			Code:     `“\U0000000d”`,
			Expected:  "\U0000000d",
		},
		{
			Code:     `«\U0000000d»`,
			Expected:  "\U0000000d",
		},



		{
			Code:     `'\U0000000E'`,
			Expected:  "\U0000000E",
		},
		{
			Code:     `"\U0000000E"`,
			Expected:  "\U0000000E",
		},
		{
			Code:     `‘\U0000000E’`,
			Expected:  "\U0000000E",
		},
		{
			Code:     `“\U0000000E”`,
			Expected:  "\U0000000E",
		},
		{
			Code:     `«\U0000000E»`,
			Expected:  "\U0000000E",
		},



		{
			Code:     `'\U0000000e'`,
			Expected:  "\U0000000e",
		},
		{
			Code:     `"\U0000000e"`,
			Expected:  "\U0000000e",
		},
		{
			Code:     `‘\U0000000e’`,
			Expected:  "\U0000000e",
		},
		{
			Code:     `“\U0000000e”`,
			Expected:  "\U0000000e",
		},
		{
			Code:     `«\U0000000e»`,
			Expected:  "\U0000000e",
		},



		{
			Code:     `'\U0000000F'`,
			Expected:  "\U0000000F",
		},
		{
			Code:     `"\U0000000F"`,
			Expected:  "\U0000000F",
		},
		{
			Code:     `‘\U0000000F’`,
			Expected:  "\U0000000F",
		},
		{
			Code:     `“\U0000000F”`,
			Expected:  "\U0000000F",
		},
		{
			Code:     `«\U0000000F»`,
			Expected:  "\U0000000F",
		},



		{
			Code:     `'\U0000000f'`,
			Expected:  "\U0000000f",
		},
		{
			Code:     `"\U0000000f"`,
			Expected:  "\U0000000f",
		},
		{
			Code:     `‘\U0000000f’`,
			Expected:  "\U0000000f",
		},
		{
			Code:     `“\U0000000f”`,
			Expected:  "\U0000000f",
		},
		{
			Code:     `«\U0000000f»`,
			Expected:  "\U0000000f",
		},



//@TODO: ###############################################



		{
			Code:     `'\U00000010'`,
			Expected:  "\U00000010",
		},
		{
			Code:     `"\U00000010"`,
			Expected:  "\U00000010",
		},
		{
			Code:     `‘\U00000010’`,
			Expected:  "\U00000010",
		},
		{
			Code:     `“\U00000010”`,
			Expected:  "\U00000010",
		},
		{
			Code:     `«\U00000010»`,
			Expected:  "\U00000010",
		},



		{
			Code:     `'\U00000020'`,
			Expected:  "\U00000020",
		},
		{
			Code:     `"\U00000020"`,
			Expected:  "\U00000020",
		},
		{
			Code:     `‘\U00000020’`,
			Expected:  "\U00000020",
		},
		{
			Code:     `“\U00000020”`,
			Expected:  "\U00000020",
		},
		{
			Code:     `«\U00000020»`,
			Expected:  "\U00000020",
		},



		{
			Code:     `'\U00000030'`,
			Expected:  "\U00000030",
		},
		{
			Code:     `"\U00000030"`,
			Expected:  "\U00000030",
		},
		{
			Code:     `‘\U00000030’`,
			Expected:  "\U00000030",
		},
		{
			Code:     `“\U00000030”`,
			Expected:  "\U00000030",
		},
		{
			Code:     `«\U00000030»`,
			Expected:  "\U00000030",
		},



		{
			Code:     `'\U00000040'`,
			Expected:  "\U00000040",
		},
		{
			Code:     `"\U00000040"`,
			Expected:  "\U00000040",
		},
		{
			Code:     `‘\U00000040’`,
			Expected:  "\U00000040",
		},
		{
			Code:     `“\U00000040”`,
			Expected:  "\U00000040",
		},
		{
			Code:     `«\U00000040»`,
			Expected:  "\U00000040",
		},



		{
			Code:     `'\U00000050'`,
			Expected:  "\U00000050",
		},
		{
			Code:     `"\U00000050"`,
			Expected:  "\U00000050",
		},
		{
			Code:     `‘\U00000050’`,
			Expected:  "\U00000050",
		},
		{
			Code:     `“\U00000050”`,
			Expected:  "\U00000050",
		},
		{
			Code:     `«\U00000050»`,
			Expected:  "\U00000050",
		},



		{
			Code:     `'\U00000060'`,
			Expected:  "\U00000060",
		},
		{
			Code:     `"\U00000060"`,
			Expected:  "\U00000060",
		},
		{
			Code:     `‘\U00000060’`,
			Expected:  "\U00000060",
		},
		{
			Code:     `“\U00000060”`,
			Expected:  "\U00000060",
		},
		{
			Code:     `«\U00000060»`,
			Expected:  "\U00000060",
		},



		{
			Code:     `'\U00000070'`,
			Expected:  "\U00000070",
		},
		{
			Code:     `"\U00000070"`,
			Expected:  "\U00000070",
		},
		{
			Code:     `‘\U00000070’`,
			Expected:  "\U00000070",
		},
		{
			Code:     `“\U000000070”`,
			Expected:  "\U000000070",
		},
		{
			Code:     `«\U000000070»`,
			Expected:  "\U000000070",
		},



		{
			Code:     `'\U00000080'`,
			Expected:  "\U00000080",
		},
		{
			Code:     `"\U00000080"`,
			Expected:  "\U00000080",
		},
		{
			Code:     `‘\U00000080’`,
			Expected:  "\U00000080",
		},
		{
			Code:     `“\U00000080”`,
			Expected:  "\U00000080",
		},
		{
			Code:     `«\U00000080»`,
			Expected:  "\U00000080",
		},



		{
			Code:     `'\U00000090'`,
			Expected:  "\U00000090",
		},
		{
			Code:     `"\U00000090"`,
			Expected:  "\U00000090",
		},
		{
			Code:     `‘\U00000090’`,
			Expected:  "\U00000090",
		},
		{
			Code:     `“\U00000090”`,
			Expected:  "\U00000090",
		},
		{
			Code:     `«\U00000090»`,
			Expected:  "\U00000090",
		},



		{
			Code:     `'\U000000A0'`,
			Expected:  "\U000000A0",
		},
		{
			Code:     `"\U000000A0"`,
			Expected:  "\U000000A0",
		},
		{
			Code:     `‘\U000000A0’`,
			Expected:  "\U000000A0",
		},
		{
			Code:     `“\U000000A0”`,
			Expected:  "\U000000A0",
		},
		{
			Code:     `«\U000000A0»`,
			Expected:  "\U000000A0",
		},



		{
			Code:     `'\U000000a0'`,
			Expected:  "\U000000a0",
		},
		{
			Code:     `"\U000000a0"`,
			Expected:  "\U000000a0",
		},
		{
			Code:     `‘\U000000a0’`,
			Expected:  "\U000000a0",
		},
		{
			Code:     `“\U000000a0”`,
			Expected:  "\U000000a0",
		},
		{
			Code:     `«\U000000a0»`,
			Expected:  "\U000000a0",
		},



		{
			Code:     `'\U000000B0'`,
			Expected:  "\U000000B0",
		},
		{
			Code:     `"\U000000B0"`,
			Expected:  "\U000000B0",
		},
		{
			Code:     `‘\U000000B0’`,
			Expected:  "\U000000B0",
		},
		{
			Code:     `“\U000000B0”`,
			Expected:  "\U000000B0",
		},
		{
			Code:     `«\U000000B0»`,
			Expected:  "\U000000B0",
		},



		{
			Code:     `'\U000000b0'`,
			Expected:  "\U000000b0",
		},
		{
			Code:     `"\U000000b0"`,
			Expected:  "\U000000b0",
		},
		{
			Code:     `‘\U000000b0’`,
			Expected:  "\U000000b0",
		},
		{
			Code:     `“\U000000b0”`,
			Expected:  "\U000000b0",
		},
		{
			Code:     `«\U000000b0»`,
			Expected:  "\U000000b0",
		},



		{
			Code:     `'\U000000C0'`,
			Expected:  "\U000000C0",
		},
		{
			Code:     `"\U000000C0"`,
			Expected:  "\U000000C0",
		},
		{
			Code:     `‘\U000000C0’`,
			Expected:  "\U000000C0",
		},
		{
			Code:     `“\U000000C0”`,
			Expected:  "\U000000C0",
		},
		{
			Code:     `«\U000000C0»`,
			Expected:  "\U000000C0",
		},



		{
			Code:     `'\U000000c0'`,
			Expected:  "\U000000c0",
		},
		{
			Code:     `"\U000000c0"`,
			Expected:  "\U000000c0",
		},
		{
			Code:     `‘\U000000c0’`,
			Expected:  "\U000000c0",
		},
		{
			Code:     `“\U000000c0”`,
			Expected:  "\U000000c0",
		},
		{
			Code:     `«\U000000c0»`,
			Expected:  "\U000000c0",
		},



		{
			Code:     `'\U000000D0'`,
			Expected:  "\U000000D0",
		},
		{
			Code:     `"\U000000D0"`,
			Expected:  "\U000000D0",
		},
		{
			Code:     `‘\U000000D0’`,
			Expected:  "\U000000D0",
		},
		{
			Code:     `“\U000000D0”`,
			Expected:  "\U000000D0",
		},
		{
			Code:     `«\U000000D0»`,
			Expected:  "\U000000D0",
		},



		{
			Code:     `'\U000000d0'`,
			Expected:  "\U000000d0",
		},
		{
			Code:     `"\U000000d0"`,
			Expected:  "\U000000d0",
		},
		{
			Code:     `‘\U000000d0’`,
			Expected:  "\U000000d0",
		},
		{
			Code:     `“\U000000d0”`,
			Expected:  "\U000000d0",
		},
		{
			Code:     `«\U000000d0»`,
			Expected:  "\U000000d0",
		},



		{
			Code:     `'\U000000E0'`,
			Expected:  "\U000000E0",
		},
		{
			Code:     `"\U000000E0"`,
			Expected:  "\U000000E0",
		},
		{
			Code:     `‘\U000000E0’`,
			Expected:  "\U000000E0",
		},
		{
			Code:     `“\U000000E0”`,
			Expected:  "\U000000E0",
		},
		{
			Code:     `«\U000000E0»`,
			Expected:  "\U000000E0",
		},



		{
			Code:     `'\U000000e0'`,
			Expected:  "\U000000e0",
		},
		{
			Code:     `"\U000000e0"`,
			Expected:  "\U000000e0",
		},
		{
			Code:     `‘\U000000e0’`,
			Expected:  "\U000000e0",
		},
		{
			Code:     `“\U000000e0”`,
			Expected:  "\U000000e0",
		},
		{
			Code:     `«\U000000e0»`,
			Expected:  "\U000000e0",
		},



		{
			Code:     `'\U000000F0'`,
			Expected:  "\U000000F0",
		},
		{
			Code:     `"\U000000F0"`,
			Expected:  "\U000000F0",
		},
		{
			Code:     `‘\U000000F0’`,
			Expected:  "\U000000F0",
		},
		{
			Code:     `“\U000000F0”`,
			Expected:  "\U000000F0",
		},
		{
			Code:     `«\U000000F0»`,
			Expected:  "\U000000F0",
		},



		{
			Code:     `'\U000000f0'`,
			Expected:  "\U000000f0",
		},
		{
			Code:     `"\U000000f0"`,
			Expected:  "\U000000f0",
		},
		{
			Code:     `‘\U000000f0’`,
			Expected:  "\U000000f0",
		},
		{
			Code:     `“\U000000f0”`,
			Expected:  "\U000000f0",
		},
		{
			Code:     `«\U000000f0»`,
			Expected:  "\U000000f0",
		},



//@TODO: ###############################################



		{
			Code:     `'\U01000000'`,
			Expected:  "\U01000000",
		},
		{
			Code:     `"\U023bCdEf"`,
			Expected:  "\U023bCdEf",
		},
		{
			Code:     `‘\U023bCdEf’`,
			Expected:  "\U023bCdEf",
		},
		{
			Code:     `“\U023bCdEf”`,
			Expected:  "\U023bCdEf",
		},
		{
			Code:     `«\U023bCdDf»`,
			Expected:  "\U023bCdDf",
		},



//@TODO: ###############################################



		{
			Code:     `'\U109aBcDE'`,
			Expected:  "\U109aBcDE",
		},
		{
			Code:     `"\U109aBcDE"`,
			Expected:  "\U109aBcDE",
		},
		{
			Code:     `‘\U109aBcDE’`,
			Expected:  "\U109aBcDE",
		},
		{
			Code:     `“\U109aBcDE”`,
			Expected:  "\U109aBcDE",
		},
		{
			Code:     `«\U109aBcDE»`,
			Expected:  "\U109aBcDE",
		},



//@TODO: ###############################################



		{
			Code:     `'\'\''`,
			Expected: `''`,
		},
		{
			Code:     `"\"\""`,
			Expected: `""`,
		},



		{
			Code:     `'\'\'\''`,
			Expected: `'''`,
		},
		{
			Code:     `"\"\"\""`,
			Expected: `"""`,
		},



		{
			Code:     `'\'\'\'\''`,
			Expected: `''''`,
		},
		{
			Code:     `"\"\"\"\""`,
			Expected: `""""`,
		},



		{
			Code:     `'\'\'\'\'\''`,
			Expected: `'''''`,
		},
		{
			Code:     `"\"\"\"\"\""`,
			Expected: `"""""`,
		},
//@TODO: ###############################################



		{
			Code:     `'\a\b\f\n\r\t\v'`,
			Expected:  "\a\b\f\n\r\t\v",
		},
		{
			Code:     `"\a\b\f\n\r\t\v"`,
			Expected:  "\a\b\f\n\r\t\v",
		},
		{
			Code:     `‘\a\b\f\n\r\t\v’`,
			Expected:  "\a\b\f\n\r\t\v",
		},
		{
			Code:     `“\a\b\f\n\r\t\v”`,
			Expected:  "\a\b\f\n\r\t\v",
		},
		{
			Code:     `«\a\b\f\n\r\t\v»`,
			Expected:  "\a\b\f\n\r\t\v",
		},



		{
			Code:     `'\aC\bh\fe\nr\rR\ty\v!'`,
			Expected:  "\aC\bh\fe\nr\rR\ty\v!",
		},
		{
			Code:     `"\aC\bh\fe\nr\rR\ty\v!"`,
			Expected:  "\aC\bh\fe\nr\rR\ty\v!",
		},
		{
			Code:     `‘\aC\bh\fe\nr\rR\ty\v!’`,
			Expected:  "\aC\bh\fe\nr\rR\ty\v!",
		},
		{
			Code:     `“\aC\bh\fe\nr\rR\ty\v!”`,
			Expected:  "\aC\bh\fe\nr\rR\ty\v!",
		},
		{
			Code:     `«\aC\bh\fe\nr\rR\ty\v!»`,
			Expected:  "\aC\bh\fe\nr\rR\ty\v!",
		},



		{
			Code:     `'"apple" ‘banana’ “cherry” «grape»'`,
			Expected:  `"apple" ‘banana’ “cherry” «grape»`,
		},
		{
			Code:     `"‘apple’ “banana” «cherry» 'grape'"`,
			Expected:  `‘apple’ “banana” «cherry» 'grape'`,
		},
		{
			Code:     `‘“apple” «banana» 'cherry' "grape"’`,
			Expected:  `“apple” «banana» 'cherry' "grape"`,
		},
		{
			Code:     `“«apple» 'banana' "cherry" ‘grape’”`,
			Expected:  `«apple» 'banana' "cherry" ‘grape’`,
		},
		{
			Code:     `«'apple' "banana" ‘cherry’ “grape”»`,
			Expected:  `'apple' "banana" ‘cherry’ “grape”`,
		},
	}


	for testNumber, test := range tests {

		actualCompiled, err := Compile( strings.NewReader(test.Code) )

		if nil != err {
			t.Errorf("For test #%d, did not expected an error, but actually got one: %v", testNumber, err)
			continue
		}

		if expected, actual := test.Expected, actualCompiled.String(); expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, expected, actual)
			continue
		}

		if expected, actual := []rune(test.Expected), actualCompiled.Runes(); len(expected) != len(expected) {
			t.Errorf("For test #%d, expected %d runes, but actually got %d runes.", testNumber, expected, actual)
			continue
		} else {
			for runeNumber, expectedDatum := range expected {
				actualDatum := actual[runeNumber]

				if expectedDatum != actualDatum {
					t.Errorf("For test #%d an rune #%d, expected rune to be %q, but actually was %q.", testNumber, runeNumber, expectedDatum, actualDatum)
					continue
				}
			}
		}

		if expected, actual := []byte(test.Expected), actualCompiled.Bytes(); len(expected) != len(expected) {
			t.Errorf("For test #%d, expected %d bytes, but actually got %d bytes.", testNumber, expected, actual)
			continue
		} else {
			for byteNumber, expectedDatum := range expected {
				actualDatum := actual[byteNumber]

				if expectedDatum != actualDatum {
					t.Errorf("For test #%d an byte #%d, expected byte to be %d, but actually was %d.", testNumber, byteNumber, expectedDatum, actualDatum)
					continue
				}
			}
		}
	}
}


func TestCompileSyntaxError(t *testing.T) {

	tests := []struct{
		Code         string
		ExpectedCode string
	}{
		{
			Code:         `'`,
			ExpectedCode: `'`,
		},
		{
			Code:         `"`,
			ExpectedCode: `"`,
		},
		{
			Code:         `‘`,
			ExpectedCode: `‘`,
		},
		{
			Code:         `’`,
			ExpectedCode: ``,
		},
		{
			Code:         `“`,
			ExpectedCode: `“`,
		},
		{
			Code:         `”`,
			ExpectedCode: ``,
		},
		{
			Code:         `«`,
			ExpectedCode: `«`,
		},
		{
			Code:         `»`,
			ExpectedCode: ``,
		},



		{
			Code:         `'''`,
			ExpectedCode: `'''`,
		},
		{
			Code:         `"""`,
			ExpectedCode: `"""`,
		},



		{
			Code:         `'''''`,
			ExpectedCode: `'''''`,
		},
		{
			Code:         `"""""`,
			ExpectedCode: `"""""`,
		},



		{
			Code:         `'''''''`,
			ExpectedCode: `'''''''`,
		},
		{
			Code:         `"""""""`,
			ExpectedCode: `"""""""`,
		},



		{
			Code:         `'''''''''`,
			ExpectedCode: `'''''''''`,
		},
		{
			Code:         `"""""""""`,
			ExpectedCode: `"""""""""`,
		},



		{
			Code:         `'\'`,
			ExpectedCode: `'\'`,
		},
		{
			Code:         `"\"`,
			ExpectedCode: `"\"`,
		},
		{
			Code:         `‘\’`,
			ExpectedCode: `‘\’`,
		},
		{
			Code:         `“\”`,
			ExpectedCode: `“\”`,
		},
		{
			Code:         `«\»`,
			ExpectedCode: `«\»`,
		},



		{
			Code:         `'\\\'`,
			ExpectedCode: `'\\\'`,
		},
		{
			Code:         `"\\\"`,
			ExpectedCode: `"\\\"`,
		},
		{
			Code:         `‘\\\’`,
			ExpectedCode: `‘\\\’`,
		},
		{
			Code:         `“\\\”`,
			ExpectedCode: `“\\\”`,
		},
		{
			Code:         `«\\\»`,
			ExpectedCode: `«\\\»`,
		},



		{
			Code:         `apple banana cherry`,
			ExpectedCode: ``,
		},



		{
			Code:         `'apple banana cherry`,
			ExpectedCode: `'apple banana cherry`,
		},
		{
			Code:         `"apple banana cherry`,
			ExpectedCode: `"apple banana cherry`,
		},
		{
			Code:         `‘apple banana cherry`,
			ExpectedCode: `‘apple banana cherry`,
		},
		{
			Code:         `’apple banana cherry`,
			ExpectedCode: ``,
		},
		{
			Code:         `“apple banana cherry`,
			ExpectedCode: `“apple banana cherry`,
		},
		{
			Code:         `”apple banana cherry`,
			ExpectedCode: ``,
		},
		{
			Code:         `«apple banana cherry`,
			ExpectedCode: `«apple banana cherry`,
		},
		{
			Code:         `»apple banana cherry`,
			ExpectedCode: ``,
		},



		{
			Code:         `apple banana cherry'`,
			ExpectedCode: ``,
		},
		{
			Code:         `apple banana cherry"`,
			ExpectedCode: ``,
		},
		{
			Code:         `apple banana cherry‘`,
			ExpectedCode: ``,
		},
		{
			Code:         `apple banana cherry’`,
			ExpectedCode: ``,
		},
		{
			Code:         `apple banana cherry“`,
			ExpectedCode: ``,
		},
		{
			Code:         `apple banana cherry”`,
			ExpectedCode: ``,
		},
		{
			Code:         `apple banana cherry«`,
			ExpectedCode: ``,
		},
		{
			Code:         `apple banana cherry»`,
			ExpectedCode: ``,
		},



		{
			Code:         `'\A'`,
			ExpectedCode: `'\A`,
		},
		{
			Code:         `"\A"`,
			ExpectedCode: `"\A`,
		},
		{
			Code:         `‘\A’`,
			ExpectedCode: `‘\A`,
		},
		{
			Code:         `“\A”`,
			ExpectedCode: `“\A`,
		},
		{
			Code:         `«\A»`,
			ExpectedCode: `«\A`,
		},



		{
			Code:         `'\B'`,
			ExpectedCode: `'\B`,
		},
		{
			Code:         `"\B"`,
			ExpectedCode: `"\B`,
		},
		{
			Code:         `‘\B’`,
			ExpectedCode: `‘\B`,
		},
		{
			Code:         `“\B”`,
			ExpectedCode: `“\B`,
		},
		{
			Code:         `«\B»`,
			ExpectedCode: `«\B`,
		},



		{
			Code:         `'\C'`,
			ExpectedCode: `'\C`,
		},
		{
			Code:         `"\C"`,
			ExpectedCode: `"\C`,
		},
		{
			Code:         `‘\C’`,
			ExpectedCode: `‘\C`,
		},
		{
			Code:         `“\C”`,
			ExpectedCode: `“\C`,
		},
		{
			Code:         `«\C»`,
			ExpectedCode: `«\C`,
		},



		{
			Code:         `'\D'`,
			ExpectedCode: `'\D`,
		},
		{
			Code:         `"\D"`,
			ExpectedCode: `"\D`,
		},
		{
			Code:         `‘\D’`,
			ExpectedCode: `‘\D`,
		},
		{
			Code:         `“\D”`,
			ExpectedCode: `“\D`,
		},
		{
			Code:         `«\D»`,
			ExpectedCode: `«\D`,
		},



		{
			Code:         `'\E'`,
			ExpectedCode: `'\E`,
		},
		{
			Code:         `"\E"`,
			ExpectedCode: `"\E`,
		},
		{
			Code:         `‘\E’`,
			ExpectedCode: `‘\E`,
		},
		{
			Code:         `“\E”`,
			ExpectedCode: `“\E`,
		},
		{
			Code:         `«\E»`,
			ExpectedCode: `«\E`,
		},



		{
			Code:         `'\F'`,
			ExpectedCode: `'\F`,
		},
		{
			Code:         `"\F"`,
			ExpectedCode: `"\F`,
		},
		{
			Code:         `‘\F’`,
			ExpectedCode: `‘\F`,
		},
		{
			Code:         `“\F”`,
			ExpectedCode: `“\F`,
		},
		{
			Code:         `«\F»`,
			ExpectedCode: `«\F`,
		},



		{
			Code:         `'\G'`,
			ExpectedCode: `'\G`,
		},
		{
			Code:         `"\G"`,
			ExpectedCode: `"\G`,
		},
		{
			Code:         `‘\G’`,
			ExpectedCode: `‘\G`,
		},
		{
			Code:         `“\G”`,
			ExpectedCode: `“\G`,
		},
		{
			Code:         `«\G»`,
			ExpectedCode: `«\G`,
		},



		{
			Code:         `'\H'`,
			ExpectedCode: `'\H`,
		},
		{
			Code:         `"\H"`,
			ExpectedCode: `"\H`,
		},
		{
			Code:         `‘\H’`,
			ExpectedCode: `‘\H`,
		},
		{
			Code:         `“\H”`,
			ExpectedCode: `“\H`,
		},
		{
			Code:         `«\H»`,
			ExpectedCode: `«\H`,
		},



		{
			Code:         `'\I'`,
			ExpectedCode: `'\I`,
		},
		{
			Code:         `"\I"`,
			ExpectedCode: `"\I`,
		},
		{
			Code:         `‘\I’`,
			ExpectedCode: `‘\I`,
		},
		{
			Code:         `“\I”`,
			ExpectedCode: `“\I`,
		},
		{
			Code:         `«\I»`,
			ExpectedCode: `«\I`,
		},



		{
			Code:         `'\J'`,
			ExpectedCode: `'\J`,
		},
		{
			Code:         `"\J"`,
			ExpectedCode: `"\J`,
		},
		{
			Code:         `‘\J’`,
			ExpectedCode: `‘\J`,
		},
		{
			Code:         `“\J”`,
			ExpectedCode: `“\J`,
		},
		{
			Code:         `«\J»`,
			ExpectedCode: `«\J`,
		},



		{
			Code:         `'\K'`,
			ExpectedCode: `'\K`,
		},
		{
			Code:         `"\K"`,
			ExpectedCode: `"\K`,
		},
		{
			Code:         `‘\K’`,
			ExpectedCode: `‘\K`,
		},
		{
			Code:         `“\K”`,
			ExpectedCode: `“\K`,
		},
		{
			Code:         `«\K»`,
			ExpectedCode: `«\K`,
		},



		{
			Code:         `'\L'`,
			ExpectedCode: `'\L`,
		},
		{
			Code:         `"\L"`,
			ExpectedCode: `"\L`,
		},
		{
			Code:         `‘\L’`,
			ExpectedCode: `‘\L`,
		},
		{
			Code:         `“\L”`,
			ExpectedCode: `“\L`,
		},
		{
			Code:         `«\L»`,
			ExpectedCode: `«\L`,
		},



		{
			Code:         `'\M'`,
			ExpectedCode: `'\M`,
		},
		{
			Code:         `"\M"`,
			ExpectedCode: `"\M`,
		},
		{
			Code:         `‘\M’`,
			ExpectedCode: `‘\M`,
		},
		{
			Code:         `“\M”`,
			ExpectedCode: `“\M`,
		},
		{
			Code:         `«\M»`,
			ExpectedCode: `«\M`,
		},



		{
			Code:         `'\N'`,
			ExpectedCode: `'\N`,
		},
		{
			Code:         `"\N"`,
			ExpectedCode: `"\N`,
		},
		{
			Code:         `‘\N’`,
			ExpectedCode: `‘\N`,
		},
		{
			Code:         `“\N”`,
			ExpectedCode: `“\N`,
		},
		{
			Code:         `«\N»`,
			ExpectedCode: `«\N`,
		},



		{
			Code:         `'\O'`,
			ExpectedCode: `'\O`,
		},
		{
			Code:         `"\O"`,
			ExpectedCode: `"\O`,
		},
		{
			Code:         `‘\O’`,
			ExpectedCode: `‘\O`,
		},
		{
			Code:         `“\O”`,
			ExpectedCode: `“\O`,
		},
		{
			Code:         `«\O»`,
			ExpectedCode: `«\O`,
		},



		{
			Code:         `'\P'`,
			ExpectedCode: `'\P`,
		},
		{
			Code:         `"\P"`,
			ExpectedCode: `"\P`,
		},
		{
			Code:         `‘\P’`,
			ExpectedCode: `‘\P`,
		},
		{
			Code:         `“\P”`,
			ExpectedCode: `“\P`,
		},
		{
			Code:         `«\P»`,
			ExpectedCode: `«\P`,
		},



		{
			Code:         `'\Q'`,
			ExpectedCode: `'\Q`,
		},
		{
			Code:         `"\Q"`,
			ExpectedCode: `"\Q`,
		},
		{
			Code:         `‘\Q’`,
			ExpectedCode: `‘\Q`,
		},
		{
			Code:         `“\Q”`,
			ExpectedCode: `“\Q`,
		},
		{
			Code:         `«\Q»`,
			ExpectedCode: `«\Q`,
		},



		{
			Code:         `'\R'`,
			ExpectedCode: `'\R`,
		},
		{
			Code:         `"\R"`,
			ExpectedCode: `"\R`,
		},
		{
			Code:         `‘\R’`,
			ExpectedCode: `‘\R`,
		},
		{
			Code:         `“\R”`,
			ExpectedCode: `“\R`,
		},
		{
			Code:         `«\R»`,
			ExpectedCode: `«\R`,
		},



		{
			Code:         `'\S'`,
			ExpectedCode: `'\S`,
		},
		{
			Code:         `"\S"`,
			ExpectedCode: `"\S`,
		},
		{
			Code:         `‘\S’`,
			ExpectedCode: `‘\S`,
		},
		{
			Code:         `“\S”`,
			ExpectedCode: `“\S`,
		},
		{
			Code:         `«\S»`,
			ExpectedCode: `«\S`,
		},



		{
			Code:         `'\T'`,
			ExpectedCode: `'\T`,
		},
		{
			Code:         `"\T"`,
			ExpectedCode: `"\T`,
		},
		{
			Code:         `‘\T’`,
			ExpectedCode: `‘\T`,
		},
		{
			Code:         `“\T”`,
			ExpectedCode: `“\T`,
		},
		{
			Code:         `«\T»`,
			ExpectedCode: `«\T`,
		},



		// \U IS VALID! Ex: \U1234ABCD



		{
			Code:         `'\V'`,
			ExpectedCode: `'\V`,
		},
		{
			Code:         `"\V"`,
			ExpectedCode: `"\V`,
		},
		{
			Code:         `‘\V’`,
			ExpectedCode: `‘\V`,
		},
		{
			Code:         `“\V”`,
			ExpectedCode: `“\V`,
		},
		{
			Code:         `«\V»`,
			ExpectedCode: `«\V`,
		},



		{
			Code:         `'\W'`,
			ExpectedCode: `'\W`,
		},
		{
			Code:         `"\W"`,
			ExpectedCode: `"\W`,
		},
		{
			Code:         `‘\W’`,
			ExpectedCode: `‘\W`,
		},
		{
			Code:         `“\W”`,
			ExpectedCode: `“\W`,
		},
		{
			Code:         `«\W»`,
			ExpectedCode: `«\W`,
		},



		{
			Code:         `'\X'`,
			ExpectedCode: `'\X`,
		},
		{
			Code:         `"\X"`,
			ExpectedCode: `"\X`,
		},
		{
			Code:         `‘\X’`,
			ExpectedCode: `‘\X`,
		},
		{
			Code:         `“\X”`,
			ExpectedCode: `“\X`,
		},
		{
			Code:         `«\X»`,
			ExpectedCode: `«\X`,
		},



		{
			Code:         `'\Y'`,
			ExpectedCode: `'\Y`,
		},
		{
			Code:         `"\Y"`,
			ExpectedCode: `"\Y`,
		},
		{
			Code:         `‘\Y’`,
			ExpectedCode: `‘\Y`,
		},
		{
			Code:         `“\Y”`,
			ExpectedCode: `“\Y`,
		},
		{
			Code:         `«\Y»`,
			ExpectedCode: `«\Y`,
		},



		{
			Code:         `'\Z'`,
			ExpectedCode: `'\Z`,
		},
		{
			Code:         `"\Z"`,
			ExpectedCode: `"\Z`,
		},
		{
			Code:         `‘\Z’`,
			ExpectedCode: `‘\Z`,
		},
		{
			Code:         `“\Z”`,
			ExpectedCode: `“\Z`,
		},
		{
			Code:         `«\Z»`,
			ExpectedCode: `«\Z`,
		},



		// \a IS VALID!



		// \b IS VALID!



		{
			Code:         `'\c'`,
			ExpectedCode: `'\c`,
		},
		{
			Code:         `"\c"`,
			ExpectedCode: `"\c`,
		},
		{
			Code:         `‘\c’`,
			ExpectedCode: `‘\c`,
		},
		{
			Code:         `“\c”`,
			ExpectedCode: `“\c`,
		},
		{
			Code:         `«\c»`,
			ExpectedCode: `«\c`,
		},



		{
			Code:         `'\d'`,
			ExpectedCode: `'\d`,
		},
		{
			Code:         `"\d"`,
			ExpectedCode: `"\d`,
		},
		{
			Code:         `‘\d’`,
			ExpectedCode: `‘\d`,
		},
		{
			Code:         `“\d”`,
			ExpectedCode: `“\d`,
		},
		{
			Code:         `«\d»`,
			ExpectedCode: `«\d`,
		},



		// \e IS VALID!



		// \f IS VALID!



		{
			Code:         `'\g'`,
			ExpectedCode: `'\g`,
		},
		{
			Code:         `"\g"`,
			ExpectedCode: `"\g`,
		},
		{
			Code:         `‘\g’`,
			ExpectedCode: `‘\g`,
		},
		{
			Code:         `“\g”`,
			ExpectedCode: `“\g`,
		},
		{
			Code:         `«\g»`,
			ExpectedCode: `«\g`,
		},



		{
			Code:         `'\h'`,
			ExpectedCode: `'\h`,
		},
		{
			Code:         `"\h"`,
			ExpectedCode: `"\h`,
		},
		{
			Code:         `‘\h’`,
			ExpectedCode: `‘\h`,
		},
		{
			Code:         `“\h”`,
			ExpectedCode: `“\h`,
		},
		{
			Code:         `«\h»`,
			ExpectedCode: `«\h`,
		},



		{
			Code:         `'\i'`,
			ExpectedCode: `'\i`,
		},
		{
			Code:         `"\i"`,
			ExpectedCode: `"\i`,
		},
		{
			Code:         `‘\i’`,
			ExpectedCode: `‘\i`,
		},
		{
			Code:         `“\i”`,
			ExpectedCode: `“\i`,
		},
		{
			Code:         `«\i»`,
			ExpectedCode: `«\i`,
		},



		{
			Code:         `'\j'`,
			ExpectedCode: `'\j`,
		},
		{
			Code:         `"\j"`,
			ExpectedCode: `"\j`,
		},
		{
			Code:         `‘\j’`,
			ExpectedCode: `‘\j`,
		},
		{
			Code:         `“\j”`,
			ExpectedCode: `“\j`,
		},
		{
			Code:         `«\j»`,
			ExpectedCode: `«\j`,
		},



		{
			Code:         `'\k'`,
			ExpectedCode: `'\k`,
		},
		{
			Code:         `"\k"`,
			ExpectedCode: `"\k`,
		},
		{
			Code:         `‘\k’`,
			ExpectedCode: `‘\k`,
		},
		{
			Code:         `“\k”`,
			ExpectedCode: `“\k`,
		},
		{
			Code:         `«\k»`,
			ExpectedCode: `«\k`,
		},



		{
			Code:         `'\l'`,
			ExpectedCode: `'\l`,
		},
		{
			Code:         `"\l"`,
			ExpectedCode: `"\l`,
		},
		{
			Code:         `‘\l’`,
			ExpectedCode: `‘\l`,
		},
		{
			Code:         `“\l”`,
			ExpectedCode: `“\l`,
		},
		{
			Code:         `«\l»`,
			ExpectedCode: `«\l`,
		},



		{
			Code:         `'\m'`,
			ExpectedCode: `'\m`,
		},
		{
			Code:         `"\m"`,
			ExpectedCode: `"\m`,
		},
		{
			Code:         `‘\m’`,
			ExpectedCode: `‘\m`,
		},
		{
			Code:         `“\m”`,
			ExpectedCode: `“\m`,
		},
		{
			Code:         `«\m»`,
			ExpectedCode: `«\m`,
		},



		// \n IS VALID!



		{
			Code:         `'\o'`,
			ExpectedCode: `'\o`,
		},
		{
			Code:         `"\o"`,
			ExpectedCode: `"\o`,
		},
		{
			Code:         `‘\o’`,
			ExpectedCode: `‘\o`,
		},
		{
			Code:         `“\o”`,
			ExpectedCode: `“\o`,
		},
		{
			Code:         `«\o»`,
			ExpectedCode: `«\o`,
		},



		{
			Code:         `'\p'`,
			ExpectedCode: `'\p`,
		},
		{
			Code:         `"\p"`,
			ExpectedCode: `"\p`,
		},
		{
			Code:         `‘\p’`,
			ExpectedCode: `‘\p`,
		},
		{
			Code:         `“\p”`,
			ExpectedCode: `“\p`,
		},
		{
			Code:         `«\p»`,
			ExpectedCode: `«\p`,
		},



		{
			Code:         `'\q'`,
			ExpectedCode: `'\q`,
		},
		{
			Code:         `"\q"`,
			ExpectedCode: `"\q`,
		},
		{
			Code:         `‘\q’`,
			ExpectedCode: `‘\q`,
		},
		{
			Code:         `“\q”`,
			ExpectedCode: `“\q`,
		},
		{
			Code:         `«\q»`,
			ExpectedCode: `«\q`,
		},



		// \r IS VALID!



		{
			Code:         `'\s'`,
			ExpectedCode: `'\s`,
		},
		{
			Code:         `"\s"`,
			ExpectedCode: `"\s`,
		},
		{
			Code:         `‘\s’`,
			ExpectedCode: `‘\s`,
		},
		{
			Code:         `“\s”`,
			ExpectedCode: `“\s`,
		},
		{
			Code:         `«\s»`,
			ExpectedCode: `«\s`,
		},



		// \t IS VALID!



		// \u IS VALID! Ex: \u1234



		// \v IS VALID!



		{
			Code:         `'\w'`,
			ExpectedCode: `'\w`,
		},
		{
			Code:         `"\w"`,
			ExpectedCode: `"\w`,
		},
		{
			Code:         `‘\w’`,
			ExpectedCode: `‘\w`,
		},
		{
			Code:         `“\w”`,
			ExpectedCode: `“\w`,
		},
		{
			Code:         `«\w»`,
			ExpectedCode: `«\w`,
		},



		// \x IS VALID! Ex: \x1B



		{
			Code:         `'\y'`,
			ExpectedCode: `'\y`,
		},
		{
			Code:         `"\y"`,
			ExpectedCode: `"\y`,
		},
		{
			Code:         `‘\y’`,
			ExpectedCode: `‘\y`,
		},
		{
			Code:         `“\y”`,
			ExpectedCode: `“\y`,
		},
		{
			Code:         `«\y»`,
			ExpectedCode: `«\y`,
		},



		{
			Code:         `'\z'`,
			ExpectedCode: `'\z`,
		},
		{
			Code:         `"\z"`,
			ExpectedCode: `"\z`,
		},
		{
			Code:         `‘\z’`,
			ExpectedCode: `‘\z`,
		},
		{
			Code:         `“\z”`,
			ExpectedCode: `“\z`,
		},
		{
			Code:         `«\z»`,
			ExpectedCode: `«\z`,
		},



		{
			Code:         `'\xYZ'`,
			ExpectedCode: `'\xYZ`,
		},
		{
			Code:         `"\xYZ"`,
			ExpectedCode: `"\xYZ`,
		},
		{
			Code:         `‘\xYZ’`,
			ExpectedCode: `‘\xYZ`,
		},
		{
			Code:         `“\xYZ”`,
			ExpectedCode: `“\xYZ`,
		},
		{
			Code:         `«\xYZ»`,
			ExpectedCode: `«\xYZ`,
		},



		{
			Code:         `'\xYz'`,
			ExpectedCode: `'\xYz`,
		},
		{
			Code:         `"\xYz"`,
			ExpectedCode: `"\xYz`,
		},
		{
			Code:         `‘\xYz’`,
			ExpectedCode: `‘\xYz`,
		},
		{
			Code:         `“\xYz”`,
			ExpectedCode: `“\xYz`,
		},
		{
			Code:         `«\xYz»`,
			ExpectedCode: `«\xYz`,
		},



		{
			Code:         `'\xyZ'`,
			ExpectedCode: `'\xyZ`,
		},
		{
			Code:         `"\xyZ"`,
			ExpectedCode: `"\xyZ`,
		},
		{
			Code:         `‘\xyZ’`,
			ExpectedCode: `‘\xyZ`,
		},
		{
			Code:         `“\xyZ”`,
			ExpectedCode: `“\xyZ`,
		},
		{
			Code:         `«\xyZ»`,
			ExpectedCode: `«\xyZ`,
		},



		{
			Code:         `'\xyz'`,
			ExpectedCode: `'\xyz`,
		},
		{
			Code:         `"\xyz"`,
			ExpectedCode: `"\xyz`,
		},
		{
			Code:         `‘\xyz’`,
			ExpectedCode: `‘\xyz`,
		},
		{
			Code:         `“\xyz”`,
			ExpectedCode: `“\xyz`,
		},
		{
			Code:         `«\xyz»`,
			ExpectedCode: `«\xyz`,
		},



		{
			Code:         `'\uVWXY'`,
			ExpectedCode: `'\uVWXY`,
		},
		{
			Code:         `"\uVWXY"`,
			ExpectedCode: `"\uVWXY`,
		},
		{
			Code:         `‘\uVWXY’`,
			ExpectedCode: `‘\uVWXY`,
		},
		{
			Code:         `“\uVWXY”`,
			ExpectedCode: `“\uVWXY`,
		},
		{
			Code:         `«\uVWXY»`,
			ExpectedCode: `«\uVWXY`,
		},




		{
			Code:         `'\uvwxy'`,
			ExpectedCode: `'\uvwxy`,
		},
		{
			Code:         `"\uvwxy"`,
			ExpectedCode: `"\uvwxy`,
		},
		{
			Code:         `‘\uvwxy’`,
			ExpectedCode: `‘\uvwxy`,
		},
		{
			Code:         `“\uvwxy”`,
			ExpectedCode: `“\uvwxy`,
		},
		{
			Code:         `«\uvwxy»`,
			ExpectedCode: `«\uvwxy`,
		},




		{
			Code:         `'\uDEFG'`,
			ExpectedCode: `'\uDEFG`,
		},
		{
			Code:         `"\uDEFG"`,
			ExpectedCode: `"\uDEFG`,
		},
		{
			Code:         `‘\uDEFG’`,
			ExpectedCode: `‘\uDEFG`,
		},
		{
			Code:         `“\uDEFG”`,
			ExpectedCode: `“\uDEFG`,
		},
		{
			Code:         `«\uDEFG»`,
			ExpectedCode: `«\uDEFG`,
		},



		{
			Code:         `'\udefg'`,
			ExpectedCode: `'\udefg`,
		},
		{
			Code:         `"\udefg"`,
			ExpectedCode: `"\udefg`,
		},
		{
			Code:         `‘\udefg’`,
			ExpectedCode: `‘\udefg`,
		},
		{
			Code:         `“\udefg”`,
			ExpectedCode: `“\udefg`,
		},
		{
			Code:         `«\udefg»`,
			ExpectedCode: `«\udefg`,
		},



		{
			Code:         `'\u123'`,
			ExpectedCode: `'\u123'`,
		},
		{
			Code:         `"\u123"`,
			ExpectedCode: `"\u123"`,
		},
		{
			Code:         `‘\u123’`,
			ExpectedCode: `‘\u123’`,
		},
		{
			Code:         `“\u123”`,
			ExpectedCode: `“\u123”`,
		},
		{
			Code:         `«\u123»`,
			ExpectedCode: `«\u123»`,
		},



		{
			Code:         `'\u12'`,
			ExpectedCode: `'\u12'`,
		},
		{
			Code:         `"\u12"`,
			ExpectedCode: `"\u12"`,
		},
		{
			Code:         `‘\u12’`,
			ExpectedCode: `‘\u12’`,
		},
		{
			Code:         `“\u12”`,
			ExpectedCode: `“\u12”`,
		},
		{
			Code:         `«\u12»`,
			ExpectedCode: `«\u12»`,
		},



		{
			Code:         `'\u1'`,
			ExpectedCode: `'\u1'`,
		},
		{
			Code:         `"\u1"`,
			ExpectedCode: `"\u1"`,
		},
		{
			Code:         `‘\u1’`,
			ExpectedCode: `‘\u1’`,
		},
		{
			Code:         `“\u1”`,
			ExpectedCode: `“\u1”`,
		},
		{
			Code:         `«\u1»`,
			ExpectedCode: `«\u1»`,
		},



		{
			Code:         `'\u'`,
			ExpectedCode: `'\u'`,
		},
		{
			Code:         `"\u"`,
			ExpectedCode: `"\u"`,
		},
		{
			Code:         `‘\u’`,
			ExpectedCode: `‘\u’`,
		},
		{
			Code:         `“\u”`,
			ExpectedCode: `“\u”`,
		},
		{
			Code:         `«\u»`,
			ExpectedCode: `«\u»`,
		},
	}


	for testNumber, test := range tests {

		compiled, err := Compile( strings.NewReader(test.Code) )
		if nil == err {
			t.Errorf("For test #%d, expected an error, but did not actually get one. Code: %q.", testNumber, compiled.String())
			continue
		}
		if complainer, ok := err.(SyntaxErrorComplainer); !ok {
			t.Errorf("For test #%d, expected an error to be of type SyntaxErrorComplainer, but actually was: %T.", testNumber, err)
			continue
		} else if expected, actual := test.ExpectedCode, complainer.Code(); expected != actual {
			t.Errorf("For test #%d, expected code in error to be %q, but actually was: %q.", testNumber, expected, actual)
			continue
		}
	}
}
