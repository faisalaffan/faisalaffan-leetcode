# 3282 — Reach End Of Array With Max Score

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func findMaximumScore(nums []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n) Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3282: Reach End of Array With Max Score
// https://leetcode.com/problems/reach-end-of-array-with-max-score/
// Difficulty: Medium
// Time: O(n) Space: O(1)

import "fmt"

func main() {
	fmt.Println(findMaximumScore([]int{1, 3, 1, 5}))   // 7
	fmt.Println(findMaximumScore([]int{4, 3, 1, 3, 2})) // 16
	fmt.Println(findMaximumScore([]int{2, 2, 2, 2}))    // 6
}

func findMaximumScore(nums []int) int64 {
	var res int64 = 0
	ma := nums[0]
	for i := 1; i < len(nums); i++ {
		res += int64(ma)
		if nums[i] > ma {
			ma = nums[i]
		}
	}
	return res
}
```
