package main

import (
	"testing"
	"fmt"
	"time"
	//"math/rand"
	//"errors"
)

func TestReturns(t *testing.T) {

	// Test case 1: simple
	triangle := [][]int{[]int{1}, []int{1, 2}, []int{1, 2, 3}, []int{-1, -2, -3, -4}}

	fmt.Println("Running test case 1...")
	start := time.Now()
	if v, err := LeastTriangle(triangle); err != nil {
		t.Errorf(err.Error())
	} else if v != -4 {
		t.Errorf("Expected: -4, Got: %v", v)
	} else {
		fmt.Printf("Correct! Least triangle: %v\n", v)
	}
	end := time.Since(start)
	fmt.Printf("Runtime: %v\n\n", end)

	fmt.Println("Running test case 2...")
	// Test case 2: more complex
	triangle = [][]int{[]int{15}, []int{-14, -7}, []int{20, -13, -5}, []int{-3, 8, 23, -26}, []int{1, -4, -5, -18, 5}, []int{-16, 31, 2, 9, 28, 3}}

	start = time.Now()
	if v, err := LeastTriangle(triangle); err != nil {
		t.Errorf(err.Error())
	} else if v != -42 {
		t.Errorf("Expected: -42, Got: %v", v)
	} else {
		fmt.Printf("Correct! Least triangle: %v\n", v)
	}
	end = time.Since(start)
	fmt.Printf("Runtime: %v\n\n", end)

	fmt.Println("Running test case 3...")
	// Test 3: Really really big triangle runtime test, if you want to max out your CPU for a few minutes...
	/*var bigTriangle [][]int
	for i := 0; i < 1000; i++ {
		var t []int
		for j:= 0; j < i+1; j++ {
			t = append(t, rand.Int())
		}
		bigTriangle = append(bigTriangle, t)
	}

	start = time.Now()
	if v, err := LeastTriangle(bigTriangle); err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Printf("Least triangle: %v\n", v)
	}
	end = time.Since(start)
	fmt.Printf("Runtime: %v\n\n", end)*/
}

func TestError(t *testing.T) {

	// Test case 1: nil triangle
	triangle := make([][]int, 0)

	fmt.Println("Running test error case 1...")
	if _, err := LeastTriangle(triangle); err != nil {
		if err.Error() == "Triangle is nil!" {
			fmt.Println(err.Error())
			fmt.Println()
		} else {
			t.Errorf("Expected: \"Triangle is nil!\", Got: %v", err.Error())
		}
	} else {
		t.Errorf("Should have received error due to empty triangle!")
	}

	// Test case 2: first level of triangle is null
	triangle = [][]int{nil, []int{}}

	fmt.Println("Running test error case 2...")
	if _, err := LeastTriangle(triangle); err != nil {
		if err.Error() == "First level is nil!" {
			fmt.Println(err.Error())
			fmt.Println()
		} else {
			t.Errorf("Expected: \"First level is nil!\", Got: %v", err.Error())
		}
	} else {
		t.Errorf("Should have received error due to malformed triangle!")
	}

	// Test case 3: no second level
	triangle = [][]int{[]int{15}}

	fmt.Println("Running test error case 3...")
	if _, err := LeastTriangle(triangle); err != nil {
		if err.Error() == "Triangle has less than 2 levels!" {
			fmt.Println(err.Error())
			fmt.Println()
		} else {
			t.Errorf("Expected: \"Triangle has less than 2 levels!\", Got: %v", err.Error())
		}
	} else {
		t.Errorf("Should have received error due to only one level!")
	}

	// Test case 4: first level contains != 1 int
	triangle = [][]int{[]int{15, 35},[]int{45,50}}

	fmt.Println("Running test error case 4...")
	if _, err := LeastTriangle(triangle); err != nil {
		if err.Error() == "First level of triangle doest not contain one int!" {
			fmt.Println(err.Error())
			fmt.Println()
		} else {
			t.Errorf("Expected: \"First level of triangle doest not contain one int!\", Got: %v", err.Error())
		}
	} else {
		t.Errorf("Should have received error due to wrong first level size!")
	}

	// Test case 5: second level doesn't contain two ints
	triangle = [][]int{[]int{15}, []int{1, 2, 3}}

	fmt.Println("Running test error case 5...")
	if _, err := LeastTriangle(triangle); err != nil {
		if err.Error() == "Second level of triangle does not contain two ints!" {
			fmt.Println(err.Error())
			fmt.Println()
		} else {
			t.Errorf("Expected: \"Second level of triangle does not contain two ints!\", Got: %v", err.Error())
		}
	} else {
		t.Errorf("Should have received error due to wrong second level size!")
	}

	// Test case 6: malformed triangle
	triangle = [][]int{[]int{15}, []int{1, 2}, []int{20, -13, -5, 7}, []int{-3, 8, 23, -26}, []int{1, -4, -5, -18, 5}, []int{-16, 31, 2, 9, 28, 3}}

	fmt.Println("Running test error case 6...")
	if _, err := LeastTriangle(triangle); err != nil {
		if err.Error() == "Triangle is malformed!" {
			fmt.Println(err.Error())
			fmt.Println()
		} else {
			t.Errorf("Expected: \"Triangle is malformed!\", Got: %v", err.Error())
		}
	} else {
		t.Errorf("Should have received error due to malformed triangle!")
	}
}