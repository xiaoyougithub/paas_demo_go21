package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func binarySearch(arr []int, start int, end int, target int) int {
	if start >= end {
		if target > arr[start] {
			return start + 1
		} else {
			return start
		}
	}
	
	mid := (start + end) / 2
	if arr[mid] == target {
		return mid + 1
	}
	
	if target < arr[mid] {
		return binarySearch(arr, start, mid-1, target)
	} else {
		return binarySearch(arr, mid+1, end, target)
	}
}

func binaryInsertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		temp := arr[i]
		pos := binarySearch(arr, 0, i-1, temp)
		
		for j := i; j > pos; j-- {
			arr[j] = arr[j-1]
		}
		
		arr[pos] = temp
	}
}

func main() {
	data1 := []int{64, 34, 25, 12, 22, 11, 90}
	data2 := []int{64, 34, 25, 12, 22, 11, 90}
	
	fmt.Println("原始数组:", data1)
	
	bubbleSort(data1)
	fmt.Println("冒泡排序结果:", data1)
	
	binaryInsertionSort(data2)
	fmt.Println("二分插入排序结果:", data2)
}