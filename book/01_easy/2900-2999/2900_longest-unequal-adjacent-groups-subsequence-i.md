# 2900 — Longest Unequal Adjacent Groups Subsequence I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func LongestUnequalAdjacentGroupsSubsequenceI(n int, words []string, groups []int) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2900: Longest Unequal Adjacent Groups Subsequence I
// https://leetcode.com/problems/longest-unequal-adjacent-groups-subsequence-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: getWordsInLongestSubsequence
	fmt.Println(LongestUnequalAdjacentGroupsSubsequenceI(3, []string{"e", "a", "b"}, []int{0, 0, 1})) // ["e","b"]
	fmt.Println(LongestUnequalAdjacentGroupsSubsequenceI(4, []string{"a", "b", "c", "d"}, []int{1, 0, 1, 0})) // ["a","b","c","d"]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: getWordsInLongestSubsequence
func LongestUnequalAdjacentGroupsSubsequenceI(n int, words []string, groups []int) []string {
	result := []string{words[0]}
	for i := 1; i < n; i++ {
		if groups[i] != groups[i-1] {
			result = append(result, words[i])
		}
	}
	return result
}
```
