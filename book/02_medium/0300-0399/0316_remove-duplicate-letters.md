# 0316 — Remove Duplicate Letters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func removeDuplicateLetters(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #316: Remove Duplicate Letters
// https://leetcode.com/problems/remove-duplicate-letters/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func removeDuplicateLetters(s string) string {
	// Count last occurrence of each character
	lastOccur := [26]int{}
  // Range loop: iterasi dengan indeks + nilai
	for i := range s {
		lastOccur[s[i]-'a'] = i
	}

	stack := make([]byte, 0, len(s))
	seen := [26]bool{}

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if seen[ch-'a'] {
			continue
		}

		// Pop while stack top is greater and appears later
		for len(stack) > 0 && ch < stack[len(stack)-1] && i < lastOccur[stack[len(stack)-1]-'a'] {
			seen[stack[len(stack)-1]-'a'] = false
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, ch)
		seen[ch-'a'] = true
	}

	return string(stack)
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", removeDuplicateLetters("bcabc"))
	// Expected: "abc"

	// Test case 2
	fmt.Println("Test 2:", removeDuplicateLetters("cbacdcbc"))
	// Expected: "acdb"

	// Test case 3: Single character
	fmt.Println("Test 3:", removeDuplicateLetters("aaaa"))
	// Expected: "a"
}
```
