package main

// LeetCode #3017: Count the Number of Houses at a Certain Distance II
// https://leetcode.com/problems/count-the-number-of-houses-at-a-certain-distance-ii/
// Difficulty: Hard

import "fmt"

func countOfPairs(n int, x int, y int) []int64 {
	if x > y {
		x, y = y, x
	}
	diff := make([]int64, n+1)
	add := func(l, r int, val int64) {
		if l > r { return }
		diff[l] += val
		if r+1 <= n { diff[r+1] -= val }
	}
	for i := 1; i <= n; i++ {
		if x+1 >= y {
			add(1, n-i, 2)
			continue
		}
		if i <= x {
			k := (x + y + 1) / 2
			add(1, k-i, 2)
			add(x-i+2, x-i+y-k, 2)
			add(x-i+1, x-i+1+n-y, 2)
		} else if i < (x+y)/2 {
			k := i + (y-x+1)/2
			add(1, k-i, 2)
			add(i-x+2, i-x+y-k, 2)
			add(i-x+1, i-x+1+n-y, 2)
		} else {
			add(1, n-i, 2)
		}
	}
	ans := make([]int64, n)
	cur := int64(0)
	for i := 1; i <= n; i++ {
		cur += diff[i]
		ans[i-1] = cur
	}
	return ans
}

func main() {
	fmt.Println(countOfPairs(3, 1, 3))
	fmt.Println(countOfPairs(5, 2, 4))
}
