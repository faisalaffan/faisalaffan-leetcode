# 2811 — Check If It Is Possible To Split Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CheckIfItIsPossibleToSplitArray(nums []int, m int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2811: Check if it is Possible to Split Array
// https://leetcode.com/problems/check-if-it-is-possible-to-split-array/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func CheckIfItIsPossibleToSplitArray(nums []int, m int) bool {
	n := len(nums)
	if n <= 2 {
		return true
	}

	for i := 0; i < n-1; i++ {
		if nums[i]+nums[i+1] >= m {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(CheckIfItIsPossibleToSplitArray([]int{2, 3, 3, 2, 3}, 6))
	fmt.Println(CheckIfItIsPossibleToSplitArray([]int{1, 1}, 3))
	fmt.Println(CheckIfItIsPossibleToSplitArray([]int{1, 2, 1}, 4))
}
```
