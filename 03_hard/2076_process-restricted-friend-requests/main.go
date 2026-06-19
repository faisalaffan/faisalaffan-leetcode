package main

// LeetCode #2076: Process Restricted Friend Requests
// https://leetcode.com/problems/process-restricted-friend-requests/
// Difficulty: Hard
// Approach: Union-Find + Restriction Check

import "fmt"

type unionFind struct {
	parent []int
	rank   []int
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &unionFind{parent, rank}
}

func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *unionFind) union(x, y int) {
	rx, ry := uf.find(x), uf.find(y)
	if rx == ry {
		return
	}
	if uf.rank[rx] < uf.rank[ry] {
		rx, ry = ry, rx
	}
	uf.parent[ry] = rx
	if uf.rank[rx] == uf.rank[ry] {
		uf.rank[rx]++
	}
}

func friendRequests(n int, restrictions [][]int, requests [][]int) []bool {
	uf := newUnionFind(n)
	ans := make([]bool, len(requests))

	for i, req := range requests {
		u, v := req[0], req[1]
		ru, rv := uf.find(u), uf.find(v)
		canFriend := true

		// Check all restrictions: if both u's group and v's group
		// would contain both ends of a restriction, reject
		for _, res := range restrictions {
			a, b := res[0], res[1]
			ra, rb := uf.find(a), uf.find(b)
			// After union, ru and rv would be merged.
			// If (ra, rb) == (ru, rv) or (ra, rb) == (rv, ru), restriction is violated
			if (ra == ru && rb == rv) || (ra == rv && rb == ru) {
				canFriend = false
				break
			}
		}

		if canFriend {
			uf.union(u, v)
		}
		ans[i] = canFriend
	}

	return ans
}

func main() {
	fmt.Println("2076. Process Restricted Friend Requests")

	// Example 1
	n1 := 3
	restrictions1 := [][]int{{0, 1}}
	requests1 := [][]int{{0, 2}, {2, 1}}
	fmt.Printf("n=%d restrictions=%v requests=%v → %v (expected [true, false])\n",
		n1, restrictions1, requests1, friendRequests(n1, restrictions1, requests1))

	// Example 2
	n2 := 5
	restrictions2 := [][]int{{0, 1}, {1, 2}, {2, 3}}
	requests2 := [][]int{{0, 4}, {1, 2}, {3, 1}, {3, 4}}
	fmt.Printf("n=%d restrictions=%v requests=%v → %v (expected [true, false, true, false])\n",
		n2, restrictions2, requests2, friendRequests(n2, restrictions2, requests2))
}
