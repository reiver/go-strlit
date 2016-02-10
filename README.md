# go-strlit

A library that provides tools to "compile" a string literals into the value they represent, for the Go programming language.

This supports a number of different beginning and ending characters for string literals. Namely:
* '...'
* "..."
* ‘...’
* “...”
* ‹...›
* «...»

This also supports the following string literal escape sequences:
* \0 (null)
* \a (bell)
* \b (backspace)
* \e (escape)
* \f (form feed)
* \n (line feed)
* \r (carriage return)
* \t (horizontal tab )
* \v (vertical tab)
* \\\\ (backslash)
* \'
* \"
* \?
* \‘
* \’
* \“
* \”
* \‹
* \›
* \«
* \»
* \xhh
* \uhhhh
* \Uhhhhhhhh


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-strlit

[![GoDoc](https://godoc.org/github.com/reiver/go-strlit?status.svg)](https://godoc.org/github.com/reiver/go-strlit)


## Example
```
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
```
