# 2575 — Find The Divisibility Array Of A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func divisibilityArray(word string, m int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2575: Find the Divisibility Array of a String
// https://leetcode.com/problems/find-the-divisibility-array-of-a-string/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func divisibilityArray(word string, m int) []int {
	n := len(word)
  // Alokasi slice integer
	ans := make([]int, n)
	rem := 0
	for i := 0; i < n; i++ {
		rem = (rem*10 + int(word[i]-'0')) % m
		if rem == 0 {
			ans[i] = 1
		} else {
			ans[i] = 0
		}
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", divisibilityArray("998244353", 3))
	// Expected: [1,1,0,0,0,1,1,0,0]

	// Test case 2
	fmt.Println("Test 2:", divisibilityArray("1010", 10))
	// Expected: [0,1,0,1]

	// Test case 3
	fmt.Println("Test 3:", divisibilityArray("10", 5))
	// Expected: [0,0]
}
```
