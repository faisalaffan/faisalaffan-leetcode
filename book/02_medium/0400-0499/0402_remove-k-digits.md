# 0402 — Remove K Digits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func removeKdigits(num string, k int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #402: Remove K Digits
// https://leetcode.com/problems/remove-k-digits/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func removeKdigits(num string, k int) string {
	if k >= len(num) {
		return "0"
	}

	stack := make([]byte, 0, len(num))
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(num); i++ {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > num[i] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}

	// If still need to remove, remove from end
	stack = stack[:len(stack)-k]

	// Remove leading zeros
	result := strings.TrimLeft(string(stack), "0")
	if result == "" {
		return "0"
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", removeKdigits("1432219", 3))
	// Expected: "1219"

	// Test case 2
	fmt.Println("Test 2:", removeKdigits("10200", 1))
	// Expected: "200"

	// Test case 3
	fmt.Println("Test 3:", removeKdigits("10", 2))
	// Expected: "0"
}
```
