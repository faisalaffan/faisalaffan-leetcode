# 0793 — Preimage Size Of Factorial Zeroes Function

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func preimageSizeFZF(K int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search, Monotonic Stack/Queue

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #793: Preimage Size of Factorial Zeroes Function
// https://leetcode.com/problems/preimage-size-of-factorial-zeroes-function/
// Difficulty: Hard
//
// Let f(x) be the number of trailing zeroes in x!. For a given K, find how
// many non-negative integers x have f(x) == K.
//   - f(x) is monotonic non-decreasing.
//   - f(x) jumps by more than 1 at multiples of 25, 125, etc., so some K
//     values are skipped entirely (answer 0). Otherwise the answer is 5.
//   - Binary search for the leftmost x with f(x) >= K and the leftmost x with
//     f(x) > K; the difference is the answer.

import "fmt"

func main() {
	// K=0: x in {0,1,2,3,4} → 5
	fmt.Println(preimageSizeFZF(0))
	// K=5: no x gives exactly 5 trailing zeros → 0
	fmt.Println(preimageSizeFZF(5))
	// K=3: x in {15,16,17,18,19} → 5
	fmt.Println(preimageSizeFZF(3))
	// K=1: x in {5,6,7,8,9} → 5
	fmt.Println(preimageSizeFZF(1))
	// K=1000000000: large input
	fmt.Println(preimageSizeFZF(1000000000))
}

func preimageSizeFZF(K int) int {
	// Find leftmost x such that trailingZeros(x) >= K
	lo, hi := 0, 5*K+5
	for lo < hi {
		mid := lo + (hi-lo)/2
		if trailingZeros(mid) < K {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	left := lo

	// Find leftmost x such that trailingZeros(x) > K
	lo, hi = 0, 5*K+5
	for lo < hi {
		mid := lo + (hi-lo)/2
		if trailingZeros(mid) <= K {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	right := lo

	return right - left
}

// trailingZeros returns the number of trailing zeroes in x!.
// Equals sum_{i=1..∞} floor(x / 5^i).
func trailingZeros(x int) int {
	count := 0
	for x >= 5 {
		x /= 5
		count += x
	}
	return count
}
```
