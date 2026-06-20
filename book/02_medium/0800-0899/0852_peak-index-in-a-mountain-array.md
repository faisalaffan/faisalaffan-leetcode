# 0852 — Peak Index In A Mountain Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func PeakIndexInAMountainArray(arr []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #852: Peak Index in a Mountain Array
// https://leetcode.com/problems/peak-index-in-a-mountain-array/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(PeakIndexInAMountainArray([]int{0, 1, 0}))
	fmt.Println(PeakIndexInAMountainArray([]int{0, 2, 1, 0}))
	fmt.Println(PeakIndexInAMountainArray([]int{0, 10, 5, 2}))
}

// Time: O(log n) | Space: O(1)
func PeakIndexInAMountainArray(arr []int) int {
	left, right := 1, len(arr)-2
  // Two-pointer: gerakkan kiri atau kanan
	for left < right {
		mid := (left + right) / 2
		if arr[mid] > arr[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
```
