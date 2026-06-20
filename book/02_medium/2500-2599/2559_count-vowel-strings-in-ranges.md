# 2559 — Count Vowel Strings In Ranges

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func vowelStrings(words []string, queries [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n + q)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2559: Count Vowel Strings in Ranges
// https://leetcode.com/problems/count-vowel-strings-in-ranges/
// Difficulty: Medium
// Time: O(n + q) | Space: O(n)

import "fmt"

func vowelStrings(words []string, queries [][]int) []int {
	isVowel := func(ch byte) bool {
		return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
	}

	n := len(words)
  // Alokasi slice integer
	prefix := make([]int, n+1)
	for i, w := range words {
		prefix[i+1] = prefix[i]
		if len(w) > 0 && isVowel(w[0]) && isVowel(w[len(w)-1]) {
			prefix[i+1]++
		}
	}

  // Alokasi slice integer
	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = prefix[q[1]+1] - prefix[q[0]]
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", vowelStrings([]string{"aba", "bcb", "ece", "aa", "e"}, [][]int{{0, 2}, {1, 4}, {1, 1}}))
	// Expected: [2,3,0]

	// Test case 2
	fmt.Println("Test 2:", vowelStrings([]string{"a", "e", "i"}, [][]int{{0, 2}, {0, 0}, {2, 2}}))
	// Expected: [3,1,1]

	// Test case 3: empty range
	fmt.Println("Test 3:", vowelStrings([]string{"a", "b", "c"}, [][]int{{1, 1}}))
	// Expected: [0]
}
```
