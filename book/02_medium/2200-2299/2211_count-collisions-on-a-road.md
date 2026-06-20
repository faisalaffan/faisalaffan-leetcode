# 2211 — Count Collisions On A Road

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countCollisions(directions string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2211: Count Collisions on a Road
// https://leetcode.com/problems/count-collisions-on-a-road/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func countCollisions(directions string) int {
	n := len(directions)
	left, right := 0, n-1

	for left < n && directions[left] == 'L' {
		left++
	}
	for right >= 0 && directions[right] == 'R' {
		right--
	}

	count := 0
	for i := left; i <= right; i++ {
		if directions[i] != 'S' {
			count++
		}
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println(countCollisions("RLRSLL"))
	// Expected: 5

	// Test case 2
	fmt.Println(countCollisions("LLRR"))
	// Expected: 0

	// Test case 3
	fmt.Println(countCollisions("SSRSSRLLRSLLRSRSSRLRRRRRRS"))
	// Expected: 20
}
```
