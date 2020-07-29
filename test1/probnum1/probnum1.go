package main

import "fmt"

func main() {
	// Input Value
	arr := []int32{1, 2, 3, 4, 4, 3, 2, 1, 6, 5}

	// Initialize the Hash Map
	hashMap := make(map[int32]int32)

	/*
		Loop the Array and then input to hashmap
		HashMap useful for collecting and count the number
	*/
	for _, v := range arr {
		hashMap[v]++
	}

	// Loop To Detect the duplicates number -> can detect multiple numbers
	for i, v := range hashMap {
		if v > 1 {
			fmt.Println(i, " is Duplicate")
		}
	}

}
