# 0553 — Optimal Division

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func OptimalDivision(nums []int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #553: Optimal Division
// https://leetcode.com/problems/optimal-division/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(OptimalDivision([]int{1000, 100, 10, 2}))
	fmt.Println(OptimalDivision([]int{2, 3, 4}))
	fmt.Println(OptimalDivision([]int{2}))
}

func OptimalDivision(nums []int) string {
	n := len(nums)
  // Edge case: input kosong — langsung return
	if n == 0 {
		return ""
	}
	if n == 1 {
		return strconv.Itoa(nums[0])
	}
	if n == 2 {
		return strconv.Itoa(nums[0]) + "/" + strconv.Itoa(nums[1])
	}

	result := strconv.Itoa(nums[0]) + "/(" + strconv.Itoa(nums[1])
	for i := 2; i < n; i++ {
		result += "/" + strconv.Itoa(nums[i])
	}
	result += ")"

	return result
}
```
