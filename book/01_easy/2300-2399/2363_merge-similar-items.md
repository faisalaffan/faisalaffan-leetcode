# 2363 — Merge Similar Items

## Deskripsi

**Soal:** [2363. Merge Similar Items](https://leetcode.com/problems/merge-similar-items/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2363: Merge Similar Items
// https://leetcode.com/problems/merge-similar-items/
// Difficulty: Easy
// Time O((n+m) log(n+m)) | Space O(n+m)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MergeSimilarItems([][]int{{1, 1}, {4, 5}, {3, 8}}, [][]int{{3, 1}, {1, 5}})) // [[1,6],[3,9],[4,5]]
	fmt.Println(MergeSimilarItems([][]int{{1, 1}, {3, 2}, {2, 3}}, [][]int{{2, 1}, {3, 2}, {1, 3}})) // [[1,4],[2,4],[3,4]]
}

func MergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	valueMap := map[int]int{}
	for _, item := range items1 {
		valueMap[item[0]] += item[1]
	}
	for _, item := range items2 {
		valueMap[item[0]] += item[1]
	}

  // Membuat slice 2D untuk DP/tabel
	result := make([][]int, 0, len(valueMap))
	for v, w := range valueMap {
		result = append(result, []int{v, w})
	}
	sort.Slice(result, func(i, j int) bool { return result[i][0] < result[j][0] })
	return result
}
```
