package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	} else {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
		Walk(t, ch)
	}
}
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	if <-ch1 == <-ch2 {
		return true
	} else {
		return false
	}

}
func main() {
	t1 := tree.New(1)
	t2 := tree.New(2)
	fmt.Println("Tree 1 == Tree 2,", Same(t1, t2))

	/*ch := make(chan int)
	go func() {
		Walk(tree.New(1),ch)
		close(ch)
	}()
	for i:=range ch{
		fmt.Println(i)
	}*/
}
