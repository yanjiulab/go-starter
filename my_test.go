package main

import "testing"

// calc.go
func Add(a, b int) int {
	return a + b
}

// calc_test.go
func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Fatal("加法错误")
	}
}

func TestAdd2(t *testing.T) {
	if Add(1, 23) != 24 {
		t.Fatal("加法错误")
	}
}
