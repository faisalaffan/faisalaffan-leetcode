# 3110 — Score Of A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func ScoreOfAString(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3110: Score of a String
// https://leetcode.com/problems/score-of-a-string/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: scoreOfString
	fmt.Println(ScoreOfAString("hello")) // 13
	fmt.Println(ScoreOfAString("zaz"))   // 50
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: scoreOfString
func ScoreOfAString(s string) int {
	score := 0
	for i := 1; i < len(s); i++ {
		diff := int(s[i]) - int(s[i-1])
		if diff < 0 {
			diff = -diff
		}
		score += diff
	}
	return score
}
```
