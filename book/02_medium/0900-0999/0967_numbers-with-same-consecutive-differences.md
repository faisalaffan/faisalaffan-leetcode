# 0967 — Numbers With Same Consecutive Differences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func numsSameConsecDiff(n int, k int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(n * 2^n)  
**Kompleksitas Ruang:** O(n * 2^n)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #967: Numbers With Same Consecutive Differences
// https://leetcode.com/problems/numbers-with-same-consecutive-differences/
// Difficulty: Medium

import "fmt"

// Time: O(n * 2^n) | Space: O(n * 2^n)
func numsSameConsecDiff(n int, k int) []int {
	if n == 1 {
		return []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	}

  // Alokasi slice integer
	ans := make([]int, 0)
	for d := 1; d <= 9; d++ {
		dfs(n, k, d, &ans)
	}
	return ans
}

func dfs(n, k, cur int, ans *[]int) {
	if n == 1 {
		*ans = append(*ans, cur)
		return
	}
	last := cur % 10
	if last+k <= 9 {
		dfs(n-1, k, cur*10+last+k, ans)
	}
	if k != 0 && last-k >= 0 {
		dfs(n-1, k, cur*10+last-k, ans)
	}
}

func main() {
	fmt.Println(numsSameConsecDiff(3, 7))
	fmt.Println(numsSameConsecDiff(2, 1))
	fmt.Println(numsSameConsecDiff(2, 0))
}
```
