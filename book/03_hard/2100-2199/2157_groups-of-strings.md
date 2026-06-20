# 2157 — Groups Of Strings

## Deskripsi

**Soal:** [2157. Groups Of Strings](https://leetcode.com/problems/groups-of-strings/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Bitmask (representasi himpunan dengan bit)

## Solusi Go

```go
package main

// LeetCode #2157: Groups of Strings
// https://leetcode.com/problems/groups-of-strings/
// Difficulty: Hard
//
// Union-Find + bitmask (26-bit). Two words are connected if one can become
// the other by adding, deleting, or replacing one character.

import "fmt"

func main() {
	fmt.Println(groupStrings([]string{"a", "b", "ab", "cde"}))       // [2, 3]
	fmt.Println(groupStrings([]string{"a", "ab", "abc"}))            // [1, 3]
	fmt.Println(groupStrings([]string{"abc", "acb", "bac", "bca"}))  // [1, 4]
	fmt.Println(groupStrings([]string{"ab"}))                        // [1, 1]
}

func groupStrings(words []string) []int {
	n := len(words)
  // Membuat slice untuk menyimpan hasil
	masks := make([]uint32, n)
  // Membuat map untuk pencarian O(1): key → value
	idxOf := make(map[uint32]int)

	for i, w := range words {
		var m uint32
		for _, ch := range w {
			m |= 1 << (ch - 'a')
		}
		masks[i] = m
		idxOf[m] = i
	}

	// Union-Find
  // Membuat slice untuk menyimpan hasil
	parent := make([]int, n)
  // Membuat slice untuk menyimpan hasil
	sz := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		sz[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		for parent[x] != x {
			parent[x] = parent[parent[x]]
			x = parent[x]
		}
		return x
	}
	union := func(a, b int) {
		ra, rb := find(a), find(b)
		if ra == rb {
			return
		}
		if sz[ra] < sz[rb] {
			ra, rb = rb, ra
		}
		parent[rb] = ra
		sz[ra] += sz[rb]
	}

	// Union identical masks first
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if masks[i] == masks[j] {
				union(i, j)
			}
		}
	}

	// For each unique mask, try transforms
  // Membuat map untuk pencarian O(1): key → value
	seen := make(map[uint32]bool)
	for i, m := range masks {
		if seen[m] {
			continue
		}
		seen[m] = true

		// Delete one bit
		for b := 0; b < 26; b++ {
			if m&(1<<b) != 0 {
				nm := m & ^(1 << b)
				if j, ok := idxOf[nm]; ok {
					union(i, j)
				}
			}
		}

		// Add one bit
		for b := 0; b < 26; b++ {
			if m&(1<<b) == 0 {
				nm := m | (1 << b)
				if j, ok := idxOf[nm]; ok {
					union(i, j)
				}
			}
		}

		// Replace one bit
		for b1 := 0; b1 < 26; b1++ {
			if m&(1<<b1) == 0 {
				continue
			}
			for b2 := 0; b2 < 26; b2++ {
				if b1 == b2 || m&(1<<b2) != 0 {
					continue
				}
				nm := (m & ^(1 << b1)) | (1 << b2)
				if j, ok := idxOf[nm]; ok {
					union(i, j)
				}
			}
		}
	}

  // Membuat map untuk pencarian O(1): key → value
	groupSizes := make(map[int]int)
	maxSize := 0
	for i := 0; i < n; i++ {
		r := find(i)
		groupSizes[r]++
		if groupSizes[r] > maxSize {
			maxSize = groupSizes[r]
		}
	}
	return []int{len(groupSizes), maxSize}
}

func GroupsOfStrings() any {
	return groupStrings([]string{"a", "b", "ab", "cde"})
}
```
