# 1849 — Splitting A String Into Descending Consecutive Values

## Deskripsi

**Soal:** [1849. Splitting A String Into Descending Consecutive Values](https://leetcode.com/problems/splitting-a-string-into-descending-consecutive-values/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2), Space: O(n) for recursion  
**Kompleksitas Ruang:** O(n) for recursion

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman)

## Solusi Go

```go
package main

// LeetCode #1849: Splitting a String Into Descending Consecutive Values
// https://leetcode.com/problems/splitting-a-string-into-descending-consecutive-values/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(SplitString("1234"))
	fmt.Println(SplitString("050043"))
	fmt.Println(SplitString("9080701"))
}

// Time: O(n^2), Space: O(n) for recursion
func SplitString(s string) bool {
	var dfs func(idx int, prev int64, count int) bool
	dfs = func(idx int, prev int64, count int) bool {
		if idx == len(s) {
			return count >= 2
		}
		num := int64(0)
		for i := idx; i < len(s); i++ {
			num = num*10 + int64(s[i]-'0')
			if num > 1<<62 {
				break
			}
			if count == 0 || prev-num == 1 {
				if dfs(i+1, num, count+1) {
					return true
				}
			}
			if num == 0 {
				break
			}
		}
		return false
	}
	return dfs(0, 0, 0)
}
```
