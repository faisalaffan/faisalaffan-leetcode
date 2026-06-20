# 3315 — Construct The Minimum Bitwise Array Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minBitwiseArray(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log m) Space: O(1) (excluding output)  
**Kompleksitas Ruang:** O(1) (excluding output)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3315: Construct the Minimum Bitwise Array II
// https://leetcode.com/problems/construct-the-minimum-bitwise-array-ii/
// Difficulty: Medium
// Time: O(n log m) Space: O(1) (excluding output)

import "fmt"

func main() {
	fmt.Println(minBitwiseArray([]int{11, 13, 31})) // [9 12 15]
	fmt.Println(minBitwiseArray([]int{2, 3, 5}))    // [-1 1 4]
	fmt.Println(minBitwiseArray([]int{7}))           // [3]
}

func minBitwiseArray(nums []int) []int {
  // Alokasi slice integer
	ans := make([]int, len(nums))
	for i, num := range nums {
		if num == 2 {
			ans[i] = -1
			continue
		}
		// Find rightmost block of 1s in binary
		p := 0
		for (num>>p)&1 == 1 {
			p++
		}
		ans[i] = num ^ (1 << (p - 1))
	}
	return ans
}
```
