# 0674 — Longest Continuous Increasing Subsequence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func findLengthOfLCIS(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #674: Longest Continuous Increasing Subsequence
// https://leetcode.com/problems/longest-continuous-increasing-subsequence/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 4, 7}))    // 3
	fmt.Println(findLengthOfLCIS([]int{2, 2, 2, 2, 2}))    // 1
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 7}))       // 4
}

// findLengthOfLCIS finds the length of the longest continuous increasing subsequence.
// Time: O(n). Space: O(1).
func findLengthOfLCIS(nums []int) int {
  // Edge case: input kosong — langsung return
	if len(nums) == 0 {
		return 0
	}
	maxLen, curr := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			curr++
			if curr > maxLen {
				maxLen = curr
			}
		} else {
			curr = 1
		}
	}
	return maxLen
}
```
