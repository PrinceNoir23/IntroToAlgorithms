package main

import (
	// prob "github.com/PrinceNoir23/IntroToAlgorithms/Go/Joel/ch01/prob"
	ex2 "github.com/PrinceNoir23/IntroToAlgorithms/Go/Joel/ch02/ex"
	ex1 "github.com/PrinceNoir23/IntroToAlgorithms/Go/Joel/ch01/ex"
	"fmt"

	alg "github.com/PrinceNoir23/IntroToAlgorithms/Go/Joel/ch02/algorithms"
)

func main() {
	ex1.Ex1_2()
	arr := []int{5, 2, 4, 6, 1, 3}
	res := alg.InsertionSort(arr, true)
	fmt.Println(res)
	fmt.Println("-----------------------------------------------------------------------------------------")

	ex2.Ex1_1()
	fmt.Println("-----------------------------------------------------------------------------------------")
	ex2.Ex1_2()
	fmt.Println("-----------------------------------------------------------------------------------------")
	ex2.Ex1_3()
}
