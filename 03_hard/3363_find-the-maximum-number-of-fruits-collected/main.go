package main

// LeetCode #3363: Find the Maximum Number of Fruits Collected
// https://leetcode.com/problems/find-the-maximum-number-of-fruits-collected/
// Difficulty: Hard
//
// Three children start from different corners and all end at (n-1,n-1).
// Child 1: (0,0)->(n-1,n-1) in n-1 moves → only possible on main diagonal.
// Child 2: (0,n-1)->(n-1,n-1) in upper triangle (i < j).
// Child 3: (n-1,0)->(n-1,n-1) in lower triangle (i > j).
// Since regions are disjoint, solve each independently.

import "fmt"

func main() {
	// Example: fruits=[[1,2,3],[4,5,6],[7,8,9]] -> 29
	fmt.Println(maxCollectedFruits([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))

	// n=1: [[5]] -> 5
	fmt.Println(maxCollectedFruits([][]int{{5}}))

	// n=2: [[1,2],[3,4]]
	// Child 1: 1+4 = 5
	// Child 2: (0,1)->(1,1): 2
	// Child 3: (1,0)->(1,1): 3
	// Total: 10
	fmt.Println(maxCollectedFruits([][]int{{1, 2}, {3, 4}}))

	// n=4 example from LeetCode
	fmt.Println(maxCollectedFruits([][]int{
		{1, 2, 3, 4},
		{5, 6, 8, 7},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}))
}

func maxCollectedFruits(fruits [][]int) int {
	n := len(fruits)
	if n == 1 {
		return fruits[0][0]
	}

	// Child 1: sum of main diagonal (only possible path with n-1 moves)
	ans := 0
	for i := 0; i < n; i++ {
		ans += fruits[i][i]
	}

	// Child 2: DP from (0, n-1) in upper triangle (i < j)
	// dp2[i][j] = max fruits collected from (i,j) to destination
	// Moving: (i+1, j-1), (i+1, j), (i+1, j+1)
	dp2 := make([][]int, n)
	for i := 0; i < n; i++ {
		dp2[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp2[i][j] = -1
		}
	}

	// Initialize valid end positions for child 2: one step before destination
	// Dest is (n-1, n-1). Valid prev: (n-2, n-2), (n-2, n-1)
	// But (n-2, n-2) is on diagonal (child 1 territory), so only (n-2, n-1)
	if n >= 2 {
		dp2[n-2][n-1] = fruits[n-2][n-1]
	}

	for i := n - 3; i >= 0; i-- {
		for j := n - 1; j > i; j-- {
			best := -1
			// From (i,j) we can go to (i+1, j-1), (i+1, j), (i+1, j+1)
			for _, dj := range []int{-1, 0, 1} {
				nj := j + dj
				if nj > i+1 && nj < n && dp2[i+1][nj] >= 0 {
					if dp2[i+1][nj] > best {
						best = dp2[i+1][nj]
					}
				}
			}
			if best >= 0 {
				dp2[i][j] = fruits[i][j] + best
			}
		}
	}

	if dp2[0][n-1] >= 0 {
		ans += dp2[0][n-1]
	}

	// Child 3: DP from (n-1, 0) in lower triangle (i > j)
	// Moving: (i-1, j+1), (i, j+1), (i+1, j+1)
	dp3 := make([][]int, n)
	for i := 0; i < n; i++ {
		dp3[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp3[i][j] = -1
		}
	}

	// Valid prev: (n-1, n-2)
	if n >= 2 {
		dp3[n-1][n-2] = fruits[n-1][n-2]
	}

	for j := n - 3; j >= 0; j-- {
		for i := n - 1; i > j; i-- {
			best := -1
			for _, di := range []int{-1, 0, 1} {
				ni := i + di
				if ni > j+1 && ni < n && dp3[ni][j+1] >= 0 {
					if dp3[ni][j+1] > best {
						best = dp3[ni][j+1]
					}
				}
			}
			if best >= 0 {
				dp3[i][j] = fruits[i][j] + best
			}
		}
	}

	if dp3[n-1][0] >= 0 {
		ans += dp3[n-1][0]
	}

	return ans
}
