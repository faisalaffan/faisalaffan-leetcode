# 2327 — Number Of People Aware Of A Secret

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func peopleAwareOfSecret(n int, delay int, forget int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2327: Number of People Aware of a Secret
// https://leetcode.com/problems/number-of-people-aware-of-a-secret/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func peopleAwareOfSecret(n int, delay int, forget int) int {
	const mod = 1_000_000_007
	// dp[i] = number of people who discovered the secret on day i
  // Alokasi slice integer
	dp := make([]int, n+1)
	dp[1] = 1

	var sharing int64 = 0 // people currently sharing
	for i := 2; i <= n; i++ {
		// New sharers: people who reached delay threshold (i-delay)
		if i-delay >= 1 {
			sharing = (sharing + int64(dp[i-delay])) % mod
		}
		// People who forget: reached forget threshold (i-forget)
		if i-forget >= 1 {
			sharing = (sharing - int64(dp[i-forget]) + mod) % mod
		}
		dp[i] = int(sharing)
	}

	var total int64 = 0
	for i := n - forget + 1; i <= n; i++ {
		if i >= 1 {
			total = (total + int64(dp[i])) % mod
		}
	}
	return int(total)
}

func main() {
	// Test case 1
	fmt.Println(peopleAwareOfSecret(6, 2, 4))
	// Expected: 5

	// Test case 2
	fmt.Println(peopleAwareOfSecret(4, 1, 3))
	// Expected: 6

	// Test case 3
	fmt.Println(peopleAwareOfSecret(10, 3, 5))
	// Expected: 12
}
```
