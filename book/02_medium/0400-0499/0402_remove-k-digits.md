# 0402 — Remove K Digits

## Deskripsi

**Soal:** [0402. Remove K Digits](https://leetcode.com/problems/remove-k-digits/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Stack (tumpukan LIFO)

**Fungsi Solusi:** `func removeKdigits(num string, k int) string`

## Solusi Go

```go
package main

// LeetCode #402: Remove K Digits
// https://leetcode.com/problems/remove-k-digits/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func removeKdigits(num string, k int) string {
	if k >= len(num) {
		return "0"
	}

  // Membuat slice untuk menyimpan hasil
	stack := make([]byte, 0, len(num))
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(num); i++ {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > num[i] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}

	// If still need to remove, remove from end
	stack = stack[:len(stack)-k]

	// Remove leading zeros
	result := strings.TrimLeft(string(stack), "0")
	if result == "" {
		return "0"
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", removeKdigits("1432219", 3))
	// Expected: "1219"

	// Test case 2
	fmt.Println("Test 2:", removeKdigits("10200", 1))
	// Expected: "200"

	// Test case 3
	fmt.Println("Test 3:", removeKdigits("10", 2))
	// Expected: "0"
}
```
