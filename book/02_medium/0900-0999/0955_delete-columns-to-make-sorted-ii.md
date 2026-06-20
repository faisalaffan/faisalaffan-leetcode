# 0955 — Delete Columns To Make Sorted Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data yang perlu diurutkan dengan aturan tertentu. Tugasmu adalah mengurutkan data tersebut dan mungkin melakukan operasi tambahan setelah terurut.

Mengurutkan data adalah operasi fundamental di computer science. Go menyediakan `sort.Ints()` untuk integer, `sort.Strings()` untuk string, dan `sort.Slice()` untuk custom sorting dengan closure.

**Konsep kunci:** comparator, ascending/descending, stable sort, custom sort key.

**Fungsi yang perlu kamu implementasikan:**
```go
func minDeletionSize(strs []string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * m)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

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
