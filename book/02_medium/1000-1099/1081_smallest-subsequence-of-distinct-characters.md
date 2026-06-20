# 1081 — Smallest Subsequence Of Distinct Characters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func smallestSubsequence(s string) string
```

> **💡 Hint:** Monotonic stack (greedy). Track last occurrence and used set.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack, Monotonic Stack/Queue

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(26) = O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1081: Smallest Subsequence of Distinct Characters
// https://leetcode.com/problems/smallest-subsequence-of-distinct-characters/
// Difficulty: Medium
//
// Approach: Monotonic stack (greedy). Track last occurrence and used set.
// Time: O(n)
// Space: O(26) = O(1)

import "fmt"

func main() {
	fmt.Println(smallestSubsequence("bcabc"))  // "abc"
	fmt.Println(smallestSubsequence("cbacdcbc")) // "acdb"
}

func smallestSubsequence(s string) string {
  // Alokasi slice integer
	lastOccur := make([]int, 26)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		lastOccur[s[i]-'a'] = i
	}

	used := make([]bool, 26)
	stack := make([]byte, 0)

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if used[ch-'a'] {
			continue
		}

		for len(stack) > 0 && stack[len(stack)-1] > ch && lastOccur[stack[len(stack)-1]-'a'] > i {
			used[stack[len(stack)-1]-'a'] = false
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, ch)
		used[ch-'a'] = true
	}

	return string(stack)
}
```
