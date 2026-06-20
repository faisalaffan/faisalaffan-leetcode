# 0932 — Beautiful Array

## Deskripsi

**Soal:** [0932. Beautiful Array](https://leetcode.com/problems/beautiful-array/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman)

**Fungsi Solusi:** `func beautifulArray(n int) []int`

## Solusi Go

```go
package main

// LeetCode #932: Beautiful Array
// https://leetcode.com/problems/beautiful-array/
// Difficulty: Medium

import "fmt"

// Time: O(n log n) | Space: O(n)
func beautifulArray(n int) []int {
  // Membuat map untuk pencarian O(1): key → value
	memo := make(map[int][]int)
	var dfs func(int) []int
	dfs = func(n int) []int {
		if v, ok := memo[n]; ok {
			return v
		}
  // Membuat slice untuk menyimpan hasil
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
