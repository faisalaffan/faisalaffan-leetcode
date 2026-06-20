# 3307 — Find The K Th Character In String Game Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func kthCharacter(k int, operations []int) byte
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3307: Find the K-th Character in String Game II
// https://leetcode.com/problems/find-the-k-th-character-in-string-game-ii/
// Difficulty: Hard
//
// Binary search from the end. Simulate the string lengths (doubling each
// operation). Work backwards: at each operation, if k is in the newly added
// half, map it back to the corresponding position in the first half and
// accumulate the shift. The shift per operation is operations[i] + 1.
//
// The string game: start with "a". For operation v:
//   new_str = str + shift(str, v+1)  where shift(s, n) shifts each char by n.
// The length doubles each operation.

import "fmt"

func main() {
	// Example 1: k=5, operations=[0,0,0] => "b"
	//   (0-indexed k=5 = position 5; base char 'a' shifted by 1 step = 'b')
	fmt.Println(kthCharacter(5, []int{0, 0, 0}))
	// Example 2: single operation
	fmt.Println(kthCharacter(2, []int{1}))
	// Example 3: k=1 always returns "a"
	fmt.Println(kthCharacter(1, []int{0, 0, 0, 0}))
	// Example 4: mixed operations
	fmt.Println(kthCharacter(10, []int{0, 1, 0, 1}))
	// Example 5: larger shift
	fmt.Println(kthCharacter(4, []int{1, 0}))
}

func kthCharacter(k int, operations []int) byte {
	// k is 1-indexed
	m := len(operations)

	// Precompute lengths: len[i] = length after operation i
  // Alokasi slice integer
	lengths := make([]int, m+1)
	lengths[0] = 1 // initial string "a"
	for i := 1; i <= m; i++ {
		lengths[i] = lengths[i-1] * 2
	}

	totalShift := 0
	pos := k

	// Work backwards through operations
	for i := m - 1; i >= 0; i-- {
		half := lengths[i] // length before this operation
		if pos > half {
			// k is in the newly added half
			pos -= half
			totalShift += operations[i] + 1
		}
	}

	// Base character is 'a'
	return byte('a' + (totalShift % 26))
}
```
