# 2700 — Differences Between Two Objects

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func objectDiff(obj1, obj2 any) map[string]any`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2700: Differences Between Two Objects
// https://leetcode.com/problems/differences-between-two-objects/
// Difficulty: Medium [Paid] (JS problem)
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"reflect"
)

func objectDiff(obj1, obj2 any) map[string]any {
  // HashMap: O(1) lookup
	result := make(map[string]any)
	diff(obj1, obj2, "", result)
	return result
}

func diff(a, b any, path string, result map[string]any) {
	if reflect.DeepEqual(a, b) {
		return
	}

	if a == nil || b == nil {
		result[path] = []any{a, b}
		return
	}

	map1, ok1 := a.(map[string]any)
	map2, ok2 := b.(map[string]any)
	if ok1 && ok2 {
  // HashMap: O(1) lookup
		allKeys := make(map[string]bool)
		for k := range map1 {
			allKeys[k] = true
		}
		for k := range map2 {
			allKeys[k] = true
		}
		for k := range allKeys {
			newPath := path
			if newPath == "" {
				newPath = k
			} else {
				newPath = path + "." + k
			}
			diff(map1[k], map2[k], newPath, result)
		}
		return
	}

	result[path] = []any{a, b}
}

func main() {
	// Test case 1: simple diff
	o1 := map[string]any{"a": 1, "b": 2}
	o2 := map[string]any{"a": 1, "b": 3}
	fmt.Println("Test 1:", objectDiff(o1, o2))
	// Expected: map[b:[2 3]]

	// Test case 2: missing key
	o3 := map[string]any{"a": 1, "c": 3}
	fmt.Println("Test 2:", objectDiff(o1, o3))
	// Expected: map[b:[2 <nil>] c:[<nil> 3]]

	// Test case 3: same objects
	fmt.Println("Test 3:", objectDiff(o1, o1))
	// Expected: map[]
}
```
