# 0424 — Longest Repeating Character Replacement

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func characterReplacement(s string, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Sliding Window

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #424: Longest Repeating Character Replacement
// https://leetcode.com/problems/longest-repeating-character-replacement/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func characterReplacement(s string, k int) int {
	freq := [26]int{}
	left, maxFreq, maxLen := 0, 0, 0

	for right := 0; right < len(s); right++ {
		freq[s[right]-'A']++
		if freq[s[right]-'A'] > maxFreq {
			maxFreq = freq[s[right]-'A']
		}

		// Window size - maxFreq = chars to replace
		for right-left+1-maxFreq > k {
			freq[s[left]-'A']--
			left++
			// Recompute maxFreq (or keep old - it's safe)
		}

		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", characterReplacement("ABAB", 2))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", characterReplacement("AABABBA", 1))
	// Expected: 4

	// Test case 3
	fmt.Println("Test 3:", characterReplacement("AAAA", 2))
	// Expected: 4
}
```
