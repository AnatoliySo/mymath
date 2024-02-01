package main

import (
	"testing"
)

func TestSqrt(t *testing.T) {
	const in, out = 4, 2
	if x := Sqrt(in); x != out {
		t.Errorf("Sqrt(%v) = %v, want %v", in, x, out)
	}
}

func TestAbs(t *testing.T) {
	const in, out = -5, 5
	if x := Abs(in); x != out {
		t.Errorf("Abs(%v) = %v, want %v", in, x, out)
	}
}

func TestPow(t *testing.T) {
	const in1, in2, out = 2, 3, 8
	if x := Pow(in1, in2); x != out {
		t.Errorf("Pow(%v, %v) = %v, want %v", in1, in2, x, out)
	}
}
