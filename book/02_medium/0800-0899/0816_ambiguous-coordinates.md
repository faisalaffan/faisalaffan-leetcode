# 0816 — Ambiguous Coordinates

## Deskripsi

**Soal:** [0816. Ambiguous Coordinates](https://leetcode.com/problems/ambiguous-coordinates/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^3)  
**Kompleksitas Ruang:** O(n^2)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #816: Ambiguous Coordinates
// https://leetcode.com/problems/ambiguous-coordinates/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(AmbiguousCoordinates("(123)"))
	fmt.Println(AmbiguousCoordinates("(00011)"))
	fmt.Println(AmbiguousCoordinates("(0123)"))
}

// Time: O(n^3) | Space: O(n^2)
func AmbiguousCoordinates(s string) []string {
	s = s[1 : len(s)-1]
	n := len(s)

	var result []string

	for i := 1; i < n; i++ {
		left := validNums(s[:i])
		right := validNums(s[i:])
		for _, a := range left {
			for _, b := range right {
				result = append(result, "("+a+", "+b+")")
			}
		}
	}

	return result
}

func validNums(s string) []string {
	var res []string

	if len(s) == 0 {
		return res
	}

	if s[0] == '0' && s[len(s)-1] == '0' {
		if len(s) == 1 {
			res = append(res, s)
		}
		return res
	}

	if s[0] == '0' {
		res = append(res, "0."+s[1:])
		return res
	}

	res = append(res, s)
	if s[len(s)-1] == '0' {
		return res
	}

	for i := 1; i < len(s); i++ {
		res = append(res, s[:i]+"."+s[i:])
	}

	return res
}
```
