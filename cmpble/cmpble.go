package cmpble

import "fmt"

type Runner interface {
	RunSomething() // Interface is comparable
}

type Tester interface {
	TestSomething() error // Interface is not comparage
}

type BobRunner struct{}
type AlisaTester struct{}

func ComparableTypes() {
	var b bool
	var i int
	var f float64
	var s string
	var p *int // Two pointer values are equal if they point to the same variable or if both have value nil
	var ch chan int

	fmt.Printf("%t %d %f %s %p %v\n", b, i, f, s, p, ch)

	// Interface types that are not type parameters are comparable.
	// Two interface values are equal if they have identical dynamic types and equal dynamic values or if both have value nil
	bob := BobRunner{}
	bob_2 := BobRunner{}
	fmt.Println(bob == bob_2)
}

func OrderedTypes() {
	var i int
	var f float64
	var s string

	fmt.Printf("%d %f %s\n", i, f, s)
}

func (b BobRunner) RunSomething() {
	fmt.Println("BobRunner run something")
}

func (a AlisaTester) TestSomething() error {
	return nil
}
