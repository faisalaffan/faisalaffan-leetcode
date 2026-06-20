# 0017 — Letter Combinations Of A Phone Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func letterCombinations(digits string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(4^n)  
**Kompleksitas Ruang:** O(4^n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #17: Letter Combinations of a Phone Number
// https://leetcode.com/problems/letter-combinations-of-a-phone-number/
// Difficulty: Medium

import "fmt"

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	phone := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	result := []string{""}
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(digits); i++ {
		letters := phone[digits[i]]
		var temp []string
		for _, prefix := range result {
			for j := 0; j < len(letters); j++ {
				temp = append(temp, prefix+string(letters[j]))
			}
		}
		result = temp
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println(letterCombinations("23")) // ["ad","ae","af","bd","be","bf","cd","ce","cf"]

	// Test case 2
	fmt.Println(letterCombinations("")) // []

	// Test case 3
	fmt.Println(letterCombinations("2")) // ["a","b","c"]
}

// Time: O(4^n) | Space: O(4^n)
```
