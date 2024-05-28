package main

import "fmt"

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}
	floats := map[string]float64{
		"first":  34.12,
		"second": 12.98,
	}
	fmt.Printf("Non-Generic Sums: %v and %v\n",
		// SumInts(ints),
		// SumFloats(floats))
		// SumIntsOrFloats[string, int64](ints),
		// SumIntsOrFloats[string, float64](floats))
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats)) //这里不是必需，可以从参数推导
}

//定义类型约束，和TS一样的
type Number interface {
	int64 | float64
}

func SumIntsOrFloats[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}
