# 0610 — Triangle Judgement

## Deskripsi

**Soal:** [0610. Triangle Judgement](https://leetcode.com/problems/triangle-judgement/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func TriangleJudgement() string`

## Solusi Go

```go
package main

// LeetCode #610: Triangle Judgement
// https://leetcode.com/problems/triangle-judgement/
// Difficulty: Easy

import "fmt"

func TriangleJudgement() string {
	return "SELECT x, y, z, CASE WHEN x + y > z AND x + z > y AND y + z > x THEN 'Yes' ELSE 'No' END AS triangle FROM Triangle"
}

func main() {
	fmt.Println(TriangleJudgement())
}
```
