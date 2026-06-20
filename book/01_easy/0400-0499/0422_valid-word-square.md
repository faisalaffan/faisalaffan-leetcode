# 0422 — Valid Word Square

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func ValidWordSquare(words []string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n*m), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #422: Valid Word Square
// https://leetcode.com/problems/valid-word-square/
// Difficulty: Easy [Paid]

import "fmt"

// Time: O(n*m), Space: O(1)
func ValidWordSquare(words []string) bool {
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			if j >= len(words) || i >= len(words[j]) || words[i][j] != words[j][i] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(ValidWordSquare([]string{"abcd", "bnrt", "crmy", "dtyx"}))
	fmt.Println(ValidWordSquare([]string{"abcd", "bnrt", "crm", "dt"}))
	fmt.Println(ValidWordSquare([]string{"ball", "area", "lead", "lady"}))
}
```
