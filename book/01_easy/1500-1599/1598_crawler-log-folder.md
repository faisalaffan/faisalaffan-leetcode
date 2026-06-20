# 1598 — Crawler Log Folder

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func minOperations(logs []string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1598: Crawler Log Folder
// https://leetcode.com/problems/crawler-log-folder/
// Difficulty: Easy
//
// LeetCode submission: func minOperations(logs []string) int

import "fmt"

func main() {
	fmt.Println(CrawlerLogFolder([]string{"d1/", "d2/", "../", "d21/", "./"}))            // 2
	fmt.Println(CrawlerLogFolder([]string{"d1/", "../", "../", "../"}))                   // 0
	fmt.Println(CrawlerLogFolder([]string{"./", "../", "./"}))                            // 0
}

// Time: O(n), Space: O(1)
func CrawlerLogFolder(logs []string) int {
	depth := 0
	for _, log := range logs {
		switch log {
		case "../":
			if depth > 0 {
				depth--
			}
		case "./":
			// do nothing
		default:
			depth++
		}
	}
	return depth
}
```
