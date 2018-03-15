package src

import (
	"fmt"
	"sync"
)

type argError struct {
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%s", e.prob)
}

func reader(i *int64, sums chan int64) {
	for {
		t := <-sums
		if *i > t {
			*i = t
		}
	}
}

func LeastTriangle(tri [][]int) (int64, error) {

	if tri == nil {
		return -1, &argError{"arg is nil!"}
	}
	if tri[0] == nil {
		return -1, &argError{"first arg slice is empty!"}
	}

	sums := make(chan int64)
	var least *int64 = new(int64)
	*least = 4294967295			// Largest possible int32 value
	var wg sync.WaitGroup

	go reader(least, sums)

	for {
		curr = <-nodes
		if curr.left == nil || curr.right == nil {
			break
		}
		wg.Add(1)
		go sumTree(curr, sums, func() { wg.Done() })
		nodes <-curr.left
		nodes <-curr.right
	}

	wg.Wait()

	return *least, nil
}

func main() {

	var triangle [][]int
	if v, err := LeastTriangle(triangle); err != nil{
		err.Error()
	} else {
		fmt.Printf("Least triangle: %v", v)
	}
}