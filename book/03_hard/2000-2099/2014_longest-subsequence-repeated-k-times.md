# 2014 — Longest Subsequence Repeated K Times

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func longestSubsequenceRepeatedK(s string, k int) string
```

> **💡 Hint:** BFS generate candidate strings in order of length.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2014: Longest Subsequence Repeated k Times
// https://leetcode.com/problems/longest-subsequence-repeated-k-times/
// Difficulty: Hard
// Approach: BFS generate candidate strings in order of length.
// Count character frequencies, max_uses = freq / k.
// Generate all possible strings up to n/k length.
// For each candidate, check if repeated k times is a subsequence of s.
// Keep the longest.

import "fmt"

func longestSubsequenceRepeatedK(s string, k int) string {
	// Count frequencies
  // Alokasi slice integer
	freq := make([]int, 26)
	for _, ch := range s {
		freq[ch-'a']++
	}

	// Max uses for each character
  // Alokasi slice integer
	maxUses := make([]int, 26)
	for i := 0; i < 26; i++ {
		maxUses[i] = freq[i] / k
	}

	// Check if str is a subsequence of s
	isSubseq := func(str string) bool {
		j := 0
  // Loop linear O(n): iterasi setiap elemen
		for i := 0; i < len(s) && j < len(str); i++ {
			if s[i] == str[j] {
				j++
			}
		}
		return j == len(str)
	}

	// Check if t repeated k times is a subsequence of s
	check := func(t string) bool {
		if len(t) == 0 {
			return false
		}
		concat := ""
		for i := 0; i < k; i++ {
			concat += t
		}
		return isSubseq(concat)
	}

	// BFS to generate candidates
	queue := []string{""}
	best := ""

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for c := 0; c < 26; c++ {
			if maxUses[c] == 0 {
				continue
			}
			next := cur + string(rune('a'+c))
			if !check(next) {
				continue
			}
			queue = append(queue, next)
			if len(next) > len(best) || (len(next) == len(best) && next > best) {
				best = next
			}
		}
	}

	return best
}

func main() {
	// Example: "letsleetcode", k=2 -> "let"
	fmt.Println(longestSubsequenceRepeatedK("letsleetcode", 2))

	// Additional tests
	fmt.Println(longestSubsequenceRepeatedK("aabbaabbaabb", 3))
	fmt.Println(longestSubsequenceRepeatedK("abcd", 2))
}
```
