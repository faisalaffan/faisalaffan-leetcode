# 1087 — Brace Expansion

## Deskripsi

**Soal:** [1087. Brace Expansion](https://leetcode.com/problems/brace-expansion/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * k) where k is number of expansions  
**Kompleksitas Ruang:** O(n * k)

**Algoritma:** Backtracking (pelacakan mundur)

> **Ide Kunci:** Backtracking - parse braces and generate all expansions

## Solusi Go

```go
package main

// LeetCode #1087: Brace Expansion
// https://leetcode.com/problems/brace-expansion/
// Difficulty: Medium
//
// Approach: Backtracking - parse braces and generate all expansions
// Time: O(n * k) where k is number of expansions
// Space: O(n * k)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(expand("{a,b}c{d,e}f")) // ["acdf","acef","bcdf","bcef"]
	fmt.Println(expand("abcd"))          // ["abcd"]
}

func expand(s string) []string {
  // Membuat slice untuk menyimpan hasil
	result := make([]string, 0)
	backtrack(s, 0, "", &result)
	sort.Strings(result)
	return result
}

func backtrack(s string, idx int, cur string, result *[]string) {
	if idx == len(s) {
		*result = append(*result, cur)
		return
	}

	if s[idx] == '{' {
		// Find the closing brace
		end := idx + 1
		for s[end] != '}' {
			end++
		}
		// Parse options
  // Membuat slice untuk menyimpan hasil
		options := make([]byte, 0)
		for k := idx + 1; k < end; k++ {
			if s[k] != ',' {
				options = append(options, s[k])
			}
		}
		sort.Slice(options, func(i, j int) bool {
			return options[i] < options[j]
		})
		for _, opt := range options {
			backtrack(s, end+1, cur+string(opt), result)
		}
	} else {
		backtrack(s, idx+1, cur+string(s[idx]), result)
	}
}
```
