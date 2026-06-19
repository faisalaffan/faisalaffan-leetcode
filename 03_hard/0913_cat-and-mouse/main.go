package main

// LeetCode #913: Cat and Mouse
// https://leetcode.com/problems/cat-and-mouse/
// Difficulty: Hard
// Game theory DP: mouse at 1, cat at 2, hole at 0.
// Mouse moves first. 0=draw, 1=mouse wins, 2=cat wins.
// Minimax with degree-based DP (queue approach) to handle cycles.

import "fmt"

const (
	DRAW = 0
	MOUSE_WIN = 1
	CAT_WIN = 2
)

type state struct{ m, c, t int }

func catMouseGame(graph [][]int) int {
	n := len(graph)
	// dp[m][c][turn]: result for mouse=m, cat=c, turn=0(mouse)/1(cat)
	dp := make([][][]int, n)
	degree := make([][][]int, n)
	for m := 0; m < n; m++ {
		dp[m] = make([][]int, n)
		degree[m] = make([][]int, n)
		for c := 0; c < n; c++ {
			dp[m][c] = []int{DRAW, DRAW}
			degree[m][c] = []int{0, 0}
			// Mouse turn: mouse can move to neighbors
			degree[m][c][0] = len(graph[m])
			// Cat turn: cat can move to neighbors except hole 0
			deg := 0
			for _, nb := range graph[c] {
				if nb != 0 {
					deg++
				}
			}
			degree[m][c][1] = deg
		}
	}

	q := make([]state, 0)

	// Initialize terminal states
	for m := 0; m < n; m++ {
		for t := 0; t < 2; t++ {
			// mouse at hole -> mouse wins
			dp[0][m][t] = MOUSE_WIN
			q = append(q, state{0, m, t})
		}
		for c := 0; c < n; c++ {
			// cat meets mouse (not at hole) -> cat wins
			if c == 1 && m == 0 { continue }
			if c == m && m != 0 {
				dp[m][c][0] = CAT_WIN
				dp[m][c][1] = CAT_WIN
				q = append(q, state{m, c, 0})
				q = append(q, state{m, c, 1})
			}
		}
	}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		m, c, t := cur.m, cur.c, cur.t
		res := dp[m][c][t]

		// Get all parent states that can transition to this state
		parents := getParents(graph, m, c, t)

		for _, p := range parents {
			pm, pc, pt := p.m, p.c, p.t
			if dp[pm][pc][pt] != DRAW {
				continue
			}
			if pt == 0 { // mouse's turn
				if res == MOUSE_WIN {
					// mouse can move to a winning state
					dp[pm][pc][pt] = MOUSE_WIN
					q = append(q, state{pm, pc, pt})
				} else {
					// decrement degree
					degree[pm][pc][pt]--
					if degree[pm][pc][pt] == 0 {
						dp[pm][pc][pt] = CAT_WIN
						q = append(q, state{pm, pc, pt})
					}
				}
			} else { // cat's turn
				if res == CAT_WIN {
					// cat can move to a winning state
					dp[pm][pc][pt] = CAT_WIN
					q = append(q, state{pm, pc, pt})
				} else {
					degree[pm][pc][pt]--
					if degree[pm][pc][pt] == 0 {
						dp[pm][pc][pt] = MOUSE_WIN
						q = append(q, state{pm, pc, pt})
					}
				}
			}
		}
	}

	return dp[1][2][0]
}

func getParents(graph [][]int, m, c, t int) []state {
	var res []state
	if t == 0 { // current is mouse's turn, so mouse just moved -> previous was cat's turn
		// cat moved from some node to c
		for _, pc := range graph[c] {
			if pc == 0 { continue }
			res = append(res, state{m, pc, 1})
		}
	} else { // current is cat's turn, so cat just moved -> previous was mouse's turn
		for _, pm := range graph[m] {
			res = append(res, state{pm, c, 0})
		}
	}
	return res
}

func main() {
	graph := [][]int{{2,5},{3},{0,4,5},{1,4,5},{2,3},{0,2,3}}
	fmt.Println(catMouseGame(graph)) // Expected: 0
}
