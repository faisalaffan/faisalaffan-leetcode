# 1499 — Max Value Of Equation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func findMaxValueOfEquation(points [][]int, k int) int
```

> **💡 Hint:** Monotonic Deque

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Sliding Window, Monotonic Stack/Queue

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Sliding Window** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1499: Max Value of Equation
// https://leetcode.com/problems/max-value-of-equation/
// Difficulty: Hard
//
// Approach: Monotonic Deque
// For points i < j: yi + yj + |xi - xj| = (yi - xi) + (xj + yj)
// Since xi < xj, we fix j and want max of (yi - xi) for points i where
// xj - xi <= k (i.e., xi >= xj - k).
// Use a deque storing pairs (xi, yi - xi) sorted by (yi - xi) descending.
// Pop front when xi < xj - k (out of window).
// Pop back when new value is larger (keep decreasing order).

import (
	"container/list"
	"fmt"
	"math"
)

func main() {
	// Example 1
	fmt.Println(findMaxValueOfEquation([][]int{{1, 3}, {2, 0}, {5, 10}, {6, -10}}, 1))
	// Expected: 4

	// Example 2
	fmt.Println(findMaxValueOfEquation([][]int{{0, 0}, {3, 0}, {9, 2}}, 3))
	// Expected: 3
}

type pair struct {
	x, diff int // diff = y - x
}

func findMaxValueOfEquation(points [][]int, k int) int {
	ans := math.MinInt32
	dq := list.New()

	for _, p := range points {
		xj, yj := p[0], p[1]

		// Remove points out of window (xi < xj - k)
		for dq.Len() > 0 && dq.Front().Value.(pair).x < xj-k {
			dq.Remove(dq.Front())
		}

		// If deque not empty, front has max (yi - xi)
		if dq.Len() > 0 {
			bestDiff := dq.Front().Value.(pair).diff
			val := bestDiff + xj + yj
			if val > ans {
				ans = val
			}
		}

		// Insert current point, maintain decreasing order
		diff := yj - xj
		for dq.Len() > 0 && dq.Back().Value.(pair).diff <= diff {
			dq.Remove(dq.Back())
		}
		dq.PushBack(pair{x: xj, diff: diff})
	}

	return ans
}
```
