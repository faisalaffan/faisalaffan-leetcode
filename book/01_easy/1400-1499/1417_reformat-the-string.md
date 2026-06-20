# 1417 — Reformat The String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func reformat(s string) string

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1417: Reformat The String
// https://leetcode.com/problems/reformat-the-string/
// Difficulty: Easy
//
// LeetCode submission: func reformat(s string) string

import "fmt"

func main() {
	fmt.Println(ReformatTheString("a0b1c2")) // "a0b1c2"
	fmt.Println(ReformatTheString("leetcode")) // ""
	fmt.Println(ReformatTheString("1229857369")) // ""
}

// Time: O(n), Space: O(n)
func ReformatTheString(s string) string {
	letters := make([]byte, 0, len(s))
	digits := make([]byte, 0, len(s))
  // Range loop: iterasi dengan indeks + nilai
	for i := range s {
		if s[i] >= 'a' && s[i] <= 'z' {
			letters = append(letters, s[i])
		} else {
			digits = append(digits, s[i])
		}
	}
	if abs(len(letters)-len(digits)) > 1 {
		return ""
	}
	res := make([]byte, len(s))
	var first, second []byte
	if len(letters) >= len(digits) {
		first, second = letters, digits
	} else {
		first, second = digits, letters
	}
	idx := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(first); i++ {
		res[idx] = first[i]
		idx++
		if i < len(second) {
			res[idx] = second[i]
			idx++
		}
	}
	return string(res)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
