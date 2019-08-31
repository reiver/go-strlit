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

Paired String Literals

With strlit.Paired, package strlit supports decoding from, and encoding to Paired String Literals.

Some examples of pairs string literals might be:

	‘Hello world!’

	“Hello world!”

	‹Hello world!›

	«Hello world!»

	「Hello world!」

	『Hello world!』

Note that this did NOT include these:

	'Hello world!'

	"Hello world!"

(Look very closely if you think these are the same as the first 2 examples in with the (valid) paired string literals.)

Paired String Literals have a different symbol for the Begin Symbol, and the End Symbol.

Here is a simple example using strlit.Paired.Decode():

	var src []byte = []byte(`«Hello world!»`) // <---- This is the paired string literal we will decode.
	
	// ...
	
	var paired strlit.Paried
	paired.BeginSymbol = '«' // <---- Note this is the same as the beginning symbol of our paired string literal above.
	paired.EndSymbol   = '»' // <---- Note this is the same as the ending symbol of our paired string literal above.
	
	var dst []byte = make([]byte, len(src)) // <---- We make ‘dst’ the same size as ‘src’. The result will be shorter than this. But we will shorten it later.
	
	bytesWritten, bytesRead, err := paried.Decode(dst, src)
	if nil != err {
		return err
	}
	
	// ...
	
	dst = dst[:bytesWritten] // <---- We shorten ‘dst’ so that we only look at the bytes that we actully wrote to.
	
	// ...
	
	src = src[:bytesRead] // <--- We shortern ‘src’ to get past the bytes we already read (and already decoded).

Balanced Paired String Literals

This package's strlit.Paired supports, and enforces Nesting of the beginning symbol, and ending symbol.

So, for example, these are also Paired String Literals, but have nesting:

	‘And she said, ‘do you know him?’.’

	“And she said, “do you know him?”.”

	‹And she said, ‹do you know him?›.›

	«And she said, «do you know him?».»

	「And she said, 「do you know him?」.」

	『And she said, 『do you know him?』.』

Note that this functions as a way of including the beginning symbol, and ending symbol of the Paired String Literal
within the string literal. However, there are rules to this: they too have to be paired correctly.

Paired String Literal Symbols

Really, almost any 2 different symbols could be use as the beginning symbol, and ending symbol, of a Paired String Literal.

But Unicode was provide a number of symbols that are naturally paired with each other.

Here are some of them:

Left Single Quotation Mark,
Right Single Quotation Mark:

	‘...’

Left Double Quotation Mark,
Right Double Quotation Mark:

	“...”

Single Left-Pointing Angle Quotation Mark,
Single Right-Pointing Angle Quotation Mark:

	‹...›

Left-Pointing Double Angle Quotation Mark,
Right-Pointing Double Angle Quotation Mark:

	«...»

Left Corner Bracket,
Right Corner Bracket:

	「...」

Left White Corner Bracket,
Right White Corner Bracket:

	『...』

Left Square Bracket,
Right Square Bracket:

	[...]

Left Square Bracket with Quill,
Right Square Bracket with Quill:

	⁅...⁆

Mathematical Left White Square Bracket,
Mathematical Right White Square Bracket:

	⟦...⟧

Left Parenthesis,
Right Parenthesis:

	(...)

Medium Left Parenthesis Ornament,
Medium Right Parenthesis Ornament:

	❨...❩

Medium Flattened Left Parenthesis Ornament,
Medium Flattened Right Parenthesis Ornament:

	❪...❫

Left White Parenthesis,
Right White Parenthesis:

	⦅...⦆

Left Double Parenthesis,
Right Double Parenthesis:

	⸨...⸩

Ornate Left Parenthesis,
Ornate Right Parenthesis:

	﴾...﴿

Small Left Parenthesis,
Small Right Parenthesis:

	﹙...﹚

Fullwidth Left Parenthesis,
Fullwidth Right Parenthesis:

	（...）

Fullwidth Left White Parenthesis,
Fullwidth Right White Parenthesis:

	｟...｠

Left Curly Bracket,
Right Curly Bracket:

	{...}

Medium Left Curly Bracket Ornament,
Medium Right Curly Bracket Ornament:

	❴...❵

Left White Curly Bracket,
Right White Curly Bracket:

	⦃...⦄

Small Left Curly Bracket,
Small Right Curly Bracket:

	﹛...﹜

Left-Pointing Angle Bracket,
Right-Pointing Angle Bracket:

	〈...〉

Mathematical Left Angle Bracket,
Mathematical Right Angle Bracket:

	⟨...⟩

Left Tortoise Shell Bracket,
Right Tortoise Shell Bracket:

	〔...〕

Left White Tortoise Shell Bracket,
Right White Tortoise Shell Bracket:

	〘...〙

Small Left Tortoise Shell Bracket,
Small Right Tortoise Shell Bracket:

	﹝...﹞

Tibetan Mark Gug Rtags Gyon,
Tibetan Mark Gug Rtags Gyas:

	༺...༻

Tibetan Mark Ang Khang Gyon,
Tibetan Mark Ang Khang Gyas:

	༼...༽
*/
package strlit
