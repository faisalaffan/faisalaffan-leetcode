# 2221 — Find Triangular Sum Of An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func triangularSum(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2221: Find Triangular Sum of an Array
// https://leetcode.com/problems/find-triangular-sum-of-an-array/
// Difficulty: Medium
// Time: O(n^2) | Space: O(1)

import "fmt"

func triangularSum(nums []int) int {
	n := len(nums)
	for n > 1 {
		for i := 0; i < n-1; i++ {
			nums[i] = (nums[i] + nums[i+1]) % 10
		}
		n--
	}
	return nums[0]
}

func main() {
	// Test case 1
	fmt.Println(triangularSum([]int{1, 2, 3, 4, 5}))
	// Expected: 8

	// Test case 2
	fmt.Println(triangularSum([]int{5}))
	// Expected: 5

	// Test case 3
	fmt.Println(triangularSum([]int{2, 6, 6, 6}))
	// Expected: 4
}
```
