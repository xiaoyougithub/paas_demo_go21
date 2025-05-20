package main

import "fmt"

// bubbleSort implements the bubble sort algorithm for sorting integer arrays
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Last i elements are already sorted
		for j := 0; j < n-i-1; j++ {
			// Compare adjacent elements and swap if they are in wrong order
			if arr[j] > arr[j+1] {
				// Swap arr[j] and arr[j+1]
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	// Create a sample integer array
	numbers := []int{64, 34, 25, 12, 22, 11, 90}
	
	// Print array before sorting
	fmt.Println("Array before sorting:", numbers)
	
	// Call bubbleSort to sort the array
	bubbleSort(numbers)
	
	// Print array after sorting
	fmt.Println("Array after sorting:", numbers)
}