# 0380 — Insert Delete Getrandom O1

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func Constructor() RandomizedSet`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(1) per operation  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #380: Insert Delete GetRandom O(1)
// https://leetcode.com/problems/insert-delete-getrandom-o1/
// Difficulty: Medium
// Time: O(1) per operation | Space: O(n)

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	nums  []int
	pos   map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{pos: make(map[int]int)}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.pos[val]; ok {
		return false
	}
	rs.pos[val] = len(rs.nums)
	rs.nums = append(rs.nums, val)
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	idx, ok := rs.pos[val]
	if !ok {
		return false
	}

	// Swap with last element
	last := len(rs.nums) - 1
	lastVal := rs.nums[last]
	rs.nums[idx] = lastVal
	rs.pos[lastVal] = idx
	rs.nums = rs.nums[:last]
	delete(rs.pos, val)
	return true
}

func (rs *RandomizedSet) GetRandom() int {
	return rs.nums[rand.Intn(len(rs.nums))]
}

func main() {
	rs := Constructor()
	fmt.Println("Insert 1:", rs.Insert(1)) // true
	fmt.Println("Remove 2:", rs.Remove(2)) // false
	fmt.Println("Insert 2:", rs.Insert(2)) // true
	fmt.Println("GetRandom:", rs.GetRandom()) // 1 or 2
	fmt.Println("Remove 1:", rs.Remove(1)) // true
	fmt.Println("Insert 2:", rs.Insert(2)) // false (already present)
	fmt.Println("GetRandom:", rs.GetRandom()) // 2
}
```
