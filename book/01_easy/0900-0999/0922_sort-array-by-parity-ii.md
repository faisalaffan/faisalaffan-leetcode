# 0922 — Sort Array By Parity Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func sortArrayByParityII(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #922: Sort Array By Parity II
// https://leetcode.com/problems/sort-array-by-parity-ii/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(sortArrayByParityII([]int{4, 2, 5, 7})) // [4,5,2,7] or [4,7,2,5]
	fmt.Println(sortArrayByParityII([]int{2, 3}))        // [2,3]
}

// sortArrayByParityII puts even numbers at even indices, odd numbers at odd indices.
// Time: O(n). Space: O(1).
func sortArrayByParityII(nums []int) []int {
	j := 1 // odd pointer
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(nums); i += 2 {
		if nums[i]%2 == 1 {
			for nums[j]%2 == 1 {
				j += 2
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return nums
}
```
