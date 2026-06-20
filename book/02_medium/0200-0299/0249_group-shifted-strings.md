# 0249 — Group Shifted Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func groupStrings(strs []string) [][]string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * m), Space: O(n * m)  
**Kompleksitas Ruang:** O(n * m)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #249: Group Shifted Strings
// https://leetcode.com/problems/group-shifted-strings/
// Difficulty: Medium [Paid]
// Time: O(n * m), Space: O(n * m)

import (
	"fmt"
	"strings"
)

func groupStrings(strs []string) [][]string {
  // Membuat map (HashMap) — pencarian O(1)
	groups := make(map[string][]string)

	for _, s := range strs {
		key := getKey(s)
		groups[key] = append(groups[key], s)
	}

  // Membuat matriks/slice 2D untuk DP
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}

func getKey(s string) string {
	if len(s) == 0 {
		return ""
	}

	shift := s[0] - 'a'
	var sb strings.Builder

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		diff := (int(s[i]-'a') - int(shift) + 26) % 26
		sb.WriteByte(byte(diff + 'a'))
	}

	return sb.String()
}

func main() {
	fmt.Println(groupStrings([]string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"}))
	fmt.Println(groupStrings([]string{"a"}))
	fmt.Println(groupStrings([]string{"ab", "bc", "cd"}))
}
```
