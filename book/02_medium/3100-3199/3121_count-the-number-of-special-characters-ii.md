# 3121 — Count The Number Of Special Characters Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func numberOfSpecialChars(word string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(26) = O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3121: Count the Number of Special Characters II
// https://leetcode.com/problems/count-the-number-of-special-characters-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(26) = O(1)

import "fmt"

func numberOfSpecialChars(word string) int {
  // Alokasi slice integer
	firstLower := make([]int, 26)
  // Alokasi slice integer
	lastUpper := make([]int, 26)
  // Range loop: iterasi dengan indeks + nilai
	for i := range firstLower {
		firstLower[i] = -1
		lastUpper[i] = -1
	}

	for i, ch := range word {
		if ch >= 'a' && ch <= 'z' {
			idx := ch - 'a'
			if firstLower[idx] == -1 {
				firstLower[idx] = i
			}
		} else {
			idx := ch - 'A'
			lastUpper[idx] = i
		}
	}

	ans := 0
	for i := 0; i < 26; i++ {
		if firstLower[i] != -1 && lastUpper[i] != -1 && firstLower[i] > lastUpper[i] {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfSpecialChars("aaAbcBC"))  // Expected: 3
	fmt.Println(numberOfSpecialChars("abc"))       // Expected: 0
	fmt.Println(numberOfSpecialChars("AbBCab"))    // Expected: 0
}
```
