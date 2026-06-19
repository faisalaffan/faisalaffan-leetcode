package main

// LeetCode #1086: High Five
// https://leetcode.com/problems/high-five/
// Difficulty: Easy [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	items := [][]int{
		{1, 91}, {1, 92}, {2, 93}, {2, 97}, {1, 60},
		{2, 77}, {1, 65}, {1, 87}, {1, 100}, {2, 100},
		{2, 76},
	}
	fmt.Println(highFive(items)) // [[1,87],[2,88]]
}

// LeetCode submission: highFive
func highFive(items [][]int) [][]int {
	scores := make(map[int][]int)
	for _, item := range items {
		id, score := item[0], item[1]
		scores[id] = append(scores[id], score)
	}
	var ids []int
	for id := range scores {
		ids = append(ids, id)
	}
	sort.Ints(ids)
	var ans [][]int
	for _, id := range ids {
		s := scores[id]
		sort.Sort(sort.Reverse(sort.IntSlice(s)))
		sum := 0
		for i := 0; i < 5; i++ {
			sum += s[i]
		}
		ans = append(ans, []int{id, sum / 5})
	}
	return ans
}
