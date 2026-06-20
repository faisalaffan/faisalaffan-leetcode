# 0846 — Hand Of Straights

## Deskripsi

**Soal:** [0846. Hand Of Straights](https://leetcode.com/problems/hand-of-straights/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #846: Hand of Straights
// https://leetcode.com/problems/hand-of-straights/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(HandOfStraights([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3))
	fmt.Println(HandOfStraights([]int{1, 2, 3, 4, 5}, 4))
	fmt.Println(HandOfStraights([]int{2, 1}, 2))
}

// Time: O(n log n) | Space: O(n)
func HandOfStraights(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}

  // Membuat map untuk pencarian O(1): key → value
	count := make(map[int]int)
	for _, card := range hand {
		count[card]++
	}

  // Membuat slice untuk menyimpan hasil
	unique := make([]int, 0, len(count))
	for card := range count {
		unique = append(unique, card)
	}
	sort.Ints(unique)

	for _, card := range unique {
		if count[card] > 0 {
			freq := count[card]
			for i := 0; i < groupSize; i++ {
				if count[card+i] < freq {
					return false
				}
				count[card+i] -= freq
			}
		}
	}

	return true
}
```
