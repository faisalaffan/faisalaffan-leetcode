# 2562 — Find The Array Concatenation Value

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindTheArrayConcatenationValue(nums []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2562: Find the Array Concatenation Value
// https://leetcode.com/problems/find-the-array-concatenation-value/
// Difficulty: Easy
// Time O(n) | Space O(1)

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(FindTheArrayConcatenationValue([]int{7, 52, 2, 4})) // 596
	fmt.Println(FindTheArrayConcatenationValue([]int{5, 14, 13, 8, 12})) // 673
}

func FindTheArrayConcatenationValue(nums []int) int64 {
	sum := int64(0)
	i, j := 0, len(nums)-1
	for i < j {
		concat := int64(nums[i])
		n := nums[j]
		digits := 0
  // Edge case: input kosong — langsung return
		if n == 0 {
			digits = 1
		} else {
			digits = int(math.Log10(float64(n))) + 1
		}
		concat = concat * pow10(digits) + int64(n)
		sum += concat
		i++
		j--
	}
	if i == j {
		sum += int64(nums[i])
	}
	return sum
}

func pow10(n int) int64 {
	p := int64(1)
	for i := 0; i < n; i++ {
		p *= 10
	}
	return p
}
```
