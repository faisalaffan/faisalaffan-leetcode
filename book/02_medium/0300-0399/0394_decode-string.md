# 0394 — Decode String

## Deskripsi

**Soal:** [0394. Decode String](https://leetcode.com/problems/decode-string/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Stack (tumpukan LIFO)

**Fungsi Solusi:** `func decodeString(s string) string`

## Solusi Go

```go
package main

// LeetCode #394: Decode String
// https://leetcode.com/problems/decode-string/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func decodeString(s string) string {
  // Membuat slice untuk menyimpan hasil
	numStack := make([]int, 0)
  // Membuat slice untuk menyimpan hasil
	strStack := make([]string, 0)
	curNum := 0
	curStr := ""

	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			curNum = curNum*10 + int(ch-'0')
		} else if ch == '[' {
			numStack = append(numStack, curNum)
			strStack = append(strStack, curStr)
			curNum = 0
			curStr = ""
		} else if ch == ']' {
			num := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			prevStr := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			curStr = prevStr + strings.Repeat(curStr, num)
		} else {
			curStr += string(ch)
		}
	}
	return curStr
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", decodeString("3[a]2[bc]"))
	// Expected: "aaabcbc"

	// Test case 2
	fmt.Println("Test 2:", decodeString("3[a2[c]]"))
	// Expected: "accaccacc"

	// Test case 3
	fmt.Println("Test 3:", decodeString("2[abc]3[cd]ef"))
	// Expected: "abcabccdcdcdef"
}
```
