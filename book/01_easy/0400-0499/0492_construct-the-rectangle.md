# 0492 — Construct The Rectangle

## Deskripsi

**Soal:** [0492. Construct The Rectangle](https://leetcode.com/problems/construct-the-rectangle/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(sqrt(n)), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func ConstructTheRectangle(area int) []int`

## Solusi Go

```go
package main

// LeetCode #492: Construct the Rectangle
// https://leetcode.com/problems/construct-the-rectangle/
// Difficulty: Easy

import "fmt"

// Time: O(sqrt(n)), Space: O(1)
func ConstructTheRectangle(area int) []int {
	w := 1
	for i := 1; i*i <= area; i++ {
		if area%i == 0 {
			w = i
		}
	}
	return []int{area / w, w}
}

func main() {
	fmt.Println(ConstructTheRectangle(4))
	fmt.Println(ConstructTheRectangle(37))
	fmt.Println(ConstructTheRectangle(122122))
}
```
