# 1065 — Index Pairs Of A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func indexPairs(text string, words []string) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Trie

**Kompleksitas Waktu:** O(n^2 + m*k)  
**Kompleksitas Ruang:** O(m*k) for trie

> **Untuk fresh graduate:** Kuasai dulu teknik **Trie** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1065: Index Pairs of a String
// https://leetcode.com/problems/index-pairs-of-a-string/
// Difficulty: Easy [Paid]
// Time: O(n^2 + m*k) | Space: O(m*k) for trie

import "fmt"

func main() {
	fmt.Println(indexPairs("thestoryofleetcodeandme", []string{"story", "fleet", "leetcode"}))
	// [[1,5],[3,7],[10,13],[10,18]]
	fmt.Println(indexPairs("ababa", []string{"aba", "ab"}))
	// [[0,1],[0,2],[2,3],[2,4]]
}

// LeetCode submission: indexPairs
func indexPairs(text string, words []string) [][]int {
  // Membuat map (HashMap) — pencarian O(1)
	wordSet := make(map[string]bool)
	for _, w := range words {
		wordSet[w] = true
	}
	var ans [][]int
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(text); i++ {
		for j := i; j < len(text); j++ {
			if wordSet[text[i:j+1]] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}
```
