# 2586 — Count The Number Of Vowel Strings In Range

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func isVowel(ch byte) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2586: Count the Number of Vowel Strings in Range
// https://leetcode.com/problems/count-the-number-of-vowel-strings-in-range/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CountTheNumberOfVowelStringsInRange([]string{"are", "amy", "u"}, 0, 2)) // 2
	fmt.Println(CountTheNumberOfVowelStringsInRange([]string{"hey", "aeo", "mu", "ooo", "artro"}, 1, 4)) // 3
}

func isVowel(ch byte) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
}

func CountTheNumberOfVowelStringsInRange(words []string, left int, right int) int {
	count := 0
	for i := left; i <= right; i++ {
		if len(words[i]) > 0 && isVowel(words[i][0]) && isVowel(words[i][len(words[i])-1]) {
			count++
		}
	}
	return count
}
```
