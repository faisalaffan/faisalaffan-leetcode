# 1165 — Single Row Keyboard

## Deskripsi

**Soal:** [1165. Single Row Keyboard](https://leetcode.com/problems/single-row-keyboard/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1165: Single-Row Keyboard
// https://leetcode.com/problems/single-row-keyboard/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(calculateTime("abcdefghijklmnopqrstuvwxyz", "cba")) // 4
	fmt.Println(calculateTime("pqrstuvwxyzabcdefghijklmno", "leetcode")) // 73
}

// LeetCode submission: calculateTime
func calculateTime(keyboard, word string) int {
	pos := [26]int{}
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(keyboard); i++ {
		pos[keyboard[i]-'a'] = i
	}
	ans, cur := 0, 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(word); i++ {
		next := pos[word[i]-'a']
		if next > cur {
			ans += next - cur
		} else {
			ans += cur - next
		}
		cur = next
	}
	return ans
}
```
