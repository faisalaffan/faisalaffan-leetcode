# 3846 — Total Distance To Type A String Using One Finger

## Deskripsi

**Soal:** [3846. Total Distance To Type A String Using One Finger](https://leetcode.com/problems/total-distance-to-type-a-string-using-one-finger/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func TotalDistanceToTypeAStringUsingOneFinger(s string) int`

> **Ide Kunci:** Precompute keyboard positions, simulate typing from 'a'.

## Solusi Go

```go
package main

// LeetCode #3846: Total Distance to Type a String Using One Finger
// https://leetcode.com/problems/total-distance-to-type-a-string-using-one-finger/
// Difficulty: Medium [Paid]
// Time: O(N) | Space: O(1)
// Approach: Precompute keyboard positions, simulate typing from 'a'.

import "fmt"

func TotalDistanceToTypeAStringUsingOneFinger(s string) int {
	// Keyboard layout (row, col)
	keyboard := []string{
		"qwertyuiop",
		"asdfghjkl",
		"zxcvbnm",
	}

  // Membuat map untuk pencarian O(1): key → value
	pos := make(map[byte][2]int)
	for r, row := range keyboard {
		for c, ch := range row {
			pos[byte(ch)] = [2]int{r, c}
		}
	}

	total := 0
	cur := pos['a']
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		next := pos[s[i]]
		dist := abs(cur[0]-next[0]) + abs(cur[1]-next[1])
		total += dist
		cur = next
	}

	return total
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Example 1
	fmt.Println(TotalDistanceToTypeAStringUsingOneFinger("hello")) // Expected: 17

	// Example 2
	fmt.Println(TotalDistanceToTypeAStringUsingOneFinger("a")) // Expected: 0

	// Example 3
	fmt.Println(TotalDistanceToTypeAStringUsingOneFinger("qaz")) // q: (0,0), a: (1,0), z: (2,0) = |0-1|+|0-0| + |1-2|+|0-0| = 1+1 = 2
}
```
