# 1773 — Count Items Matching A Rule

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountMatches(items [][]string, ruleKey string, ruleValue string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1773: Count Items Matching a Rule
// https://leetcode.com/problems/count-items-matching-a-rule/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func CountMatches(items [][]string, ruleKey string, ruleValue string) int {
	idx := 0
	switch ruleKey {
	case "color":
		idx = 1
	case "name":
		idx = 2
	}
	count := 0
	for _, item := range items {
		if item[idx] == ruleValue {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountMatches([][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "iphone"}}, "color", "silver"))
	fmt.Println(CountMatches([][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "phone"}, {"phone", "gold", "iphone"}}, "type", "phone"))
}
```
