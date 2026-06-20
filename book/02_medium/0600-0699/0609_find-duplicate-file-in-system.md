# 0609 — Find Duplicate File In System

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func FindDuplicate(paths []string) [][]string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n * L) where n = number of files, L = max content length  |  **Ruang:** O(n * L)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #609: Find Duplicate File in System
// https://leetcode.com/problems/find-duplicate-file-in-system/
// Difficulty: Medium
// Time: O(n * L) where n = number of files, L = max content length
// Space: O(n * L)

import (
	"fmt"
	"strings"
)

func main() {
	paths := []string{
		"root/a 1.txt(abcd) 2.txt(efgh)",
		"root/c 3.txt(abcd)",
		"root/c/d 4.txt(efgh)",
	}
	fmt.Println(FindDuplicate(paths))
}

func FindDuplicate(paths []string) [][]string {
  // HashMap: O(1) lookup
	contentMap := make(map[string][]string)

	for _, path := range paths {
		parts := strings.Split(path, " ")
		dir := parts[0]
		for i := 1; i < len(parts); i++ {
			fileStr := parts[i]
			parenIdx := strings.Index(fileStr, "(")
			fileName := fileStr[:parenIdx]
			content := fileStr[parenIdx+1 : len(fileStr)-1] // remove closing ')'
			fullPath := dir + "/" + fileName
			contentMap[content] = append(contentMap[content], fullPath)
		}
	}

	result := [][]string{}
	for _, files := range contentMap {
		if len(files) > 1 {
			result = append(result, files)
		}
	}

	return result
}
```
