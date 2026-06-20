# 2980 — Check If Bitwise Or Has Trailing Zeros

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CheckIfBitwiseOrHasTrailingZeros(nums []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2980: Check if Bitwise OR Has Trailing Zeros
// https://leetcode.com/problems/check-if-bitwise-or-has-trailing-zeros/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: hasTrailingZeros
	fmt.Println(CheckIfBitwiseOrHasTrailingZeros([]int{1, 2, 3, 4, 5})) // true
	fmt.Println(CheckIfBitwiseOrHasTrailingZeros([]int{2, 4, 8, 16}))   // true
	fmt.Println(CheckIfBitwiseOrHasTrailingZeros([]int{1, 3, 5, 7, 9})) // false
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: hasTrailingZeros
func CheckIfBitwiseOrHasTrailingZeros(nums []int) bool {
	evenCount := 0
	for _, num := range nums {
		if num%2 == 0 {
			evenCount++
			if evenCount >= 2 {
				return true
			}
		}
	}
	return false
}
```
