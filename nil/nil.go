package nilvalue

import "fmt"

// Init zero values
var b bool     // false
var i int      // 0
var f float64  // 0.000000
var str string // empty string
var r rune     // \x00

func ZeroValues() {
	fmt.Printf("Init values:\nbool: %t\nint: %d\nfloat: %f\nstring: %s\nrune: %q\n", b, i, f, str, r)
}
