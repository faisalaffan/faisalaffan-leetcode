# 2735 — Collecting Chocolates

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CollectingChocolates(nums []int, x int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2735: Collecting Chocolates
// https://leetcode.com/problems/collecting-chocolates/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func CollectingChocolates(nums []int, x int) int64 {
	n := len(nums)
  // Alokasi slice integer
	minCost := make([]int, n)
	copy(minCost, nums)

	var best int64
	for i := 0; i < n; i++ {
		best += int64(nums[i])
	}

	for shift := 1; shift < n; shift++ {
		var total int64 = int64(shift) * int64(x)
		for i := 0; i < n; i++ {
			idx := (i + shift) % n
			if nums[idx] < minCost[i] {
				minCost[i] = nums[idx]
			}
			total += int64(minCost[i])
		}
		if total < best {
			best = total
		}
	}

	return best
}

func main() {
	fmt.Println(CollectingChocolates([]int{20, 1, 15}, 5))
	fmt.Println(CollectingChocolates([]int{1, 2, 3}, 4))
}
```
