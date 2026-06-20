# 1416 — Restore The Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func numberOfArrays(s string, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1416: Restore The Array
// https://leetcode.com/problems/restore-the-array/
// Difficulty: Hard

import "fmt"

const mod1416 = 1_000_000_007

func numberOfArrays(s string, k int) int {
	n := len(s)
  // Alokasi slice integer
	dp := make([]int, n+1)
	dp[n] = 1
	for i := n - 1; i >= 0; i-- {
		if s[i] == '0' {
			continue
		}
		var num int
		for j := i; j < n; j++ {
			num = num*10 + int(s[j]-'0')
			if num > k {
				break
			}
			dp[i] = (dp[i] + dp[j+1]) % mod1416
		}
	}
	return dp[0]
}

func main() {
	// Example: "1317", 2000 -> 8
	fmt.Println(numberOfArrays("1317", 2000))
}
```
