# 3498 — Reverse Degree Of A String

## Deskripsi

**Soal:** [3498. Reverse Degree Of A String](https://leetcode.com/problems/reverse-degree-of-a-string/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3498: Reverse Degree of a String
// https://leetcode.com/problems/reverse-degree-of-a-string/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ReverseDegreeOfAString("abc"))
	fmt.Println(ReverseDegreeOfAString("zaba"))
}

// ReverseDegreeOfAString computes the sum of (position_in_reversed_alphabet * (i+1)) for each character.
// Reverse: a=26, b=25, ..., z=1.
// Time: O(n). Space: O(1).
func ReverseDegreeOfAString(s string) int {
	sum := 0
	for i, ch := range s {
		revPos := 26 - int(ch-'a')
		sum += revPos * (i + 1)
	}
	return sum
}
```
