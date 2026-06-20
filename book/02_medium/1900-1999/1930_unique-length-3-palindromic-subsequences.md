# 1930 — Unique Length 3 Palindromic Subsequences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CountPalindromicSubsequence(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * 26) = O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1930: Unique Length-3 Palindromic Subsequences
// https://leetcode.com/problems/unique-length-3-palindromic-subsequences/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(CountPalindromicSubsequence("aabca"))
	fmt.Println(CountPalindromicSubsequence("adc"))
	fmt.Println(CountPalindromicSubsequence("bbcbaba"))
}

// Time: O(n * 26) = O(n), Space: O(1)
func CountPalindromicSubsequence(s string) int {
	// For each character, find first and last occurrence
  // Alokasi slice integer
	first := make([]int, 26)
  // Alokasi slice integer
	last := make([]int, 26)
	for i := 0; i < 26; i++ {
		first[i] = -1
		last[i] = -1
	}
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		idx := int(s[i] - 'a')
		if first[idx] == -1 {
			first[idx] = i
		}
		last[idx] = i
	}

	count := 0
	for c := 0; c < 26; c++ {
		if first[c] != -1 && last[c]-first[c] > 1 {
			// Count unique characters between first and last occurrence
			seen := make([]bool, 26)
			for i := first[c] + 1; i < last[c]; i++ {
				seen[s[i]-'a'] = true
			}
			for _, v := range seen {
				if v {
					count++
				}
			}
		}
	}
	return count
}
```
