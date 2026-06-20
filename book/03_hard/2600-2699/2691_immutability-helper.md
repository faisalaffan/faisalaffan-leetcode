# 2691 — Immutability Helper

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func immutableUpdate(obj any, spec map[string]any) any
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2691: Immutability Helper
// https://leetcode.com/problems/immutability-helper/
// Difficulty: Hard [Paid]
//
// Implement an immutability helper similar to React's update() or
// ImmutableJS API. Given an object and a spec describing mutations
// ({$set: val}, {$push: [items]}, {$merge: {obj}}), return a new object
// with the mutations applied, without modifying the original.

import (
	"fmt"
)

func main() {
	// Example: $set
	obj1 := map[string]any{"a": 1, "b": 2}
	spec1 := map[string]any{"a": map[string]any{"$set": 99}}
	fmt.Println(immutableUpdate(obj1, spec1))

	// Example: nested $push
	obj2 := map[string]any{"items": []any{1, 2, 3}}
	spec2 := map[string]any{"items": map[string]any{"$push": []any{4, 5}}}
	fmt.Println(immutableUpdate(obj2, spec2))

	// Example: $merge
	obj3 := map[string]any{"a": 1, "b": 2, "c": 3}
	spec3 := map[string]any{"$merge": map[string]any{"b": 99, "d": 4}}
	fmt.Println(immutableUpdate(obj3, spec3))
}

// immutableUpdate applies a mutation spec to an object and returns a new copy.
// Supported commands: $set, $push, $merge.
func immutableUpdate(obj any, spec map[string]any) any {
	// Deep copy the original
	result := deepCopy(obj).(map[string]any)

	for key, val := range spec {
		specMap, ok := val.(map[string]any)
		if !ok {
			// Direct assignment (non-command key)
			result[key] = deepCopy(val)
			continue
		}

		// Check for special commands
		if setVal, hasSet := specMap["$set"]; hasSet {
			result[key] = deepCopy(setVal)
		} else if pushVal, hasPush := specMap["$push"]; hasPush {
			if existing, ok := result[key].([]any); ok {
				toPush, ok2 := pushVal.([]any)
				if ok2 {
					newSlice := make([]any, len(existing)+len(toPush))
					copy(newSlice, existing)
					copy(newSlice[len(existing):], toPush)
					result[key] = newSlice
				}
			}
		} else if mergeVal, hasMerge := specMap["$merge"]; hasMerge {
			if existing, ok := result[key].(map[string]any); ok {
				mergeMap, ok2 := mergeVal.(map[string]any)
				if ok2 {
					merged := deepCopy(existing).(map[string]any)
					for mk, mv := range mergeMap {
						merged[mk] = deepCopy(mv)
					}
					result[key] = merged
				}
			}
		} else {
			// Nested spec: recurse
			if existing, ok := result[key].(map[string]any); ok {
				result[key] = immutableUpdate(existing, specMap)
			}
		}
	}

	return result
}

// deepCopy creates a deep copy of a value.
func deepCopy(v any) any {
	switch val := v.(type) {
	case map[string]any:
  // Membuat map (HashMap) — pencarian O(1)
		result := make(map[string]any, len(val))
		for k, vv := range val {
			result[k] = deepCopy(vv)
		}
		return result
	case []any:
		result := make([]any, len(val))
		for i, vv := range val {
			result[i] = deepCopy(vv)
		}
		return result
	default:
		return v
	}
}
```
