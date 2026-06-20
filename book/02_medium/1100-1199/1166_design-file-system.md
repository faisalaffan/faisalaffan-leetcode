# 1166 — Design File System

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diminta untuk mendesain (merancang) sebuah struktur data kustom dengan operasi tertentu (insert, delete, search, update). Tugasmu adalah memilih representasi data yang tepat agar setiap operasi berjalan efisien — biasanya O(1) atau O(log n).

Ini adalah soal yang paling sering muncul di interview sistem desain. Kamu perlu memilih kombinasi struktur data yang tepat (HashMap + Heap + LinkedList) untuk mencapai kompleksitas yang diminta.

**Konsep kunci:** HashMap (O(1) lookup), Heap (priority), Doubly Linked List (O(1) remove), TreeMap (ordered keys).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor() FileSystem
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** createPath O(L) where L = path depth, get O(L)  
**Kompleksitas Ruang:** O(P) where P = number of paths

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"strings"
)

// LeetCode #1166: Design File System
// https://leetcode.com/problems/design-file-system/
// Difficulty: Medium [Paid]

// FileSystem supports createPath(path, value) and get(path).
// Paths are absolute, use "/" as delimiter.

// Time: createPath O(L) where L = path depth, get O(L)
// Space: O(P) where P = number of paths

type FileSystem struct {
	paths map[string]int
}

func Constructor() FileSystem {
	return FileSystem{paths: make(map[string]int)}
}

func (this *FileSystem) CreatePath(path string, value int) bool {
	if path == "" || path == "/" {
		return false
	}
	if _, exists := this.paths[path]; exists {
		return false
	}

	// Check parent exists (unless parent is root)
	lastSlash := strings.LastIndex(path, "/")
	if lastSlash > 0 {
		parent := path[:lastSlash]
		if _, exists := this.paths[parent]; !exists {
			return false
		}
	}

	this.paths[path] = value
	return true
}

func (this *FileSystem) Get(path string) int {
	if val, exists := this.paths[path]; exists {
		return val
	}
	return -1
}

func main() {
	fs := Constructor()
	fmt.Printf("%t (expected: true)\n", fs.CreatePath("/a", 1))
	fmt.Printf("%d (expected: 1)\n", fs.Get("/a"))
	fmt.Printf("%t (expected: true)\n", fs.CreatePath("/a/b", 2))
	fmt.Printf("%d (expected: 2)\n", fs.Get("/a/b"))
	fmt.Printf("%t (expected: false)\n", fs.CreatePath("/a/b", 3)) // already exists
	fmt.Printf("%t (expected: false)\n", fs.CreatePath("/c/d", 4)) // parent doesn't exist
	fmt.Printf("%d (expected: -1)\n", fs.Get("/c"))
}
```
