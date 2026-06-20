# 0960 — Delete Columns To Make Sorted Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan data yang perlu diurutkan dengan aturan tertentu. Tugasmu adalah mengurutkan data tersebut dan mungkin melakukan operasi tambahan setelah terurut.

Mengurutkan data adalah operasi fundamental di computer science. Go menyediakan `sort.Ints()` untuk integer, `sort.Strings()` untuk string, dan `sort.Slice()` untuk custom sorting dengan closure.

**Konsep kunci:** comparator, ascending/descending, stable sort, custom sort key.

**Fungsi yang perlu kamu implementasikan:**
```go
func minDeletionSize(A []string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #960: Delete Columns to Make Sorted III
// https://leetcode.com/problems/delete-columns-to-make-sorted-iii/
// Difficulty: Hard

import "fmt"

func minDeletionSize(A []string) int {
	if len(A) == 0 {
		return 0
	}
	m, n := len(A), len(A[0])

	// dp[j] = longest increasing subsequence ending at column j
  // Alokasi slice integer
	dp := make([]int, n)
	for j := range dp {
		dp[j] = 1
	}

	for j := 0; j < n; j++ {
		for k := 0; k < j; k++ {
			// Check if we can place column j after column k
			ok := true
			for i := 0; i < m; i++ {
				if A[i][k] > A[i][j] {
					ok = false
					break
				}
			}
			if ok {
				dp[j] = max(dp[j], dp[k]+1)
			}
		}
	}

	// max columns we can keep
	maxKeep := 0
	for _, v := range dp {
		maxKeep = max(maxKeep, v)
	}
	// min columns to delete = total - maxKeep
	return n - maxKeep
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("Example 1:")
	fmt.Println(minDeletionSize([]string{"babca", "bbazb"}))
	// Expected: 3

	fmt.Println("Example 2:")
	fmt.Println(minDeletionSize([]string{"edcba"}))
	// Expected: 4

	fmt.Println("Example 3:")
	fmt.Println(minDeletionSize([]string{"ghi", "def", "abc"}))
	// Expected: 0
}
```
