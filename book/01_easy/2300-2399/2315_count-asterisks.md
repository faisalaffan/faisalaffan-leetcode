# 2315 — Count Asterisks

## Deskripsi

**Soal:** [2315. Count Asterisks](https://leetcode.com/problems/count-asterisks/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2315: Count Asterisks
// https://leetcode.com/problems/count-asterisks/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CountAsterisks("l|*e*et|c**o|*de|")) // 2
	fmt.Println(CountAsterisks("iamprogrammer"))      // 0
}

func CountAsterisks(s string) int {
	count := 0
	bar := false
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		if s[i] == '|' {
			bar = !bar
		} else if s[i] == '*' && !bar {
			count++
		}
	}
	return count
}
```
