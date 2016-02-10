package strlit


import (
	"fmt"
	"strconv"
	"strings"

	"testing"
)


func TestCompile(t *testing.T) {

	testClass := "0-9 '...'"
	for r := '0'; r <= '9'; r++ {

		codeTmpl     := `'%s'`
                expectedTmpl :=  `%s`

		code     := fmt.Sprintf(codeTmpl,     string(r))
		expected := fmt.Sprintf(expectedTmpl, string(r))

		testNumber := int(r - '0')
		doTestOnCompile(t, testClass, testNumber, code, expected)

	}

	testClass = "0-9 \"...\""
	for r := '0'; r <= '9'; r++ {

		codeTmpl     := `"%s"`
                expectedTmpl :=  `%s`

		code     := fmt.Sprintf(codeTmpl,     string(r))
		expected := fmt.Sprintf(expectedTmpl, string(r))

		testNumber := int(r - '0')
		doTestOnCompile(t, testClass, testNumber, code, expected)

	}

	testClass = "0-9 ‘...’"
	for r := '0'; r <= '9'; r++ {

		codeTmpl     := `‘%s’`
                expectedTmpl :=  `%s`

		code     := fmt.Sprintf(codeTmpl,     string(r))
		expected := fmt.Sprintf(expectedTmpl, string(r))

		testNumber := int(r - '0')
		doTestOnCompile(t, testClass, testNumber, code, expected)

	}

	testClass = "0-9 “...”"
	for r := '0'; r <= '9'; r++ {

		codeTmpl     := `“%s”`
                expectedTmpl :=  `%s`

		code     := fmt.Sprintf(codeTmpl,     string(r))
		expected := fmt.Sprintf(expectedTmpl, string(r))

		testNumber := int(r - '0')
		doTestOnCompile(t, testClass, testNumber, code, expected)

	}

	testClass = "0-9 ‹...›"
	for r := '0'; r <= '9'; r++ {

		codeTmpl     := `‹%s›`
                expectedTmpl :=  `%s`

		code     := fmt.Sprintf(codeTmpl,     string(r))
		expected := fmt.Sprintf(expectedTmpl, string(r))

		testNumber := int(r - '0')
		doTestOnCompile(t, testClass, testNumber, code, expected)

	}

	testClass = "0-9 «...»"
	for r := '0'; r <= '9'; r++ {

		codeTmpl     := `«%s»`
                expectedTmpl :=  `%s`

		code     := fmt.Sprintf(codeTmpl,     string(r))
		expected := fmt.Sprintf(expectedTmpl, string(r))

		testNumber := int(r - '0')
		doTestOnCompile(t, testClass, testNumber, code, expected)

	}



	testClass = "a-z '...'"
	for r := 'a'; r <= 'z'; r++ {

		codeTmpl     := `'%s'`
                expectedTmpl :=  `%s`

		code     := fmt.Sprintf(codeTmpl,     string(r))
		expected := fmt.Sprintf(expectedTmpl, string(r))

		testNumber := int(r - 'a')
		doTestOnCompile(t, testClass, testNumber, code, expected)

	}

	testClass = "a-z \"...\""
	for r := 'a'; r <= 'z'; r++ {

		codeTmpl     := `"%s"`
                expectedTmpl :=  `%s`

		code     := fmt.Sprintf(codeTmpl,     string(r))
		expected := fmt.Sprintf(expectedTmpl, string(r))

		testNumber := int(r - 'a')
		doTestOnCompile(t, testClass, testNumber, code, expected)

	}

	testClass = "a-z ‘...’"
	for r := 'a'; r <= 'z'; r++ {

		codeTmpl     := `‘%s’`
                expectedTmpl :=  `%s`

		code     := fmt.Sprintf(codeTmpl,     string(r))
		expected := fmt.Sprintf(expectedTmpl, string(r))

		testNumber := int(r - 'a')
		doTestOnCompile(t, testClass, testNumber, code, expected)

	}

	testClass = "a-z “...”"
	for r := 'a'; r <= 'z'; r++ {

		codeTmpl     := `“%s”`
                expectedTmpl :=  `%s`

		code     := fmt.Sprintf(codeTmpl,     string(r))
		expected := fmt.Sprintf(expectedTmpl, string(r))

		testNumber := int(r - 'a')
		doTestOnCompile(t, testClass, testNumber, code, expected)

	}

	testClass = "a-z ‹...›"
	for r := 'a'; r <= 'z'; r++ {

		codeTmpl     := `‹%s›`
                expectedTmpl :=  `%s`

		code     := fmt.Sprintf(codeTmpl,     string(r))
		expected := fmt.Sprintf(expectedTmpl, string(r))

		testNumber := int(r - 'a')
		doTestOnCompile(t, testClass, testNumber, code, expected)

	}

	testClass = "a-z «...»"
	for r := 'a'; r <= 'z'; r++ {

		codeTmpl     := `«%s»`
                expectedTmpl :=  `%s`

		code     := fmt.Sprintf(codeTmpl,     string(r))
		expected := fmt.Sprintf(expectedTmpl, string(r))

		testNumber := int(r - 'a')
		doTestOnCompile(t, testClass, testNumber, code, expected)

	}



	testClass = "A-Z '...'"
	for r := 'a'; r <= 'Z'; r++ {

		codeTmpl     := `'%s'`
                expectedTmpl :=  `%s`

		code     := fmt.Sprintf(codeTmpl,     string(r))
		expected := fmt.Sprintf(expectedTmpl, string(r))

		testNumber := int(r - 'A')
		doTestOnCompile(t, testClass, testNumber, code, expected)

	}

	testClass = "A-Z \"...\""
	for r := 'A'; r <= 'Z'; r++ {

		codeTmpl     := `"%s"`
                expectedTmpl :=  `%s`

		code     := fmt.Sprintf(codeTmpl,     string(r))
		expected := fmt.Sprintf(expectedTmpl, string(r))

		testNumber := int(r - 'A')
		doTestOnCompile(t, testClass, testNumber, code, expected)

	}

	testClass = "A-Z ‘...’"
	for r := 'A'; r <= 'Z'; r++ {

		codeTmpl     := `‘%s’`
                expectedTmpl :=  `%s`

		code     := fmt.Sprintf(codeTmpl,     string(r))
		expected := fmt.Sprintf(expectedTmpl, string(r))

		testNumber := int(r - 'A')
		doTestOnCompile(t, testClass, testNumber, code, expected)

	}

	testClass = "A-Z “...”"
	for r := 'A'; r <= 'Z'; r++ {

		codeTmpl     := `“%s”`
                expectedTmpl :=  `%s`

		code     := fmt.Sprintf(codeTmpl,     string(r))
		expected := fmt.Sprintf(expectedTmpl, string(r))

		testNumber := int(r - 'A')
		doTestOnCompile(t, testClass, testNumber, code, expected)

	}

	testClass = "A-Z ‹...›"
	for r := 'A'; r <= 'Z'; r++ {

		codeTmpl     := `‹%s›`
                expectedTmpl :=  `%s`

		code     := fmt.Sprintf(codeTmpl,     string(r))
		expected := fmt.Sprintf(expectedTmpl, string(r))

		testNumber := int(r - 'A')
		doTestOnCompile(t, testClass, testNumber, code, expected)

	}

	testClass = "A-Z «...»"
	for r := 'A'; r <= 'Z'; r++ {

		codeTmpl     := `«%s»`
                expectedTmpl :=  `%s`

		code     := fmt.Sprintf(codeTmpl,     string(r))
		expected := fmt.Sprintf(expectedTmpl, string(r))

		testNumber := int(r - 'A')
		doTestOnCompile(t, testClass, testNumber, code, expected)

	}



	hexadecimalRunes := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'a', 'B', 'b', 'C', 'c', 'D', 'd', 'E', 'e', 'F', 'f'}

	testClass = "\\xhh '...'"
	testNumber := 0
	for _, r1 := range hexadecimalRunes {
		for _, r0 := range hexadecimalRunes {
			hh := fmt.Sprintf("%s%s", string(r1), string(r0))
			s := fmt.Sprintf("\\x%s%s", string(r1), string(r0))

			n, err := strconv.ParseInt(hh, 16, 16)
			if nil != err {
				t.Errorf("THIS SHOULD NOT HAPPEN! %v", err)
				continue
			}

			codeTmpl     := `'%s'`
			code     := fmt.Sprintf(codeTmpl, s)

			expected := string(rune(n))

			doTestOnCompile(t, testClass, testNumber, code, expected)
			testNumber++
		}
	}

	testClass = "\\xhh \"...\""
	testNumber = 0
	for _, r1 := range hexadecimalRunes {
		for _, r0 := range hexadecimalRunes {
			hh := fmt.Sprintf("%s%s", string(r1), string(r0))
			s := fmt.Sprintf("\\x%s%s", string(r1), string(r0))

			n, err := strconv.ParseInt(hh, 16, 16)
			if nil != err {
				t.Errorf("THIS SHOULD NOT HAPPEN! %v", err)
				continue
			}

			codeTmpl     := `"%s"`
			code     := fmt.Sprintf(codeTmpl, s)

			expected := string(rune(n))

			doTestOnCompile(t, testClass, testNumber, code, expected)
			testNumber++
		}
	}

	testClass = "\\xhh ‘...’"
	testNumber = 0
	for _, r1 := range hexadecimalRunes {
		for _, r0 := range hexadecimalRunes {
			hh := fmt.Sprintf("%s%s", string(r1), string(r0))
			s := fmt.Sprintf("\\x%s%s", string(r1), string(r0))

			n, err := strconv.ParseInt(hh, 16, 16)
			if nil != err {
				t.Errorf("THIS SHOULD NOT HAPPEN! %v", err)
				continue
			}

			codeTmpl     := `‘%s’`
			code     := fmt.Sprintf(codeTmpl, s)

			expected := string(rune(n))

			doTestOnCompile(t, testClass, testNumber, code, expected)
			testNumber++
		}
	}

	testClass = "\\xhh “...”"
	testNumber = 0
	for _, r1 := range hexadecimalRunes {
		for _, r0 := range hexadecimalRunes {
			hh := fmt.Sprintf("%s%s", string(r1), string(r0))
			s := fmt.Sprintf("\\x%s%s", string(r1), string(r0))

			n, err := strconv.ParseInt(hh, 16, 16)
			if nil != err {
				t.Errorf("THIS SHOULD NOT HAPPEN! %v", err)
				continue
			}

			codeTmpl     := `“%s”`
			code     := fmt.Sprintf(codeTmpl, s)

			expected := string(rune(n))

			doTestOnCompile(t, testClass, testNumber, code, expected)
			testNumber++
		}
	}

	testClass = "\\xhh ‹...›"
	testNumber = 0
	for _, r1 := range hexadecimalRunes {
		for _, r0 := range hexadecimalRunes {
			hh := fmt.Sprintf("%s%s", string(r1), string(r0))
			s := fmt.Sprintf("\\x%s%s", string(r1), string(r0))

			n, err := strconv.ParseInt(hh, 16, 16)
			if nil != err {
				t.Errorf("THIS SHOULD NOT HAPPEN! %v", err)
				continue
			}

			codeTmpl     := `‹%s›`
			code     := fmt.Sprintf(codeTmpl, s)

			expected := string(rune(n))

			doTestOnCompile(t, testClass, testNumber, code, expected)
			testNumber++
		}
	}

	testClass = "\\xhh «...»"
	testNumber = 0
	for _, r1 := range hexadecimalRunes {
		for _, r0 := range hexadecimalRunes {
			hh := fmt.Sprintf("%s%s", string(r1), string(r0))
			s := fmt.Sprintf("\\x%s%s", string(r1), string(r0))

			n, err := strconv.ParseInt(hh, 16, 16)
			if nil != err {
				t.Errorf("THIS SHOULD NOT HAPPEN! %v", err)
				continue
			}

			codeTmpl     := `«%s»`
			code     := fmt.Sprintf(codeTmpl, s)

			expected := string(rune(n))

			doTestOnCompile(t, testClass, testNumber, code, expected)
			testNumber++
		}
	}



	testClass = "\\uhhhh '...'"
	testNumber = 0
	for _, r3 := range hexadecimalRunes {
		for _, r2 := range hexadecimalRunes {
			for _, r1 := range hexadecimalRunes {
				for _, r0 := range hexadecimalRunes {
					hhhh := fmt.Sprintf("%s%s%s%s", string(r3), string(r2), string(r1), string(r0))
					s := fmt.Sprintf("\\u%s%s%s%s", string(r3), string(r2), string(r1), string(r0))

					n, err := strconv.ParseInt(hhhh, 16, 32)
					if nil != err {
						t.Errorf("THIS SHOULD NOT HAPPEN! %v", err)
						continue
					}

					codeTmpl := `'%s'`
					code     := fmt.Sprintf(codeTmpl, s)

					expected := string(rune(n))

					doTestOnCompile(t, testClass, testNumber, code, expected)
					testNumber++
				}
			}
		}
	}

	testClass = "\\uhhhh \"...\""
	testNumber = 0
	for _, r3 := range hexadecimalRunes {
		for _, r2 := range hexadecimalRunes {
			for _, r1 := range hexadecimalRunes {
				for _, r0 := range hexadecimalRunes {
					hhhh := fmt.Sprintf("%s%s%s%s", string(r3), string(r2), string(r1), string(r0))
					s := fmt.Sprintf("\\u%s%s%s%s", string(r3), string(r2), string(r1), string(r0))

					n, err := strconv.ParseInt(hhhh, 16, 32)
					if nil != err {
						t.Errorf("THIS SHOULD NOT HAPPEN! %v", err)
						continue
					}

					codeTmpl := `"%s"`
					code     := fmt.Sprintf(codeTmpl, s)

					expected := string(rune(n))

					doTestOnCompile(t, testClass, testNumber, code, expected)
					testNumber++
				}
			}
		}
	}

	testClass = "\\uhhhh ‘...’"
	testNumber = 0
	for _, r3 := range hexadecimalRunes {
		for _, r2 := range hexadecimalRunes {
			for _, r1 := range hexadecimalRunes {
				for _, r0 := range hexadecimalRunes {
					hhhh := fmt.Sprintf("%s%s%s%s", string(r3), string(r2), string(r1), string(r0))
					s := fmt.Sprintf("\\u%s%s%s%s", string(r3), string(r2), string(r1), string(r0))

					n, err := strconv.ParseInt(hhhh, 16, 32)
					if nil != err {
						t.Errorf("THIS SHOULD NOT HAPPEN! %v", err)
						continue
					}

					codeTmpl := `‘%s’`
					code     := fmt.Sprintf(codeTmpl, s)

					expected := string(rune(n))

					doTestOnCompile(t, testClass, testNumber, code, expected)
					testNumber++
				}
			}
		}
	}

	testClass = "\\uhhhh “...”"
	testNumber = 0
	for _, r3 := range hexadecimalRunes {
		for _, r2 := range hexadecimalRunes {
			for _, r1 := range hexadecimalRunes {
				for _, r0 := range hexadecimalRunes {
					hhhh := fmt.Sprintf("%s%s%s%s", string(r3), string(r2), string(r1), string(r0))
					s := fmt.Sprintf("\\u%s%s%s%s", string(r3), string(r2), string(r1), string(r0))

					n, err := strconv.ParseInt(hhhh, 16, 32)
					if nil != err {
						t.Errorf("THIS SHOULD NOT HAPPEN! %v", err)
						continue
					}

					codeTmpl := `“%s”`
					code     := fmt.Sprintf(codeTmpl, s)

					expected := string(rune(n))

					doTestOnCompile(t, testClass, testNumber, code, expected)
					testNumber++
				}
			}
		}
	}

	testClass = "\\uhhhh ‹...›"
	testNumber = 0
	for _, r3 := range hexadecimalRunes {
		for _, r2 := range hexadecimalRunes {
			for _, r1 := range hexadecimalRunes {
				for _, r0 := range hexadecimalRunes {
					hhhh := fmt.Sprintf("%s%s%s%s", string(r3), string(r2), string(r1), string(r0))
					s := fmt.Sprintf("\\u%s%s%s%s", string(r3), string(r2), string(r1), string(r0))

					n, err := strconv.ParseInt(hhhh, 16, 32)
					if nil != err {
						t.Errorf("THIS SHOULD NOT HAPPEN! %v", err)
						continue
					}

					codeTmpl := `‹%s›`
					code     := fmt.Sprintf(codeTmpl, s)

					expected := string(rune(n))

					doTestOnCompile(t, testClass, testNumber, code, expected)
					testNumber++
				}
			}
		}
	}

	testClass = "\\uhhhh «...»"
	testNumber = 0
	for _, r3 := range hexadecimalRunes {
		for _, r2 := range hexadecimalRunes {
			for _, r1 := range hexadecimalRunes {
				for _, r0 := range hexadecimalRunes {
					hhhh := fmt.Sprintf("%s%s%s%s", string(r3), string(r2), string(r1), string(r0))
					s := fmt.Sprintf("\\u%s%s%s%s", string(r3), string(r2), string(r1), string(r0))

					n, err := strconv.ParseInt(hhhh, 16, 32)
					if nil != err {
						t.Errorf("THIS SHOULD NOT HAPPEN! %v", err)
						continue
					}

					codeTmpl := `«%s»`
					code     := fmt.Sprintf(codeTmpl, s)

					expected := string(rune(n))

					doTestOnCompile(t, testClass, testNumber, code, expected)
					testNumber++
				}
			}
		}
	}


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
			Code:     `‹›`,
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
			Code:     `‹ ›`,
			Expected:  ` `,
		},
		{
			Code:     `« »`,
			Expected:  ` `,
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
			Code:     `‹123456789›`,
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
			Code:     `‹a1b2c3d4e5f6g7h8i9›`,
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
			Code:     `‹apple banana cherry›`,
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
			Code:     `‹\0›`,
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



//		{
//			Code:     `'\U01000000'`,
//			Expected:  "\U01000000",
//		},
//		{
//			Code:     `"\U023bCdEf"`,
//			Expected:  "\U023bCdEf",
//		},
//		{
//			Code:     `‘\U023bCdEf’`,
//			Expected:  "\U023bCdEf",
//		},
//		{
//			Code:     `“\U023bCdEf”`,
//			Expected:  "\U023bCdEf",
//		},
//		{
//			Code:     `«\U023bCdDf»`,
//			Expected:  "\U023bCdDf",
//		},



//@TODO: ###############################################



//		{
//			Code:     `'\U109aBcDE'`,
//			Expected:  "\U109aBcDE",
//		},
//		{
//			Code:     `"\U109aBcDE"`,
//			Expected:  "\U109aBcDE",
//		},
//		{
//			Code:     `‘\U109aBcDE’`,
//			Expected:  "\U109aBcDE",
//		},
//		{
//			Code:     `“\U109aBcDE”`,
//			Expected:  "\U109aBcDE",
//		},
//		{
//			Code:     `«\U109aBcDE»`,
//			Expected:  "\U109aBcDE",
//		},



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


//	tests = append(tests, hexadecialEscapeTests...)
//	tests = append(tests, lowercaseUEscapeTests...)


	testClass = "specific"
	for testNumber, test := range tests {
		doTestOnCompile(t, testClass, testNumber, test.Code, test.Expected)
	}

}


func doTestOnCompile(t *testing.T, testClass string, testNumber int, testCode string, testExpected string) {

	actualCompiled, err := Compile( strings.NewReader(testCode) )

	if nil != err {
		t.Errorf("[%s] For test #%d, did not expected an error, but actually got one: %v", testClass, testNumber, err)
		return
	}

	if expected, actual := testExpected, actualCompiled.String(); expected != actual {
		t.Errorf("[%s] For test #%d, expected %q, but actually got %q.", testClass, testNumber, expected, actual)
		return
	}

	if expected, actual := []rune(testExpected), actualCompiled.Runes(); len(expected) != len(expected) {
		t.Errorf("[%s] For test #%d, expected %d runes, but actually got %d runes.", testClass, testNumber, expected, actual)
		return
	} else {
		for runeNumber, expectedDatum := range expected {
			actualDatum := actual[runeNumber]

			if expectedDatum != actualDatum {
				t.Errorf("[%s] For test #%d an rune #%d, expected rune to be %q, but actually was %q.", testClass, testNumber, runeNumber, expectedDatum, actualDatum)
				continue
			}
		}
	}

	if expected, actual := []byte(testExpected), actualCompiled.Bytes(); len(expected) != len(expected) {
		t.Errorf("[%s] For test #%d, expected %d bytes, but actually got %d bytes.", testClass, testNumber, expected, actual)
		return
	} else {
		for byteNumber, expectedDatum := range expected {
			actualDatum := actual[byteNumber]

			if expectedDatum != actualDatum {
				t.Errorf("[%s] For test #%d an byte #%d, expected byte to be %d, but actually was %d.", testClass, testNumber, byteNumber, expectedDatum, actualDatum)
				continue
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
