# 2810 — Faulty Keyboard

## Deskripsi

**Soal:** [2810. Faulty Keyboard](https://leetcode.com/problems/faulty-keyboard/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2810: Faulty Keyboard
// https://leetcode.com/problems/faulty-keyboard/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import (
	"fmt"
)

func main() {
	fmt.Println(FaultyKeyboard("string"))
	fmt.Println(FaultyKeyboard("poiinter"))
}

func FaultyKeyboard(s string) string {
	result := []rune{}
	for _, ch := range s {
		if ch == 'i' {
			// Reverse the current result
			for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
				result[i], result[j] = result[j], result[i]
			}
		} else {
			result = append(result, ch)
		}
	}
	return string(result)
}
```
