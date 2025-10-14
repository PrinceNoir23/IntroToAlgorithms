package algorithms

func InsertionSort(arr []int) []int {
	// the loop will start at the second element of the array
	for j := 1; j < len(arr); j++ {
		// we go to the KEY in the array 
		key := arr[j]
		// we compare the key with the elements before it
		i := j - 1
		for i >= 0 && arr[i] > key {
			// then we shift the element to the right and we now look for a smaller element
			arr[i+1] = arr[i]
			i--
		}
		// we insert the key in the correct position
		arr[i+1] = key
	}
	return arr
}