# 3205 — Maximum Array Hopping Score I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxScore(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3205: Maximum Array Hopping Score I
// https://leetcode.com/problems/maximum-array-hopping-score-i/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func maxScore(nums []int) int {
	ans := 0
	mx := 0
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > mx {
			mx = nums[i]
		}
		ans += mx
	}
	return ans
}

func main() {
	fmt.Println(maxScore([]int{1, 2, 3, 4, 5})) // Expected: 14 (2+3+4+5)
	fmt.Println(maxScore([]int{5, 4, 3, 2, 1})) // Expected: 4 (1+1+1+1)
	fmt.Println(maxScore([]int{1, 5, 2, 6, 3})) // Expected: 15 (5+6+6+6)
}
```
