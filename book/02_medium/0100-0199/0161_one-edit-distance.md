# 0161 — One Edit Distance

## Deskripsi

**Soal:** [0161. One Edit Distance](https://leetcode.com/problems/one-edit-distance/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func isOneEditDistance(s string, t string) bool`

## Solusi Go

```go
package main

// LeetCode #161: One Edit Distance
// https://leetcode.com/problems/one-edit-distance/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1)

import "fmt"

func isOneEditDistance(s string, t string) bool {
	ns, nt := len(s), len(t)
	if abs(ns-nt) > 1 {
		return false
	}

	if ns > nt {
		s, t = t, s
		ns, nt = nt, ns
	}

	for i := 0; i < ns; i++ {
		if s[i] != t[i] {
			if ns == nt {
				return s[i+1:] == t[i+1:]
			}
			return s[i:] == t[i+1:]
		}
	}

	return ns+1 == nt
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(isOneEditDistance("ab", "acb"))
	fmt.Println(isOneEditDistance("", ""))
	fmt.Println(isOneEditDistance("a", ""))
}
```
