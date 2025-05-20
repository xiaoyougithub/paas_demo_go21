package main

import "fmt"

// bubbleSort implements the bubble sort algorithm
// It sorts the input slice in ascending order
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// Compare adjacent elements
			if arr[j] > arr[j+1] {
				// Swap if they are in the wrong order
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	// Create a sample integer array
	numbers := []int{64, 34, 25, 12, 22, 11, 90}
	
	// Print the array before sorting
	fmt.Println("Array before sorting:")
	fmt.Println(numbers)
	
	// Call bubbleSort function to sort the array
	bubbleSort(numbers)
	
	// Print the array after sorting
	fmt.Println("Array after sorting:")
	fmt.Println(numbers)
}