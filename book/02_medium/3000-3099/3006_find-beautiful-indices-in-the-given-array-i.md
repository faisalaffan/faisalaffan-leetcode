# 3006 — Find Beautiful Indices In The Given Array I

## Deskripsi

**Soal:** [3006. Find Beautiful Indices In The Given Array I](https://leetcode.com/problems/find-beautiful-indices-in-the-given-array-i/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n+m)  
**Kompleksitas Ruang:** O(n+m)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3006: Find Beautiful Indices in the Given Array I
// https://leetcode.com/problems/find-beautiful-indices-in-the-given-array-i/
// Difficulty: Medium
// Time: O(n+m) | Space: O(n+m)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(beautifulIndices("abcd", "a", "da", 1))
	fmt.Println(beautifulIndices("isawsquirrelnearmysquirrelhouseohmy", "my", "squirrel", 15))
	fmt.Println(beautifulIndices("bcd", "a", "a", 1))
}

func beautifulIndices(s string, a string, b string, k int) (ans []int) {
	posA := kmpSearch(s, a)
	posB := kmpSearch(s, b)
	for _, i := range posA {
		idx := sort.SearchInts(posB, i)
		if idx < len(posB) && posB[idx]-i <= k {
			ans = append(ans, i)
		} else if idx > 0 && i-posB[idx-1] <= k {
			ans = append(ans, i)
		}
	}
	return
}

func kmpSearch(text, pattern string) (pos []int) {
	if pattern == "" {
		return
	}
	n, m := len(text), len(pattern)
  // Membuat slice untuk menyimpan hasil
	pi := make([]int, m)
	for i := 1; i < m; i++ {
		j := pi[i-1]
		for j > 0 && pattern[i] != pattern[j] {
			j = pi[j-1]
		}
		if pattern[i] == pattern[j] {
			j++
		}
		pi[i] = j
	}
	j := 0
	for i := 0; i < n; i++ {
		for j > 0 && text[i] != pattern[j] {
			j = pi[j-1]
		}
		if text[i] == pattern[j] {
			j++
		}
		if j == m {
			pos = append(pos, i-m+1)
			j = pi[j-1]
		}
	}
	return
}
```
