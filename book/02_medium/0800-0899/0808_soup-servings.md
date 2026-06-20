# 0808 — Soup Servings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func SoupServings(n int) float64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS, Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #808: Soup Servings
// https://leetcode.com/problems/soup-servings/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(SoupServings(50))
	fmt.Println(SoupServings(100))
	fmt.Println(SoupServings(800))
}

func SoupServings(n int) float64 {
	if n > 4800 {
		return 1.0
	}
	n = (n + 24) / 25

  // Membuat matriks/slice 2D untuk DP
	memo := make([][]float64, n+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range memo {
		memo[i] = make([]float64, n+1)
	}
	var dfs func(int, int) float64
	dfs = func(a, b int) float64 {
		if a <= 0 && b <= 0 {
			return 0.5
		}
		if a <= 0 {
			return 1.0
		}
		if b <= 0 {
			return 0.0
		}
		if memo[a][b] > 0 {
			return memo[a][b]
		}
		memo[a][b] = 0.25 * (dfs(a-4, b) + dfs(a-3, b-1) + dfs(a-2, b-2) + dfs(a-1, b-3))
		return memo[a][b]
	}

	return dfs(n, n)
}
```
