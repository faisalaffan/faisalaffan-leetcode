# 0556 — Next Greater Element Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func NextGreaterElementIii(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(log n) = O(number of digits)  
**Kompleksitas Ruang:** O(log n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #556: Next Greater Element III
// https://leetcode.com/problems/next-greater-element-iii/
// Difficulty: Medium
// Time: O(log n) = O(number of digits)
// Space: O(log n)

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(NextGreaterElementIii(12))
	fmt.Println(NextGreaterElementIii(21))
	fmt.Println(NextGreaterElementIii(1234))
}

func NextGreaterElementIii(n int) int {
	s := []byte(strconv.Itoa(n))

	// Find first decreasing digit from right
	i := len(s) - 2
	for i >= 0 && s[i] >= s[i+1] {
		i--
	}
	if i < 0 {
		return -1
	}

	// Find smallest digit larger than s[i] from right
	j := len(s) - 1
	for s[j] <= s[i] {
		j--
	}

	s[i], s[j] = s[j], s[i]

	// Reverse suffix
	left, right := i+1, len(s)-1
  // Two-pointer: gerakkan kiri atau kanan
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}

	result, _ := strconv.Atoi(string(s))
	if result > math.MaxInt32 {
		return -1
	}
	return result
}
```
