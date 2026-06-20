# 1865 — Finding Pairs With A Certain Sum

## Deskripsi

**Soal:** [1865. Finding Pairs With A Certain Sum](https://leetcode.com/problems/finding-pairs-with-a-certain-sum/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func Constructor(nums1 []int, nums2 []int) FindSumPairs`

## Solusi Go

```go
package main

// LeetCode #1865: Finding Pairs With a Certain Sum
// https://leetcode.com/problems/finding-pairs-with-a-certain-sum/
// Difficulty: Medium

import "fmt"

type FindSumPairs struct {
	nums1 []int
	nums2 []int
	freq  map[int]int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[int]int)
	for _, v := range nums2 {
		freq[v]++
	}
	return FindSumPairs{nums1: nums1, nums2: nums2, freq: freq}
}

func (this *FindSumPairs) Add(index int, val int) {
	old := this.nums2[index]
	this.freq[old]--
	this.nums2[index] += val
	this.freq[this.nums2[index]]++
}

func (this *FindSumPairs) Count(tot int) int {
	ans := 0
	for _, v := range this.nums1 {
		ans += this.freq[tot-v]
	}
	return ans
}

func main() {
	obj := Constructor([]int{1, 1, 2, 2, 2, 3}, []int{1, 4, 5, 2, 5, 4})
	fmt.Println(obj.Count(7))
	obj.Add(3, 2)
	fmt.Println(obj.Count(8))
	fmt.Println(obj.Count(4))
	obj.Add(0, 1)
	obj.Add(1, 1)
	fmt.Println(obj.Count(7))
}
```
