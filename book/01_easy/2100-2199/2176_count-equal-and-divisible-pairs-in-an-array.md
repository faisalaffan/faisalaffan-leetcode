# 2176 — Count Equal And Divisible Pairs In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CountEqualAndDivisiblePairsInAnArray(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2176: Count Equal and Divisible Pairs in an Array
// https://leetcode.com/problems/count-equal-and-divisible-pairs-in-an-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountEqualAndDivisiblePairsInAnArray([]int{3, 1, 2, 2, 2, 1, 3}, 2)) // 4
	fmt.Println(CountEqualAndDivisiblePairsInAnArray([]int{1, 2, 3, 4}, 1))           // 0
}

// Time: O(n^2), Space: O(1)
func CountEqualAndDivisiblePairsInAnArray(nums []int, k int) int {
	n := len(nums)
	count := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] == nums[j] && (i*j)%k == 0 {
				count++
			}
		}
	}
	return count
}
```
