package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

/*type Tree struct{
	Left *Tree
	Value int
	Rigth *Tree
} */

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	} else if t.Left == nil {
		ch <- t.Value
		if t.Right != nil {
			Walk(t.Right, ch)
		}
		return
	} else {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Walker to avoid panic
func Walker(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	//close the channel to avoid panic
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) (isSame bool) {
	var br bool
	done := make(chan bool)
	c1 := make(chan int)
	c2 := make(chan int)
	go Walker(t1, c1)
	go Walker(t2, c2)
	go func() {
		for i := range c1 {
			if i == <-c2 {
				br = true
			} else {
				br = false
				break
			}
		}
		done <- true
	}()
	<-done
	return br
}

func main() {
	t1 := tree.New(5)
	t2 := tree.New(5)
	fmt.Println("Tree 1:", t1)
	fmt.Println("Tree 2:", t2)
	fmt.Println("Are they same? - ", Same(t1, t2))
	fmt.Println("------")
}
