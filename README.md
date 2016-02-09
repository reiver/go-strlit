# go-strlit

A library that provides tools to "compile" a string literals into the value they represent, for the Go programming language.

This supports a number of different beginning and ending characters for string literals. Namely:
* '...'
* "..."
* ‘...’
* “...”
* «...»


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-strlit

[![GoDoc](https://godoc.org/github.com/reiver/go-strlit?status.svg)](https://godoc.org/github.com/reiver/go-strlit)


## Example
```
  stringLiteral := `"This is a string literal.\tStill going."`
//stringLiteral := `'This is a string literal.\tStill going.'`
//stringLiteral := `‘This is a string literal.\tStill going.’`
//stringLiteral := `“This is a string literal.\tStill going.”`
//stringLiteral := `«This is a string literal.\tStill going.»`

compiled, err := strlit.Compile( strings.NewReader(stringLiteral) )
if nil != err {
	//@TODO
}

fmt.Printf("The value of the string literal is: %s\n", compiled.String())
```
