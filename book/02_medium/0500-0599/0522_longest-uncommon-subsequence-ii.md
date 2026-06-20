# 0522 — Longest Uncommon Subsequence Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindLUSlength(strs []string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2 * L) where L is max length  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #522: Longest Uncommon Subsequence II
// https://leetcode.com/problems/longest-uncommon-subsequence-ii/
// Difficulty: Medium
// Time: O(n^2 * L) where L is max length
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(FindLUSlength([]string{"aba", "cdc", "eae"}))
	fmt.Println(FindLUSlength([]string{"aaa", "aaa", "aa"}))
}

func FindLUSlength(strs []string) int {
	maxLen := -1

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(strs); i++ {
		isUnique := true
		for j := 0; j < len(strs); j++ {
			if i != j && isSubseq(strs[i], strs[j]) {
				isUnique = false
				break
			}
		}
		if isUnique && len(strs[i]) > maxLen {
			maxLen = len(strs[i])
		}
	}

	return maxLen
}

func isSubseq(a, b string) bool {
	i := 0
	for j := 0; i < len(a) && j < len(b); j++ {
		if a[i] == b[j] {
			i++
		}
	}
	return i == len(a)
}
```
