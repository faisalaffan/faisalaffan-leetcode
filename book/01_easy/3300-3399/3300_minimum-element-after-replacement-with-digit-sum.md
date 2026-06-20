# 3300 — Minimum Element After Replacement With Digit Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func digitSum(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * log n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3300: Minimum Element After Replacement With Digit Sum
// https://leetcode.com/problems/minimum-element-after-replacement-with-digit-sum/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumElementAfterReplacementWithDigitSum([]int{10, 12, 13, 14}))
	fmt.Println(MinimumElementAfterReplacementWithDigitSum([]int{1, 2, 3, 4}))
	fmt.Println(MinimumElementAfterReplacementWithDigitSum([]int{999, 19, 199}))
}

// digitSum returns the sum of digits of n.
func digitSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

// MinimumElementAfterReplacementWithDigitSum returns the minimum element after replacing each element with its digit sum.
// Time: O(n * log n). Space: O(1).
func MinimumElementAfterReplacementWithDigitSum(nums []int) int {
	minVal := int(^uint(0) >> 1) // MaxInt
	for _, num := range nums {
		s := digitSum(num)
		if s < minVal {
			minVal = s
		}
	}
	return minVal
}
```
