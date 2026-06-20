# 3702 — Longest Subsequence With Non Zero Bitwise Xor

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func longestSubsequenceWithNonZeroBitwiseXor(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3702: Longest Subsequence With Non-Zero Bitwise XOR
// https://leetcode.com/problems/longest-subsequence-with-non-zero-bitwise-xor/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func longestSubsequenceWithNonZeroBitwiseXor(nums []int) int {
	xorSum := 0
	allZero := true
	for _, num := range nums {
		xorSum ^= num
		if num != 0 {
			allZero = false
		}
	}
	if allZero {
		return 0
	}
	if xorSum != 0 {
		return len(nums)
	}
	return len(nums) - 1
}

func main() {
	fmt.Println(longestSubsequenceWithNonZeroBitwiseXor([]int{1, 2, 3}))
	fmt.Println(longestSubsequenceWithNonZeroBitwiseXor([]int{2, 3, 4}))
	fmt.Println(longestSubsequenceWithNonZeroBitwiseXor([]int{0, 0, 0}))
}
```
