# 1996 — The Number Of Weak Characters In The Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func TheNumberOfWeakCharactersInTheGame(properties [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n), Space: O(1) (ignoring sort space)  
**Kompleksitas Ruang:** O(1) (ignoring sort space)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1996: The Number of Weak Characters in the Game
// https://leetcode.com/problems/the-number-of-weak-characters-in-the-game/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(TheNumberOfWeakCharactersInTheGame([][]int{{5, 5}, {6, 3}, {3, 6}}))
	fmt.Println(TheNumberOfWeakCharactersInTheGame([][]int{{2, 2}, {3, 3}}))
	fmt.Println(TheNumberOfWeakCharactersInTheGame([][]int{{1, 5}, {10, 4}, {4, 3}}))
}

// Time: O(n log n), Space: O(1) (ignoring sort space)
func TheNumberOfWeakCharactersInTheGame(properties [][]int) int {
  // Custom sort dengan comparator
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] == properties[j][0] {
			return properties[i][1] < properties[j][1]
		}
		return properties[i][0] > properties[j][0]
	})

	ans := 0
	maxDef := 0
	for _, p := range properties {
		if p[1] < maxDef {
			ans++
		} else {
			maxDef = p[1]
		}
	}
	return ans
}
```
