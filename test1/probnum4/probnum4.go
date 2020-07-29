package main

import "fmt"

func main() {
	// Input Value
	arr := []int32{-6, 4, -5, 8, -10, 0, 8}

	/*
		The Max Product function will iterate every number,
		compare the max value in every number and save which number that has the max value
	*/
	maxProduct, subArr := maxProduct(arr)

	fmt.Printf("the maximum product sub-array is %v having products %v", subArr, maxProduct)
}

func maxProduct(arr []int32) (int32, []int32) {
	// Initialize the helper variable
	tempArr := []int32{}

	// End Process if Array is nil
	if len(arr) == 0 {
		return -1, tempArr // There is no array(Nil)
	}

	// Initialize the point32er variable
	currentMax := arr[0]
	currentMin := arr[0]
	result := arr[0]

	// Iterate the Array
	for i := 1; i < len(arr); i++ {
		tempNum := currentMax                                               // Store the old current max, for comparing the mininum value
		currentMax = max(arr[i], max(currentMax*arr[i], currentMin*arr[i])) // get the max value based on the current turn(i)
		currentMin = min(arr[i], min(tempNum*arr[i], currentMin*arr[i]))    // get the min value, for comparing and get the current max value

		if currentMax > result {
			tempArr = append(tempArr, arr[i]) // save the selected array value
			result = currentMax
		}
	}
	return result, tempArr
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}
