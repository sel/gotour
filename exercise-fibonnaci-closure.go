package main

import (
	"fmt"
)

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var last []int
	return func() int {
		switch len(last) {
		case 0:
			last = append(last, 0)
		case 1:
			last = append(last, 1)
		default:
			last[0], last[1] = last[1], last[0]+last[1]
		}
		return last[len(last)-1]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 50; i++ {
		fmt.Println(f())
	}
}
