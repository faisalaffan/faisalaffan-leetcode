# 3460 — Longest Common Prefix After At Most One Removal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func longestCommonPrefix(s string, t string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(min(n,m)) Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3460: Longest Common Prefix After at Most One Removal
// https://leetcode.com/problems/longest-common-prefix-after-at-most-one-removal/
// Difficulty: Medium [Paid]
// Time: O(min(n,m)) Space: O(1)

import "fmt"

func longestCommonPrefix(s string, t string) int {
	n, m := len(s), len(t)
	maxLen := 0

	// without removal
	i, j := 0, 0
	for i < n && j < m && s[i] == t[j] {
		i++
		j++
	}
	maxLen = i

	// with one removal from s
	i, j = 0, 0
	removed := false
	for i < n && j < m {
		if s[i] == t[j] {
			i++
			j++
		} else if !removed {
			i++
			removed = true
		} else {
			break
		}
	}
	if j > maxLen {
		maxLen = j
	}

	return maxLen
}

func main() {
	fmt.Println(longestCommonPrefix("abcde", "abfde")) // 2 (ab)
	fmt.Println(longestCommonPrefix("abc", "abc"))     // 3
	fmt.Println(longestCommonPrefix("abcd", "abxd"))   // 3 (abx vs abc — remove c from s -> abxd vs abxd... actually ab)
}
```
