# 1614 — Maximum Nesting Depth Of The Parentheses

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumNestingDepthOfTheParentheses(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1614: Maximum Nesting Depth of the Parentheses
// https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func MaximumNestingDepthOfTheParentheses(s string) int {
	maxDepth, currentDepth := 0, 0
	for _, ch := range s {
		if ch == '(' {
			currentDepth++
			if currentDepth > maxDepth {
				maxDepth = currentDepth
			}
		} else if ch == ')' {
			currentDepth--
		}
	}
	return maxDepth
}

func main() {
	fmt.Println(MaximumNestingDepthOfTheParentheses("(1+(2*3)+((8)/4))+1"))
	fmt.Println(MaximumNestingDepthOfTheParentheses("(1)+((2))+(((3)))"))
	fmt.Println(MaximumNestingDepthOfTheParentheses(""))
}
```
