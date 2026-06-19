package main

// LeetCode #2326: Spiral Matrix IV
// https://leetcode.com/problems/spiral-matrix-iv/
// Difficulty: Medium
// Time: O(m * n) | Space: O(1)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result[i][j] = -1
		}
	}

	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dir := 0
	r, c := 0, 0

	for head != nil {
		result[r][c] = head.Val
		head = head.Next

		nr, nc := r+dirs[dir][0], c+dirs[dir][1]
		if nr < 0 || nr >= m || nc < 0 || nc >= n || result[nr][nc] != -1 {
			dir = (dir + 1) % 4
			nr, nc = r+dirs[dir][0], c+dirs[dir][1]
		}
		r, c = nr, nc
	}
	return result
}

func main() {
	// Test case 1: head = [3,0,2,6,8,1,7,9,4,2,5,5,0], m = 3, n = 5
	head1 := &ListNode{3, &ListNode{0, &ListNode{2, &ListNode{6, &ListNode{8, &ListNode{1, &ListNode{7, &ListNode{9, &ListNode{4, &ListNode{2, &ListNode{5, &ListNode{5, &ListNode{0, nil}}}}}}}}}}}}}
	fmt.Println(spiralMatrix(3, 5, head1))

	// Test case 2: head = [0,1,2], m = 1, n = 4
	head2 := &ListNode{0, &ListNode{1, &ListNode{2, nil}}}
	fmt.Println(spiralMatrix(1, 4, head2))
}
