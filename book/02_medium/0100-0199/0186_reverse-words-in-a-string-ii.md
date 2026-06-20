# 0186 — Reverse Words In A String Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func reverseWords(s []byte) 
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1) in-place  
**Kompleksitas Ruang:** O(1) in-place

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #186: Reverse Words in a String II
// https://leetcode.com/problems/reverse-words-in-a-string-ii/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1) in-place

import "fmt"

func reverseWords(s []byte) {
	reverse := func(arr []byte, l, r int) {
  // Two-pointer: gerakkan kiri atau kanan
		for l < r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}

	reverse(s, 0, len(s)-1)

	start := 0
	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == ' ' {
			reverse(s, start, i-1)
			start = i + 1
		}
	}
}

func main() {
	s1 := []byte("the sky is blue")
	reverseWords(s1)
	fmt.Println(string(s1))

	s2 := []byte("hello world")
	reverseWords(s2)
	fmt.Println(string(s2))

	s3 := []byte("a")
	reverseWords(s3)
	fmt.Println(string(s3))
}
```
