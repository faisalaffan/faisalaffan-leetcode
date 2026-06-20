# 0955 — Delete Columns To Make Sorted Ii

## Deskripsi

**Soal:** [0955. Delete Columns To Make Sorted Ii](https://leetcode.com/problems/delete-columns-to-make-sorted-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * m)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func minDeletionSize(strs []string) int`

## Solusi Go

```go
package main

// LeetCode #955: Delete Columns to Make Sorted II
// https://leetcode.com/problems/delete-columns-to-make-sorted-ii/
// Difficulty: Medium

import "fmt"

// Time: O(n * m) | Space: O(n)
func minDeletionSize(strs []string) int {
	m := len(strs)
	n := len(strs[0])
  // Membuat slice untuk menyimpan hasil
	cut := make([]bool, m)
	ans := 0

	for col := 0; col < n; col++ {
		ok := true
		for row := 0; row+1 < m; row++ {
			if !cut[row] && strs[row][col] > strs[row+1][col] {
				ans++
				ok = false
				break
			}
		}
		if ok {
			for row := 0; row+1 < m; row++ {
				if strs[row][col] < strs[row+1][col] {
					cut[row] = true
				}
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(minDeletionSize([]string{"ca", "bb", "ac"}))
	fmt.Println(minDeletionSize([]string{"xc", "yb", "za"}))
	fmt.Println(minDeletionSize([]string{"zyx", "wvu", "tsr"}))
}
```
