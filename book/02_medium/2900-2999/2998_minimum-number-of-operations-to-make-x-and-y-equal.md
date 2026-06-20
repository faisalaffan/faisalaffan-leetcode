# 2998 — Minimum Number Of Operations To Make X And Y Equal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumOperationsToMakeEqual(x int, y int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, DFS

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(log n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2998: Minimum Number of Operations to Make X and Y Equal
// https://leetcode.com/problems/minimum-number-of-operations-to-make-x-and-y-equal/
// Difficulty: Medium
// Time: O(log n) | Space: O(log n)

import "fmt"

func main() {
	fmt.Println(minimumOperationsToMakeEqual(26, 1))
	fmt.Println(minimumOperationsToMakeEqual(3, 10))
	fmt.Println(minimumOperationsToMakeEqual(54, 2))
}

func minimumOperationsToMakeEqual(x int, y int) int {
	f := map[int]int{}
	var dfs func(int) int
	dfs = func(x int) int {
		if y >= x {
			return y - x
		}
		if v, ok := f[x]; ok {
			return v
		}
		a := x%5 + 1 + dfs(x/5)
		b := 5 - x%5 + 1 + dfs(x/5+1)
		c := x%11 + 1 + dfs(x/11)
		d := 11 - x%11 + 1 + dfs(x/11+1)
		res := x - y
		if a < res {
			res = a
		}
		if b < res {
			res = b
		}
		if c < res {
			res = c
		}
		if d < res {
			res = d
		}
		f[x] = res
		return res
	}
	return dfs(x)
}
```
