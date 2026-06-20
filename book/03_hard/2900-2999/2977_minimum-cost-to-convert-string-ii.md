# 2977 — Minimum Cost To Convert String Ii

## Deskripsi

**Soal:** [2977. Minimum Cost To Convert String Ii](https://leetcode.com/problems/minimum-cost-to-convert-string-ii/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP), Trie (pohon awalan), Floyd-Warshall (lintasan semua pasangan)

**Fungsi Solusi:** `func minimumCost(source string, target string, original []string, changed []string, cost []int) int64`

> **Ide Kunci:** // 1. Trie to assign integer IDs to all substrings in the dictionary.

## Solusi Go

```go
package main

// LeetCode #2977: Minimum Cost to Convert String II
// https://leetcode.com/problems/minimum-cost-to-convert-string-ii/
// Difficulty: Hard
//
// Given source, target strings, and a dictionary of substring conversions
// (original[i] -> changed[i] at cost[i]), find minimum total cost to
// convert source to target by converting substrings.
//
// Approach:
// 1. Trie to assign integer IDs to all substrings in the dictionary.
// 2. Floyd-Warshall to find shortest conversion path between any two substrings.
// 3. DP[i] = min cost to convert source[i:] to target[i:].
//    Try matching substrings of all lengths, using dp[j+1] + conversion_cost.

import (
	"fmt"
	"math"
)

const alphabetSize = 26

type trieNode struct {
	children [alphabetSize]*trieNode
	idx      int // -1 means no ID assigned
}

func minimumCost(source string, target string, original []string, changed []string, cost []int) int64 {
	n := len(source)

	// Step 1: Build trie and assign IDs to all dictionary substrings
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

	m := len(original)
  // Membuat slice untuk menyimpan hasil
	origIDs := make([]int, m)
  // Membuat slice untuk menyimpan hasil
	changedIDs := make([]int, m)
	for i := 0; i < m; i++ {
		origIDs[i] = insert(original[i])
		changedIDs[i] = insert(changed[i])
	}

	// Step 2: Floyd-Warshall for shortest conversion paths
	const big = math.MaxInt64 / 2
  // Membuat slice 2D untuk DP/tabel
	dist := make([][]int64, id)
  // Iterasi seluruh elemen
	for i := range dist {
		dist[i] = make([]int64, id)
		for j := range dist[i] {
			dist[i][j] = big
		}
		dist[i][i] = 0
	}
	for i := 0; i < m; i++ {
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

	// Step 3: DP from right to left
  // Membuat slice untuk menyimpan hasil
	dp := make([]int64, n+1)
  // Iterasi seluruh elemen
	for i := range dp {
		dp[i] = big
	}
	dp[n] = 0

	for i := n - 1; i >= 0; i-- {
		// Option: skip matching characters
		if source[i] == target[i] {
			dp[i] = dp[i+1]
		}

		// Try all substring pairs starting at i
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
	// Example
	fmt.Println(minimumCost("abcd", "abef", []string{"cd"}, []string{"ef"}, []int{2}))

	// LeetCode example
	fmt.Println(minimumCost("abcdef", "abcefg", []string{"abc", "def"}, []string{"abc", "efg"}, []int{1, 2}))

	// Same string
	fmt.Println(minimumCost("abcd", "abcd", []string{"a"}, []string{"b"}, []int{5}))

	// Single char
	fmt.Println(minimumCost("a", "b", []string{"a"}, []string{"b"}, []int{10}))

	// Impossible
	fmt.Println(minimumCost("a", "b", []string{"c"}, []string{"d"}, []int{5}))
}
```
