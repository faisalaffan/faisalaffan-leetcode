# 2390 — Removing Stars From A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func removeStars(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2390: Removing Stars From a String
// https://leetcode.com/problems/removing-stars-from-a-string/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Stack: push non-star, pop on star.

import "fmt"

func main() {
	fmt.Println(removeStars("leet**cod*e")) // "lecoe"
	fmt.Println(removeStars("erase*****"))  // ""
}

func removeStars(s string) string {
	res := make([]byte, 0, len(s))
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			res = res[:len(res)-1]
		} else {
			res = append(res, s[i])
		}
	}
	return string(res)
}
```
