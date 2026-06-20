# 3938 — Maximum Path Intersection Sum In A Grid

## Deskripsi

**Soal:** [3938. Maximum Path Intersection Sum In A Grid](https://leetcode.com/problems/maximum-path-intersection-sum-in-a-grid/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(C * R^2)  
**Kompleksitas Ruang:** O(R^2) where R = min(m,n), C = max(m,n)

**Algoritma:** Dynamic Programming (DP), Prefix Sum (jumlah kumulatif)

**Fungsi Solusi:** `func maxPathIntersectionSum(grid [][]int) int`

> **Ide Kunci:** Column-by-column DP with 2D prefix max optimization.

## Solusi Go

```go
package main

// LeetCode #3938: Maximum Path Intersection Sum in a Grid
// https://leetcode.com/problems/maximum-path-intersection-sum-in-a-grid/
// Difficulty: Medium
// Time: O(C * R^2) | Space: O(R^2) where R = min(m,n), C = max(m,n)
// Approach: Column-by-column DP with 2D prefix max optimization.
// dp[r1][r2] = max shared sum with P1 at exit row r1, P2 at exit row r2.
// For each column, compute shared interval overlap between the two paths
// within the column. Transition via 4 overlap cases, each O(1) using
// precomputed 2D prefix max arrays. Transpose grid so DP dim = min(m,n).

import (
	"fmt"
	"math"
)

const negInf = math.MinInt32 / 2

func maxPathIntersectionSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// Transpose if m > n to make DP dimension the smaller one
	if m > n {
  // Membuat slice 2D untuk DP/tabel
		t := make([][]int, n)
		for i := 0; i < n; i++ {
			t[i] = make([]int, m)
			for j := 0; j < m; j++ {
				t[i][j] = grid[j][i]
			}
		}
		grid = t
		m, n = n, m
	}
	// Now m <= n. DP dim = m.
	// P1 starts at (0,0), ends at (m-1,n-1), moves right/down.
	// P2 starts at (m-1,0), ends at (0,n-1), moves right/up.

	// Column prefix sums
  // Membuat slice 2D untuk DP/tabel
	pref := make([][]int, n)
	for c := 0; c < n; c++ {
		pref[c] = make([]int, m+1)
		for r := 0; r < m; r++ {
			pref[c][r+1] = pref[c][r] + grid[r][c]
		}
	}

	// dp[r1][r2] after column 0
  // Membuat slice 2D untuk DP/tabel
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = negInf
		}
	}

	// Column 0: P1 enters at row 0, P2 enters at row m-1
	// P1 visits [0, r1], P2 visits [r2, m-1]
	// Overlap = [max(0, r2), min(r1, m-1)] = [r2, r1] if r2 <= r1
	for r1 := 0; r1 < m; r1++ {
		for r2 := 0; r2 <= r1; r2++ {
			dp[r1][r2] = pref[0][r1+1] - pref[0][r2]
		}
	}

	// Process columns 1..n-1
	for c := 1; c < n; c++ {
  // Membuat slice 2D untuk DP/tabel
		newdp := make([][]int, m)
		for i := 0; i < m; i++ {
			newdp[i] = make([]int, m)
			for j := 0; j < m; j++ {
				newdp[i][j] = negInf
			}
		}

		// 2D prefix max of dp (upper-left quadrant: p1 <= i, p2 >= j)
  // Membuat slice 2D untuk DP/tabel
		pmax := make([][]int, m)
		for i := 0; i < m; i++ {
			pmax[i] = make([]int, m+1)
			for j := 0; j <= m; j++ {
				pmax[i][j] = negInf
			}
		}
		// Row-wise suffix max
		for i := 0; i < m; i++ {
			best := negInf
			for j := m - 1; j >= 0; j-- {
				if dp[i][j] > best {
					best = dp[i][j]
				}
				pmax[i][j] = best
			}
		}
		// Column-wise prefix max
		for j := 0; j < m; j++ {
			for i := 1; i < m; i++ {
				if pmax[i-1][j] > pmax[i][j] {
					pmax[i][j] = pmax[i-1][j]
				}
			}
		}

		// Global best (no overlap in this column)
		globalBest := negInf
		for i := 0; i < m; i++ {
			for j := 0; j < m; j++ {
				if dp[i][j] > globalBest {
					globalBest = dp[i][j]
				}
			}
		}

		// For each (r1, r2), compute best transition
		for r1 := 0; r1 < m; r1++ {
			for r2 := 0; r2 < m; r2++ {
				best := globalBest

				// Case 1: p1 <= r2 AND r1 <= p2
				// overlap = [r2, r1], sum = pref[r1+1] - pref[r2]
				if r2 < m && r1 < m {
					val := pmax[r2][r1]
					if val > negInf/2 {
						cand := val + pref[c][r1+1] - pref[c][r2]
						if cand > best {
							best = cand
						}
					}
				}

				// Case 2 & 4: scan p1 from r2+1 to r1
				// Need max_{p2 >= r1} dp[p1][p2] for each p1
				if r1 > r2 {
					// Precompute suffix max per p1 (computed lazily)
					for p1 := r2 + 1; p1 <= r1; p1++ {
						// suffix max over p2 >= r1
						f2 := negInf
						for p2 := r1; p2 < m; p2++ {
							if dp[p1][p2] > f2 {
								f2 = dp[p1][p2]
							}
						}
						if f2 > negInf/2 {
							// Case 2: p1 > r2 AND r1 <= p2 (p2 >= r1 always)
							// overlap = [p1, r1], sum = pref[r1+1] - pref[p1]
							cand2 := f2 + pref[c][r1+1] - pref[c][p1]
							if cand2 > best {
								best = cand2
							}

							// Case 3: p1 <= r2 (already covered by Case 1)
							// Case 4: p1 > r2 AND r1 > p2 (need p2 < r1)
							// overlap = [p1, p2] where p1 <= p2 < r1
							// max over p2 in [p1, r1-1]
							for p2 := p1; p2 < r1 && p2 < m; p2++ {
								if dp[p1][p2] > negInf/2 {
									cand4 := dp[p1][p2] + pref[c][p2+1] - pref[c][p1]
									if cand4 > best {
										best = cand4
									}
								}
							}

							// Case 3 also: p1 <= r2 AND r1 > p2
							// Already covered by Case 1 for the overlap sum
							// But Case 3 = [r2, p2], which differs from Case 1 = [r2, r1]
							// Need: max over p1 <= r2, p2 < r1
							// This is separate from the p1 > r2 loop
						}
					}

					// Case 3: p1 <= r2 AND r1 > p2 (p2 < r1)
					// overlap = [r2, p2], sum = pref[p2+1] - pref[r2]
					// Need max over p1 <= r2, p2 < r1
					if r1 > 0 {
						for p1 := 0; p1 <= r2 && p1 < m; p1++ {
							for p2 := r2; p2 < r1 && p2 < m; p2++ {
								if dp[p1][p2] > negInf/2 {
									cand3 := dp[p1][p2] + pref[c][p2+1] - pref[c][r2]
									if cand3 > best {
										best = cand3
									}
								}
							}
						}
					}
				}

				newdp[r1][r2] = best
			}
		}

		dp = newdp
	}

	// Answer: max over all exit pairs
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			if dp[i][j] > ans {
				ans = dp[i][j]
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(maxPathIntersectionSum([][]int{{1, 2, 0, -3}, {1, -2, 1, 0}, {-4, 2, -1, 3}, {3, -3, 3, -2}, {-1, -5, 0, 1}}))
	fmt.Println(maxPathIntersectionSum([][]int{{4, -2, -3}, {-1, -3, -1}, {-4, 2, -1}}))
}
```
