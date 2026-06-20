# 0244 — Shortest Word Distance Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor(wordsDict []string) WordDistance
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n) for init, O(m+n) for shortest, Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
  // Membuat map (HashMap) — pencarian O(1)
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
