# 0651 — 4 Keys Keyboard

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxA(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #651: 4 Keys Keyboard
// https://leetcode.com/problems/4-keys-keyboard/
// Difficulty: Medium [Paid]
// Time: O(n^2)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(maxA(3))
	fmt.Println(maxA(7))
	fmt.Println(maxA(10))
}

func maxA(n int) int {
  // Alokasi slice integer
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		for j := 1; j < i-1; j++ {
			dp[i] = max(dp[i], dp[j]*(i-j-1))
		}
	}
	return dp[n]
}
```
