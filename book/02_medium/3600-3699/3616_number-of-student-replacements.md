# 3616 — Number Of Student Replacements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func NumberOfStudentReplacements(grades []int, threshold int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3616: Number of Student Replacements
// https://leetcode.com/problems/number-of-student-replacements/
// Difficulty: Medium [Paid]
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	grades := []int{1, 2, 3, 4, 5}
	fmt.Println("Test 1:", NumberOfStudentReplacements(grades, 3))
	// Test case 2
	grades2 := []int{5, 4, 3, 2, 1}
	fmt.Println("Test 2:", NumberOfStudentReplacements(grades2, 3))
	// Test case 3
	grades3 := []int{1, 1, 1}
	fmt.Println("Test 3:", NumberOfStudentReplacements(grades3, 2))
}

func NumberOfStudentReplacements(grades []int, threshold int) int {
	replacements := 0
	for _, g := range grades {
		if g < threshold {
			replacements++
		}
	}
	return replacements
}
```
