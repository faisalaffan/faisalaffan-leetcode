# 3178 — Find The Child Who Has The Ball After K Seconds

## Deskripsi

**Soal:** [3178. Find The Child Who Has The Ball After K Seconds](https://leetcode.com/problems/find-the-child-who-has-the-ball-after-k-seconds/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(1). Space: O(1).  
**Kompleksitas Ruang:** O(1).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3178: Find the Child Who Has the Ball After K Seconds
// https://leetcode.com/problems/find-the-child-who-has-the-ball-after-k-seconds/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindTheChildWhoHasTheBallAfterKSeconds(3, 5))
	fmt.Println(FindTheChildWhoHasTheBallAfterKSeconds(5, 6))
	fmt.Println(FindTheChildWhoHasTheBallAfterKSeconds(4, 2))
}

// FindTheChildWhoHasTheBallAfterKSeconds returns the child who has the ball after k seconds.
// Children pass the ball left-to-right, then right-to-left, repeatedly.
// Time: O(1). Space: O(1).
func FindTheChildWhoHasTheBallAfterKSeconds(n int, k int) int {
	cycleLen := 2 * (n - 1)
	k %= cycleLen
	if k < n {
		return k
	}
	return cycleLen - k
}
```
