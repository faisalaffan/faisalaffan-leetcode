# 1228 — Missing Number In Arithmetic Progression

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func missingNumber(arr []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1228: Missing Number In Arithmetic Progression
// https://leetcode.com/problems/missing-number-in-arithmetic-progression/
// Difficulty: Easy [Paid]
// Time: O(log n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{5, 7, 11, 13}))  // 9
	fmt.Println(missingNumber([]int{15, 13, 12}))    // 14
}

// LeetCode submission: missingNumber
func missingNumber(arr []int) int {
	n := len(arr)
	diff := (arr[n-1] - arr[0]) / n
	lo, hi := 0, n-1
	for lo < hi {
		mid := (lo + hi) >> 1
		if arr[mid] == arr[0]+mid*diff {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return arr[0] + diff*lo
}
```
