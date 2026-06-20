# 0944 — Delete Columns To Make Sorted

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func minDeletionSize(strs []string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n * m). Space: O(1).  |  **Ruang:** O(1).


## 💻 Solusi Go

```go
package main

// LeetCode #944: Delete Columns to Make Sorted
// https://leetcode.com/problems/delete-columns-to-make-sorted/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(minDeletionSize([]string{"cba", "daf", "ghi"})) // 1
	fmt.Println(minDeletionSize([]string{"a", "b"}))             // 0
	fmt.Println(minDeletionSize([]string{"zyx", "wvu", "tsr"}))  // 3
}

// minDeletionSize counts columns to delete so the remaining columns are sorted.
// Time: O(n * m). Space: O(1).
func minDeletionSize(strs []string) int {
	if len(strs) == 0 {
		return 0
	}
	n, m := len(strs), len(strs[0])
	count := 0
	for col := 0; col < m; col++ {
		for row := 1; row < n; row++ {
			if strs[row][col] < strs[row-1][col] {
				count++
				break
			}
		}
	}
	return count
}
```
