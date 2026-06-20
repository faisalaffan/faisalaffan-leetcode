# 0677 — Map Sum Pairs

## Deskripsi

**Soal:** [0677. Map Sum Pairs](https://leetcode.com/problems/map-sum-pairs/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(L) for insert, O(L) for sum  
**Kompleksitas Ruang:** O(n * L) where n is number of keys

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #677: Map Sum Pairs
// https://leetcode.com/problems/map-sum-pairs/
// Difficulty: Medium
// Time: O(L) for insert, O(L) for sum
// Space: O(n * L) where n is number of keys

import (
	"fmt"
)

func main() {
	ms := MapSumConstructor()
	ms.Insert("apple", 3)
	fmt.Println(ms.Sum("ap"))
	ms.Insert("app", 2)
	fmt.Println(ms.Sum("ap"))
}

type MapSum struct {
	prefixSum map[string]int
	values    map[string]int
}

func MapSumConstructor() MapSum {
	return MapSum{
		prefixSum: make(map[string]int),
		values:    make(map[string]int),
	}
}

func (m *MapSum) Insert(key string, val int) {
	delta := val
	if oldVal, ok := m.values[key]; ok {
		delta = val - oldVal
	}
	m.values[key] = val

	for i := 1; i <= len(key); i++ {
		m.prefixSum[key[:i]] += delta
	}
}

func (m *MapSum) Sum(prefix string) int {
	return m.prefixSum[prefix]
}
```
