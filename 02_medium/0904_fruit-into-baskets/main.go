package main

// LeetCode #904: Fruit Into Baskets
// https://leetcode.com/problems/fruit-into-baskets/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(FruitIntoBaskets([]int{1, 2, 1}))
	fmt.Println(FruitIntoBaskets([]int{0, 1, 2, 2}))
	fmt.Println(FruitIntoBaskets([]int{1, 2, 3, 2, 2}))
}

// Time: O(n) | Space: O(1)
func FruitIntoBaskets(fruits []int) int {
	cnt := make(map[int]int)
	left, ans := 0, 0

	for right, fruit := range fruits {
		cnt[fruit]++
		for len(cnt) > 2 {
			cnt[fruits[left]]--
			if cnt[fruits[left]] == 0 {
				delete(cnt, fruits[left])
			}
			left++
		}
		if right-left+1 > ans {
			ans = right - left + 1
		}
	}

	return ans
}
