package main

import "fmt"

// bubbleSort implements the bubble sort algorithm to sort an integer slice
// It repeatedly steps through the slice, compares adjacent elements and swaps
// them if they are in the wrong order until the slice is fully sorted
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Flag to optimize if no swaps occur in a pass
		swapped := false
		for j := 0; j < n-i-1; j++ {
			// Compare adjacent elements
			if arr[j] > arr[j+1] {
				// Swap elements if they are in wrong order
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		// If no swaps occurred in this pass, array is sorted
		if !swapped {
			break
		}
	}
}

func main() {
	// Create a sample integer array
	numbers := []int{64, 34, 25, 12, 22, 11, 90}
	
	// Print array before sorting
	fmt.Println("Before sorting:", numbers)
	
	// Call bubbleSort to sort the array
	bubbleSort(numbers)
	
	// Print array after sorting
	fmt.Println("After sorting:", numbers)
}