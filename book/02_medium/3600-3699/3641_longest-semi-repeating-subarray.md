# 3641 — Longest Semi Repeating Subarray

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func longestSemiRepeatingSubarray(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3641: Longest Semi-Repeating Subarray
// https://leetcode.com/problems/longest-semi-repeating-subarray/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func longestSemiRepeatingSubarray(nums []int, k int) int {
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	dupCount := 0
	maxLen := 0
	l := 0

	for r, x := range nums {
		freq[x]++
		if freq[x] == 2 {
			dupCount++
		}

		for dupCount > k {
			left := nums[l]
			freq[left]--
			if freq[left] == 1 {
				dupCount--
			}
			l++
		}

		if r-l+1 > maxLen {
			maxLen = r - l + 1
		}
	}

	return maxLen
}

func main() {
	fmt.Println(longestSemiRepeatingSubarray([]int{1, 2, 3, 1, 2, 3, 4}, 2))
	fmt.Println(longestSemiRepeatingSubarray([]int{1, 1, 1, 1, 1}, 4))
	fmt.Println(longestSemiRepeatingSubarray([]int{1, 1, 1, 1, 1}, 0))
}
```
