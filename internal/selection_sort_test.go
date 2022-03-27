package internal_test

import (
	algos "algos/internal"
	"testing"
)

func TestSelectionsort(t *testing.T) {

	type testInputs[T algos.Number] struct {
		Input  []T
		Answer []T
	}

	var inputs = []testInputs[int]{
		{[]int{5, 1, 8, 0, 4, 5, 6, 7, 11, 9}, []int{0, 1, 4, 5, 5, 6, 7, 8, 9, 11}},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{-15, 15, -45, 0, 2, 3, 22}, []int{-45, -15, 0, 2, 3, 15, 22}},
	}

	var inputsB = []testInputs[float64]{
		{[]float64{5, 1, 8, 0, 4, 5, 6, 7, 11, 9}, []float64{0, 1, 4, 5, 5, 6, 7, 8, 9, 11}},
		{[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]float64{-15, 15, -45, 0, 2, 3, 22}, []float64{-45, -15, 0, 2, 3, 15, 22}},
		{[]float64{0.123, 0.1, 0.7, 0.222, 222, 7, 0.001}, []float64{0.001, 0.1, 0.123, 0.222, 0.7, 7, 222}},
	}

	for _, cases := range inputs {
		result := algos.SelectionSort(cases.Input)
		algos.AssertEqualSlices(t, cases.Answer, result)

	}
	for _, cases := range inputsB {
		result := algos.SelectionSort(cases.Input)
		algos.AssertEqualSlices(t, cases.Answer, result)

	}
}

func BenchmarkSelectionsort10(b *testing.B) {

	inputs := algos.GenerateSlice(10)

	for i := 0; i < b.N; i++ {
		algos.SelectionSort(inputs)
	}

}

func BenchmarkSelectionsort100(b *testing.B) {

	inputs := algos.GenerateSlice(100)

	for i := 0; i < b.N; i++ {
		algos.SelectionSort(inputs)
	}

}

func BenchmarkSelectionsort1000(b *testing.B) {

	inputs := algos.GenerateSlice(1000)

	for i := 0; i < b.N; i++ {
		algos.SelectionSort(inputs)
	}

}

func BenchmarkSelectionSort10000(b *testing.B) {

	inputs := algos.GenerateSlice(10000)

	for i := 0; i < b.N; i++ {
		algos.SelectionSort(inputs)
	}

}

func BenchmarkSelectionSort100000(b *testing.B) {

	inputs := algos.GenerateSlice(100000)

	for i := 0; i < b.N; i++ {
		algos.SelectionSort(inputs)
	}

}
