# 1520 — Maximum Number Of Non Overlapping Substrings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxNumOfSubstrings(s string) []string
```

> **💡 Hint:** Greedy Interval

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1520: Maximum Number of Non-Overlapping Substrings
// https://leetcode.com/problems/maximum-number-of-non-overlapping-substrings/
// Difficulty: Hard
//
// Approach: Greedy Interval
// 1. For each character, find its first and last occurrence in s.
// 2. For each character, expand its interval until all chars in the interval
//    have their full range within the interval.
// 3. Sort intervals by end ascending. Greedy pick: if start > last_end, take it.

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(maxNumOfSubstrings("adefaddaccc"))
	// Expected: ["e","f","ccc"] or similar valid order

	// Example 2
	fmt.Println(maxNumOfSubstrings("abbaccd"))
	// Expected: ["d","bb","cc"] or similar
}

type iv struct{ l, r int }

func maxNumOfSubstrings(s string) []string {
	n := len(s)

	// First and last occurrence of each char
  // Alokasi slice integer
	first := make([]int, 26)
  // Alokasi slice integer
	last := make([]int, 26)
	for i := 0; i < 26; i++ {
		first[i] = n
		last[i] = -1
	}
	for i, ch := range s {
		c := int(ch - 'a')
		if i < first[c] {
			first[c] = i
		}
		if i > last[c] {
			last[c] = i
		}
	}

	// Compute minimal interval for each character that appears
	minIntervals := make([]iv, 0)
	for c := 0; c < 26; c++ {
		if first[c] == n {
			continue
		}
		l, r := first[c], last[c]
		changed := true
		for changed {
			changed = false
			for j := l; j <= r; j++ {
				ch := int(s[j] - 'a')
				if first[ch] < l {
					l = first[ch]
					changed = true
				}
				if last[ch] > r {
					r = last[ch]
					changed = true
				}
			}
		}
		minIntervals = append(minIntervals, iv{l, r})
	}

	// Dedup by (l,r)
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[int]map[int]bool)
	unique := make([]iv, 0)
	for _, inv := range minIntervals {
		if seen[inv.l] == nil {
			seen[inv.l] = make(map[int]bool)
		}
		if !seen[inv.l][inv.r] {
			seen[inv.l][inv.r] = true
			unique = append(unique, inv)
		}
	}

	// Sort by end ascending
  // Custom sort dengan comparator
	sort.Slice(unique, func(i, j int) bool {
		return unique[i].r < unique[j].r
	})

	// Greedy pick non-overlapping substrings
	result := make([]string, 0)
	end := -1
	for _, inv := range unique {
		if inv.l > end {
			result = append(result, s[inv.l:inv.r+1])
			end = inv.r
		}
	}

	return result
}
```
