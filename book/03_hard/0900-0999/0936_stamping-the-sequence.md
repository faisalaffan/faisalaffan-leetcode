# 0936 — Stamping The Sequence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func movesToStamp(stamp string, target string) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #936: Stamping The Sequence
// https://leetcode.com/problems/stamping-the-sequence/
// Difficulty: Hard
// Reverse simulation: work backwards from target to "????...",
// trying to match stamp at each position. When matched, replace
// characters with '?'. Record stamped positions in reverse order.

import "fmt"

func movesToStamp(stamp string, target string) []int {
	s := []byte(stamp)
	t := []byte(target)
	m, n := len(s), len(t)
  // Alokasi slice integer
	result := make([]int, 0)

	// Count how many characters are not '?'
	remaining := n
	visited := make([]bool, n-m+1)

	// Check if stamp matches at position pos (consider '?' as wildcard)
	canStamp := func(pos int) bool {
		for i := 0; i < m; i++ {
			if t[pos+i] != '?' && t[pos+i] != s[i] {
				return false
			}
		}
		return true
	}

	// Apply stamp at position pos, return number of newly stamped characters
	applyStamp := func(pos int) int {
		cnt := 0
		for i := 0; i < m; i++ {
			if t[pos+i] != '?' {
				t[pos+i] = '?'
				cnt++
			}
		}
		return cnt
	}

	for remaining > 0 {
		found := false
		for pos := 0; pos <= n-m; pos++ {
			if visited[pos] {
				continue
			}
			if canStamp(pos) {
				visited[pos] = true
				cnt := applyStamp(pos)
				remaining -= cnt
				result = append(result, pos)
				found = true
				break
			}
		}
		if !found {
			return []int{}
		}
	}

	// Reverse result (we built from last to first)
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

func main() {
	fmt.Println(movesToStamp("abc", "ababc")) // Expected: [0,2]
	fmt.Println(movesToStamp("abca", "aabcaca")) // Expected: possible
	fmt.Println(movesToStamp("aye", "eyeye")) // Expected: [] (impossible?)
}
```
