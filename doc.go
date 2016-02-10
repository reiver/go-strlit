/*
Package strlit provides a way to "compile" a string literal into the value the it represents.

For example, here are some string different string literals:

'apple banana cherry\tgrape'

"apple banana cherry\tgrape"

‘apple banana cherry\tgrape’

“apple banana cherry\tgrape”

‹apple banana cherry\tgrape›

«apple banana cherry\tgrape»

Note that 6 different styles of quotes are supported by this package.

It also supports a number of escape sequences. Namely:

\0 (null)

\a (bell)

\b (backspace)

\e (escape)

\f (form feed)

\n (line feed)

\r (carriage return)

\t (horizontal tab )

\v (vertical tab)

\\ (backslash)

\'

\"

\?

\‘

\’

\“

\”

\‹

\›

\«

\»

\xhh

\uhhhh

\Uhhhhhhhh

Example

An example usage of this package looks like the following.

	  stringLiteral := `"This is a string literal.\tStill going."`
	//stringLiteral := `'This is a string literal.\tStill going.'`
	//stringLiteral := `‘This is a string literal.\tStill going.’`
	//stringLiteral := `“This is a string literal.\tStill going.”`
	//stringLiteral := `‹This is a string literal.\tStill going.›`
	//stringLiteral := `«This is a string literal.\tStill going.»`

	compiled, code, err := strlit.Compile( strings.NewReader(stringLiteral) )
	if nil != err {
		//@TODO
	}

	fmt.Printf("The value of the string literal is: %s\n", compiled.String())
	fmt.Printf("The original pre-compiled string literal is: %s\n", code.String())

*/
package strlit
