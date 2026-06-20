# 1221 — Split A String In Balanced Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func balancedStringSplit(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1221: Split a String in Balanced Strings
// https://leetcode.com/problems/split-a-string-in-balanced-strings/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(balancedStringSplit("RLRRLLRLRL")) // 4
	fmt.Println(balancedStringSplit("RLLLLRRRLR")) // 3
	fmt.Println(balancedStringSplit("LLLLRRRR"))   // 1
}

// LeetCode submission: balancedStringSplit
func balancedStringSplit(s string) int {
	count, ans := 0, 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		if s[i] == 'R' {
			count++
		} else {
			count--
		}
		if count == 0 {
			ans++
		}
	}
	return ans
}
```
