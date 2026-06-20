# 1220 — Count Vowels Permutation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countVowelPermutation(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1220: Count Vowels Permutation
// https://leetcode.com/problems/count-vowels-permutation/
// Difficulty: Hard

import "fmt"

const mod = 1_000_000_007

// Rules:
//   a -> e
//   e -> a, i
//   i -> a, e, o, u
//   o -> i, u
//   u -> a
// Order: 0=a, 1=e, 2=i, 3=o, 4=u
// next[i] = list of vowels that can follow vowel i
var next = [5][]int{
	{1},       // a -> e
	{0, 2},    // e -> a, i
	{0, 1, 3, 4}, // i -> a, e, o, u
	{2, 4},    // o -> i, u
	{0},       // u -> a
}

func countVowelPermutation(n int) int {
	dp := [5]int{1, 1, 1, 1, 1} // length 1

	for length := 2; length <= n; length++ {
		ndp := [5]int{}
		for v := 0; v < 5; v++ {
			for _, nxt := range next[v] {
				ndp[nxt] = (ndp[nxt] + dp[v]) % mod
			}
		}
		dp = ndp
	}

	total := 0
	for _, v := range dp {
		total = (total + v) % mod
	}
	return total
}

func main() {
	// Test case 1: n=1 -> 5 (a, e, i, o, u)
	fmt.Println(countVowelPermutation(1)) // 5

	// Test case 2: n=2 -> 10
	fmt.Println(countVowelPermutation(2)) // 10

	// Test case 3: n=3
	fmt.Println(countVowelPermutation(3)) // 19

	// Test case 4: n=5
	fmt.Println(countVowelPermutation(5)) // 68

	// Test case 5: n=20000 (large)
	fmt.Println(countVowelPermutation(20000)) // 759959057
}
```
