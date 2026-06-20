# 1250 — Check If It Is A Good Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func isGoodArray(nums []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** GCD / Matematika

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **GCD / Matematika** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1250: Check If It Is a Good Array
// https://leetcode.com/problems/check-if-it-is-a-good-array/
// Difficulty: Hard
//
// Given an array of positive integers nums, return true if for every integer x
// that can be formed as a linear combination of the elements of nums with
// integer coefficients, there is a subset of nums whose GCD is 1.
//
// By Bezout's identity, a subset of numbers has GCD 1 iff we can form 1 as a
// linear combination. Therefore, the condition is equivalent to: the GCD of
// the entire array is 1.

import "fmt"

func main() {
	// Example 1: gcd(12,5,7,23) = 1 -> true
	fmt.Println(isGoodArray([]int{12, 5, 7, 23})) // true

	// Example 2: gcd(29,6,10) = 1 -> true
	fmt.Println(isGoodArray([]int{29, 6, 10})) // true

	// Example 3: gcd(3,6) = 3 != 1 -> false
	fmt.Println(isGoodArray([]int{3, 6})) // false

	// Single element that is 1 -> true
	fmt.Println(isGoodArray([]int{1})) // true

	// Coprime numbers
	fmt.Println(isGoodArray([]int{6, 10, 15})) // true (gcd=1)

	// All even -> false
	fmt.Println(isGoodArray([]int{4, 8, 12})) // false
}

// isGoodArray returns true if the GCD of the entire array is 1.
func isGoodArray(nums []int) bool {
  // Edge case: input kosong — langsung return
	if len(nums) == 0 {
		return false
	}

	g := nums[0]
	for i := 1; i < len(nums); i++ {
		g = gcd(g, nums[i])
		if g == 1 {
			return true // early exit
		}
	}

	return g == 1
}

// gcd computes the greatest common divisor using Euclidean algorithm.
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
```
