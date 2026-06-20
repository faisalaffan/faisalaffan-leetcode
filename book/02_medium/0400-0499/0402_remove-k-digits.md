# 0402 — Remove K Digits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func removeKdigits(num string, k int) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Stack

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Stack** — sering muncul di interview!

## 💻 Solusi Go

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

	stack := make([]byte, 0, len(num))
  // Linear scan O(n)
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
