# 2815 — Max Pair Sum In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxDigit(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2815: Max Pair Sum in an Array
// https://leetcode.com/problems/max-pair-sum-in-an-array/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(MaxPairSumInAnArray([]int{51, 71, 17, 24, 42}))
	fmt.Println(MaxPairSumInAnArray([]int{1, 2, 3, 4}))
}

func maxDigit(n int) int {
	maxD := 0
	for n > 0 {
		d := n % 10
		if d > maxD {
			maxD = d
		}
		n /= 10
	}
	return maxD
}

func MaxPairSumInAnArray(nums []int) int {
  // Alokasi slice integer
	maxVal := make([]int, 10) // digits 0-9
	ans := -1
	for _, n := range nums {
		md := maxDigit(n)
		if maxVal[md] > 0 {
			sum := maxVal[md] + n
			if sum > ans {
				ans = sum
			}
		}
		if n > maxVal[md] {
			maxVal[md] = n
		}
	}
	return ans
}
```
