# 3412 — Find Mirror Score Of A String

## Deskripsi

**Soal:** [3412. Find Mirror Score Of A String](https://leetcode.com/problems/find-mirror-score-of-a-string/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n) Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Stack (tumpukan LIFO)

**Fungsi Solusi:** `func calculateScore(s string) int64`

## Solusi Go

```go
package main

// LeetCode #3412: Find Mirror Score of a String
// https://leetcode.com/problems/find-mirror-score-of-a-string/
// Difficulty: Medium
// Time: O(n) Space: O(n)

import "fmt"

func calculateScore(s string) int64 {
  // Membuat slice 2D untuk DP/tabel
	stacks := make([][]int, 26)
	var ans int64
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		ch := int(s[i] - 'a')
		mirror := 25 - ch
		if len(stacks[mirror]) > 0 {
			j := stacks[mirror][len(stacks[mirror])-1]
			stacks[mirror] = stacks[mirror][:len(stacks[mirror])-1]
			ans += int64(i - j)
		} else {
			stacks[ch] = append(stacks[ch], i)
		}
	}
	return ans
}

func main() {
	fmt.Println(calculateScore("aczzx")) // 5
	fmt.Println(calculateScore("abcdef")) // 0
}
```
