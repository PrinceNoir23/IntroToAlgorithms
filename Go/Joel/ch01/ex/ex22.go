package exercise

import (
	"fmt"
	"math"
)

func Ex2_2() {
	fmt.Println("The insertion sort algorithm has a time complexity of c1(n^2) in the worst case.")
	fmt.Println("The merge sort algorithm has a time complexity of c2(nlogn) in the worst case.")
	fmt.Println("Where c1 and c2 are constants that depend on the implementation of the algorithms.")
	fmt.Println("In general, the merge sort beat the insertion sort for large inputs.")
	for n := 0; n <= 100; n += 1{
		fmt.Printf("n = %d, insertion sort = %f, merge sort = %f\n", n, insertionSort(8, make([]int, n)), mergeSort(64, make([]int, n)))
	}
	fmt.Println("Above n aproximately 50, the merge sort is always faster than the insertion sort for this implementation.")
}

func insertionSort(constant uint8, n []int) float64 {
	return float64(constant) * math.Pow(float64(len(n)), 2)
}

func mergeSort(constant uint8, n []int) float64 {
	return float64(constant) * float64(len(n)) * math.Log2(float64(len(n)))
}
