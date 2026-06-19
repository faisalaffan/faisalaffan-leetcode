package main

// LeetCode #2977: Minimum Cost to Convert String II
// https://leetcode.com/problems/minimum-cost-to-convert-string-ii/
// Difficulty: Hard
//
// Approach: Trie + DP + Floyd-Warshall
// 1. Use a trie to assign integer IDs to all original/changed substrings.
// 2. Build a graph of substring-to-substring costs and run Floyd-Warshall.
// 3. DP[i] = min cost to convert source[i:] to target[i:].
//    - If source[i]==target[i], DP[i] = DP[i+1] (skip matching char).
//    - For every substring pair (source[i:j+1], target[i:j+1]) that exists
//      in the dictionary, try converting that substring.

import (
	"fmt"
	"math"
)

type trieNode struct {
	children [26]*trieNode
	idx      int
}

func minimumCost(source string, target string, original []string, changed []string, cost []int) int64 {
	n := len(source)
	root := &trieNode{}
	id := 0
	insert := func(s string) int {
		cur := root
		for _, ch := range s {
			c := ch - 'a'
			if cur.children[c] == nil {
				cur.children[c] = &trieNode{idx: -1}
			}
			cur = cur.children[c]
		}
		if cur.idx == -1 {
			cur.idx = id
			id++
		}
		return cur.idx
	}
	origIDs := make([]int, len(original))
	changedIDs := make([]int, len(changed))
	for i, s := range original {
		origIDs[i] = insert(s)
	}
	for i, s := range changed {
		changedIDs[i] = insert(s)
	}
	const big = math.MaxInt64 / 2
	dist := make([][]int64, id)
	for i := range dist {
		dist[i] = make([]int64, id)
		for j := range dist[i] {
			dist[i][j] = big
		}
		dist[i][i] = 0
	}
	for i := 0; i < len(original); i++ {
		u, v := origIDs[i], changedIDs[i]
		if int64(cost[i]) < dist[u][v] {
			dist[u][v] = int64(cost[i])
		}
	}
	for k := 0; k < id; k++ {
		for i := 0; i < id; i++ {
			if dist[i][k] == big {
				continue
			}
			for j := 0; j < id; j++ {
				if nd := dist[i][k] + dist[k][j]; nd < dist[i][j] {
					dist[i][j] = nd
				}
			}
		}
	}
	dp := make([]int64, n+1)
	for i := range dp {
		dp[i] = big
	}
	dp[n] = 0
	for i := n - 1; i >= 0; i-- {
		if source[i] == target[i] {
			dp[i] = dp[i+1]
		}
		curS := root
		curT := root
		for j := i; j < n; j++ {
			cs := source[j] - 'a'
			ct := target[j] - 'a'
			if curS.children[cs] == nil || curT.children[ct] == nil {
				break
			}
			curS = curS.children[cs]
			curT = curT.children[ct]
			if curS.idx != -1 && curT.idx != -1 && dist[curS.idx][curT.idx] != big {
				if nd := dist[curS.idx][curT.idx] + dp[j+1]; nd < dp[i] {
					dp[i] = nd
				}
			}
		}
	}
	if dp[0] >= big {
		return -1
	}
	return dp[0]
}

func main() {
	// Example: source="abcd" -> 2
	// Convert "cd" to "ef" cost 2: "ab"+"ef" = "abef"
	fmt.Println(minimumCost("abcd", "abef", []string{"cd"}, []string{"ef"}, []int{2}))

	// LeetCode example
	fmt.Println(minimumCost("abcdef", "abcefg", []string{"abc", "def"}, []string{"abc", "efg"}, []int{1, 2}))

	// Same string, cost 0
	fmt.Println(minimumCost("abcd", "abcd", []string{"a"}, []string{"b"}, []int{5}))

	// Single char conversion
	fmt.Println(minimumCost("a", "b", []string{"a"}, []string{"b"}, []int{10}))
}
