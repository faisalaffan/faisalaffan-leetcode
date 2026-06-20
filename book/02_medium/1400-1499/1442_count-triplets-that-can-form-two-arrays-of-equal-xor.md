# 1442 — Count Triplets That Can Form Two Arrays Of Equal Xor

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func countTriplets(arr []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2) where n = len(arr)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1442: Count Triplets That Can Form Two Arrays of Equal XOR
// https://leetcode.com/problems/count-triplets-that-can-form-two-arrays-of-equal-xor/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(countTriplets([]int{2, 3, 1, 6, 7})) // 4

	// Test case 2
	fmt.Println(countTriplets([]int{1, 1, 1, 1, 1})) // 10

	// Test case 3
	fmt.Println(countTriplets([]int{1, 2, 3})) // 2

	// Test case 4
	fmt.Println(countTriplets([]int{1})) // 0
}

// Time: O(n^2) where n = len(arr)
// Space: O(1)
func countTriplets(arr []int) int {
	n := len(arr)
	count := 0

	// For pairs (i, k) where arr[i]^...^arr[k] == 0,
	// any j between i+1 and k works, giving (k-i) triplets
	for i := 0; i < n; i++ {
		xor := arr[i]
		for k := i + 1; k < n; k++ {
			xor ^= arr[k]
			if xor == 0 {
				count += k - i
			}
		}
	}

	return count
}
```
