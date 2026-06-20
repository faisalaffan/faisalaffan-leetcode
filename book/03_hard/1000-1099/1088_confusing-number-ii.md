# 1088 — Confusing Number Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func confusingNumberII(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, DFS

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1088: Confusing Number II
// https://leetcode.com/problems/confusing-number-ii/
// Difficulty: Hard [Paid]
//
// A confusing number is one that when rotated 180 degrees becomes a different
// valid number. Valid digits: 0→0, 1→1, 6→9, 8→8, 9→6. DFS-generate all
// numbers using these digits ≤ n and count the confusing ones.

import "fmt"

var confusingDigits = []int{0, 1, 6, 8, 9}
var rotateMap = map[int]int{0: 0, 1: 1, 6: 9, 8: 8, 9: 6}

func main() {
	fmt.Println(confusingNumberII(20))
	fmt.Println(confusingNumberII(100))
}

func confusingNumberII(n int) int {
	count := 0
	var dfs func(int)
	dfs = func(cur int) {
		if cur > n {
			return
		}
		if cur != 0 && isConfusing(cur) {
			count++
		}
		for _, d := range confusingDigits {
			if cur == 0 && d == 0 {
				continue
			}
			dfs(cur*10 + d)
		}
	}
	dfs(0)
	return count
}

func isConfusing(num int) bool {
	original, rotated := num, 0
	for num > 0 {
		d := num % 10
		rotated = rotated*10 + rotateMap[d]
		num /= 10
	}
	return rotated != original
}
```
