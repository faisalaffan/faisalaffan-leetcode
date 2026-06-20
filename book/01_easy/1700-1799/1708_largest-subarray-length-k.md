# 1708 — Largest Subarray Length K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func LargestSubarray(nums []int, k int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(k) (for output)  
**Kompleksitas Ruang:** O(k) (for output)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1708: Largest Subarray Length K
// https://leetcode.com/problems/largest-subarray-length-k/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(k) (for output)
func LargestSubarray(nums []int, k int) []int {
	bestIdx := 0
	for i := 1; i <= len(nums)-k; i++ {
		if nums[i] > nums[bestIdx] {
			bestIdx = i
		}
	}
	return nums[bestIdx : bestIdx+k]
}

func main() {
	fmt.Println(LargestSubarray([]int{1, 4, 5, 2, 3}, 3))
	fmt.Println(LargestSubarray([]int{1, 4, 5, 2, 3}, 4))
	fmt.Println(LargestSubarray([]int{1, 2, 3, 4, 5}, 2))
}
```
