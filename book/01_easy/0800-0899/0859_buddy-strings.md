# 0859 — Buddy Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func buddyStrings(s string, goal string) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n). Space: O(1).  |  **Ruang:** O(1).

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #859: Buddy Strings
// https://leetcode.com/problems/buddy-strings/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(buddyStrings("ab", "ba"))   // true
	fmt.Println(buddyStrings("ab", "ab"))   // false
	fmt.Println(buddyStrings("aa", "aa"))   // true
	fmt.Println(buddyStrings("abcd", "badc")) // false
}

// buddyStrings checks if swapping two letters in s makes it equal to goal.
// Time: O(n). Space: O(1).
func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	if s == goal {
		// Need at least one duplicate character to swap
  // HashMap: O(1) lookup
		seen := make(map[byte]bool)
  // Range loop
		for i := range s {
			if seen[s[i]] {
				return true
			}
			seen[s[i]] = true
		}
		return false
	}
  // Alokasi slice
	diff := make([]int, 0)
  // Range loop
	for i := range s {
		if s[i] != goal[i] {
			diff = append(diff, i)
			if len(diff) > 2 {
				return false
			}
		}
	}
	return len(diff) == 2 && s[diff[0]] == goal[diff[1]] && s[diff[1]] == goal[diff[0]]
}
```
