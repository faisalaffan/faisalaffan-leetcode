# 2698 — Find The Punishment Number Of An Integer

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func punishmentNumber(n int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(n * 2^len(num))  |  **Ruang:** O(log n)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2698: Find the Punishment Number of an Integer
// https://leetcode.com/problems/find-the-punishment-number-of-an-integer/
// Difficulty: Medium
// Time: O(n * 2^len(num)) | Space: O(log n)

import (
	"fmt"
	"strconv"
)

func punishmentNumber(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		sq := i * i
		sqStr := strconv.Itoa(sq)
		if canPartition(sqStr, i) {
			total += sq
		}
	}
	return total
}

func canPartition(s string, target int) bool {
	if len(s) == 0 {
		return target == 0
	}
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		prefix, _ := strconv.Atoi(s[:i+1])
		if prefix > target {
			break
		}
		if canPartition(s[i+1:], target-prefix) {
			return true
		}
	}
	return false
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", punishmentNumber(10))
	// Expected: 182 (1+81+100)

	// Test case 2
	fmt.Println("Test 2:", punishmentNumber(37))
	// Expected: 1478

	// Test case 3
	fmt.Println("Test 3:", punishmentNumber(1))
	// Expected: 1
}
```
