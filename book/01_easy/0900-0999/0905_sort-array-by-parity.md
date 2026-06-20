# 0905 — Sort Array By Parity

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func sortArrayByParity(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #905: Sort Array By Parity
// https://leetcode.com/problems/sort-array-by-parity/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4})) // [2,4,3,1] or [4,2,1,3] etc.
	fmt.Println(sortArrayByParity([]int{0}))           // [0]
	fmt.Println(sortArrayByParity([]int{1, 3, 5}))     // [1,3,5]
}

// sortArrayByParity moves all even numbers to the front, odd to the back.
// Time: O(n). Space: O(1).
func sortArrayByParity(nums []int) []int {
	l, r := 0, len(nums)-1
  // Two-pointer: gerakkan kiri atau kanan
	for l < r {
		if nums[l]%2 == 0 {
			l++
		} else {
			nums[l], nums[r] = nums[r], nums[l]
			r--
		}
	}
	return nums
}
```
