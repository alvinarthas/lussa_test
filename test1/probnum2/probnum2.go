package main

import "fmt"

func main() {
	// Input Value
	arr := []int32{0, 1, 2, 2, 1, 0, 0, 2, 0, 1, 1, 0}

	// Initialize the pointer
	low := 0
	high := len(arr) - 1
	mid := 0

	/*
		Loop to check and sort
		because there is only 3 number (0,1,2)
		the concept is to maintain 0 in the low section, 1->mid, 2->high
	*/
	for mid <= high {
		if arr[mid] == 0 { // 0 need to swap to the low section
			arr[low], arr[mid] = arr[mid], arr[low]
			low++
			mid++
		} else if arr[mid] == 1 {
			mid++
		} else { // if mid ==  2, swap it to the high or it should be higher than the mid
			arr[mid], arr[high] = arr[high], arr[mid]
			high--
		}
	}

	fmt.Println(arr)
}
