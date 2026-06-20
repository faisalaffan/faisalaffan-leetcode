# 2899 — Last Visited Integers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func LastVisitedIntegers(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2899: Last Visited Integers
// https://leetcode.com/problems/last-visited-integers/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: lastVisitedIntegers
	fmt.Println(LastVisitedIntegers([]int{1, 2, -1, -1, -1})) // [2, 1, -1]
	fmt.Println(LastVisitedIntegers([]int{1, -1, 2, -1, -1})) // [1, 2, 1]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: lastVisitedIntegers
func LastVisitedIntegers(nums []int) []int {
	seen := []int{}
	result := []int{}
	k := 0

	for _, num := range nums {
		if num != -1 {
			seen = append(seen, num)
			k = 0
		} else {
			k++
			if k <= len(seen) {
				result = append(result, seen[len(seen)-k])
			} else {
				result = append(result, -1)
			}
		}
	}
	return result
}
```
