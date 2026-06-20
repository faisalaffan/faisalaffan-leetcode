# 2131 — Longest Palindrome By Concatenating Two Letter Words

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func longestPalindrome(words []string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2131: Longest Palindrome by Concatenating Two-Letter Words
// https://leetcode.com/problems/longest-palindrome-by-concatenating-two-letter-words/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func longestPalindrome(words []string) int {
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[string]int)
	for _, w := range words {
		freq[w]++
	}

	length := 0
	centerUsed := false

	for w, count := range freq {
		if count == 0 {
			continue
		}
		rev := string([]byte{w[1], w[0]})

		if w == rev {
			// Same letter pair like "aa"
			pairs := count / 2
			length += pairs * 4
			if count%2 == 1 && !centerUsed {
				length += 2
				centerUsed = true
			}
			freq[w] = 0
		} else if revCount, ok := freq[rev]; ok && revCount > 0 {
			pairs := count
			if revCount < pairs {
				pairs = revCount
			}
			length += pairs * 4
			freq[w] -= pairs
			freq[rev] -= pairs
		}
	}

	return length
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", longestPalindrome([]string{"lc", "cl", "gg"}))
	// Expected: 6

	// Test case 2
	fmt.Println("Test 2:", longestPalindrome([]string{"ab", "ty", "yt", "lc", "cl", "ab"}))
	// Expected: 8

	// Test case 3
	fmt.Println("Test 3:", longestPalindrome([]string{"cc", "ll", "xx"}))
	// Expected: 2
}
```
