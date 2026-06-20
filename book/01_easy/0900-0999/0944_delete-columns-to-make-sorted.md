# 0944 — Delete Columns To Make Sorted

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data yang perlu diurutkan dengan aturan tertentu. Tugasmu adalah mengurutkan data tersebut dan mungkin melakukan operasi tambahan setelah terurut.

Mengurutkan data adalah operasi fundamental di computer science. Go menyediakan `sort.Ints()` untuk integer, `sort.Strings()` untuk string, dan `sort.Slice()` untuk custom sorting dengan closure.

**Konsep kunci:** comparator, ascending/descending, stable sort, custom sort key.

**Fungsi yang perlu kamu implementasikan:**
```go
func minDeletionSize(strs []string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * m). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
