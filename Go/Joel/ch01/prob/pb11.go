package prob

import (
	"fmt"
	"math"

)


func Prob1_1() {
	times := []float64{1000000, 60000, 3600000, 86400000, 2629800000, 31557600000, 3155760000000}
	fmt.Println("Comparison of running times")
	fmt.Println("f(n)\t=1\tsecond\tminute\thour\tday\tmonth\tyear\tcentury")
	fmt.Println("------------------------------------------------------------------------")
	fmt.Printf("result is %f\n", log(0, times[0]))
	fmt.Println("lg n\\\\\\\\")
	fmt.Println("sqrt n\\\\\\\\")
	fmt.Println("n\\\\\\\\")
	fmt.Println("n lg n\\\\\\\\")
	fmt.Println("n^2\\\\\\\\")
	fmt.Println("n^3\\\\\\\\")
	fmt.Println("2^n\\\\\\\\")
	fmt.Println("n!\\\\\\\\")
}
func log(n float64, time float64) float64 {
    if time <= 0 {
        return 0
    }
    if n < 1 {
        n = 1
    }
    for n*math.Log2(n) < time {
        n += 1
    }
    return n
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