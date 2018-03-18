package main

import (
	// "fmt"
	"sync"
	"errors"
)

type Triangle [][]int				// Holds all triangle values

type errorString struct {
	prob string
}

func (e *errorString) Error() string {
	return e.prob
}

// Reads incoming sums on the sums channel and compares them to the current least sum
func reader(i *int64, sums chan int64, mut sync.Mutex) {
	for {
		t := <-sums
		if *i > t {
			mut.Lock()
			*i = t
			mut.Unlock()
		}
	}
}

// For taking the sum of one 'level' of the triangle
func sumSlice(s []int) int64 {
	var sum int64
	for _, v := range(s) {
		sum += int64(v)
	}
	return sum
}

func (t Triangle) sumTriangle(index int, sumChan chan int64, callback func()) {

	defer callback()

	var sum int64 = int64(t[0][index] + t[1][index] + t[1][index+1])

	sumChan <-sum				  // Start off by adding the smallest triangle to the sum channel

	for i := 2; i < len(t); i++ {
		sum += sumSlice(t[i][index:index+i+1])
		sumChan <-sum
	}
}

/*
	Visual depiction of triangle as slice of int slices
				[		 [0],
					   [0,  1],
					 [0,  1,  2],
				   [0,  1,  2,  3],
				 [0,  1,  2,  3,  4]	]
					..........
*/

func LeastTriangle(tri Triangle) (int64, error) {

	// Error handling
	switch {
		case tri == nil || len(tri) == 0:
			return -1, errors.New("Triangle is nil!")
		case tri[0] == nil:
			return -1, errors.New("First level is nil!")
		case len(tri) < 2:
			return -1, errors.New("Triangle has less than 2 levels!")
		case len(tri[0]) != 1:
			return -1, errors.New("First level of triangle doest not contain one int!")
		case len(tri[1]) != 2:
			return -1, errors.New("Second level of triangle does not contain two ints!")
	}

	sums := make(chan int64)
	var least *int64 = new(int64)
	*least = int64(tri[0][0] + tri[1][0] + tri[1][1]) 		// initial value is top small triangle

	var wg sync.WaitGroup
	var mut = &sync.Mutex{}

	go reader(least, sums, *mut)

	for i := 0; i < len(tri) - 1; i++ {
		if len(tri[i]) != i + 1 {
			return -1, errors.New("Triangle is malformed!")
		}
		for j := 0; j < len(tri[i]); j++ {
			wg.Add(1)
			go tri[i:].sumTriangle(j, sums, func() { wg.Done() }) // Callback lets workgroup know that the goroutine has finished
		}
	}

	wg.Wait()

	return *least, nil
}

func main() {

	LeastTriangle(nil)
}