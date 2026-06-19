package main

// LeetCode #2307: Check for Contradictions in Equations
// https://leetcode.com/problems/check-for-contradictions-in-equations/
// Difficulty: Hard [Paid]
//
// Approach: Weighted Union-Find. Each equation a / b = val means a = val * b.
// Store parent and ratio (value[a] / value[parent[a]]). When unioning a and b
// with val (where a/b = val):
//   - Find roots ra, rb and ratios ra, rb
//   - If ra == rb, check consistency: ratio[b]/ratio[a] == val
//   - Otherwise, union ra under rb (or vice versa) with appropriate ratio

import "fmt"

func main() {
	// Example 1: [["a","b"],["b","c"],["a","c"]], [2.0,3.0,6.0] => true (no contradiction)
	fmt.Println(checkContradictions([][]string{{"a", "b"}, {"b", "c"}, {"a", "c"}}, []float64{2.0, 3.0, 6.0}))
	// Example 2: [["a","b"],["b","c"],["a","c"]], [2.0,3.0,5.0] => false (contradiction: a=2b, b=3c, a=6c but a/c=5)
	fmt.Println(checkContradictions([][]string{{"a", "b"}, {"b", "c"}, {"a", "c"}}, []float64{2.0, 3.0, 5.0}))
	// Example 3: [["a","b"],["c","d"]], [2.0,3.0] => true
	fmt.Println(checkContradictions([][]string{{"a", "b"}, {"c", "d"}}, []float64{2.0, 3.0}))
	// Edge: single equation
	fmt.Println(checkContradictions([][]string{{"a", "b"}}, []float64{2.0}))
	// Edge: self-loop (a/a = 2.0 is contradiction since a/a must be 1.0)
	fmt.Println(checkContradictions([][]string{{"a", "a"}}, []float64{2.0}))
}

type WeightedUF struct {
	parent map[string]string
	ratio  map[string]float64 // value[key] / value[parent[key]]
}

func NewWeightedUF() *WeightedUF {
	return &WeightedUF{
		parent: make(map[string]string),
		ratio:  make(map[string]float64),
	}
}

func (uf *WeightedUF) find(x string) (string, float64) {
	if _, ok := uf.parent[x]; !ok {
		uf.parent[x] = x
		uf.ratio[x] = 1.0
	}
	if uf.parent[x] == x {
		return x, 1.0
	}
	root, r := uf.find(uf.parent[x])
	uf.ratio[x] *= r
	uf.parent[x] = root
	return root, uf.ratio[x]
}

func checkContradictions(equations [][]string, values []float64) bool {
	// Tolerance for floating point comparison
	const eps = 1e-9

	uf := NewWeightedUF()

	for i, eq := range equations {
		a, b := eq[0], eq[1]
		val := values[i] // a / b = val => a = val * b

		if a == b {
			if abs(val-1.0) > eps {
				return false
			}
			continue
		}

		ra, raRatio := uf.find(a) // ratio[a] / ratio[ra] = raRatio
		rb, rbRatio := uf.find(b) // ratio[b] / ratio[rb] = rbRatio
		// a = val * b
		// raRatio * ra = a, mbaumRatio * rb = b
		// raRatio * ra = val * rbRatio * rb
		// ra/rb = val * rbRatio / raRatio

		if ra == rb {
			// Both already in same set. Check: ratio[a]/ratio[b] == val
			// ratio[a] = raRatio, ratio[b] = rbRatio (since both relative to same root)
			if abs(raRatio/rbRatio-val) > eps {
				return false
			}
		} else {
			// Union: make ra point to rb
			// We want: ratio[a] / ratio[b] = val
			// Let new ratio for ra: ratio[ra] / ratio[rb] = ?
			// ratio[a] = raRatio * uf.ratio[ra] (after union)
			// We want: raRatio * x / rbRatio = val
			// x = val * rbRatio / raRatio
			uf.parent[ra] = rb
			uf.ratio[ra] = val * rbRatio / raRatio
		}
	}
	return true
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
