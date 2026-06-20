# 0207 — Course Schedule

## Deskripsi

**Soal:** [0207. Course Schedule](https://leetcode.com/problems/course-schedule/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(V+E), Space: O(V+E)  
**Kompleksitas Ruang:** O(V+E)

**Algoritma:** Queue (antrian FIFO)

**Fungsi Solusi:** `func canFinish(numCourses int, prerequisites [][]int) bool`

## Solusi Go

```go
package main

// LeetCode #207: Course Schedule
// https://leetcode.com/problems/course-schedule/
// Difficulty: Medium
// Time: O(V+E), Space: O(V+E)

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
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

	count := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		count++

		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return count == numCourses
}

func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
	fmt.Println(canFinish(5, [][]int{{1, 0}, {2, 1}, {3, 2}}))
}
```
