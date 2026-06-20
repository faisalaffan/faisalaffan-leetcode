# 1055 — Shortest Way To Form String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func shortestWay(source string, target string) int
```

> **💡 Hint:** Greedy with two pointers. Count subsequence matches.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * m) worst case, O(n + m) with precomputed indices  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1055: Shortest Way to Form String
// https://leetcode.com/problems/shortest-way-to-form-string/
// Difficulty: Medium
//
// Approach: Greedy with two pointers. Count subsequence matches.
// Time: O(n * m) worst case, O(n + m) with precomputed indices
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(shortestWay("abc", "abcbc"))   // 2
	fmt.Println(shortestWay("abc", "acdbc"))   // -1
	fmt.Println(shortestWay("xyz", "xzyxz"))   // 2
}

func shortestWay(source string, target string) int {
	count := 0
	i := 0 // index in target

	// Pre-check: all chars in target must exist in source
  // Membuat map (HashMap) — pencarian O(1)
	sourceSet := make(map[byte]bool)
	for k := 0; k < len(source); k++ {
		sourceSet[source[k]] = true
	}
	for k := 0; k < len(target); k++ {
		if !sourceSet[target[k]] {
			return -1
		}
	}

	for i < len(target) {
		count++
		j := 0 // index in source
		for j < len(source) && i < len(target) {
			if source[j] == target[i] {
				i++
			}
			j++
		}
	}

	return count
}
```
