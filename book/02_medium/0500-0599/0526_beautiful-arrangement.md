# 0526 — Beautiful Arrangement

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CountArrangement(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Backtracking

**Kompleksitas Waktu:** O(k) where k = number of valid permutations  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Backtracking** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #526: Beautiful Arrangement
// https://leetcode.com/problems/beautiful-arrangement/
// Difficulty: Medium
// Time: O(k) where k = number of valid permutations
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(CountArrangement(2))
	fmt.Println(CountArrangement(1))
}

func CountArrangement(n int) int {
	used := make([]bool, n+1)
	count := 0

	var backtrack func(pos int)
	backtrack = func(pos int) {
		if pos > n {
			count++
			return
		}
		for i := 1; i <= n; i++ {
			if !used[i] && (i%pos == 0 || pos%i == 0) {
				used[i] = true
				backtrack(pos + 1)
				used[i] = false
			}
		}
	}

	backtrack(1)
	return count
}
```
