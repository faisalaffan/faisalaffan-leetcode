# 2825 — Make String A Subsequence Using Cyclic Increments

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func MakeStringASubsequenceUsingCyclicIncrements(str1 string, str2 string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2825: Make String a Subsequence Using Cyclic Increments
// https://leetcode.com/problems/make-string-a-subsequence-using-cyclic-increments/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func MakeStringASubsequenceUsingCyclicIncrements(str1 string, str2 string) bool {
	j := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(str1) && j < len(str2); i++ {
		if str1[i] == str2[j] || (str1[i]-'a'+1)%26 == str2[j]-'a' {
			j++
		}
	}
	return j == len(str2)
}

func main() {
	fmt.Println(MakeStringASubsequenceUsingCyclicIncrements("abc", "bcd"))
	fmt.Println(MakeStringASubsequenceUsingCyclicIncrements("abc", "ad"))
}
```
