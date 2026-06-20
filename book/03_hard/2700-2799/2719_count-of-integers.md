# 2719 — Count Of Integers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countOfIntegers(num1 string, num2 string, minSum int, maxSum int) int
```

> **💡 Hint:** Digit DP. Count integers in [num1, num2] with digit sum in [minSum, maxSum].

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS, Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2719: Count of Integers
// https://leetcode.com/problems/count-of-integers/
// Difficulty: Hard
//
// Approach: Digit DP. Count integers in [num1, num2] with digit sum in [minSum, maxSum].
// Result = f(num2) - f(num1-1) modulo 1e9+7.

import "fmt"

const mod = 1000000007

func countOfIntegers(num1 string, num2 string, minSum int, maxSum int) int {
	var count func(num string) int
	count = func(num string) int {
		n := len(num)
  // Membuat matriks/slice 2D untuk DP
		memo := make([][][]int, n)
  // Range loop: iterasi dengan indeks + nilai
		for i := range memo {
			memo[i] = make([][]int, maxSum+1)
			for j := range memo[i] {
				memo[i][j] = make([]int, 2)
				memo[i][j][0] = -1
				memo[i][j][1] = -1
			}
		}

		var dfs func(pos, sum int, tight bool) int
		dfs = func(pos, sum int, tight bool) int {
			if sum > maxSum {
				return 0
			}
			if pos == n {
				if sum >= minSum {
					return 1
				}
				return 0
			}
			t := 0
			if tight {
				t = 1
			}
			if memo[pos][sum][t] != -1 {
				return memo[pos][sum][t]
			}

			limit := 9
			if tight {
				limit = int(num[pos] - '0')
			}

			total := 0
			for d := 0; d <= limit; d++ {
				nextTight := tight && (d == limit)
				total = (total + dfs(pos+1, sum+d, nextTight)) % mod
			}

			memo[pos][sum][t] = total
			return total
		}

		return dfs(0, 0, true)
	}

	ans := count(num2)
	sub := subtractOne(num1)
	ans = (ans - count(sub) + mod) % mod
	return ans
}

func subtractOne(s string) string {
	b := []byte(s)
	i := len(b) - 1
	for i >= 0 && b[i] == '0' {
		b[i] = '9'
		i--
	}
	if i < 0 {
		return "0"
	}
	b[i]--
	if b[0] == '0' && len(b) > 1 {
		b = b[1:]
	}
	return string(b)
}

func main() {
	// Example 1
	fmt.Println(countOfIntegers("1", "12", 1, 8))
	// Example 2
	fmt.Println(countOfIntegers("1", "5", 1, 5))
	// Single number
	fmt.Println(countOfIntegers("10", "10", 1, 10))
	// Large range
	fmt.Println(countOfIntegers("1", "1000", 1, 100))
}
```
