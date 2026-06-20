# 2937 — Make Three Strings Equal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func MakeThreeStringsEqual(s1 string, s2 string, s3 string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(min(len(s1), len(s2), len(s3)))  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2937: Make Three Strings Equal
// https://leetcode.com/problems/make-three-strings-equal/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: findMinimumOperations
	fmt.Println(MakeThreeStringsEqual("abc", "abb", "ab")) // 2
	fmt.Println(MakeThreeStringsEqual("dac", "bac", "cac")) // -1
	fmt.Println(MakeThreeStringsEqual("a", "a", "a"))       // 0
}

// Time: O(min(len(s1), len(s2), len(s3))) | Space: O(1)
// LeetCode submission name: findMinimumOperations
func MakeThreeStringsEqual(s1 string, s2 string, s3 string) int {
	// Find longest common prefix length
	i := 0
	for i < len(s1) && i < len(s2) && i < len(s3) {
		if s1[i] != s2[i] || s1[i] != s3[i] {
			break
		}
		i++
	}
	if i == 0 {
		return -1
	}
	return (len(s1) - i) + (len(s2) - i) + (len(s3) - i)
}
```
