package main

// LeetCode #2677: Chunk Array
// https://leetcode.com/problems/chunk-array/
// Difficulty: Easy
// Time: O(n) | Space: O(n)
// Note: JavaScript problem, adapted to Go. Splits array into chunks of given size.

import "fmt"

func main() {
	fmt.Println(ChunkArray([]int{1, 2, 3, 4, 5}, 1))
	fmt.Println(ChunkArray([]int{1, 9, 6, 3, 2}, 3))
}

func ChunkArray(arr []int, size int) [][]int {
	var result [][]int
	for i := 0; i < len(arr); i += size {
		end := i + size
		if end > len(arr) {
			end = len(arr)
		}
		result = append(result, arr[i:end])
	}
	return result
}
