# 1869 — Longer Contiguous Segments Of Ones Than Zeros

## Deskripsi

**Soal:** [1869. Longer Contiguous Segments Of Ones Than Zeros](https://leetcode.com/problems/longer-contiguous-segments-of-ones-than-zeros/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func CheckZeroOnes(s string) bool`

## Solusi Go

```go
package main

// LeetCode #1869: Longer Contiguous Segments of Ones Than Zeros
// https://leetcode.com/problems/longer-contiguous-segments-of-ones-than-zeros/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func CheckZeroOnes(s string) bool {
	maxOnes, maxZeros := 0, 0
	curOnes, curZeros := 0, 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			curOnes++
			curZeros = 0
			if curOnes > maxOnes {
				maxOnes = curOnes
			}
		} else {
			curZeros++
			curOnes = 0
			if curZeros > maxZeros {
				maxZeros = curZeros
			}
		}
	}
	return maxOnes > maxZeros
}

func main() {
	fmt.Println(CheckZeroOnes("1101"))
	fmt.Println(CheckZeroOnes("111000"))
	fmt.Println(CheckZeroOnes("110100010"))
}
```
