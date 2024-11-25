package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()

	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	for v1 := range ch1 {
		v2, ok := <-ch2
		if !ok || v1 != v2 {
			return false
		}
	}

	// Проверим, что ch2 тоже завершён
	// Это на случай, если количество элементов в t1 больше, чем в t2
	_, ok := <-ch2
	return !ok
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(1)
	t3 := tree.New(2)
	res1 := Same(t1, t2)
	res2 := Same(t1, t3)
	fmt.Println("Is t1 the same as t2? Answer: ", res1)
	fmt.Println("Is t1 the same as t3? Answer: ", res2)
}
