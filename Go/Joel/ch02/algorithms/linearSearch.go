package algorithms

func LinearSearch(arr []int, val int) int {
	for i := range arr {
		if arr[i] == val {
			return i
		}
	}
	return -1
}

// Loop Invariant

// - If the function has not returned yet and the loop is at index i , then val is not equal to any element in arr[0..i-1] .
// Why it holds

// - Initialization: Before the first iteration ( i = 0 ), the set arr[0..i-1] is empty, so the invariant is trivially true.
// - Maintenance: At iteration i , you check arr[i] == val .
//   - If true, you return i immediately and stop (correct index found).
//   - If false, you continue; thus val is still not in arr[0..i] , which means at the next iteration i+1 the invariant (“not in arr[0..i] ”) remains true for the new prefix arr[0..(i+1)-1] .
// - Termination: If the loop completes without returning, the invariant implies val is not present in any arr[0..len(arr)-1] , so returning -1 is correct.
// Restating formally

// - For every iteration where the loop variable has value k , if the function has not returned yet, then for all j with 0 ≤ j < k , arr[j] ≠ val .
// Context to this code

// - for i := range arr iterates i = 0, 1, ..., len(arr)-1 . The invariant is evaluated at the top of each iteration and preserved by the conditional check and continue behavior.
