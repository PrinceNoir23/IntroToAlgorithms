package prob

import (
	"fmt"

)


func Prob1_1() {
	fmt.Println("Comparison of running times")
	fmt.Println("f(n) | second | minute | hour | day | month | year | century")
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("lg n | 2^(10^6) | 2^(6*10^7) | 2^(3.6*10^9) | 2^(8.64*10^10) | 2^(2.6298*10^12) | 2^(3.15576*10^14)")
	fmt.Println("sqrt n| 10^12 | 36*10^14 | 12.96*10^18 | 746496*10^16 | 6718464*10^18 | 994519296*10^18| 995827586973696*10^16 | *10^16 ")
	fmt.Println("the idea is to set f(n)=timeInMiliseconds and then solve for n")
	fmt.Println("the other solutions are in https://ita.skanev.com/01/problems/01.html")
}

// func sqrt(n float64, time float64) float64 {
// 	return float64(math.Sqrt(float64(n)))
// }
// func n(n float64, time float64) float64 {
// 	return n
// }
// func nLgN(n float64, time float64) float64 {
// 	return n * log(n)
// }
// func nSquared(n float64, time float64) float64 {
// 	return n * n
// }
// func nCubed(n float64, time float64) float64 {
// 	return n * n * n
// }
// func twoToTheN(n float64, time float64) float64 {
// 	return float64(math.Pow(2, float64(n)))
// }
// func factorial(n float64, time float64) float64 {
// 	return float64(math.Gamma(float64(n + 1)))
// }