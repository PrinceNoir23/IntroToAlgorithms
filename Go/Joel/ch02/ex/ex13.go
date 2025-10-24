package exercise

import (
	"fmt"

	alg "github.com/PrinceNoir23/IntroToAlgorithms/Go/Joel/ch02/algorithms"
)

func Ex1_3() {
	arr := []int{31, 41, 59, 26, 41, 58}
	val := 41
	index := alg.LinearSearch(arr, val)
	if index == -1 {
		fmt.Printf("Value %d not found in array\n", val)
	} else {
		fmt.Printf("Value %d found at index %d\n", val, index)
	}
}
