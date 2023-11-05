package mtest

func IntSum(a, b int) int {
	return a + b
}

func doSomething() string {
	return "do something"
}

func dataRace() int {
	var counter int
	go func() {
		counter++
	}()
	return counter
}

func runSomething() error {
	return nil
}
