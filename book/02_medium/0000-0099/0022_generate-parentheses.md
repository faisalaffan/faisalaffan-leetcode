# 0022 — Generate Parentheses

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func generateParenthesis(n int) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Backtracking

**Kompleksitas Waktu:** O(4^n / sqrt(n))  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Backtracking** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #22: Generate Parentheses
// https://leetcode.com/problems/generate-parentheses/
// Difficulty: Medium

import "fmt"

func generateParenthesis(n int) []string {
	result := []string{}
	var backtrack func(curr string, open, close int)
	backtrack = func(curr string, open, close int) {
		if len(curr) == 2*n {
			result = append(result, curr)
			return
		}
		if open < n {
			backtrack(curr+"(", open+1, close)
		}
		if close < open {
			backtrack(curr+")", open, close+1)
		}
	}
	backtrack("", 0, 0)
	return result
}

func main() {
	// Test case 1
	fmt.Println(generateParenthesis(3)) // ["((()))","(()())","(())()","()(())","()()()"]

	// Test case 2
	fmt.Println(generateParenthesis(1)) // ["()"]

	// Test case 3
	fmt.Println(generateParenthesis(2)) // ["(())","()()"]
}

// Time: O(4^n / sqrt(n)) | Space: O(n)
```
