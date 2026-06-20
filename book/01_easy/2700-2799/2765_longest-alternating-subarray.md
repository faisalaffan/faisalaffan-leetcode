# 2765 — Longest Alternating Subarray

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func LongestAlternatingSubarray(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2765: Longest Alternating Subarray
// https://leetcode.com/problems/longest-alternating-subarray/
// Difficulty: Easy
// Time: O(n^2) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(LongestAlternatingSubarray([]int{2, 3, 4, 3, 4}))
	fmt.Println(LongestAlternatingSubarray([]int{4, 5, 6}))
}

func LongestAlternatingSubarray(nums []int) int {
	n := len(nums)
	ans := -1
	for i := 0; i < n-1; i++ {
		if nums[i+1]-nums[i] != 1 {
			continue
		}
		length := 2
		expected := -1 // next diff should be -1
		for j := i + 2; j < n; j++ {
			diff := nums[j] - nums[j-1]
			if diff != expected {
				break
			}
			length++
			expected = -expected
		}
		if length > ans {
			ans = length
		}
	}
	return ans
}
```
