# 2930 — Number Of Strings Which Can Be Rearranged To Contain Substring

## Deskripsi

**Soal:** [2930. Number Of Strings Which Can Be Rearranged To Contain Substring](https://leetcode.com/problems/number-of-strings-which-can-be-rearranged-to-contain-substring/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2930: Number of Strings Which Can Be Rearranged to Contain Substring
// https://leetcode.com/problems/number-of-strings-which-can-be-rearranged-to-contain-substring/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(stringCount(4))
	fmt.Println(stringCount(10))
}

func stringCount(n int) int {
	const mod int = 1e9 + 7
  // Membuat slice untuk menyimpan hasil
	f := make([][2][3][2]int, n+1)
  // Iterasi seluruh elemen
	for i := range f {
		for j := range f[i] {
			for k := range f[i][j] {
				for l := range f[i][j][k] {
					f[i][j][k][l] = -1
				}
			}
		}
	}
	var dfs func(i, l, e, t int) int
	dfs = func(i, l, e, t int) int {
		if i == 0 {
			if l == 1 && e == 2 && t == 1 {
				return 1
			}
			return 0
		}
		if f[i][l][e][t] == -1 {
			a := dfs(i-1, l, e, t) * 23 % mod
			b := dfs(i-1, min2(1, l+1), e, t)
			c := dfs(i-1, l, min2(2, e+1), t)
			d := dfs(i-1, l, e, min2(1, t+1))
			f[i][l][e][t] = (a + b + c + d) % mod
		}
		return f[i][l][e][t]
	}
	return dfs(n, 0, 0, 0)
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
