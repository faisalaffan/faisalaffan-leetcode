# 3211 — Generate Binary Strings Without Adjacent Zeros

## Deskripsi

**Soal:** [3211. Generate Binary Strings Without Adjacent Zeros](https://leetcode.com/problems/generate-binary-strings-without-adjacent-zeros/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(2^n)  
**Kompleksitas Ruang:** O(n) for recursion

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman)

**Fungsi Solusi:** `func validStrings(n int) []string`

## Solusi Go

```go
package main

// LeetCode #3211: Generate Binary Strings Without Adjacent Zeros
// https://leetcode.com/problems/generate-binary-strings-without-adjacent-zeros/
// Difficulty: Medium
// Time: O(2^n) | Space: O(n) for recursion

import "fmt"

func validStrings(n int) []string {
  // Membuat slice untuk menyimpan hasil
	ans := make([]string, 0)
	var dfs func(cur []byte)
	dfs = func(cur []byte) {
		if len(cur) == n {
			ans = append(ans, string(cur))
			return
		}
		// Option 1: append '1'
		cur = append(cur, '1')
		dfs(cur)
		cur = cur[:len(cur)-1]

		// Option 2: append '0' only if previous was not '0'
		if len(cur) == 0 || cur[len(cur)-1] != '0' {
			cur = append(cur, '0')
			dfs(cur)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(make([]byte, 0, n))
	return ans
}

func main() {
	fmt.Println(validStrings(3)) // Expected: ["010","011","101","110","111"]
	fmt.Println(validStrings(1)) // Expected: ["0","1"]
}
```
