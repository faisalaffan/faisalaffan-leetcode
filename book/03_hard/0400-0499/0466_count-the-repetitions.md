# 0466 — Count The Repetitions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

import "fmt"

// LeetCode #466: Count The Repetitions
// https://leetcode.com/problems/count-the-repetitions/
// Difficulty: Hard
//
// Find the maximum m such that [s2, m] is a subsequence of [s1, n1].
// Greedy matching through repeated s1, counting how many full s2 sequences
// are matched as a subsequence.

func main() {
	// Example: s1="acb", n1=4 => "acbacbacbacb"
	// s2="ab", n2=2 => "abab"
	// "abab" is a subsequence of "acbacbacbacb" => m=2
	fmt.Println("m:", getMaxRepetitions("acb", 4, "ab", 2)) // 2

	// Example 2
	fmt.Println("m:", getMaxRepetitions("abc", 4, "ab", 2)) // 2

	// Repeated char
	fmt.Println("m:", getMaxRepetitions("aaa", 3, "aa", 1)) // 4

	// No match
	fmt.Println("m:", getMaxRepetitions("a", 1, "b", 1)) // 0
}

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	if n1 == 0 {
		return 0
	}

	len1, len2 := len(s1), len(s2)
	totalChars := int64(n1) * int64(len1)

	matchCount := 0
	s2idx := 0

	for i := int64(0); i < totalChars; i++ {
		if s1[i%int64(len1)] == s2[s2idx] {
			s2idx++
			if s2idx == len2 {
				matchCount++
				s2idx = 0
			}
		}
	}
	return matchCount / n2
}
```
