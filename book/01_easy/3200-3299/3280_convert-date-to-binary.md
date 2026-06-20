# 3280 — Convert Date To Binary

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func toBinary(n int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3280: Convert Date to Binary
// https://leetcode.com/problems/convert-date-to-binary/
// Difficulty: Easy

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(ConvertDateToBinary("2080-02-29"))
	fmt.Println(ConvertDateToBinary("1900-01-01"))
}

// toBinary converts an integer to its binary string representation without leading zeros.
func toBinary(n int) string {
  // Edge case: input kosong — langsung return
	if n == 0 {
		return "0"
	}
	s := ""
	for n > 0 {
		s = string('0'+byte(n%2)) + s
		n /= 2
	}
	return s
}

// ConvertDateToBinary converts a date string to binary format.
// Time: O(1). Space: O(1).
func ConvertDateToBinary(date string) string {
	parts := strings.Split(date, "-")
	year, _ := strconv.Atoi(parts[0])
	month, _ := strconv.Atoi(parts[1])
	day, _ := strconv.Atoi(parts[2])
	return toBinary(year) + "-" + toBinary(month) + "-" + toBinary(day)
}
```
