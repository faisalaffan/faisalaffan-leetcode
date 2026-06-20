# 0210 — Course Schedule Ii

## Deskripsi

**Soal:** [0210. Course Schedule Ii](https://leetcode.com/problems/course-schedule-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(V+E), Space: O(V+E)  
**Kompleksitas Ruang:** O(V+E)

**Algoritma:** Queue (antrian FIFO)

**Fungsi Solusi:** `func findOrder(numCourses int, prerequisites [][]int) []int`

## Solusi Go

```go
package main

// LeetCode #210: Course Schedule II
// https://leetcode.com/problems/course-schedule-ii/
// Difficulty: Medium
// Time: O(V+E), Space: O(V+E)

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
  // Membuat slice 2D untuk DP/tabel
	graph := make([][]int, numCourses)
  // Membuat slice untuk menyimpan hasil
	inDegree := make([]int, numCourses)

	for _, pre := range prerequisites {
		course, prereq := pre[0], pre[1]
		graph[prereq] = append(graph[prereq], course)
		inDegree[course]++
	}

	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	result := []int{}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)

		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	if len(result) != numCourses {
		return nil
	}
	return result
}

func main() {
	fmt.Println(findOrder(2, [][]int{{1, 0}}))
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
	fmt.Println(findOrder(1, [][]int{}))
}
```
