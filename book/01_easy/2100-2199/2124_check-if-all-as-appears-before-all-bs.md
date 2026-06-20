# 2124 — Check If All As Appears Before All Bs

## Deskripsi

**Soal:** [2124. Check If All As Appears Before All Bs](https://leetcode.com/problems/check-if-all-as-appears-before-all-bs/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2124: Check if All A's Appears Before All B's
// https://leetcode.com/problems/check-if-all-as-appears-before-all-bs/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckIfAllAsAppearsBeforeAllBs("aaabbb")) // true
	fmt.Println(CheckIfAllAsAppearsBeforeAllBs("abab"))   // false
	fmt.Println(CheckIfAllAsAppearsBeforeAllBs("bbb"))    // true
}

// Time: O(n), Space: O(1)
func CheckIfAllAsAppearsBeforeAllBs(s string) bool {
	foundB := false
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		if s[i] == 'b' {
			foundB = true
		} else if s[i] == 'a' && foundB {
			return false
		}
	}
	return true
}
```
