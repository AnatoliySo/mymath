package mymath

import "math"

func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func Abs(x float64) float64 {
	return math.Abs(x)
}

func Pow(x, y float64) float64 {
	return math.Pow(x, y)
}

func Max(x, y float64) float64 {
	return math.Max(x, y)
}

func Yn(y int, x float64) float64 {
	return math.Yn(y, x)
}
