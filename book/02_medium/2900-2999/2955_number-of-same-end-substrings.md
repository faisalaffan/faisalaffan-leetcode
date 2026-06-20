# 2955 — Number Of Same End Substrings

## Deskripsi

**Soal:** [2955. Number Of Same End Substrings](https://leetcode.com/problems/number-of-same-end-substrings/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * 26 + q * 26) = O(n + q)  
**Kompleksitas Ruang:** O(26 * n)

**Algoritma:** —

**Fungsi Solusi:** `func numberOfSameEndSubstrings(s string, queries [][]int) []int`

## Solusi Go

```go
package main

// LeetCode #2955: Number of Same-End Substrings
// https://leetcode.com/problems/number-of-same-end-substrings/
// Difficulty: Medium

import "fmt"

func numberOfSameEndSubstrings(s string, queries [][]int) []int {
	n := len(s)

	// prefix[c][i+1] = count of character c in s[0..i]
  // Membuat slice 2D untuk DP/tabel
	prefix := make([][]int, 26)
	for c := 0; c < 26; c++ {
		prefix[c] = make([]int, n+1)
	}
	for i := 0; i < n; i++ {
		for c := 0; c < 26; c++ {
			prefix[c][i+1] = prefix[c][i]
		}
		prefix[s[i]-'a'][i+1]++
	}

  // Membuat slice untuk menyimpan hasil
	ans := make([]int, len(queries))
	for qi, q := range queries {
		l, r := q[0], q[1]
		total := 0
		for c := 0; c < 26; c++ {
			x := prefix[c][r+1] - prefix[c][l]
			if x > 0 {
				total += x * (x + 1) / 2
			}
		}
		ans[qi] = total
	}

	return ans
}

func main() {
	// Test case 1: s = "abca", queries = [[0,3]]
	// 'a' count=2: 2*3/2=3, 'b' count=1: 1, 'c' count=1: 1 -> total=5
	fmt.Println(numberOfSameEndSubstrings("abca", [][]int{{0, 3}})) // [5]

	// Test case 2: single character
	fmt.Println(numberOfSameEndSubstrings("a", [][]int{{0, 0}})) // [1]

	// Test case 3: all same characters
	fmt.Println(numberOfSameEndSubstrings("aaaa", [][]int{{0, 3}})) // [10]
	// 'a' count=4: 4*5/2=10

	// Test case 4: multiple queries
	// "abc": [0,0]->a=1->1, [0,1]->a=1,b=1->2, [0,2]->a=1,b=1,c=1->3
	fmt.Println(numberOfSameEndSubstrings("abc", [][]int{{0, 0}, {0, 1}, {0, 2}}))
	// [1, 2, 3]
	fmt.Println(numberOfSameEndSubstrings("aba", [][]int{{0, 2}})) // [4]
	// 'a' at 0 and 2: count=2 -> 2*3/2=3, 'b' at 1: count=1 -> 1. total=4
	// Substrings: "a"(0), "b"(1), "a"(2), "aba"(0..2) -> 4
}

// Time: O(n * 26 + q * 26) = O(n + q) | Space: O(26 * n)
```
