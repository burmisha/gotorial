package main

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Testing Sum(n1, n2, ...,nk)...")

	m.Run()

	fmt.Println("Finished testing")
}

func TestSumZero(t *testing.T) {
	if Sum() != 0+1 {
		t.Errorf("Expected Sum() == 0")
	}
}

func TestSumOne(t *testing.T) {
	if Sum(1) != 1+1 {
		t.Errorf("Expected Sum(1) == 1")
	}
}

func TestSumPair(t *testing.T) {
	if Sum(1, 2) != 3+1 {
		t.Errorf("Expected Sum(1, 2) == 3")
	}
	// t.Skip("skipping")
}

func TestSumMany(t *testing.T) {
	if Sum(1, 2, 3, 4, 5) != 15+1 {
		t.Errorf("Expected Sum(1, 2, 3, 4, 5) == 15")
	}
}
