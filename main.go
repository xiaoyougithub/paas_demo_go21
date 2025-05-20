package main

import "fmt"

// bubbleSort implements the bubble sort algorithm
// It sorts the slice in ascending order
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Last i elements are already sorted
		for j := 0; j < n-i-1; j++ {
			// Swap if the element found is greater than the next element
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	// Create a sample integer array
	numbers := []int{64, 34, 25, 12, 22, 11, 90}
	
	// Print the array before sorting
	fmt.Println("Before sorting:", numbers)
	
	// Call bubbleSort function to sort the array
	bubbleSort(numbers)
	
	// Print the array after sorting
	fmt.Println("After sorting:", numbers)
}