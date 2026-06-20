# 2612 — Minimum Reverse Operations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minReverseOperations(n int, p int, banned []int, k int) []int
```

> **💡 Hint:** BFS + DSU skip list.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Two Pointer, Union-Find (DSU)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2612: Minimum Reverse Operations
// https://leetcode.com/problems/minimum-reverse-operations/
// Difficulty: Hard
//
// Approach: BFS + DSU skip list.
// From each position x, reversing any length-k subarray containing x
// moves it to position y = 2*i + k - 1 - x. The reachable y values
// form a contiguous range with step 2 (same parity). We use DSU
// parent pointers to quickly skip already-visited positions.

import (
	"fmt"
)

func main() {
	// Example 1: n=4, p=0, banned=[1,2], k=4 -> [0,-1,-1,1]
	fmt.Println(minReverseOperations(4, 0, []int{1, 2}, 4))
	// Example 2: n=5, p=0, banned=[2,4], k=3
	fmt.Println(minReverseOperations(5, 0, []int{2, 4}, 3))
}

func minReverseOperations(n int, p int, banned []int, k int) []int {
  // Alokasi slice integer
	ans := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range ans {
		ans[i] = -1
	}

  // Membuat map (HashMap) — pencarian O(1)
	bannedSet := make(map[int]bool)
	for _, b := range banned {
		bannedSet[b] = true
	}

	// DSU parent: next unvisited position with same parity
	// n+2 sentinel to avoid bounds checking
  // Alokasi slice integer
	parent := make([]int, n+2)
  // Range loop: iterasi dengan indeks + nilai
	for i := range parent {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	// Mark positions used by the DSU by unioning with i+2
	markUsed := func(x int) {
		parent[x] = find(x + 2)
	}

	// Mark banned positions as used
	markUsed(p)
	for _, b := range banned {
		markUsed(b)
	}

	q := []int{p}
	ans[p] = 0

	for len(q) > 0 {
		x := q[0]
		q = q[1:]

		// Compute valid subarray range: i = subarray start
		// y = 2*i + k - 1 - x
		// Constraints: 0 <= i <= n-k and i <= x <= i+k-1
		left := max(0, x-k+1)
		right := min(x, n-k)
		if left > right {
			continue
		}
		L := 2*left + k - 1 - x
		R := 2*right + k - 1 - x
		if L > R {
			L, R = R, L
		}

		targetParity := (k - 1 - x) & 1
		start := L
		if (start & 1) != targetParity {
			start++
		}

		for y := find(start); y <= R; y = find(y + 2) {
			if !bannedSet[y] && ans[y] == -1 {
				ans[y] = ans[x] + 1
				q = append(q, y)
				markUsed(y)
			} else {
				// Even if banned/visited, mark as used for DSU
				markUsed(y)
			}
		}
	}

	return ans
}
```
