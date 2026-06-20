# 1071 — Greatest Common Divisor Of Strings

## Deskripsi

**Soal:** [1071. Greatest Common Divisor Of Strings](https://leetcode.com/problems/greatest-common-divisor-of-strings/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n + m)  
**Kompleksitas Ruang:** O(n + m)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1071: Greatest Common Divisor of Strings
// https://leetcode.com/problems/greatest-common-divisor-of-strings/
// Difficulty: Easy
// Time: O(n + m) | Space: O(n + m)

import "fmt"

func main() {
	fmt.Println(gcdOfStrings("ABCABC", "ABC")) // "ABC"
	fmt.Println(gcdOfStrings("ABABAB", "ABAB")) // "AB"
	fmt.Println(gcdOfStrings("LEET", "CODE"))   // ""
}

// LeetCode submission: gcdOfStrings
func gcdOfStrings(str1, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	return str1[:gcd(len(str1), len(str2))]
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
```
