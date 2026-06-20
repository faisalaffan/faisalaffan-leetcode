# 3024 — Type Of Triangle

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func TypeOfTriangle(nums []int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3024: Type of Triangle
// https://leetcode.com/problems/type-of-triangle/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	// LeetCode name: triangleType
	fmt.Println(TypeOfTriangle([]int{3, 3, 3})) // equilateral
	fmt.Println(TypeOfTriangle([]int{3, 4, 5})) // scalene
	fmt.Println(TypeOfTriangle([]int{3, 3, 5})) // isosceles
	fmt.Println(TypeOfTriangle([]int{1, 2, 3})) // none
}

// Time: O(1) | Space: O(1)
// LeetCode submission name: triangleType
func TypeOfTriangle(nums []int) string {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	a, b, c := nums[0], nums[1], nums[2]

	// Check if valid triangle
	if a+b <= c {
		return "none"
	}

	if a == b && b == c {
		return "equilateral"
	}
	if a == b || b == c || a == c {
		return "isosceles"
	}
	return "scalene"
}
```
