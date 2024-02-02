package mymath

import (
	"math"
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

func TestMax(t *testing.T) {
	const in1, in2, out = 2, 3, 3
	if x := Max(in1, in2); x != out {
		t.Errorf("Max(%v, %v) = %v, want %v", in1, in2, x, out)
	}
}

func TestYn(t *testing.T) {
	const in1, in2 = 1, 2
	out := math.Yn(in1, in2)
	if x := Yn(in1, in2); x != out {
		t.Errorf("Yn(%v, %v) = %v, want %v", in1, in2, x, out)
	}
}

func TestCeil(t *testing.T) {
	const in, out = 2.7, 3
	if x := Ceil(in); x != out {
		t.Errorf("Ceil(%v) = %v, want %v", in, x, out)
	}
}

func TestFloor(t *testing.T) {
	const in, out = 2.7, 2
	if x := Floor(in); x != out {
		t.Errorf("Floor(%v) = %v, want %v", in, x, out)
	}
}

func TestMin(t *testing.T) {
	const in1, in2, out = 2, 3, 2
	if x := Min(in1, in2); x != out {
		t.Errorf("Min(%v, %v) = %v, want %v", in1, in2, x, out)
	}
}
