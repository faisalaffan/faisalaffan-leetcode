# 1784 — Check If Binary String Has At Most One Segment Of Ones

## Deskripsi

**Soal:** [1784. Check If Binary String Has At Most One Segment Of Ones](https://leetcode.com/problems/check-if-binary-string-has-at-most-one-segment-of-ones/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func CheckOnesSegment(s string) bool`

## Solusi Go

```go
package main

// LeetCode #1784: Check if Binary String Has at Most One Segment of Ones
// https://leetcode.com/problems/check-if-binary-string-has-at-most-one-segment-of-ones/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

// Time: O(n), Space: O(1)
func CheckOnesSegment(s string) bool {
	return !strings.Contains(s, "01")
}

func main() {
	fmt.Println(CheckOnesSegment("1001"))
	fmt.Println(CheckOnesSegment("110"))
	fmt.Println(CheckOnesSegment("1"))
}
```
