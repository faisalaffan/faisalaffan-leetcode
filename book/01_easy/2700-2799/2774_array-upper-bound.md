# 2774 — Array Upper Bound

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ArrayUpperBound(nums []int, target int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2774: Array Upper Bound
// https://leetcode.com/problems/array-upper-bound/
// Difficulty: Easy [Paid]
// Time: O(log n) | Space: O(1)
// Note: JS problem, adapted to Go. Returns upper bound of target in sorted array.

import "fmt"

func main() {
	fmt.Println(ArrayUpperBound([]int{1, 2, 2, 2, 3}, 2))
	fmt.Println(ArrayUpperBound([]int{1, 3, 5}, 4))
}

func ArrayUpperBound(nums []int, target int) int {
	left, right := 0, len(nums)
  // Two-pointer: gerakkan kiri atau kanan
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left - 1
}
```
