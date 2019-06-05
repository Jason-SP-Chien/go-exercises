package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	f0, f1 := 0, 0
	return func() int {
		f2 := f0+f1
		if f2 == 0 {
			f2 = 1
		}
		f0, f1 = f1, f2
		return f0
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 30; i++ {
		fmt.Println(f())
	}
}