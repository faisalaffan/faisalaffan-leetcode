# 0721 — Accounts Merge

## Deskripsi

**Soal:** [0721. Accounts Merge](https://leetcode.com/problems/accounts-merge/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(nk * alpha(nk))  
**Kompleksitas Ruang:** O(nk)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #721: Accounts Merge
// https://leetcode.com/problems/accounts-merge/
// Difficulty: Medium
// Time: O(nk * alpha(nk))
// Space: O(nk)

import (
	"fmt"
	"sort"
)

func main() {
	accounts := [][]string{
		{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"John", "johnnybravo@mail.com"},
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"Mary", "mary@mail.com"},
	}
	fmt.Println(accountsMerge(accounts))
}

func accountsMerge(accounts [][]string) [][]string {
  // Membuat map untuk pencarian O(1): key → value
	parent := make(map[string]string)
  // Membuat map untuk pencarian O(1): key → value
	owner := make(map[string]string)

	var find func(x string) string
	find = func(x string) string {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(x, y string) {
		px, py := find(x), find(y)
		if px != py {
			parent[px] = py
		}
	}

	for _, acc := range accounts {
		name := acc[0]
		firstEmail := acc[1]
		for _, email := range acc[1:] {
			parent[email] = email
			owner[email] = name
			union(firstEmail, email)
		}
	}

  // Membuat map untuk pencarian O(1): key → value
	groups := make(map[string][]string)
	for email := range parent {
		root := find(email)
		groups[root] = append(groups[root], email)
	}

  // Membuat slice 2D untuk DP/tabel
	result := make([][]string, 0, len(groups))
	for root, emails := range groups {
		sort.Strings(emails)
		merged := append([]string{owner[root]}, emails...)
		result = append(result, merged)
	}

	return result
}
```
