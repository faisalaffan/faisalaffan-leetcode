# 2869 — Minimum Operations To Collect Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumOperationsToCollectElements(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(k)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2869: Minimum Operations to Collect Elements
// https://leetcode.com/problems/minimum-operations-to-collect-elements/
// Difficulty: Easy
// Time: O(n) | Space: O(k)

import "fmt"

func main() {
	fmt.Println(MinimumOperationsToCollectElements([]int{3, 1, 5, 4, 2}, 2))
	fmt.Println(MinimumOperationsToCollectElements([]int{3, 1, 5, 4, 2}, 5))
}

func MinimumOperationsToCollectElements(nums []int, k int) int {
	seen := make([]bool, k+1)
	collected := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] >= 1 && nums[i] <= k && !seen[nums[i]] {
			seen[nums[i]] = true
			collected++
		}
		if collected == k {
			return len(nums) - i
		}
	}
	return len(nums)
}
```
