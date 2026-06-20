# 1128 — Number Of Equivalent Domino Pairs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func numEquivDominoPairs(dominoes [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1128: Number of Equivalent Domino Pairs
// https://leetcode.com/problems/number-of-equivalent-domino-pairs/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(numEquivDominoPairs([][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}})) // 1
	fmt.Println(numEquivDominoPairs([][]int{{1, 2}, {1, 2}, {1, 1}, {1, 2}}))  // 3
}

// LeetCode submission: numEquivDominoPairs
func numEquivDominoPairs(dominoes [][]int) int {
  // Membuat map (HashMap) — pencarian O(1)
	count := make(map[[2]int]int)
	ans := 0
	for _, d := range dominoes {
		a, b := d[0], d[1]
		if a > b {
			a, b = b, a
		}
		key := [2]int{a, b}
		ans += count[key]
		count[key]++
	}
	return ans
}
```
