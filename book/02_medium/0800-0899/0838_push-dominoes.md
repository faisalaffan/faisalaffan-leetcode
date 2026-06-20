# 0838 — Push Dominoes

## Deskripsi

**Soal:** [0838. Push Dominoes](https://leetcode.com/problems/push-dominoes/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** LIS (Longest Increasing Subsequence)

## Solusi Go

```go
package main

// LeetCode #838: Push Dominoes
// https://leetcode.com/problems/push-dominoes/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(PushDominoes("RR.L"))
	fmt.Println(PushDominoes(".L.R...LR..L.."))
	fmt.Println(PushDominoes("L.R"))
}

// Time: O(n) | Space: O(n)
func PushDominoes(dominoes string) string {
	n := len(dominoes)
	res := []byte(dominoes)

	for i := 0; i < n; i++ {
		if res[i] == 'R' {
			// Find the next non-dot character
			j := i + 1
			for j < n && res[j] == '.' {
				j++
			}
			if j == n || res[j] == 'R' {
				// All dots between i and j become 'R'
				for k := i + 1; k < j; k++ {
					res[k] = 'R'
				}
			} else if res[j] == 'L' {
				// Collision: left and right meet in the middle
				left, right := i+1, j-1
  // Loop two-pointer: kiri vs kanan
				for left < right {
					res[left] = 'R'
					res[right] = 'L'
					left++
					right--
				}
			}
			i = j
		} else if res[i] == 'L' {
			// Propagate L leftwards
			j := i - 1
			for j >= 0 && res[j] == '.' {
				res[j] = 'L'
				j--
			}
		}
	}

	// Handle initial dots before first 'L'
	for i := 0; i < n && res[i] == '.'; i++ {
		if i+1 < n && res[i+1] == 'L' {
			res[i] = 'L'
		}
	}

	return string(res)
}
```
