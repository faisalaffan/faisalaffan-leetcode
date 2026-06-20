# 2760 — Longest Even Odd Subarray With Threshold

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func LongestEvenOddSubarrayWithThreshold(nums []int, threshold int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2760: Longest Even Odd Subarray With Threshold
// https://leetcode.com/problems/longest-even-odd-subarray-with-threshold/
// Difficulty: Easy
// Time: O(n^2) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(LongestEvenOddSubarrayWithThreshold([]int{3, 2, 5, 4}, 5))
	fmt.Println(LongestEvenOddSubarrayWithThreshold([]int{4, 5, 2, 1}, 4))
}

func LongestEvenOddSubarrayWithThreshold(nums []int, threshold int) int {
	maxLen := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 != 0 || nums[i] > threshold {
			continue
		}
		length := 1
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > threshold {
				break
			}
			if nums[j]%2 == nums[j-1]%2 {
				break
			}
			length++
		}
		if length > maxLen {
			maxLen = length
		}
	}
	return maxLen
}
```
