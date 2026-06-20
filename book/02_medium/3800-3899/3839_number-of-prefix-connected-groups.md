# 3839 — Number Of Prefix Connected Groups

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func NumberOfPrefixConnectedGroups(words []string, k int) int
```

> **💡 Hint:** Group words by first k characters (prefix). Count groups with >= 2 words.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(N * K)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3839: Number of Prefix Connected Groups
// https://leetcode.com/problems/number-of-prefix-connected-groups/
// Difficulty: Medium
// Time: O(N * K) | Space: O(N)
// Approach: Group words by first k characters (prefix). Count groups with >= 2 words.

import "fmt"

func NumberOfPrefixConnectedGroups(words []string, k int) int {
  // Membuat map (HashMap) — pencarian O(1)
	prefixCount := make(map[string]int)

	for _, w := range words {
		if len(w) < k {
			continue
		}
		prefixCount[w[:k]]++
	}

	ans := 0
	for _, cnt := range prefixCount {
		if cnt >= 2 {
			ans++
		}
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println(NumberOfPrefixConnectedGroups([]string{"apple", "apply", "banana", "bandit"}, 2)) // Expected: 2

	// Example 2
	fmt.Println(NumberOfPrefixConnectedGroups([]string{"car", "cat", "cartoon"}, 3)) // Expected: 1

	// Example 3
	fmt.Println(NumberOfPrefixConnectedGroups([]string{"bat", "dog", "dog", "doggy", "bat"}, 3)) // Expected: 2
}
```
