# 0893 — Groups Of Special Equivalent Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func GroupsOfSpecialEquivalentStrings(words []string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n * m) where n = len(words), m = avg word length  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #893: Groups of Special-Equivalent Strings
// https://leetcode.com/problems/groups-of-special-equivalent-strings/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(GroupsOfSpecialEquivalentStrings([]string{"abcd", "cdab", "cbad", "xyzz", "zzxy", "zzyx"}))
	fmt.Println(GroupsOfSpecialEquivalentStrings([]string{"abc", "acb", "bac", "bca", "cab", "cba"}))
	fmt.Println(GroupsOfSpecialEquivalentStrings([]string{"a"}))
}

// Time: O(n * m) where n = len(words), m = avg word length | Space: O(n)
func GroupsOfSpecialEquivalentStrings(words []string) int {
  // Membuat map (HashMap) — pencarian O(1)
	groups := make(map[[52]int]bool)

	for _, word := range words {
		var key [52]int
		for i, c := range word {
			// Even indices: 0-25, Odd indices: 26-51
			key[int(c-'a')+26*(i%2)]++
		}
		groups[key] = true
	}

	return len(groups)
}
```
