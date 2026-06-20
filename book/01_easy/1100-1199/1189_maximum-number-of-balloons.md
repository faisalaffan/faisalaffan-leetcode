# 1189 — Maximum Number Of Balloons

## Deskripsi

**Soal:** [1189. Maximum Number Of Balloons](https://leetcode.com/problems/maximum-number-of-balloons/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1189: Maximum Number of Balloons
// https://leetcode.com/problems/maximum-number-of-balloons/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(maxNumberOfBalloons("nlaebolko"))           // 1
	fmt.Println(maxNumberOfBalloons("loonbalxballpoon"))    // 2
	fmt.Println(maxNumberOfBalloons("leetcode"))            // 0
}

// LeetCode submission: maxNumberOfBalloons
func maxNumberOfBalloons(text string) int {
	count := [26]int{}
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(text); i++ {
		count[text[i]-'a']++
	}
	ans := count[1]           // b
	ans = min(ans, count[0])  // a
	ans = min(ans, count[11]/2) // l (needs 2)
	ans = min(ans, count[14]/2) // o (needs 2)
	ans = min(ans, count[13]) // n
	return ans
}
```
