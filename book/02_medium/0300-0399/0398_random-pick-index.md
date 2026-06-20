# 0398 — Random Pick Index

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor(nums []int) Solution
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n) for init, O(1) for pick  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #398: Random Pick Index
// https://leetcode.com/problems/random-pick-index/
// Difficulty: Medium
// Time: O(n) for init, O(1) for pick | Space: O(n)

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums: nums}
}

func (s *Solution) Pick(target int) int {
	// Reservoir sampling
	count := 0
	result := 0
	for i, num := range s.nums {
		if num == target {
			count++
			if rand.Intn(count) == 0 {
				result = i
			}
		}
	}
	return result
}

func main() {
	sol := Constructor([]int{1, 2, 3, 3, 3})
	counts := map[int]int{}
	for i := 0; i < 30000; i++ {
		counts[sol.Pick(3)]++
	}
	fmt.Println("Counts for target=3:", counts)
	// Expected: roughly 10000 each for indices 2,3,4
}
```
