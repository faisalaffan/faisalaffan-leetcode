# 0763 — Partition Labels

## Deskripsi

**Soal:** [0763. Partition Labels](https://leetcode.com/problems/partition-labels/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #763: Partition Labels
// https://leetcode.com/problems/partition-labels/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
	fmt.Println(partitionLabels("eccbbbbdec"))
}

func partitionLabels(s string) []int {
  // Membuat slice untuk menyimpan hasil
	last := make([]int, 26)
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		last[s[i]-'a'] = i
	}

  // Membuat slice untuk menyimpan hasil
	result := make([]int, 0)
	start, end := 0, 0

  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		if last[s[i]-'a'] > end {
			end = last[s[i]-'a']
		}
		if i == end {
			result = append(result, end-start+1)
			start = i + 1
		}
	}

	return result
}
```
