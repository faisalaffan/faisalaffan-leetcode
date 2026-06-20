# 0932 — Beautiful Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func beautifulArray(n int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, DP

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #932: Beautiful Array
// https://leetcode.com/problems/beautiful-array/
// Difficulty: Medium

import "fmt"

// Time: O(n log n) | Space: O(n)
func beautifulArray(n int) []int {
  // HashMap: O(1) lookup
	memo := make(map[int][]int)
	var dfs func(int) []int
	dfs = func(n int) []int {
		if v, ok := memo[n]; ok {
			return v
		}
  // Alokasi slice
		res := make([]int, n)
		if n == 1 {
			res[0] = 1
		} else {
			left := dfs((n + 1) / 2)
			right := dfs(n / 2)
			for i, v := range left {
				res[i] = 2*v - 1
			}
			for i, v := range right {
				res[(n+1)/2+i] = 2 * v
			}
		}
		memo[n] = res
		return res
	}
	return dfs(n)
}

func main() {
	fmt.Println(beautifulArray(4))
	fmt.Println(beautifulArray(5))
	fmt.Println(beautifulArray(1))
}
```
