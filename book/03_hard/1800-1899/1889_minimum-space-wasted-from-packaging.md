# 1889 — Minimum Space Wasted From Packaging

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minWastedSpace(packages []int, boxes [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1889: Minimum Space Wasted From Packaging
// https://leetcode.com/problems/minimum-space-wasted-from-packaging/
// Difficulty: Hard

import (
	"fmt"
	"math"
	"sort"
)

func minWastedSpace(packages []int, boxes [][]int) int {
	const mod = 1_000_000_007
  // Urutkan secara ascending — O(n log n)
	sort.Ints(packages)
	n := len(packages)
  // Alokasi slice integer
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + packages[i]
	}

	ans := math.MaxInt64

	for _, supplier := range boxes {
  // Urutkan secara ascending — O(n log n)
		sort.Ints(supplier)
		if supplier[len(supplier)-1] < packages[n-1] {
			continue // cannot fit the largest package
		}
		waste := 0
		prevIdx := 0
		for _, box := range supplier {
			// find the last package that fits in this box
			idx := sort.Search(n-prevIdx, func(k int) bool {
				return packages[prevIdx+k] > box
			}) + prevIdx
			if idx > prevIdx {
				count := idx - prevIdx
				// waste = box * count - sum of packages in [prevIdx, idx)
				sum := prefix[idx] - prefix[prevIdx]
				waste += box*count - sum
				prevIdx = idx
				if waste > ans { // early break
					break
				}
			}
		}
		if prevIdx == n && waste < ans {
			ans = waste
		}
	}

	if ans == math.MaxInt64 {
		return -1
	}
	return ans % mod
}

func main() {
	// Example: packages=[2,3,5], boxes=[[4,8],[2,8]] -> 6
	fmt.Println(minWastedSpace([]int{2, 3, 5}, [][]int{{4, 8}, {2, 8}}))

	// Additional test
	fmt.Println(minWastedSpace([]int{3, 5, 8, 10, 11, 12}, [][]int{{12}, {11, 9}, {10, 5, 14}}))
}
```
