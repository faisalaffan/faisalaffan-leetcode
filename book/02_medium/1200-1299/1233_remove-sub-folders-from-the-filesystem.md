# 1233 — Remove Sub Folders From The Filesystem

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func removeSubfolders(folder []string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n log n * L) where L = average path length  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

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
