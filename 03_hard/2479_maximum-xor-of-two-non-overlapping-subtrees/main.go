package main

// LeetCode #2479: Maximum XOR of Two Non-Overlapping Subtrees
// https://leetcode.com/problems/maximum-xor-of-two-non-overlapping-subtrees/
// Difficulty: Hard [Paid]
//
// Given a rooted tree (0 is root) with values at each node, find the
// maximum XOR value of two non-overlapping subtrees. Two subtrees
// are non-overlapping if they don't share any node.
//
// Approach: Compute XOR of each subtree via DFS. Use a binary trie
// to find max XOR while avoiding overlapping subtrees. Process nodes
// post-order, removing a subtree's XOR from trie after processing.

import (
	"fmt"
)

func main() {
	// Example 1
	fmt.Println(maxXor(6, [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}}, []int{2, 3, 5, 7, 1, 4}))
	// Example 2
	fmt.Println(maxXor(3, [][]int{{0, 1}, {1, 2}}, []int{1, 2, 3}))
	// Edge: simple chain
	fmt.Println(maxXor(2, [][]int{{0, 1}}, []int{5, 3}))
}

type TrieNode struct {
	children [2]*TrieNode
	cnt      int
}

func maxXor(n int, edges [][]int, values []int) int64 {
	if n < 2 {
		return 0
	}

	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// Compute subtree XORs
	subXor := make([]int, n)
	var dfs func(u, p int) int
	dfs = func(u, p int) int {
		xor := values[u]
		for _, v := range adj[u] {
			if v != p {
				xor ^= dfs(v, u)
			}
		}
		subXor[u] = xor
		return xor
	}
	dfs(0, -1)

	// Trie operations
	root := &TrieNode{}
	insert := func(x int) {
		node := root
		for i := 60; i >= 0; i-- {
			bit := (x >> i) & 1
			if node.children[bit] == nil {
				node.children[bit] = &TrieNode{}
			}
			node = node.children[bit]
			node.cnt++
		}
	}
	remove := func(x int) {
		node := root
		for i := 60; i >= 0; i-- {
			bit := (x >> i) & 1
			node = node.children[bit]
			node.cnt--
		}
	}
	maxXor := func(x int) int {
		node := root
		ans := 0
		for i := 60; i >= 0; i-- {
			bit := (x >> i) & 1
			want := 1 - bit
			if node.children[want] != nil && node.children[want].cnt > 0 {
				ans |= (1 << i)
				node = node.children[want]
			} else {
				node = node.children[bit]
			}
		}
		return ans
	}

	// Insert all subtree XORs initially
	for i := 0; i < n; i++ {
		insert(subXor[i])
	}

	ans := 0
	var dfs2 func(u, p int)
	dfs2 = func(u, p int) {
		remove(subXor[u])
		for _, v := range adj[u] {
			if v != p {
				dfs2(v, u)
			}
		}
		// Now trie contains only XORs of subtrees that don't overlap with u's subtree
		best := maxXor(subXor[u])
		if best > ans {
			ans = best
		}
		insert(subXor[u])
	}
	dfs2(0, -1)

	return int64(ans)
}
