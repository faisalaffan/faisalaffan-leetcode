# 1805 — Number Of Different Integers In A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func NumDifferentIntegers(word string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1805: Number of Different Integers in a String
// https://leetcode.com/problems/number-of-different-integers-in-a-string/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func NumDifferentIntegers(word string) int {
  // HashMap: O(1) lookup
	seen := make(map[string]bool)
	i := 0
	for i < len(word) {
		if word[i] >= '0' && word[i] <= '9' {
			j := i
			for j < len(word) && word[j] >= '0' && word[j] <= '9' {
				j++
			}
			for i < j && word[i] == '0' {
				i++
			}
			num := word[i:j]
			if !seen[num] {
				seen[num] = true
			}
			i = j
		} else {
			i++
		}
	}
	return len(seen)
}

func main() {
	fmt.Println(NumDifferentIntegers("a123bc34d8ef34"))
	fmt.Println(NumDifferentIntegers("leet1234code234"))
	fmt.Println(NumDifferentIntegers("a1b01c001"))
}
```
