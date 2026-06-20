# 2944 — Minimum Number Of Coins For Fruits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumCoins(prices []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2944: Minimum Number of Coins for Fruits
// https://leetcode.com/problems/minimum-number-of-coins-for-fruits/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(minimumCoins([]int{3, 1, 2}))
	fmt.Println(minimumCoins([]int{1, 10, 1, 1}))
	fmt.Println(minimumCoins([]int{26, 53, 10, 24, 25, 20, 63, 51}))
}

func minimumCoins(prices []int) int {
	n := len(prices)
  // Alokasi slice integer
	f := make([]int, n+1)
	var dfs func(int) int
	dfs = func(i int) int {
		if i*2 >= n {
			return prices[i-1]
		}
		if f[i] == 0 {
			f[i] = 1 << 30
			for j := i + 1; j <= i*2+1; j++ {
				cost := dfs(j) + prices[i-1]
				if cost < f[i] {
					f[i] = cost
				}
			}
		}
		return f[i]
	}
	return dfs(1)
}
```
