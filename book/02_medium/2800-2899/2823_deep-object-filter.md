# 2823 — Deep Object Filter

## Deskripsi

**Soal:** [2823. Deep Object Filter](https://leetcode.com/problems/deep-object-filter/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func DeepObjectFilter(obj FilterObj, predicate func(string, interface{}) bool) FilterObj`

## Solusi Go

```go
package main

// LeetCode #2823: Deep Object Filter
// https://leetcode.com/problems/deep-object-filter/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

type FilterObj map[string]interface{}

func DeepObjectFilter(obj FilterObj, predicate func(string, interface{}) bool) FilterObj {
	result := make(FilterObj)
	for k, v := range obj {
		if !predicate(k, v) {
			continue
		}
		if m, ok := v.(map[string]interface{}); ok {
			filtered := DeepObjectFilter(m, predicate)
			if len(filtered) > 0 {
				result[k] = filtered
			}
		} else {
			result[k] = v
		}
	}
	return result
}

func main() {
	obj := FilterObj{
		"a": 1,
		"b": 0,
		"c": FilterObj{"d": 2, "e": 0},
	}
	// Filter out values that are 0
	result := DeepObjectFilter(obj, func(k string, v interface{}) bool {
		if num, ok := v.(int); ok && num == 0 {
			return false
		}
		return true
	})
	fmt.Println(result)

	// Filter keys starting with 'a'
	result2 := DeepObjectFilter(obj, func(k string, v interface{}) bool {
		return k >= "b"
	})
	fmt.Println(result2)
}
```
