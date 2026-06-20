# 0650 — 2 Keys Keyboard

## Deskripsi

**Soal:** [0650. 2 Keys Keyboard](https://leetcode.com/problems/2-keys-keyboard/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n sqrt(n))  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #650: 2 Keys Keyboard
// https://leetcode.com/problems/2-keys-keyboard/
// Difficulty: Medium
// Time: O(n sqrt(n))
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(minSteps(3))
	fmt.Println(minSteps(1))
	fmt.Println(minSteps(10))
}

func minSteps(n int) int {
	result := 0
	d := 2

	for n > 1 {
		for n%d == 0 {
			result += d
			n /= d
		}
		d++
	}

	return result
}
```
