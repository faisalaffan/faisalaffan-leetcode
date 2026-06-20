# 2755 — Deep Merge Of Two Objects

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func DeepMergeOfTwoObjects(obj1, obj2 JSONObj) JSONObj`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2755: Deep Merge of Two Objects
// https://leetcode.com/problems/deep-merge-of-two-objects/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

type JSONObj map[string]interface{}

func DeepMergeOfTwoObjects(obj1, obj2 JSONObj) JSONObj {
	result := make(JSONObj)
	for k, v := range obj1 {
		result[k] = v
	}
	for k, v2 := range obj2 {
		if v1, ok := result[k]; ok {
			m1, ok1 := v1.(map[string]interface{})
			m2, ok2 := v2.(map[string]interface{})
			if ok1 && ok2 {
				result[k] = DeepMergeOfTwoObjects(JSONObj(m1), JSONObj(m2))
			} else {
				result[k] = v2
			}
		} else {
			result[k] = v2
		}
	}
	return result
}

func main() {
	merged := DeepMergeOfTwoObjects(
		JSONObj{"a": 1, "b": JSONObj{"c": 2}},
		JSONObj{"b": JSONObj{"d": 3}, "e": 4},
	)
	fmt.Println(merged)

	merged2 := DeepMergeOfTwoObjects(
		JSONObj{"x": 1},
		JSONObj{"x": 2},
	)
	fmt.Println(merged2)
}
```
