# 0670 — Maximum Swap

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumSwap(num int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #670: Maximum Swap
// https://leetcode.com/problems/maximum-swap/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(maximumSwap(2736))
	fmt.Println(maximumSwap(9973))
	fmt.Println(maximumSwap(98368))
}

func maximumSwap(num int) int {
	s := []byte(strconv.Itoa(num))
	n := len(s)

	// Track last occurrence of each digit
  // Alokasi slice integer
	last := make([]int, 10)
	for i := 0; i < n; i++ {
		last[s[i]-'0'] = i
	}

	for i := 0; i < n; i++ {
		for d := 9; d > int(s[i]-'0'); d-- {
			if last[d] > i {
				s[i], s[last[d]] = s[last[d]], s[i]
				result, _ := strconv.Atoi(string(s))
				return result
			}
		}
	}

	return num
}
```
