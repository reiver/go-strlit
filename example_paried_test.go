package strlit_test

import (
	"github.com/reiver/go-strlit"

	"fmt"
	"strings"
)

func ExamplePaired() {

	var src []byte = []byte(`«Hi» + «how are you?»`)

	var dst strings.Builder

	var paired strlit.Paired
	paired.BeginSymbol = '«'
	paired.EndSymbol   = '»'

	numWritten, numRead, err := paired.Decode(&dst, src)
	if nil != err {
		fmt.Printf("ERROR: problem decoding: %s\n", err)
		return
	}

	fmt.Printf("%d bytes were read.\n", numRead)
	fmt.Printf("%d bytes were written.\n", numWritten)
	fmt.Println("The value of the paired string literal...")
	fmt.Printf("%s\n", src[:numRead])
	fmt.Println("... is...")
	fmt.Printf("%s\n", dst.String())

	// Output:
	// 6 bytes were read.
	// 2 bytes were written.
	// The value of the paired string literal...
	// «Hi»
	// ... is...
	// Hi
}
