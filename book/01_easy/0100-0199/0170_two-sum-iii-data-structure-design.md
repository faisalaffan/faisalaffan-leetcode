# 0170 — Two Sum Iii Data Structure Design

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor() TwoSum
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #170: Two Sum III - Data structure design
// https://leetcode.com/problems/two-sum-iii-data-structure-design/
// Difficulty: Easy [Paid]

import "fmt"

type TwoSum struct {
	nums map[int]int
}

func Constructor() TwoSum {
	return TwoSum{nums: make(map[int]int)}
}

func (t *TwoSum) Add(number int) {
	t.nums[number]++
}

// Time: O(n) | Space: O(n)
func (t *TwoSum) Find(value int) bool {
	for num := range t.nums {
		want := value - num
		if want == num && t.nums[num] > 1 {
			return true
		}
		if want != num && t.nums[want] > 0 {
			return true
		}
	}
	return false
}

func main() {
	t := Constructor()
	t.Add(1)
	t.Add(3)
	t.Add(5)
	fmt.Println(t.Find(4))
	fmt.Println(t.Find(7))
}
```
