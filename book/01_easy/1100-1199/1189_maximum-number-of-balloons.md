# 1189 — Maximum Number Of Balloons

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxNumberOfBalloons(text string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1189: Maximum Number of Balloons
// https://leetcode.com/problems/maximum-number-of-balloons/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(maxNumberOfBalloons("nlaebolko"))           // 1
	fmt.Println(maxNumberOfBalloons("loonbalxballpoon"))    // 2
	fmt.Println(maxNumberOfBalloons("leetcode"))            // 0
}

// LeetCode submission: maxNumberOfBalloons
func maxNumberOfBalloons(text string) int {
	count := [26]int{}
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(text); i++ {
		count[text[i]-'a']++
	}
	ans := count[1]           // b
	ans = min(ans, count[0])  // a
	ans = min(ans, count[11]/2) // l (needs 2)
	ans = min(ans, count[14]/2) // o (needs 2)
	ans = min(ans, count[13]) // n
	return ans
}
```
