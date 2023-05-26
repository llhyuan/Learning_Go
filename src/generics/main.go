package main

import "fmt"

type Number interface {
	int64 | float64
}

func main() {
	ints := map[string]int64{
		"first":  45,
		"second": 85,
	}

	floats := map[string]float64{
		"first":  45.97,
		"second": 42.90,
	}

	fmt.Printf("The sums are %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Generic sums: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

	fmt.Printf("Contrained sums: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))

}

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var sum int64
	for _, v := range m {
		sum += v
	}
	return sum
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var sum float64
	for _, v := range m {
		sum += v
	}
	return sum
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
