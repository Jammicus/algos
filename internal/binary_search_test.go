package internal_test

import (
	algos "algos/internal"
	"testing"
)

func TestBinarySearch(t *testing.T) {

	var inputs = []struct {
		Input   []int
		Integer int
		Answer  int
	}{
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 9, 9},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 5, 5},
		{[]int{0, 1, 3, 5, 7, 9, 11}, 11, 6},
		{[]int{0, 1, 1, 5, 7, 9, 11}, 1, 2},
		{[]int{0, 1, 1, 5, 7, 9, 11, 15, 19, 22}, 19, 8},
	}

	var inputsB = []struct {
		Input   []float64
		Integer float64
		Answer  float64
	}{
		{[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 9, 9},
		{[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 5, 5},
		{[]float64{0, 1, 3, 5, 7, 9, 11}, 11, 6},
		{[]float64{0, 1, 1, 5, 7, 9, 11}, 1, 2},
		{[]float64{0, 1, 1, 5, 7, 9, 11, 15, 19, 22}, 19, 8},
	}

	for _, cases := range inputs {
		result := algos.BinarySearch(cases.Input, cases.Integer)
		algos.AssertEqual(t, cases.Answer, result)

	}
	for _, cases := range inputsB {
		result := algos.BinarySearch(cases.Input, cases.Integer)
		algos.AssertEqual(t, cases.Answer, result)

	}
}

func BenchmarkBinarySearch10EntryMiddle(b *testing.B) {

	inputs := []int{}
	for i := 0; i < 10; i++ {

		inputs = append(inputs, i)
	}
	for i := 0; i < b.N; i++ {
		algos.BinarySearch(inputs, len(inputs)/2)
	}

}

func BenchmarkBinarySearch100EntryMiddle(b *testing.B) {

	inputs := []int{}
	for i := 0; i < 100; i++ {

		inputs = append(inputs, i)
	}
	for i := 0; i < b.N; i++ {
		algos.BinarySearch(inputs, len(inputs)/2)
	}

}

func BenchmarkBinarySearch1000EntryMiddle(b *testing.B) {

	inputs := []int{}
	for i := 0; i < 1000; i++ {

		inputs = append(inputs, i)
	}
	for i := 0; i < b.N; i++ {
		algos.BinarySearch(inputs, len(inputs)/2)
	}

}

func BenchmarkBinarySearch10000EntryMiddle(b *testing.B) {

	inputs := []int{}
	for i := 0; i < 10000; i++ {

		inputs = append(inputs, i)
	}
	for i := 0; i < b.N; i++ {
		algos.BinarySearch(inputs, len(inputs)/2)
	}

}

func BenchmarkBinarySearch100000EntryMiddle(b *testing.B) {

	inputs := []int{}
	for i := 0; i < 100000; i++ {

		inputs = append(inputs, i)
	}
	for i := 0; i < b.N; i++ {
		algos.BinarySearch(inputs, len(inputs)/2)
	}

}

func BenchmarkBinarySearch10EntryIndex0(b *testing.B) {

	inputs := []int{}
	for i := 0; i < 100; i++ {

		inputs = append(inputs, i)
	}
	for i := 0; i < b.N; i++ {
		algos.BinarySearch(inputs, 0)
	}

}

func BenchmarkBinarySearch100EntryIndex0(b *testing.B) {

	inputs := []int{}
	for i := 0; i < 100; i++ {

		inputs = append(inputs, i)
	}
	for i := 0; i < b.N; i++ {
		algos.BinarySearch(inputs, 0)
	}

}

func BenchmarkBinarySearch1000EntryIndex0(b *testing.B) {

	inputs := []int{}
	for i := 0; i < 1000; i++ {

		inputs = append(inputs, i)
	}
	for i := 0; i < b.N; i++ {
		algos.BinarySearch(inputs, 0)
	}

}

func BenchmarkBinarySearch10000EntryIndex0(b *testing.B) {

	inputs := []int{}
	for i := 0; i < 10000; i++ {

		inputs = append(inputs, i)
	}
	for i := 0; i < b.N; i++ {
		algos.BinarySearch(inputs, 0)
	}

}

func BenchmarkBinarySearch100000EntryIndex0(b *testing.B) {

	inputs := []int{}
	for i := 0; i < 100000; i++ {

		inputs = append(inputs, i)
	}
	for i := 0; i < b.N; i++ {
		algos.BinarySearch(inputs, 0)
	}

}

func BenchmarkBinarySearch10EntryEnd(b *testing.B) {

	inputs := []int{}
	for i := 0; i < 1000; i++ {

		inputs = append(inputs, i)
	}
	for i := 0; i < b.N; i++ {
		algos.BinarySearch(inputs, len(inputs)-1)
	}

}

func BenchmarkBinarySearch100EntryEnd(b *testing.B) {

	inputs := []int{}
	for i := 0; i < 100; i++ {

		inputs = append(inputs, i)
	}
	for i := 0; i < b.N; i++ {
		algos.BinarySearch(inputs, len(inputs)-1)
	}

}

func BenchmarkBinarySearch1000EntryEnd(b *testing.B) {

	inputs := []int{}
	for i := 0; i < 1000; i++ {

		inputs = append(inputs, i)
	}
	for i := 0; i < b.N; i++ {
		algos.BinarySearch(inputs, len(inputs)-1)
	}

}

func BenchmarkBinarySearch10000EntryEnd(b *testing.B) {

	inputs := []int{}
	for i := 0; i < 10000; i++ {

		inputs = append(inputs, i)
	}
	for i := 0; i < b.N; i++ {
		algos.BinarySearch(inputs, len(inputs)-1)
	}

}

func BenchmarkBinarySearch100000EntryEnd(b *testing.B) {

	inputs := []int{}
	for i := 0; i < 100000; i++ {

		inputs = append(inputs, i)
	}
	for i := 0; i < b.N; i++ {
		algos.BinarySearch(inputs, len(inputs)-1)
	}

}
