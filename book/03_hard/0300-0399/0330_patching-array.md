# 0330 — Patching Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minPatches(nums []int, n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #330: Patching Array
// https://leetcode.com/problems/patching-array/
// Difficulty: Hard

import "fmt"

func minPatches(nums []int, n int) int {
	patches := 0
	miss := int64(1) // smallest sum we cannot form
	i := 0

	for miss <= int64(n) {
		if i < len(nums) && int64(nums[i]) <= miss {
			miss += int64(nums[i])
			i++
		} else {
			// Patch with miss itself
			miss += miss
			patches++
		}
	}
	return patches
}

func main() {
	// Example 1
	fmt.Println(minPatches([]int{1, 3}, 6))
	// 1

	// Example 2
	fmt.Println(minPatches([]int{1, 5, 10}, 20))
	// 2

	// Example 3
	fmt.Println(minPatches([]int{1, 2, 2}, 5))
	// 0

	// Example 4
	fmt.Println(minPatches([]int{1, 2, 31, 33}, 2147483647))
	// 28
}
```
