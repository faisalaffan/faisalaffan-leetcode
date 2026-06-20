# 0195 — Tenth Line

## Deskripsi

**Soal:** [0195. Tenth Line](https://leetcode.com/problems/tenth-line/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func TenthLine() string`

## Solusi Go

```go
package main

// LeetCode #195: Tenth Line
// https://leetcode.com/problems/tenth-line/
// Difficulty: Easy

import "fmt"

func TenthLine() string {
	return "sed -n '10p' file.txt"
}

func main() {
	fmt.Println(TenthLine())
}
```
