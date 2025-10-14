package main

import (
	// prob "github.com/PrinceNoir23/IntroToAlgorithms/Go/Joel/ch01/prob"
	// ex "github.com/PrinceNoir23/IntroToAlgorithms/Go/Joel/ch01/ex"
	"fmt"

	alg "github.com/PrinceNoir23/IntroToAlgorithms/Go/Joel/ch02/algorithms"
)

func main() {
	arr := []int{5, 2, 4, 6, 1, 3}
	res := alg.InsertionSort(arr)
	fmt.Println(res)
}
