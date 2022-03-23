package main

import (
	algos "algos/internal"
	"fmt"
)

func main() {
	fmt.Println("Hello world")

	list := algos.GenerateSlice(10000)

	fmt.Println(algos.BinarySearch([]int{1, 2, 3, 4, 5, 6, 7}, 4))
	fmt.Println(algos.SelectionSort(list))
	fmt.Println(algos.Quicksort(list))
	fmt.Println(algos.Mergesort(list))
}
