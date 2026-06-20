# 1404 — Number Of Steps To Reduce A Number In Binary Representation To One

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func numSteps(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n) where n = length of binary string  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1404: Number of Steps to Reduce a Number in Binary Representation to One
// https://leetcode.com/problems/number-of-steps-to-reduce-a-number-in-binary-representation-to-one/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(numSteps("1101")) // 6

	// Test case 2
	fmt.Println(numSteps("10")) // 1

	// Test case 3
	fmt.Println(numSteps("1")) // 0

	// Test case 4
	fmt.Println(numSteps("1111011110000011100000110001011011110010111001010111110001"))
}

// Time: O(n) where n = length of binary string
// Space: O(1)
func numSteps(s string) int {
	steps := 0
	carry := 0

	for i := len(s) - 1; i > 0; i-- {
		digit := int(s[i]-'0') + carry
		if digit%2 == 1 {
			// Odd: add 1 (which makes it even, two operations: +1 and /2)
			steps += 2
			carry = 1
		} else {
			// Even: divide by 2 (one operation)
			steps++
			// carry stays (if we had carry, 1+0=1, but we divide by 2)
		}
	}

	return steps + carry
}
```
