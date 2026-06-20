# 3523 — Make Array Non Decreasing

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MakeArrayNonDecreasing(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3523: Make Array Non-decreasing
// https://leetcode.com/problems/make-array-non-decreasing/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", MakeArrayNonDecreasing([]int{4, 2, 3}))
	// Test case 2
	fmt.Println("Test 2:", MakeArrayNonDecreasing([]int{4, 2, 1}))
	// Test case 3
	fmt.Println("Test 3:", MakeArrayNonDecreasing([]int{1, 2, 3}))
}

func MakeArrayNonDecreasing(nums []int) int {
	ops := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			ops += nums[i-1] - nums[i]
			nums[i] = nums[i-1]
		}
	}
	return ops
}
```
