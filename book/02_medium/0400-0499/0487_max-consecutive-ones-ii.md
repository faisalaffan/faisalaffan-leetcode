# 0487 — Max Consecutive Ones Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaxConsecutiveOnesIi(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #487: Max Consecutive Ones II
// https://leetcode.com/problems/max-consecutive-ones-ii/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(MaxConsecutiveOnesIi([]int{1, 0, 1, 1, 0}))
	fmt.Println(MaxConsecutiveOnesIi([]int{1, 0, 1, 1, 0, 1}))
}

func MaxConsecutiveOnesIi(nums []int) int {
	maxLen := 0
	prevLen, curLen := 0, 0

	for _, num := range nums {
		if num == 1 {
			curLen++
		} else {
			prevLen = curLen
			curLen = 0
		}
		if prevLen+curLen+1 > maxLen {
			maxLen = prevLen + curLen + 1
		}
	}

	if maxLen > len(nums) {
		return len(nums)
	}
	return maxLen
}
```
