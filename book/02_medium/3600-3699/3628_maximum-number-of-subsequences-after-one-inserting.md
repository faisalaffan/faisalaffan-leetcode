# 3628 — Maximum Number Of Subsequences After One Inserting

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumNumberOfSubsequencesAfterOneInserting(s string, pattern string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3628: Maximum Number of Subsequences After One Inserting
// https://leetcode.com/problems/maximum-number-of-subsequences-after-one-inserting/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", MaximumNumberOfSubsequencesAfterOneInserting("abc", "abc"))
	// Test case 2
	fmt.Println("Test 2:", MaximumNumberOfSubsequencesAfterOneInserting("ab", "ab"))
	// Test case 3
	fmt.Println("Test 3:", MaximumNumberOfSubsequencesAfterOneInserting("a", "a"))
}

func MaximumNumberOfSubsequencesAfterOneInserting(s string, pattern string) int {
	// Count max subsequences of pattern in s after inserting one character
	if len(pattern) == 0 {
		return len(s) + 1
	}
	if len(pattern) == 1 {
		// Count existing, then add max possible after insert
		count := 0
		for _, c := range s {
			if byte(c) == pattern[0] {
				count++
			}
		}
		// Insert one more gives count+1
		maxAdditional := count + 1
		if maxAdditional > len(s)+1-count {
			return maxAdditional
		}
		return maxAdditional
	}

	p0, p1 := pattern[0], pattern[1]

	// Count existing
	existing := 0
	countP1 := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == p1 {
			countP1++
		} else if s[i] == p0 {
			existing += countP1
		}
	}

	// Max additional from inserting one char
	addP0 := 0
	for _, c := range s {
		if byte(c) == p1 {
			addP0++
		}
	}
	addP1 := 0
	for _, c := range s {
		if byte(c) == p0 {
			addP1++
		}
	}
	if addP0 > addP1 {
		return existing + addP0
	}
	return existing + addP1
}
```
