# 1024 — Video Stitching

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func videoStitching(clips [][]int, time int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n log n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1024: Video Stitching
// https://leetcode.com/problems/video-stitching/
// Difficulty: Medium
//
// Approach: Sort clips by start, then greedy extend reach
// Time: O(n log n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(videoStitching([][]int{{0, 2}, {4, 6}, {8, 10}, {1, 9}, {1, 5}, {5, 9}}, 10)) // 3
	fmt.Println(videoStitching([][]int{{0, 1}, {6, 8}, {0, 2}, {5, 6}, {0, 4}, {0, 3}, {6, 7}, {1, 3}, {4, 7}, {1, 4}, {2, 5}, {2, 6}, {3, 4}, {4, 5}, {5, 7}, {6, 9}}, 9)) // 3
	fmt.Println(videoStitching([][]int{{0, 2}, {2, 4}}, 4)) // 2
}

func videoStitching(clips [][]int, time int) int {
	// Sort by start time, then by end time descending
  // Linear scan O(n)
	for i := 0; i < len(clips); i++ {
		for j := i + 1; j < len(clips); j++ {
			if clips[i][0] > clips[j][0] || (clips[i][0] == clips[j][0] && clips[i][1] < clips[j][1]) {
				clips[i], clips[j] = clips[j], clips[i]
			}
		}
	}

	count := 0
	curEnd := 0
	i := 0
	n := len(clips)

	for curEnd < time {
		farthest := curEnd
		for i < n && clips[i][0] <= curEnd {
			if clips[i][1] > farthest {
				farthest = clips[i][1]
			}
			i++
		}
		if farthest == curEnd {
			return -1
		}
		curEnd = farthest
		count++
	}

	return count
}
```
