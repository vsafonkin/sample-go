package nilvalue

import "fmt"

type DoSomething func()
type SomeInterface interface{}

// Init zero values
var b bool     // false
var i int      // 0
var f float64  // 0.000000
var str string // empty string
var r rune     // \x00

// Init nil values
var slc []int
var m map[string]string
var p *string
var ch chan int
var fn DoSomething
var smint SomeInterface

type Sample struct {
	num int
	slc []int
	fn  DoSomething
}

func ZeroValues() {
	fmt.Printf("Init values:\nbool: %t\nint: %d\nfloat: %f\nstring: %s\nrune: %q\n", b, i, f, str, r)
}

func NilValues() {
	fmt.Println("Init nil values:")
	fmt.Printf("slice == nil: %t\n", slc == nil)
	fmt.Printf("map == nil: %t\n", m == nil)
	fmt.Printf("pointer == nil: %t\n", p == nil)
	fmt.Printf("channel == nil: %t\n", ch == nil)
	fmt.Printf("function == nil: %t\n", fn == nil)
	fmt.Printf("interface == nil: %t\n", smint == nil)
}

func Run() {
	sample := Sample{}
	fmt.Println(sample)
}
