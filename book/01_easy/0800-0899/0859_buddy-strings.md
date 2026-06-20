# 0859 — Buddy Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func buddyStrings(s string, goal string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
  // Membuat map (HashMap) — pencarian O(1)
		seen := make(map[byte]bool)
  // Range loop: iterasi dengan indeks + nilai
		for i := range s {
			if seen[s[i]] {
				return true
			}
			seen[s[i]] = true
		}
		return false
	}
  // Alokasi slice integer
	diff := make([]int, 0)
  // Range loop: iterasi dengan indeks + nilai
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
