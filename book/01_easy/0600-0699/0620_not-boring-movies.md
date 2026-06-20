# 0620 — Not Boring Movies

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func NotBoringMovies() string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #620: Not Boring Movies
// https://leetcode.com/problems/not-boring-movies/
// Difficulty: Easy

import "fmt"

func NotBoringMovies() string {
	return "SELECT * FROM Cinema WHERE id % 2 = 1 AND description != 'boring' ORDER BY rating DESC"
}

func main() {
	fmt.Println(NotBoringMovies())
}
```
