# 0952 — Largest Component Size By Common Factor

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func largestComponentSize(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Union-Find (DSU)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #952: Largest Component Size by Common Factor
// https://leetcode.com/problems/largest-component-size-by-common-factor/
// Difficulty: Hard

import "fmt"

func largestComponentSize(nums []int) int {
  // Edge case: input kosong — langsung return
	if len(nums) == 0 {
		return 0
	}

	maxVal := 0
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}

  // Alokasi slice integer
	parent := make([]int, maxVal+1)
  // Alokasi slice integer
	size := make([]int, maxVal+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(a, b int) {
		ra, rb := find(a), find(b)
		if ra == rb {
			return
		}
		if size[ra] < size[rb] {
			ra, rb = rb, ra
		}
		parent[rb] = ra
		size[ra] += size[rb]
	}

	// For each number, union it with its prime factors
	for _, num := range nums {
		x := num
		for p := 2; p*p <= x; p++ {
			if x%p == 0 {
				union(num, p)
				for x%p == 0 {
					x /= p
				}
			}
		}
		if x > 1 {
			union(num, x)
		}
	}

	// Count component sizes among the given numbers
  // Membuat map (HashMap) — pencarian O(1)
	compCount := make(map[int]int)
	ans := 0
	for _, num := range nums {
		root := find(num)
		compCount[root]++
		if compCount[root] > ans {
			ans = compCount[root]
		}
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println("Example 1:")
	fmt.Println(largestComponentSize([]int{4, 6, 15, 35}))
	// Expected: 4

	// Example 2
	fmt.Println("Example 2:")
	fmt.Println(largestComponentSize([]int{20, 50, 9, 63}))
	// Expected: 2

	// Example 3
	fmt.Println("Example 3:")
	fmt.Println(largestComponentSize([]int{2, 3, 6, 7, 4, 12, 21, 39}))
	// Expected: 8
}
```
