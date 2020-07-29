package main

import "fmt"

func main() {
	// Input Value
	arr := []int32{1, 2, 3}

	// Initialize the temp array
	arrTemp := []int32{}

	/*
		Using the concept of recursive tree
		we interate every number and loop over again the same number
		and print if the length is 2 (pair)
	*/
	recurPair(arr, arrTemp, 0, int32(len(arr)))
}

func recurPair(arr []int32, arrTemp []int32, i int32, n int32) { // var i to mark the current iteration
	if len(arrTemp) == 2 { // check if the temp array got a pair, then return and pop the last array
		fmt.Println(arrTemp)
		return
	}

	for j := i; j < n; j++ {
		arrTemp = append(arrTemp, arr[j]) //push to the temp array, and then check the pair again in the recurPair function
		recurPair(arr, arrTemp, j, int32(len(arr)))
		arrTemp = arrTemp[:len(arrTemp)-1]
	}
}
