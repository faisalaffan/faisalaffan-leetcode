# 1948 — Delete Duplicate Folders In System

## Deskripsi

**Soal:** [1948. Delete Duplicate Folders In System](https://leetcode.com/problems/delete-duplicate-folders-in-system/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman), Trie (pohon awalan), LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func newFolder(name string) *Folder`

## Solusi Go

```go
package main

// LeetCode #1948: Delete Duplicate Folders in System
// https://leetcode.com/problems/delete-duplicate-folders-in-system/
// Difficulty: Hard
// Build trie, serialise each subtree. If two nodes have the same serialisation,
// mark all of them for deletion. Collect remaining paths (DFS).

import (
	"fmt"
	"sort"
)

type Folder struct {
	name     string
	children map[string]*Folder
	del      bool
	// cache for serialisation to avoid recomputing
	serial string
}

func newFolder(name string) *Folder {
	return &Folder{name: name, children: make(map[string]*Folder)}
}

func deleteDuplicateFolder(paths [][]string) [][]string {
	root := newFolder("/")

	// Build trie
	for _, path := range paths {
		node := root
		for _, name := range path {
			if _, ok := node.children[name]; !ok {
				node.children[name] = newFolder(name)
			}
			node = node.children[name]
		}
	}

	// Map serialisation -> list of nodes with that serialisation
  // Membuat map untuk pencarian O(1): key → value
	hashNodes := make(map[string][]*Folder)

	var dfs func(node *Folder) string
	dfs = func(node *Folder) string {
		if node.serial != "" {
			return node.serial
		}
		if len(node.children) == 0 {
			return ""
		}
		var parts []string
		for name, child := range node.children {
			childSerial := dfs(child)
			parts = append(parts, name+"|"+childSerial)
		}
		sort.Strings(parts)
		serial := ""
		for _, p := range parts {
			serial += "(" + p + ")"
		}
		node.serial = serial
		hashNodes[serial] = append(hashNodes[serial], node)
		return serial
	}

	// Compute serialisation for all non-root, non-leaf nodes
	for _, child := range root.children {
		dfs(child)
	}

	// Mark duplicates: if serial appears more than once, mark all
	for _, nodes := range hashNodes {
		if len(nodes) > 1 {
			for _, node := range nodes {
				node.del = true
			}
		}
	}

	// Collect remaining paths (skip marked nodes and their descendants)
	var ans [][]string
	var collect func(node *Folder, path []string)
	collect = func(node *Folder, path []string) {
		if node.del {
			return
		}
		if node != root {
			ans = append(ans, append([]string{}, path...))
		}
		// Collect children in sorted order for stable output
		var names []string
		for name := range node.children {
			names = append(names, name)
		}
		sort.Strings(names)
		for _, name := range names {
			path = append(path, name)
			collect(node.children[name], path)
			path = path[:len(path)-1]
		}
	}

	collect(root, nil)
	return ans
}

func main() {
	// Example 1
	paths1 := [][]string{{"a"}, {"c"}, {"d"}, {"a", "b"}, {"c", "b"}, {"d", "a"}}
	result1 := deleteDuplicateFolder(paths1)
	fmt.Println(result1) // Expected: [[d] [d a]]

	// Example 2
	paths2 := [][]string{{"a"}, {"c"}, {"a", "b"}, {"c", "b"}, {"a", "b", "x"}, {"a", "b", "x", "y"}, {"w"}, {"w", "y"}}
	result2 := deleteDuplicateFolder(paths2)
	fmt.Println(result2)

	// Example 3
	paths3 := [][]string{{"a", "b"}, {"c", "d"}, {"c"}, {"a"}}
	result3 := deleteDuplicateFolder(paths3)
	fmt.Println(result3)
}
```
