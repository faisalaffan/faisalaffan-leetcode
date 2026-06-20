# 0244 — Shortest Word Distance Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func Constructor(wordsDict []string) WordDistance`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n) for init, O(m+n) for shortest, Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #244: Shortest Word Distance II
// https://leetcode.com/problems/shortest-word-distance-ii/
// Difficulty: Medium [Paid]
// Time: O(n) for init, O(m+n) for shortest, Space: O(n)

import (
	"fmt"
	"math"
)

type WordDistance struct {
	indices map[string][]int
}

func Constructor(wordsDict []string) WordDistance {
  // HashMap: O(1) lookup
	indices := make(map[string][]int)
	for i, w := range wordsDict {
		indices[w] = append(indices[w], i)
	}
	return WordDistance{indices}
}

func (this *WordDistance) Shortest(word1, word2 string) int {
	list1 := this.indices[word1]
	list2 := this.indices[word2]

	minDist := math.MaxInt32
	i, j := 0, 0

	for i < len(list1) && j < len(list2) {
		dist := list1[i] - list2[j]
		if dist < 0 {
			minDist = min(minDist, -dist)
			i++
		} else {
			minDist = min(minDist, dist)
			j++
		}
	}

	return minDist
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	wd := Constructor([]string{"practice", "makes", "perfect", "coding", "makes"})
	fmt.Println(wd.Shortest("coding", "practice"))
	fmt.Println(wd.Shortest("makes", "coding"))
	fmt.Println(wd.Shortest("makes", "practice"))
}
```
