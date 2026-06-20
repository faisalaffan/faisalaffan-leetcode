# 1021 — Remove Outermost Parentheses

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func removeOuterParentheses(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1021: Remove Outermost Parentheses
// https://leetcode.com/problems/remove-outermost-parentheses/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(removeOuterParentheses("(()())(())"))          // "()()()"
	fmt.Println(removeOuterParentheses("(()())(())(()(()))"))  // "()()()()(())"
	fmt.Println(removeOuterParentheses("()()"))                // ""
}

// LeetCode submission: removeOuterParentheses
func removeOuterParentheses(s string) string {
	ans := make([]byte, 0, len(s))
	depth := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			if depth > 0 {
				ans = append(ans, '(')
			}
			depth++
		} else {
			depth--
			if depth > 0 {
				ans = append(ans, ')')
			}
		}
	}
	return string(ans)
}
```
