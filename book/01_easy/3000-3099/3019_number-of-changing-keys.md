# 3019 — Number Of Changing Keys

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func NumberOfChangingKeys(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3019: Number of Changing Keys
// https://leetcode.com/problems/number-of-changing-keys/
// Difficulty: Easy

import "fmt"
import "unicode"

func main() {
	// LeetCode name: countKeyChanges
	fmt.Println(NumberOfChangingKeys("aAbBcC")) // 2
	fmt.Println(NumberOfChangingKeys("AaAaAaaA")) // 0
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: countKeyChanges
func NumberOfChangingKeys(s string) int {
	count := 0
	for i := 1; i < len(s); i++ {
		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[i-1])) {
			count++
		}
	}
	return count
}
```
