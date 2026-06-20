# 0170 — Two Sum Iii Data Structure Design

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array integer dan sebuah target. Tugasmu adalah mencari **dua angka** yang jika dijumlahkan menghasilkan target. Kembalikan **indeks** (posisi) kedua angka.

Contoh: `nums=[2,7,11,15], target=9` → `2+7=9` → `[0,1]`.

**Cara berpikir:** Gunakan HashMap. Untuk setiap angka, cek apakah `target-angka` sudah ada di map. Kalau sudah → ketemu pasangan. Kalau belum → simpan angka ke map.

**Fungsi Solusi:** `func Constructor() TwoSum`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

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
