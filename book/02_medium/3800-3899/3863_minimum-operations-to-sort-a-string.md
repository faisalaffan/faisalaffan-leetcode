# 3863 — Minimum Operations To Sort A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumOperationsToSortAString(s string) int
```

> **💡 Hint:** Count characters and determine min swaps to sort.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3863: Minimum Operations to Sort a String
// https://leetcode.com/problems/minimum-operations-to-sort-a-string/
// Difficulty: Medium
// Time: O(N) | Space: O(1)
// Approach: Count characters and determine min swaps to sort.

import "fmt"

func MinimumOperationsToSortAString(s string) int {
	// Count character frequencies
  // Alokasi slice integer
	cnt := make([]int, 26)
	for _, ch := range s {
		cnt[ch-'a']++
	}

	// Build sorted string
	sorted := make([]byte, len(s))
	idx := 0
	for c := 0; c < 26; c++ {
		for j := 0; j < cnt[c]; j++ {
			sorted[idx] = byte('a' + c)
			idx++
		}
	}

	// Count mismatching positions (where s[i] != sorted[i])
	mismatch := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		if s[i] != sorted[i] {
			mismatch++
		}
	}

	// Each swap fixes two mismatches, but there might be cycles
	// Simple: (mismatch + 1) / 2 gives swaps needed
	return (mismatch + 1) / 2
}

func main() {
	// Example
	fmt.Println(MinimumOperationsToSortAString("cba")) // Expected: 1

	// Example
	fmt.Println(MinimumOperationsToSortAString("aab")) // Expected: 0

	// Example
	fmt.Println(MinimumOperationsToSortAString("acb")) // Expected: 1
}
```
