package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	doWalk(t, ch)
	close(ch)
}

func doWalk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	doWalk(t.Left, ch)
	ch <- t.Value
	doWalk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int, 10)
	go Walk(t1, c1)
	c2 := make(chan int, 10)
	go Walk(t2, c2)
	for i := range c1 {
		if i != <-c2 {
			return false
		}
	}
	return true
}

func main() {
	s1 := Same(tree.New(1), tree.New(1))
	fmt.Println("Same(tree.New(1), tree.New(1)):", s1)
	s2 := Same(tree.New(1), tree.New(2))
	fmt.Println("Same(tree.New(1), tree.New(2)):", s2)
}
