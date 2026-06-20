# 0170 — Two Sum Iii Data Structure Design

## Deskripsi

**Soal:** [0170. Two Sum Iii Data Structure Design](https://leetcode.com/problems/two-sum-iii-data-structure-design/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func Constructor() TwoSum`

## Solusi Go

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
