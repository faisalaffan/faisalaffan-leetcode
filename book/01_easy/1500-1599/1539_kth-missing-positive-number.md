# 1539 — Kth Missing Positive Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func findKthPositive(arr []int, k int) int

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1539: Kth Missing Positive Number
// https://leetcode.com/problems/kth-missing-positive-number/
// Difficulty: Easy
//
// LeetCode submission: func findKthPositive(arr []int, k int) int

import "fmt"

func main() {
	fmt.Println(KthMissingPositiveNumber([]int{2, 3, 4, 7, 11}, 5)) // 9
	fmt.Println(KthMissingPositiveNumber([]int{1, 2, 3, 4}, 2))     // 6
	fmt.Println(KthMissingPositiveNumber([]int{1, 3, 5}, 2))        // 4
}

// Time: O(log n), Space: O(1)
func KthMissingPositiveNumber(arr []int, k int) int {
	lo, hi := 0, len(arr)
	for lo < hi {
		mid := (lo + hi) / 2
		if arr[mid]-mid-1 < k {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo + k
}
```
