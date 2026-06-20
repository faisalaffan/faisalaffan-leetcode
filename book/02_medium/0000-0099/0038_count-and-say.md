# 0038 — Count And Say

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countAndSay(n int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(2^n)  
**Kompleksitas Ruang:** O(2^n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #38: Count and Say
// https://leetcode.com/problems/count-and-say/
// Difficulty: Medium

import "fmt"

func countAndSay(n int) string {
	curr := "1"

	for i := 2; i <= n; i++ {
		var next []byte
		count := 1
		for j := 1; j < len(curr); j++ {
			if curr[j] == curr[j-1] {
				count++
			} else {
				next = append(next, byte('0'+count), curr[j-1])
				count = 1
			}
		}
		next = append(next, byte('0'+count), curr[len(curr)-1])
		curr = string(next)
	}

	return curr
}

func main() {
	// Test case 1
	fmt.Println(countAndSay(4)) // "1211"
	fmt.Println(countAndSay(1)) // "1"
	fmt.Println(countAndSay(5)) // "111221"
}

// Time: O(2^n) | Space: O(2^n)
```
