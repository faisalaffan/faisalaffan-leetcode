# 1166 — Design File System

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diminta mendesain struktur data kustom dengan operasi spesifik (insert, delete, search). Target: O(1) atau O(log n) per operasi.

**Cara berpikir:** Kombinasikan HashMap + Heap + Linked List sesuai kebutuhan.

**Fungsi Solusi:** `func Constructor() FileSystem`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** createPath O(L) where L = path depth, get O(L)  |  **Ruang:** O(P) where P = number of paths

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

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
