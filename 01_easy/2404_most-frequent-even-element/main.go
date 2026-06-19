package main

// LeetCode #2404: Most Frequent Even Element
// https://leetcode.com/problems/most-frequent-even-element/
// Difficulty: Easy
// Time O(n) | Space O(n)

import "fmt"

func main() {
	fmt.Println(MostFrequentEvenElement([]int{0, 1, 2, 2, 4, 4, 1})) // 2
	fmt.Println(MostFrequentEvenElement([]int{4, 4, 4, 9, 2, 4}))    // 4
	fmt.Println(MostFrequentEvenElement([]int{1, 3, 5, 7}))           // -1
}

func MostFrequentEvenElement(nums []int) int {
	freq := map[int]int{}
	for _, n := range nums {
		if n%2 == 0 {
			freq[n]++
		}
	}
	if len(freq) == 0 {
		return -1
	}
	bestNum := -1
	bestCount := 0
	for n, c := range freq {
		if c > bestCount || (c == bestCount && n < bestNum) {
			bestNum = n
			bestCount = c
		}
	}
	return bestNum
}
