# 0380 — Insert Delete Getrandom O1

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor() RandomizedSet
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(1) per operation  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
