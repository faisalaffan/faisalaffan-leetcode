# 2938 — Separate Black And White Balls

## Deskripsi

**Soal:** [2938. Separate Black And White Balls](https://leetcode.com/problems/separate-black-and-white-balls/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2938: Separate Black and White Balls
// https://leetcode.com/problems/separate-black-and-white-balls/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(minimumSteps("101"))
	fmt.Println(minimumSteps("100"))
	fmt.Println(minimumSteps("0111"))
}

func minimumSteps(s string) (ans int64) {
	n := len(s)
	cnt := 0
	for i := n - 1; i >= 0; i-- {
		if s[i] == '1' {
			cnt++
			ans += int64(n - i - cnt)
		}
	}
	return
}
```
