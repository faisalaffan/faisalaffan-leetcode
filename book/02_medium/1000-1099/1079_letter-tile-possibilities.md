# 1079 — Letter Tile Possibilities

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func numTilePossibilities(tiles string) int
```

> **💡 Hint:** Backtracking with frequency count

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS, Backtracking

**Kompleksitas Waktu:** O(n!) where n = len(tiles) worst case  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1079: Letter Tile Possibilities
// https://leetcode.com/problems/letter-tile-possibilities/
// Difficulty: Medium
//
// Approach: Backtracking with frequency count
// Time: O(n!) where n = len(tiles) worst case
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(numTilePossibilities("AAB")) // 8
	fmt.Println(numTilePossibilities("AAABBC")) // 188
}

func numTilePossibilities(tiles string) int {
  // Alokasi slice integer
	freq := make([]int, 26)
	for _, c := range tiles {
		freq[c-'A']++
	}

	var dfs func() int
	dfs = func() int {
		count := 0
		for i := 0; i < 26; i++ {
			if freq[i] == 0 {
				continue
			}
			count++
			freq[i]--
			count += dfs()
			freq[i]++
		}
		return count
	}

	return dfs()
}
```
