# 1700 — Number Of Students Unable To Eat Lunch

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountStudents(students []int, sandwiches []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1700: Number of Students Unable to Eat Lunch
// https://leetcode.com/problems/number-of-students-unable-to-eat-lunch/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func CountStudents(students []int, sandwiches []int) int {
	count := [2]int{0, 0}
	for _, s := range students {
		count[s]++
	}
	for _, sandwich := range sandwiches {
		if count[sandwich] == 0 {
			break
		}
		count[sandwich]--
	}
	return count[0] + count[1]
}

func main() {
	fmt.Println(CountStudents([]int{1, 1, 0, 0}, []int{0, 1, 0, 1}))
	fmt.Println(CountStudents([]int{1, 1, 1, 0, 0, 1}, []int{1, 0, 0, 0, 1, 1}))
}
```
