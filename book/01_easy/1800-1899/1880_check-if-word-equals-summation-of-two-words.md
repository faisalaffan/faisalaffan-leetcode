# 1880 — Check If Word Equals Summation Of Two Words

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func IsSumEqual(firstWord string, secondWord string, targetWord string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1880: Check if Word Equals Summation of Two Words
// https://leetcode.com/problems/check-if-word-equals-summation-of-two-words/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func IsSumEqual(firstWord string, secondWord string, targetWord string) bool {
	return wordValue(firstWord)+wordValue(secondWord) == wordValue(targetWord)
}

func wordValue(s string) int {
	val := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		val = val*10 + int(s[i]-'a')
	}
	return val
}

func main() {
	fmt.Println(IsSumEqual("acb", "cba", "cdb"))
	fmt.Println(IsSumEqual("aaa", "a", "aab"))
	fmt.Println(IsSumEqual("aaa", "a", "aaaa"))
}
```
