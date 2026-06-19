package main

// LeetCode #1196: How Many Apples Can You Put into the Basket
// https://leetcode.com/problems/how-many-apples-can-you-put-into-the-basket/
// Difficulty: Easy [Paid]
// Time: O(n log n) | Space: O(log n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxNumberOfApples([]int{100, 200, 150, 1000}))          // 4
	fmt.Println(maxNumberOfApples([]int{900, 950, 800, 1000, 700, 800})) // 5
}

// LeetCode submission: maxNumberOfApples
func maxNumberOfApples(weight []int) int {
	sort.Ints(weight)
	sum := 0
	for i, w := range weight {
		sum += w
		if sum > 5000 {
			return i
		}
	}
	return len(weight)
}
