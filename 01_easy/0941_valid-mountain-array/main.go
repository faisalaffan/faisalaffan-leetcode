package main

// LeetCode #941: Valid Mountain Array
// https://leetcode.com/problems/valid-mountain-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(validMountainArray([]int{2, 1}))          // false
	fmt.Println(validMountainArray([]int{3, 5, 5}))       // false
	fmt.Println(validMountainArray([]int{0, 3, 2, 1}))    // true
	fmt.Println(validMountainArray([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})) // false
}

// validMountainArray checks if the array is a valid mountain array.
// Time: O(n). Space: O(1).
func validMountainArray(arr []int) bool {
	n := len(arr)
	if n < 3 {
		return false
	}
	i := 0
	// Climb up
	for i+1 < n && arr[i] < arr[i+1] {
		i++
	}
	// Peak cannot be first or last
	if i == 0 || i == n-1 {
		return false
	}
	// Climb down
	for i+1 < n && arr[i] > arr[i+1] {
		i++
	}
	return i == n-1
}
