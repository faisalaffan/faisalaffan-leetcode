# 1233 — Remove Sub Folders From The Filesystem

## Deskripsi

**Soal:** [1233. Remove Sub Folders From The Filesystem](https://leetcode.com/problems/remove-sub-folders-from-the-filesystem/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n * L) where L = average path length  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func removeSubfolders(folder []string) []string`

## Solusi Go

```go
package main

import (
	"fmt"
	"sort"
	"strings"
)

// LeetCode #1233: Remove Sub-Folders from the Filesystem
// https://leetcode.com/problems/remove-sub-folders-from-the-filesystem/
// Difficulty: Medium

// Sort folders lexicographically. If a folder is a prefix of next,
// the next is a sub-folder. Add "/" to check exact prefix match.

// Time: O(n log n * L) where L = average path length
// Space: O(n)

func removeSubfolders(folder []string) []string {
	sort.Strings(folder)
  // Membuat slice untuk menyimpan hasil
	result := make([]string, 0)
	result = append(result, folder[0])

	for i := 1; i < len(folder); i++ {
		last := result[len(result)-1]
		if !strings.HasPrefix(folder[i], last+"/") {
			result = append(result, folder[i])
		}
	}

	return result
}

func main() {
	fmt.Printf("%v (expected: [/a /c/d])\n",
		removeSubfolders([]string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}))

	fmt.Printf("%v (expected: [/a])\n",
		removeSubfolders([]string{"/a", "/a/b/c", "/a/b"}))

	fmt.Printf("%v (expected: [/a/b /c /d])\n",
		removeSubfolders([]string{"/a/b", "/c", "/d"}))
}
```
