package main

import (
	"testing"
	"fmt"
)

func TestAll(t *testing.T) {

	if v, err := LeastTriangle(top); err != nil {
		err.Error()
	} else if v != 30 {
		fmt.Errorf("Wrong value returned!")
	}

	if v, err := LeastTriangle(top); err != nil {
		err.Error()
	} else if v != -15 {
		fmt.Errorf("Wrong value returned!")
	}
}