# 2546 — Apply Bitwise Operations To Make Strings Equal

## Deskripsi

**Soal:** [2546. Apply Bitwise Operations To Make Strings Equal](https://leetcode.com/problems/apply-bitwise-operations-to-make-strings-equal/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func makeStringsEqual(s string, target string) bool`

## Solusi Go

```go
package main

// LeetCode #2546: Apply Bitwise Operations to Make Strings Equal
// https://leetcode.com/problems/apply-bitwise-operations-to-make-strings-equal/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func makeStringsEqual(s string, target string) bool {
	// Operation: choose i,j, set s[i]=s[i]|s[j], s[j]=s[i]^s[j]
	// We can change any position if there's at least one '1' in either string
	hasOneS := false
	hasOneT := false
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			hasOneS = true
		}
		if target[i] == '1' {
			hasOneT = true
		}
	}
	// If target has no '1', s must have no '1' too
	if !hasOneT {
		return !hasOneS
	}
	// If target has '1', s must have at least one '1' to make it work
	return hasOneS
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", makeStringsEqual("1010", "0101"))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", makeStringsEqual("00", "11"))
	// Expected: false

	// Test case 3
	fmt.Println("Test 3:", makeStringsEqual("11", "00"))
	// Expected: false
}
```
