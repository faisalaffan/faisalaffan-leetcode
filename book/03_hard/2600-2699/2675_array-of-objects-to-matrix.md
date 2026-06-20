# 2675 — Array Of Objects To Matrix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func arrayOfObjectsToMatrix(arr []map[string]any) [][]any`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting, Prefix Sum

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2675: Array of Objects to Matrix
// https://leetcode.com/problems/array-of-objects-to-matrix/
// Difficulty: Hard [Paid]
//
// Convert an array of nested objects into a 2D matrix where each row is a
// flattened version of one object. Nested keys are joined with ".".
// The first row contains all unique flattened keys in sorted order.
// Missing values are filled with empty string.

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	// Example 1: simple array of flat objects
	obj1 := map[string]any{
		"a": 1,
		"b": 2,
	}
	obj2 := map[string]any{
		"a": 3,
		"c": 4,
	}
	fmt.Println(arrayOfObjectsToMatrix([]map[string]any{obj1, obj2}))

	// Example 2: nested objects
	obj3 := map[string]any{
		"a": map[string]any{
			"b": 1,
		},
	}
	obj4 := map[string]any{
		"a": map[string]any{
			"c": 2,
		},
	}
	fmt.Println(arrayOfObjectsToMatrix([]map[string]any{obj3, obj4}))
}

// arrayOfObjectsToMatrix converts an array of objects to a 2D matrix.
func arrayOfObjectsToMatrix(arr []map[string]any) [][]any {
	if len(arr) == 0 {
		return [][]any{}
	}

	// Collect all unique flattened keys
  // HashMap: O(1) lookup
	keySet := make(map[string]bool)
	flattened := make([]map[string]any, len(arr))

	for i, obj := range arr {
  // HashMap: O(1) lookup
		flat := make(map[string]any)
		flattenObj("", obj, flat)
		flattened[i] = flat
		for k := range flat {
			keySet[k] = true
		}
	}

	// Sort keys
	keys := make([]string, 0, len(keySet))
	for k := range keySet {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Build matrix
  // Matriks 2D
	result := make([][]any, len(arr)+1)

	// Header row
	header := make([]any, len(keys))
	for i, k := range keys {
		header[i] = k
	}
	result[0] = header

	// Data rows
	for i, flat := range flattened {
		row := make([]any, len(keys))
		for j, k := range keys {
			if v, ok := flat[k]; ok {
				row[j] = v
			} else {
				row[j] = ""
			}
		}
		result[i+1] = row
	}

	return result
}

// flattenObj recursively flattens a nested object with dot-separated keys.
func flattenObj(prefix string, obj any, result map[string]any) {
	switch v := obj.(type) {
	case map[string]any:
		for key, val := range v {
			newKey := key
			if prefix != "" {
				newKey = prefix + "." + key
			}
			flattenObj(newKey, val, result)
		}
	case []any:
		for i, val := range v {
			newKey := prefix + "." + strconv.Itoa(i)
			flattenObj(newKey, val, result)
		}
	default:
		result[prefix] = obj
	}
}
```
