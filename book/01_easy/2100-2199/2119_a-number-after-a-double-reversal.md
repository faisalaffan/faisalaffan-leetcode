# 2119 — A Number After A Double Reversal

## Deskripsi

**Soal:** [2119. A Number After A Double Reversal](https://leetcode.com/problems/a-number-after-a-double-reversal/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(1), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2119: A Number After a Double Reversal
// https://leetcode.com/problems/a-number-after-a-double-reversal/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ANumberAfterADoubleReversal(526))  // true
	fmt.Println(ANumberAfterADoubleReversal(1800)) // false
	fmt.Println(ANumberAfterADoubleReversal(0))    // true
}

// Time: O(1), Space: O(1)
func ANumberAfterADoubleReversal(num int) bool {
	// Reversing twice yields the same iff num has no trailing zeros
	return num == 0 || num%10 != 0
}
```
