# 0717 — 1 Bit And 2 Bit Characters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func isOneBitCharacter(bits []int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n). Space: O(1).  |  **Ruang:** O(1).


## 💻 Solusi Go

```go
package main

// LeetCode #717: 1-bit and 2-bit Characters
// https://leetcode.com/problems/1-bit-and-2-bit-characters/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(isOneBitCharacter([]int{1, 0, 0}))          // true
	fmt.Println(isOneBitCharacter([]int{1, 1, 1, 0}))      // false
	fmt.Println(isOneBitCharacter([]int{0}))                // true
}

// isOneBitCharacter checks if the last character must be a one-bit character.
// Time: O(n). Space: O(1).
func isOneBitCharacter(bits []int) bool {
	i := 0
	for i < len(bits)-1 {
		if bits[i] == 1 {
			i += 2
		} else {
			i++
		}
	}
	return i == len(bits)-1
}
```
