# 1111 — Maximum Nesting Depth Of Two Valid Parentheses Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxDepthAfterSplit(seq string) []int
```

> **💡 Hint:** Assign '(' to group A or B based on even/odd depth.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1111: Maximum Nesting Depth of Two Valid Parentheses Strings
// https://leetcode.com/problems/maximum-nesting-depth-of-two-valid-parentheses-strings/
// Difficulty: Medium
//
// Approach: Assign '(' to group A or B based on even/odd depth.
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(maxDepthAfterSplit("(()())")) // [0,1,1,1,1,0] or similar
	fmt.Println(maxDepthAfterSplit("()(())()")) // [0,0,0,1,1,0,0,0]
}

func maxDepthAfterSplit(seq string) []int {
  // Alokasi slice integer
	result := make([]int, len(seq))
	depth := 0

	for i, c := range seq {
		if c == '(' {
			depth++
			result[i] = depth % 2
		} else {
			result[i] = depth % 2
			depth--
		}
	}

	return result
}
```
