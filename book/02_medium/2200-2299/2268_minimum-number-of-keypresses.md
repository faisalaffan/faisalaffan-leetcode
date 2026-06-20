# 2268 — Minimum Number Of Keypresses

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumKeypresses(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + 26 log 26)  
**Kompleksitas Ruang:** O(26)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2268: Minimum Number of Keypresses
// https://leetcode.com/problems/minimum-number-of-keypresses/
// Difficulty: Medium [Paid]
// Time: O(n + 26 log 26) | Space: O(26)

import (
	"fmt"
	"sort"
)

func minimumKeypresses(s string) int {
  // Alokasi slice integer
	count := make([]int, 26)
	for _, ch := range s {
		count[ch-'a']++
	}

  // Custom sort dengan comparator
	sort.Slice(count, func(i, j int) bool {
		return count[i] > count[j]
	})

	presses := 0
	for i, c := range count {
		if c == 0 {
			break
		}
		presses += c * (i/9 + 1)
	}
	return presses
}

func main() {
	// Test case 1
	fmt.Println(minimumKeypresses("apple"))
	// Expected: 5

	// Test case 2
	fmt.Println(minimumKeypresses("abcdefghijkl"))
	// Expected: 15

	// Test case 3
	fmt.Println(minimumKeypresses("aaaaaaa"))
	// Expected: 7
}
```
