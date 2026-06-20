# 3260 — Find The Largest Palindrome Divisible By K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func largestPalindrome(n int, k int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS, Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3260: Find the Largest Palindrome Divisible by K
// https://leetcode.com/problems/find-the-largest-palindrome-divisible-by-k/
// Difficulty: Hard
//
// DP + greedy digit construction. Build the palindrome from outermost digits
// inward. For each symmetric pair (or middle digit for odd n), try digits 9..0.
// Use memoized DP to check whether a valid completion exists for the remaining
// inner positions given the current modulo remainder.
//
// dp[pos][mod] = can we complete positions [pos, half) to reach remainder 0?
// Only compute for pos = 0..half-1, where half = (n+1)/2.

import "fmt"

func main() {
	// Example 1: n=3,k=5 => "595"
	fmt.Println(largestPalindrome(3, 5))
	// Example 2: n=1,k=4 => "8"
	fmt.Println(largestPalindrome(1, 4))
	// Example 3: n=5,k=6 => "89898"
	fmt.Println(largestPalindrome(5, 6))
	// Example 4: n=2,k=2 => "88"
	fmt.Println(largestPalindrome(2, 2))
	// Example 5: n=4,k=11 => "9999"
	fmt.Println(largestPalindrome(4, 11))
}

func largestPalindrome(n int, k int) string {
	// Precompute pow10[i] = 10^i mod k
  // Alokasi slice integer
	pow10 := make([]int, n)
	pow10[0] = 1 % k
	for i := 1; i < n; i++ {
		pow10[i] = (pow10[i-1] * 10) % k
	}

	half := (n + 1) / 2

	// dp[pos][mod] = true if we can fill positions [pos, half) to reach 0 mod k
	// with the current accumulated remainder = mod
  // Membuat matriks/slice 2D untuk DP
	memo := make([][]int, half)
  // Range loop: iterasi dengan indeks + nilai
	for i := range memo {
		memo[i] = make([]int, k)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	// ans stores the digits of the result
	ans := make([]byte, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range ans {
		ans[i] = '0'
	}

	var dfs func(pos, mod int) bool
	dfs = func(pos, mod int) bool {
		if pos == half {
			return mod == 0
		}
		if memo[pos][mod] != -1 {
			return memo[pos][mod] == 1
		}

		// The symmetric index from the right
		right := n - 1 - pos

		for d := 9; d >= 0; d-- {
			inc := d * pow10[pos] % k
			if pos != right {
				inc = (inc + d*pow10[right]%k) % k
			}
			newMod := (mod + inc) % k
			if dfs(pos+1, newMod) {
				ans[pos] = byte('0' + d)
				if pos != right {
					ans[right] = byte('0' + d)
				}
				memo[pos][mod] = 1
				return true
			}
		}

		memo[pos][mod] = 0
		return false
	}

	dfs(0, 0)
	return string(ans)
}
```
