# 0961 — N Repeated Element In Size 2N Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func repeatedNTimes(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #961: N-Repeated Element in Size 2N Array
// https://leetcode.com/problems/n-repeated-element-in-size-2n-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(repeatedNTimes([]int{1, 2, 3, 3})) // 3
	fmt.Println(repeatedNTimes([]int{2, 1, 2, 5, 3, 2})) // 2
	fmt.Println(repeatedNTimes([]int{5, 1, 5, 2, 5, 3, 5, 4})) // 5
}

// repeatedNTimes finds the element repeated n times in a 2n size array.
// Time: O(n). Space: O(1).
func repeatedNTimes(nums []int) int {
	// Since the element appears n times in 2n, any two consecutive elements
	// must contain the repeated element (in most cases).
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] == nums[i+1] || nums[i] == nums[i+2] {
			return nums[i]
		}
	}
	// If not found yet, the repeated element is in the last 3 positions
	return nums[len(nums)-1]
}
```
