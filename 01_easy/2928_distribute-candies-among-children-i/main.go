package main

// LeetCode #2928: Distribute Candies Among Children I
// https://leetcode.com/problems/distribute-candies-among-children-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: distributeCandies
	fmt.Println(DistributeCandiesAmongChildrenI(5, 2)) // 3
	fmt.Println(DistributeCandiesAmongChildrenI(3, 3)) // 10
}

// Time: O(limit^2) | Space: O(1)
// LeetCode submission name: distributeCandies
func DistributeCandiesAmongChildrenI(n int, limit int) int {
	ways := 0
	for a := 0; a <= limit && a <= n; a++ {
		for b := 0; b <= limit && a+b <= n; b++ {
			c := n - a - b
			if c <= limit {
				ways++
			}
		}
	}
	return ways
}
