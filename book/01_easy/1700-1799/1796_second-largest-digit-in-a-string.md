# 1796 — Second Largest Digit In A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func SecondHighest(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1796: Second Largest Digit in a String
// https://leetcode.com/problems/second-largest-digit-in-a-string/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func SecondHighest(s string) int {
	largest := -1
	second := -1
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			digit := int(s[i] - '0')
			if digit > largest {
				second = largest
				largest = digit
			} else if digit < largest && digit > second {
				second = digit
			}
		}
	}
	return second
}

func main() {
	fmt.Println(SecondHighest("dfa12321afd"))
	fmt.Println(SecondHighest("abc1111"))
	fmt.Println(SecondHighest("ck077"))
}
```
