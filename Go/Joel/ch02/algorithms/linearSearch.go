package algorithms

func LinearSearch(arr []int, val int) int {
	for i := range arr {
		if arr[i] == val {
			return i
		}
	}
	return -1
}