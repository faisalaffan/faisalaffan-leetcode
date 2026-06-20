# 2582 — Pass The Pillow

## Deskripsi

**Soal:** [2582. Pass The Pillow](https://leetcode.com/problems/pass-the-pillow/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2582: Pass the Pillow
// https://leetcode.com/problems/pass-the-pillow/
// Difficulty: Easy
// Time O(1) | Space O(1)

import "fmt"

func main() {
	fmt.Println(PassThePillow(4, 5)) // 2
	fmt.Println(PassThePillow(3, 2)) // 3
}

func PassThePillow(n int, time int) int {
	cycle := 2 * (n - 1)
	t := time % cycle
	if t < n {
		return t + 1
	}
	return n - (t - n + 1)
}
```
