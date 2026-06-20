# 3907 — Count Smaller Elements With Opposite Parity

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func NewBIT(size int) *BIT
```

> **💡 Hint:** Process right to left. Use two BITs (even, odd) to count smaller

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Two Pointer, Fenwick Tree (BIT)

**Kompleksitas Waktu:** O(N log M)  
**Kompleksitas Ruang:** O(M) where M = max value

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3907: Count Smaller Elements With Opposite Parity
// https://leetcode.com/problems/count-smaller-elements-with-opposite-parity/
// Difficulty: Medium [Paid]
// Time: O(N log M) | Space: O(M) where M = max value
// Approach: Process right to left. Use two BITs (even, odd) to count smaller
// elements with opposite parity.

import (
	"fmt"
	"sort"
)

type BIT struct {
	tree []int
}

func NewBIT(size int) *BIT {
	return &BIT{tree: make([]int, size+2)}
}

func (b *BIT) Update(idx int, val int) {
	idx++
	for idx < len(b.tree) {
		b.tree[idx] += val
		idx += idx & -idx
	}
}

func (b *BIT) Query(idx int) int {
	idx++
	sum := 0
	for idx > 0 {
		sum += b.tree[idx]
		idx -= idx & -idx
	}
	return sum
}

func CountSmallerElementsWithOppositeParity(nums []int) []int {
	n := len(nums)
  // Alokasi slice integer
	ans := make([]int, n)

	// Coordinate compress
  // Alokasi slice integer
	sorted := make([]int, n)
	copy(sorted, nums)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(sorted)
  // Membuat map (HashMap) — pencarian O(1)
	rank := make(map[int]int)
	for i, v := range sorted {
		rank[v] = i
	}

	evenBit := NewBIT(n)
	oddBit := NewBIT(n)

	for i := n - 1; i >= 0; i-- {
		r := rank[nums[i]]
		if nums[i]%2 == 0 {
			// Count odd elements smaller than nums[i]
			ans[i] = oddBit.Query(r - 1)
			evenBit.Update(r, 1)
		} else {
			// Count even elements smaller than nums[i]
			ans[i] = evenBit.Query(r - 1)
			oddBit.Update(r, 1)
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(CountSmallerElementsWithOppositeParity([]int{5, 2, 4, 1, 3})) // Expected: [2 1 2 0 0]

	// Example 2
	fmt.Println(CountSmallerElementsWithOppositeParity([]int{4, 4, 1})) // Expected: [1 1 0]

	// Example 3
	fmt.Println(CountSmallerElementsWithOppositeParity([]int{7})) // Expected: [0]
}
```
