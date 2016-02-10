package strlit


import (
	"fmt"
	"io"
)


func Compile(runeScanner io.RuneScanner) (parcelWithResultOfCompilingStringLiteral Parcel, parcelWithPrecompiledStringLiteral Parcel, err error) {

	code := newParcel()


	var r rune


	// Make sure there is at least one rune.
	r, _, err = runeScanner.ReadRune()
	if io.EOF == err {
		runeScanner.UnreadRune()
		return nil, nil, newSyntaxErrorComplainer("", "Hit the EOF before getting any data.")
	}
	if nil != err {
		return nil, nil, err
	}


	// Make sure the first rune is a valid begin quote rune.
	var endRune rune
	switch r {
	case '\'':
		endRune = r
	case '"':
		endRune = r
	case '‘':
		endRune = '’'
	case '“':
		endRune = '”'
	case '‹':
		endRune = '›'
	case '«':
		endRune = '»'
	default:
		runeScanner.UnreadRune()
		return nil, nil, newSyntaxErrorComplainer(code.String(), "Not a string literal. Beginning rune for string literal is not valid.")
	}
	code.buffer.WriteRune(r)


	// Make sure there are more runes in the string literal code (so that we can next check it ENDS with a quote (')).
	if _, _, err := runeScanner.ReadRune(); nil != err {
		return nil, nil, newSyntaxErrorComplainer(code.String(), "Missing ending quote (') at the end of the string literal.")
	}
	runeScanner.UnreadRune()


	compiled := newParcel()
	for {
		r, _, err = runeScanner.ReadRune()
		if io.EOF == err {
			return nil, nil, newSyntaxErrorComplainer(code.String(), "Hit the EOF before being able to get to end of string literal.")
		}
		if nil != err {
			return nil, nil, err
		}
		code.buffer.WriteRune(r)

		if endRune == r {

			r2, _, err := runeScanner.ReadRune()
			if io.EOF == err {
				return compiled, code, nil
			}
			if nil != err {
				return nil, nil, newSyntaxErrorComplainer(code.String(), "Something bad happened.")
			}

			if endRune != r2 {
				runeScanner.UnreadRune()
				return compiled, code, nil
			} else {
				code.buffer.WriteRune(r2)
			}
		} else if '\\' == r {

			r2, _, err := runeScanner.ReadRune()
			if nil != err {
				return nil, nil, newSyntaxErrorComplainer(code.String(), "Backslash (\\), by itself, in the middle of a string literal.")
			}
			code.buffer.WriteRune(r2)

			switch r2 {
			case '0':
				r = '\x00'
//@TODO: Support octal escape notation.
			case 'U':
				r3, _, err := runeScanner.ReadRune()
				if nil != err {
					msg := "Problem after reading universal character escape sequence “\\x”."
					return nil, nil, newSyntaxErrorComplainer(code.String(), msg)
				}
				code.buffer.WriteRune(r3)

				r4, _, err := runeScanner.ReadRune()
				if nil != err {
					msg := fmt.Sprintf("Problem after reading universal character escape sequence “\\x%s”.", string(r3))
					return nil, nil, newSyntaxErrorComplainer(code.String(), msg)
				}
				code.buffer.WriteRune(r4)

				r5, _, err := runeScanner.ReadRune()
				if nil != err {
					msg := fmt.Sprintf("Problem after reading universal character escape sequence “\\x%s%s”.", string(r3), string(r4))
					return nil, nil, newSyntaxErrorComplainer(code.String(), msg)
				}
				code.buffer.WriteRune(r5)

				r6, _, err := runeScanner.ReadRune()
				if nil != err {
					msg := fmt.Sprintf("Problem after reading universal character escape sequence “\\x%s%s%s”.", string(r3), string(r4), string(r5))
					return nil, nil, newSyntaxErrorComplainer(code.String(), msg)
				}
				code.buffer.WriteRune(r6)



				r7, _, err := runeScanner.ReadRune()
				if nil != err {
					msg := fmt.Sprintf("Problem after reading universal character escape sequence “\\x%s%s%s%s”.", string(r3), string(r4), string(r5), string(r6))
					return nil, nil, newSyntaxErrorComplainer(code.String(), msg)
				}
				code.buffer.WriteRune(r7)

				r8, _, err := runeScanner.ReadRune()
				if nil != err {
					msg := fmt.Sprintf("Problem after reading universal character escape sequence “\\x%s%s%s%s%s”.", string(r3), string(r4), string(r5), string(r6), string(r7))
					return nil, nil, newSyntaxErrorComplainer(code.String(), msg)
				}
				code.buffer.WriteRune(r8)

				r9, _, err := runeScanner.ReadRune()
				if nil != err {
					msg := fmt.Sprintf("Problem after reading universal character escape sequence “\\x%s%s%s%s%s%s”.", string(r3), string(r4), string(r5), string(r6), string(r7), string(r8))
					return nil, nil, newSyntaxErrorComplainer(code.String(), msg)
				}
				code.buffer.WriteRune(r9)

				r10, _, err := runeScanner.ReadRune()
				if nil != err {
					msg := fmt.Sprintf("Problem after reading universal character escape sequence “\\x%s%s%s%s%s%s%s”.", string(r3), string(r4), string(r5), string(r6), string(r7), string(r8), string(r9))
					return nil, nil, newSyntaxErrorComplainer(code.String(), msg)
				}
				code.buffer.WriteRune(r10)



				if !(  ('0' <= r3 && r3 <= '9') || ('A' <= r3 && r3 <= 'F') || ('a' <= r3 && r3 <= 'f')  ) {
					msg := fmt.Sprintf("Invalid universal character escape sequence: “\\U%s%s%s%s%s%s%s%s”.", string(r3), string(r4), string(r5), string(r6), string(r7), string(r8), string(r9), string(r10))
					return nil, nil, newSyntaxErrorComplainer(code.String(), msg)
				}

				if !(  ('0' <= r4 && r4 <= '9') || ('A' <= r4 && r4 <= 'F') || ('a' <= r4 && r4 <= 'f')  ) {
					msg := fmt.Sprintf("Invalid universal character escape sequence: “\\U%s%s%s%s%s%s%s%s”.", string(r3), string(r4), string(r5), string(r6), string(r7), string(r8), string(r9), string(r10))
					return nil, nil, newSyntaxErrorComplainer(code.String(), msg)
				}

				if !(  ('0' <= r5 && r5 <= '9') || ('A' <= r5 && r5 <= 'F') || ('a' <= r5 && r5 <= 'f')  ) {
					msg := fmt.Sprintf("Invalid universal character escape sequence: “\\U%s%s%s%s%s%s%s%s”.", string(r3), string(r4), string(r5), string(r6), string(r7), string(r8), string(r9), string(r10))
					return nil, nil, newSyntaxErrorComplainer(code.String(), msg)
				}

				if !(  ('0' <= r6 && r6 <= '9') || ('A' <= r6 && r6 <= 'F') || ('a' <= r6 && r6 <= 'f')  ) {
					msg := fmt.Sprintf("Invalid universal character escape sequence: “\\U%s%s%s%s%s%s%s%s”.", string(r3), string(r4), string(r5), string(r6), string(r7), string(r8), string(r9), string(r10))
					return nil, nil, newSyntaxErrorComplainer(code.String(), msg)
				}



				if !(  ('0' <= r7 && r7 <= '9') || ('A' <= r7 && r7 <= 'F') || ('a' <= r7 && r7 <= 'f')  ) {
					msg := fmt.Sprintf("Invalid universal character escape sequence: “\\U%s%s%s%s%s%s%s%s”.", string(r3), string(r4), string(r5), string(r6), string(r7), string(r8), string(r9), string(r10))
					return nil, nil, newSyntaxErrorComplainer(code.String(), msg)
				}

				if !(  ('0' <= r8 && r8 <= '9') || ('A' <= r8 && r8 <= 'F') || ('a' <= r8 && r8 <= 'f')  ) {
					msg := fmt.Sprintf("Invalid universal character escape sequence: “\\U%s%s%s%s%s%s%s%s”.", string(r3), string(r4), string(r5), string(r6), string(r7), string(r8), string(r9), string(r10))
					return nil, nil, newSyntaxErrorComplainer(code.String(), msg)
				}

				if !(  ('0' <= r9 && r9 <= '9') || ('A' <= r9 && r9 <= 'F') || ('a' <= r9 && r9 <= 'f')  ) {
					msg := fmt.Sprintf("Invalid universal character escape sequence: “\\U%s%s%s%s%s%s%s%s”.", string(r3), string(r4), string(r5), string(r6), string(r7), string(r8), string(r9), string(r10))
					return nil, nil, newSyntaxErrorComplainer(code.String(), msg)
				}

				if !(  ('0' <= r10 && r10 <= '9') || ('A' <= r10 && r10 <= 'F') || ('a' <= r10 && r10 <= 'f')  ) {
					msg := fmt.Sprintf("Invalid universal character escape sequence: “\\U%s%s%s%s%s%s%s%s”.", string(r3), string(r4), string(r5), string(r6), string(r7), string(r8), string(r9), string(r10))
					return nil, nil, newSyntaxErrorComplainer(code.String(), msg)
				}



				r = 0

				if        '0' <= r3 && r3 <= '9' {
					r += (r3 - '0')*16*16*16*16*16*16*16
				} else if 'A' <= r3 && r3 <= 'F' {
					r += (r3 - 'A' + 10)*16*16*16*16*16*16*16
				} else if 'a' <= r3 && r3 <= 'f' {
					r += (r3 - 'a' + 10)*16*16*16*16*16*16*16
				} else {
					return nil, nil, newSyntaxErrorComplainer(code.String(), "THIS SHOULD NEVER HAPPEN!")
				}

				if        '0' <= r4 && r4 <= '9' {
					r += (r4 - '0')*16*16*16*16*16*16
				} else if 'A' <= r4 && r4 <= 'F' {
					r += (r4 - 'A' + 10)*16*16*16*16*16*16
				} else if 'a' <= r4 && r4 <= 'f' {
					r += (r4 - 'a' + 10)*16*16*16*16*16*16
				} else {
					return nil, nil, newSyntaxErrorComplainer(code.String(), "THIS SHOULD NEVER HAPPEN!")
				}

				if        '0' <= r5 && r5 <= '9' {
					r += (r5 - '0')*16*16*16*16*16
				} else if 'A' <= r5 && r5 <= 'F' {
					r += (r5 - 'A' + 10)*16*16*16*16*16
				} else if 'a' <= r5 && r5 <= 'f' {
					r += (r5 - 'a' + 10)*16*16*16*16*16
				} else {
					return nil, nil, newSyntaxErrorComplainer(code.String(), "THIS SHOULD NEVER HAPPEN!")
				}

				if        '0' <= r6 && r6 <= '9' {
					r += (r6 - '0')*16*16*16*16
				} else if 'A' <= r6 && r6 <= 'F' {
					r += (r6 - 'A' + 10)*16*16*16*16
				} else if 'a' <= r6 && r6 <= 'f' {
					r += (r6 - 'a' + 10)*16*16*16*16
				} else {
					return nil, nil, newSyntaxErrorComplainer(code.String(), "THIS SHOULD NEVER HAPPEN!")
				}



				if        '0' <= r7 && r7 <= '9' {
					r += (r7 - '0')*16*16*16
				} else if 'A' <= r7 && r7 <= 'F' {
					r += (r7 - 'A' + 10)*16*16*16
				} else if 'a' <= r7 && r7 <= 'f' {
					r += (r7 - 'a' + 10)*16*16*16
				} else {
					return nil, nil, newSyntaxErrorComplainer(code.String(), "THIS SHOULD NEVER HAPPEN!")
				}

				if        '0' <= r8 && r8 <= '9' {
					r += (r8 - '0')*16*16
				} else if 'A' <= r8 && r8 <= 'F' {
					r += (r8 - 'A' + 10)*16*16
				} else if 'a' <= r8 && r8 <= 'f' {
					r += (r8 - 'a' + 10)*16*16
				} else {
					return nil, nil, newSyntaxErrorComplainer(code.String(), "THIS SHOULD NEVER HAPPEN!")
				}

				if        '0' <= r9 && r9 <= '9' {
					r += (r9 - '0')*16
				} else if 'A' <= r9 && r9 <= 'F' {
					r += (r9 - 'A' + 10)*16
				} else if 'a' <= r9 && r9 <= 'f' {
					r += (r9 - 'a' + 10)*16
				} else {
					return nil, nil, newSyntaxErrorComplainer(code.String(), "THIS SHOULD NEVER HAPPEN!")
				}

				if        '0' <= r10 && r10 <= '9' {
					r += (r10 - '0')
				} else if 'A' <= r10 && r10 <= 'F' {
					r += (r10 - 'A' + 10)
				} else if 'a' <= r10 && r10 <= 'f' {
					r += (r10 - 'a' + 10)
				} else {
					return nil, nil, newSyntaxErrorComplainer(code.String(), "THIS SHOULD NEVER HAPPEN!")
				}
			case 'a':
				r = '\a'
			case 'b':
				r = '\b'
			case 'e':
				r = '\x1B'
			case 'f':
				r = '\f'
			case 'n':
				r = '\n'
			case 'r':
				r = '\r'
			case 't':
				r = '\t'
			case 'u':
				r3, _, err := runeScanner.ReadRune()
				if nil != err {
					msg := "Problem after reading hexadecimal escape “\\x”."
					return nil, nil, newSyntaxErrorComplainer(code.String(), msg)
				}
				code.buffer.WriteRune(r3)

				r4, _, err := runeScanner.ReadRune()
				if nil != err {
					msg := fmt.Sprintf("Problem after reading universal character escape sequence “\\u%s”.", string(r3))
					return nil, nil, newSyntaxErrorComplainer(code.String(), msg)
				}
				code.buffer.WriteRune(r4)

				r5, _, err := runeScanner.ReadRune()
				if nil != err {
					msg := fmt.Sprintf("Problem after reading universal character escape sequence “\\u%s%s”.", string(r3), string(r4))
					return nil, nil, newSyntaxErrorComplainer(code.String(), msg)
				}
				code.buffer.WriteRune(r5)

				r6, _, err := runeScanner.ReadRune()
				if nil != err {
					msg := fmt.Sprintf("Problem after reading universal character escape sequence “\\u%s%s%s”.", string(r3), string(r4), string(r5))
					return nil, nil, newSyntaxErrorComplainer(code.String(), msg)
				}
				code.buffer.WriteRune(r6)



				if !(  ('0' <= r3 && r3 <= '9') || ('A' <= r3 && r3 <= 'F') || ('a' <= r3 && r3 <= 'f')  ) {
					msg := fmt.Sprintf("Invalid universal character escape sequence: “\\u%s%s%s%s”.", string(r3), string(r4), string(r5), string(r6))
					return nil, nil, newSyntaxErrorComplainer(code.String(), msg)
				}

				if !(  ('0' <= r4 && r4 <= '9') || ('A' <= r4 && r4 <= 'F') || ('a' <= r4 && r4 <= 'f')  ) {
					msg := fmt.Sprintf("Invalid universal character escape sequence: “\\u%s%s%s%s”.", string(r3), string(r4), string(r5), string(r6))
					return nil, nil, newSyntaxErrorComplainer(code.String(), msg)
				}

				if !(  ('0' <= r5 && r5 <= '9') || ('A' <= r5 && r5 <= 'F') || ('a' <= r5 && r5 <= 'f')  ) {
					msg := fmt.Sprintf("Invalid universal character escape sequence: “\\u%s%s%s%s”.", string(r3), string(r4), string(r5), string(r6))
					return nil, nil, newSyntaxErrorComplainer(code.String(), msg)
				}

				if !(  ('0' <= r6 && r6 <= '9') || ('A' <= r6 && r6 <= 'F') || ('a' <= r6 && r6 <= 'f')  ) {
					msg := fmt.Sprintf("Invalid universal character escape sequence: “\\u%s%s%s%s”.", string(r3), string(r4), string(r5), string(r6))
					return nil, nil, newSyntaxErrorComplainer(code.String(), msg)
				}



				r = 0

				if        '0' <= r3 && r3 <= '9' {
					r += (r3 - '0')*16*16*16
				} else if 'A' <= r3 && r3 <= 'F' {
					r += (r3 - 'A' + 10)*16*16*16
				} else if 'a' <= r3 && r3 <= 'f' {
					r += (r3 - 'a' + 10)*16*16*16
				} else {
					return nil, nil, newSyntaxErrorComplainer(code.String(), "THIS SHOULD NEVER HAPPEN!")
				}

				if        '0' <= r4 && r4 <= '9' {
					r += (r4 - '0')*16*16
				} else if 'A' <= r4 && r4 <= 'F' {
					r += (r4 - 'A' + 10)*16*16
				} else if 'a' <= r4 && r4 <= 'f' {
					r += (r4 - 'a' + 10)*16*16
				} else {
					return nil, nil, newSyntaxErrorComplainer(code.String(), "THIS SHOULD NEVER HAPPEN!")
				}

				if        '0' <= r5 && r5 <= '9' {
					r += (r5 - '0')*16
				} else if 'A' <= r5 && r5 <= 'F' {
					r += (r5 - 'A' + 10)*16
				} else if 'a' <= r5 && r5 <= 'f' {
					r += (r5 - 'a' + 10)*16
				} else {
					return nil, nil, newSyntaxErrorComplainer(code.String(), "THIS SHOULD NEVER HAPPEN!")
				}

				if        '0' <= r6 && r6 <= '9' {
					r += r6 - '0'
				} else if 'A' <= r6 && r6 <= 'F' {
					r += r6 - 'A' + 10
				} else if 'a' <= r6 && r6 <= 'f' {
					r += r6 - 'a' + 10
				} else {
					return nil, nil, newSyntaxErrorComplainer(code.String(), "THIS SHOULD NEVER HAPPEN!")
				}
			case 'v':
				r = '\v'
			case 'x':
				r3, _, err := runeScanner.ReadRune()
				if nil != err {
					msg := "Problem after reading hexadecimal escape “\\x”."
					return nil, nil, newSyntaxErrorComplainer(code.String(), msg)
				}
				code.buffer.WriteRune(r3)

				r4, _, err := runeScanner.ReadRune()
				if nil != err {
					msg := fmt.Sprintf("Problem after reading hexadecimal escape “\\x%s”.", string(r3))
					return nil, nil, newSyntaxErrorComplainer(code.String(), msg)
				}
				code.buffer.WriteRune(r4)

				if !(  ('0' <= r3 && r3 <= '9') || ('A' <= r3 && r3 <= 'F') || ('a' <= r3 && r3 <= 'f')  ) {
					msg := fmt.Sprintf("Invalid hexadecimal escape sequence: “\\x%s%s”.", string(r3), string(r4))
					return nil, nil, newSyntaxErrorComplainer(code.String(), msg)
				}

				if !(  ('0' <= r4 && r4 <= '9') || ('A' <= r4 && r4 <= 'F') || ('a' <= r4 && r4 <= 'f')  ) {
					msg := fmt.Sprintf("Invalid hexadecimal escape sequence: “\\x%s%s”.", string(r3), string(r4))
					return nil, nil, newSyntaxErrorComplainer(code.String(), msg)
				}


				r = 0

				if        '0' <= r3 && r3 <= '9' {
					r += (r3 - '0')*16
				} else if 'A' <= r3 && r3 <= 'F' {
					r += (r3 - 'A' + 10)*16
				} else if 'a' <= r3 && r3 <= 'f' {
					r += (r3 - 'a' + 10)*16
				} else {
					return nil, nil, newSyntaxErrorComplainer(code.String(), "THIS SHOULD NEVER HAPPEN!")
				}

				if        '0' <= r4 && r4 <= '9' {
					r += r4 - '0'
				} else if 'A' <= r4 && r4 <= 'F' {
					r += r4 - 'A' + 10
				} else if 'a' <= r4 && r4 <= 'f' {
					r += r4 - 'a' + 10
				} else {
					return nil, nil, newSyntaxErrorComplainer(code.String(), "THIS SHOULD NEVER HAPPEN!")
				}
			case '\\':
				r = '\\'
			case '\'':
				r = '\''
			case '"':
				r = '"'
			case '?':
				r = '?'
			case '‘':
				r = '‘'
			case '’':
				r = '’'
			case '“':
				r = '“'
			case '”':
				r = '”'
			case '‹':
				r = '‹'
			case '›':
				r = '›'
			case '«':
				r = '«'
			case '»':
				r = '»'
			default:
				msg := fmt.Sprintf("Invalid blackslash escape sequence: \\%s", string(r2))
				return nil, nil, newSyntaxErrorComplainer(code.String(), msg)
			}
		}

		compiled.buffer.WriteRune(r)
	}



	return compiled, code, nil
}
