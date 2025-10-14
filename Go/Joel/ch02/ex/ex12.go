package exercise

import (
	"fmt"

	alg "github.com/PrinceNoir23/IntroToAlgorithms/Go/Joel/ch02/algorithms"
)

func Ex1_2() {
	arr := []int{31, 41, 59, 26, 41, 58}
	res := alg.InsertionSortDecreasing(arr, true)
	fmt.Printf("Sorted array: %v\n", res)
}
