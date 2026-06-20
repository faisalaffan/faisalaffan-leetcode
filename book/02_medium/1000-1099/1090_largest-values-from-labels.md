# 1090 — Largest Values From Labels

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func largestValsFromLabels(values []int, labels []int, numWanted int, useLimit int) int
```

> **💡 Hint:** Sort by value descending, pick items with label limits

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1090: Largest Values From Labels
// https://leetcode.com/problems/largest-values-from-labels/
// Difficulty: Medium
//
// Approach: Sort by value descending, pick items with label limits
// Time: O(n log n)
// Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestValsFromLabels([]int{5, 4, 3, 2, 1}, []int{1, 1, 2, 2, 3}, 3, 1)) // 9
	fmt.Println(largestValsFromLabels([]int{5, 4, 3, 2, 1}, []int{1, 3, 3, 3, 2}, 3, 2)) // 12
}

func largestValsFromLabels(values []int, labels []int, numWanted int, useLimit int) int {
	n := len(values)
	type pair struct {
		val   int
		label int
	}
	items := make([]pair, n)
	for i := 0; i < n; i++ {
		items[i] = pair{values[i], labels[i]}
	}

  // Custom sort dengan comparator
	sort.Slice(items, func(i, j int) bool {
		return items[i].val > items[j].val
	})

  // Membuat map (HashMap) — pencarian O(1)
	labelCount := make(map[int]int)
	result := 0
	selected := 0

	for _, item := range items {
		if selected >= numWanted {
			break
		}
		if labelCount[item.label] < useLimit {
			result += item.val
			labelCount[item.label]++
			selected++
		}
	}

	return result
}
```
