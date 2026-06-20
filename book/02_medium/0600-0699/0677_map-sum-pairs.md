# 0677 — Map Sum Pairs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func MapSumConstructor() MapSum
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(L) for insert, O(L) for sum  
**Kompleksitas Ruang:** O(n * L) where n is number of keys

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

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
