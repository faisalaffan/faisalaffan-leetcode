# 0609 — Find Duplicate File In System

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindDuplicate(paths []string) [][]string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * L) where n = number of files, L = max content length  
**Kompleksitas Ruang:** O(n * L)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
  // Membuat map (HashMap) — pencarian O(1)
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
