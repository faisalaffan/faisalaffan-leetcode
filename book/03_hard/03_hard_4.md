# Hard (Sulit) — Problem ��1783

## 1449 — Form Largest Integer With Digits That Add Up To Target

```go
package main

// LeetCode #1449: Form Largest Integer With Digits That Add up to Target
// https://leetcode.com/problems/form-largest-integer-with-digits-that-add-up-to-target/
// Difficulty: Hard
//
// Given an array cost where cost[i] is the cost of digit (i+1), and a target,
// find the largest integer that can be formed with total cost equal to target.
// Digits can be used multiple times.
//
// Approach: DP to find maximum length for each cost, then reconstruct
// the largest number by trying digits from 9 down to 1.

import "fmt"

func main() {
	// Example 1
	fmt.Println(largestNumber([]int{4, 3, 2, 5, 6, 7, 2, 5, 5}, 9))
	// Example 2
	fmt.Println(largestNumber([]int{7, 6, 5, 5, 5, 6, 8, 7, 8}, 12))
	// Edge: impossible
	fmt.Println(largestNumber([]int{2, 4, 6, 2, 4, 6, 4, 4, 4}, 1))
	// Edge: single digit
	fmt.Println(largestNumber([]int{1, 1, 1, 1, 1, 1, 1, 1, 1}, 3))
}

func largestNumber(cost []int, target int) string {
	const inf = 1 << 30
	f := make([][]int, 10)
	g := make([][]int, 10)
	for i := range f {
		f[i] = make([]int, target+1)
		g[i] = make([]int, target+1)
		for j := range f[i] {
			f[i][j] = -inf
		}
	}
	f[0][0] = 0
	for i := 1; i <= 9; i++ {
		c := cost[i-1]
		for j := 0; j <= target; j++ {
			if j < c || f[i][j-c]+1 < f[i-1][j] {
				f[i][j] = f[i-1][j]
				g[i][j] = j
			} else {
				f[i][j] = f[i][j-c] + 1
				g[i][j] = j - c
			}
		}
	}
	if f[9][target] < 0 {
		return "0"
	}
	ans := []byte{}
	for i, j := 9, target; i > 0; {
		if g[i][j] == j {
			i--
		} else {
			ans = append(ans, '0'+byte(i))
			j = g[i][j]
		}
	}
	return string(ans)
}
```

## 1453 — Maximum Number Of Darts Inside Of A Circular Dartboard

```go
package main

// LeetCode #1453: Maximum Number of Darts Inside of a Circular Dartboard
// https://leetcode.com/problems/maximum-number-of-darts-inside-of-a-circular-dartboard/
// Difficulty: Hard
//
// Given points on a 2D plane and a radius r, find the maximum number
// of points that can be covered by a circle of radius r.
//
// Approach: Angular sweep. For each point as center candidate, compute
// angles of all other points that lie within 2r distance, then use
// sliding window to find max points in any semicircle.

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(numPoints([][]int{{-2, 0}, {2, 0}, {0, 2}, {0, -2}}, 2))
	// Example 2
	fmt.Println(numPoints([][]int{{-3, 0}, {3, 0}, {2, 6}, {5, 4}, {0, 9}, {7, 8}}, 5))
	// Edge: single point
	fmt.Println(numPoints([][]int{{0, 0}}, 1))
}

func numPoints(darts [][]int, r int) int {
	n := len(darts)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	rr := float64(r) * float64(r)
	ans := 1

	for i := 0; i < n; i++ {
		angles := make([]float64, 0, n*2)
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			dx := float64(darts[j][0] - darts[i][0])
			dy := float64(darts[j][1] - darts[i][1])
			d2 := dx*dx + dy*dy
			if d2 > 4*rr {
				continue
			}
			if d2 == 0 {
				continue
			}
			dist := math.Sqrt(d2)
			half := dist / 2.0
			// Angle from i to j
			base := math.Atan2(dy, dx)
			// Maximum deviation from base
			delta := math.Acos(half / float64(r))
			// Ensure delta is valid
			if !math.IsNaN(delta) {
				angles = append(angles, base-delta, base+delta+2*math.Pi)
			}
		}
		if len(angles) == 0 {
			continue
		}
		sort.Float64s(angles)

		m := len(angles)
		cnt := 0
		left := 0
		for right := 0; right < m; right++ {
			if angles[right]-angles[left] > 2*math.Pi+1e-9 {
				if angles[right]-angles[left] > 2*math.Pi+1e-9 {
					left++
				}
			}
			cnt = right - left + 1
			if cnt+1 > ans {
				ans = cnt + 1
			}
		}
	}
	return ans
}
```

## 1454 — Active Users

```go
package main

// LeetCode #1454: Active Users
// https://leetcode.com/problems/active-users/
// Difficulty: Medium (listed here as Hard) [Paid]
//
// Find users who had at least 5 consecutive login days in a given month.

import (
	"fmt"
	"sort"
)

// Login represents a login record.
type Login struct {
	UserID int
	Date   string // format: "YYYY-MM-DD"
}

// findActiveUsers returns user IDs with 5+ consecutive login days.
func findActiveUsers(logins []Login) []int {
	// Group logins by user, deduplicate dates
	userDates := make(map[int]map[string]bool)
	for _, l := range logins {
		if userDates[l.UserID] == nil {
			userDates[l.UserID] = make(map[string]bool)
		}
		userDates[l.UserID][l.Date] = true
	}

	var activeUsers []int
	for userID, dates := range userDates {
		// Convert to sorted slice of ints (days since epoch)
		var dayNums []int
		for d := range dates {
			dayNums = append(dayNums, dateToDays(d))
		}
		sort.Ints(dayNums)

		// Check for 5 consecutive days
		consecutive := 1
		for i := 1; i < len(dayNums); i++ {
			if dayNums[i]-dayNums[i-1] == 1 {
				consecutive++
				if consecutive >= 5 {
					activeUsers = append(activeUsers, userID)
					break
				}
			} else if dayNums[i] != dayNums[i-1] {
				consecutive = 1
			}
		}
	}

	sort.Ints(activeUsers)
	return activeUsers
}

// dateToDays converts a YYYY-MM-DD string to days since epoch.
func dateToDays(date string) int {
	var year, month, day int
	fmt.Sscanf(date, "%d-%d-%d", &year, &month, &day)

	// Simple day count (not perfectly accurate for dates before March, but sufficient for consecutive check)
	// Use a month offset table
	daysInMonth := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	// Count days from year 0
	total := year * 365
	// Add leap years
	total += (year + 3) / 4
	if year > 0 {
		total -= (year-1)/100 - (year-1)/400
	}

	// Add months for current year
	for m := 1; m < month; m++ {
		total += daysInMonth[m]
		if m == 2 && isLeap(year) {
			total++
		}
	}

	total += day
	return total
}

func isLeap(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}

func main() {
	// Test case 1: single user with 5 consecutive logins
	logins := []Login{
		{1, "2023-01-01"},
		{1, "2023-01-02"},
		{1, "2023-01-03"},
		{1, "2023-01-04"},
		{1, "2023-01-05"},
		{2, "2023-01-01"},
		{2, "2023-01-03"},
		{2, "2023-01-05"},
	}

	active := findActiveUsers(logins)
	fmt.Printf("Test 1 - Active users: %v (expected [1])\n", active)

	// Test case 2: user with logins spanning across month boundary
	logins2 := []Login{
		{1, "2023-01-30"},
		{1, "2023-01-31"},
		{1, "2023-02-01"},
		{1, "2023-02-02"},
		{1, "2023-02-03"},
	}
	active2 := findActiveUsers(logins2)
	fmt.Printf("Test 2 - Active users (month boundary): %v (expected [1])\n", active2)

	// Test case 3: no active users
	logins3 := []Login{
		{1, "2023-06-01"},
		{1, "2023-06-03"},
		{1, "2023-06-05"},
		{2, "2023-06-02"},
	}
	active3 := findActiveUsers(logins3)
	fmt.Printf("Test 3 - Active users (none): %v (expected [])\n", active3)

	// Test case 4: multiple active users
	logins4 := []Login{
		{1, "2023-01-01"},
		{1, "2023-01-02"},
		{1, "2023-01-03"},
		{1, "2023-01-04"},
		{1, "2023-01-05"},
		{2, "2023-02-10"},
		{2, "2023-02-11"},
		{2, "2023-02-12"},
		{2, "2023-02-13"},
		{2, "2023-02-14"},
	}
	active4 := findActiveUsers(logins4)
	fmt.Printf("Test 4 - Active users (multiple): %v (expected [1 2])\n", active4)

	// Test case 5: duplicated dates
	logins5 := []Login{
		{1, "2023-01-01"},
		{1, "2023-01-01"}, // duplicate
		{1, "2023-01-02"},
		{1, "2023-01-03"},
		{1, "2023-01-04"},
		{1, "2023-01-05"},
		{1, "2023-01-05"}, // duplicate
	}
	active5 := findActiveUsers(logins5)
	fmt.Printf("Test 5 - Active users (duplicates): %v (expected [1])\n", active5)
}
```

## 1458 — Max Dot Product Of Two Subsequences

```go
package main

// LeetCode #1458: Max Dot Product of Two Subsequences
// https://leetcode.com/problems/max-dot-product-of-two-subsequences/
// Difficulty: Hard

import "fmt"

func maxDotProduct(nums1 []int, nums2 []int) int {
	n1, n2 := len(nums1), len(nums2)
	dp := make([][]int, n1)
	for i := range dp {
		dp[i] = make([]int, n2)
	}

	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			product := nums1[i] * nums2[j]
			dp[i][j] = product
			if i > 0 && dp[i-1][j] > dp[i][j] {
				dp[i][j] = dp[i-1][j]
			}
			if j > 0 && dp[i][j-1] > dp[i][j] {
				dp[i][j] = dp[i][j-1]
			}
			if i > 0 && j > 0 {
				candidate := dp[i-1][j-1]
				if candidate > 0 {
					candidate += product
				} else {
					candidate = product
				}
				if candidate > dp[i][j] {
					dp[i][j] = candidate
				}
			}
		}
	}
	return dp[n1-1][n2-1]
}

func main() {
	// Example: [2,1,-2,5], [3,0,-6] -> 18
	fmt.Println(maxDotProduct([]int{2, 1, -2, 5}, []int{3, 0, -6}))
}
```

## 1459 — Rectangles Area

```go
package main

// LeetCode #1459: Rectangles Area
// https://leetcode.com/problems/rectangles-area/
// Difficulty: Medium (listed here as Hard) [Paid]
//
// Given points (x, y) in a 2D plane, find all rectangles formed by these points.
// A rectangle has sides parallel to the axes. Return the area and the two
// opposite corner points for each rectangle.

import (
	"fmt"
	"sort"
)

// Point represents a 2D coordinate.
type Point struct {
	X, Y int
}

// Rectangle represents a rectangle with two opposite corners.
type Rectangle struct {
	Area     int
	P1, P2   Point
}

// findRectangles finds all axis-aligned rectangles from given points.
func findRectangles(points []Point) []Rectangle {
	// Build a set of points for O(1) lookup
	pointSet := make(map[Point]bool)
	for _, p := range points {
		pointSet[p] = true
	}

	var rects []Rectangle
	seen := make(map[string]bool)

	// For each pair of points, check if they can be opposite corners
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1, p2 := points[i], points[j]

			// They must be diagonal corners (different x AND different y)
			if p1.X == p2.X || p1.Y == p2.Y {
				continue
			}

			// The other two corners must exist
			c3 := Point{p1.X, p2.Y}
			c4 := Point{p2.X, p1.Y}

			if pointSet[c3] && pointSet[c4] {
				// Ensure we don't add the same rectangle twice
				// by normalizing: smaller x first
				rectP1, rectP2 := p1, p2
				if rectP1.X > rectP2.X || (rectP1.X == rectP2.X && rectP1.Y > rectP2.Y) {
					rectP1, rectP2 = rectP2, rectP1
				}
				key := fmt.Sprintf("%d,%d-%d,%d", rectP1.X, rectP1.Y, rectP2.X, rectP2.Y)
				if seen[key] {
					continue
				}
				seen[key] = true

				area := abs(p2.X-p1.X) * abs(p2.Y-p1.Y)
				if area > 0 {
					rects = append(rects, Rectangle{area, rectP1, rectP2})
				}
			}
		}
	}

	// Sort by area descending, then by p1.x, p1.y, p2.x, p2.y
	sort.Slice(rects, func(i, j int) bool {
		if rects[i].Area != rects[j].Area {
			return rects[i].Area > rects[j].Area
		}
		if rects[i].P1.X != rects[j].P1.X {
			return rects[i].P1.X < rects[j].P1.X
		}
		if rects[i].P1.Y != rects[j].P1.Y {
			return rects[i].P1.Y < rects[j].P1.Y
		}
		if rects[i].P2.X != rects[j].P2.X {
			return rects[i].P2.X < rects[j].P2.X
		}
		return rects[i].P2.Y < rects[j].P2.Y
	})

	return rects
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Test case 1: Simple rectangle
	points := []Point{
		{0, 0}, {0, 2}, {2, 0}, {2, 2},
	}
	rects := findRectangles(points)
	fmt.Println("Test 1 - Rectangles:")
	for _, r := range rects {
		fmt.Printf("  Area=%d, (%d,%d)-(%d,%d)\n", r.Area, r.P1.X, r.P1.Y, r.P2.X, r.P2.Y)
	}

	// Test case 2: Two overlapping rectangles
	points2 := []Point{
		{0, 0}, {0, 2}, {2, 0}, {2, 2},
		{1, 0}, {1, 2},
	}
	rects2 := findRectangles(points2)
	fmt.Println("\nTest 2 - Rectangles (overlapping):")
	for _, r := range rects2 {
		fmt.Printf("  Area=%d, (%d,%d)-(%d,%d)\n", r.Area, r.P1.X, r.P1.Y, r.P2.X, r.P2.Y)
	}

	// Test case 3: No rectangles (collinear)
	points3 := []Point{
		{0, 0}, {1, 1}, {2, 2},
	}
	rects3 := findRectangles(points3)
	fmt.Printf("\nTest 3 - Rectangles (collinear): %d (expected 0)\n", len(rects3))

	// Test case 4: Multiple rectangles with different areas
	points4 := []Point{
		{0, 0}, {0, 3}, {3, 0}, {3, 3},
		{1, 1}, {1, 2}, {2, 1}, {2, 2},
	}
	rects4 := findRectangles(points4)
	fmt.Println("\nTest 4 - Rectangles (nested):")
	for _, r := range rects4 {
		fmt.Printf("  Area=%d, (%d,%d)-(%d,%d)\n", r.Area, r.P1.X, r.P1.Y, r.P2.X, r.P2.Y)
	}

	// Test case 5: Single point
	points5 := []Point{{5, 5}}
	rects5 := findRectangles(points5)
	fmt.Printf("\nTest 5 - Rectangles (single point): %d (expected 0)\n", len(rects5))
}
```

## 1463 — Cherry Pickup Ii

```go
package main

// LeetCode #1463: Cherry Pickup II
// https://leetcode.com/problems/cherry-pickup-ii/
// Difficulty: Hard

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func cherryPickup(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	// dp[r][c1][c2] = max cherries at row r, robot1 col c1, robot2 col c2
	dp := make([][][]int, rows)
	for i := range dp {
		dp[i] = make([][]int, cols)
		for j := range dp[i] {
			dp[i][j] = make([]int, cols)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}
	dp[0][0][cols-1] = grid[0][0] + grid[0][cols-1]

	for r := 1; r < rows; r++ {
		for c1 := 0; c1 < cols; c1++ {
			for c2 := 0; c2 < cols; c2++ {
				prev := -1
				for d1 := -1; d1 <= 1; d1++ {
					for d2 := -1; d2 <= 1; d2++ {
						pc1 := c1 + d1
						pc2 := c2 + d2
						if pc1 >= 0 && pc1 < cols && pc2 >= 0 && pc2 < cols {
							prev = max(prev, dp[r-1][pc1][pc2])
						}
					}
				}
				if prev >= 0 {
					cherries := grid[r][c1] + grid[r][c2]
					if c1 == c2 {
						cherries -= grid[r][c1] // don't double count
					}
					dp[r][c1][c2] = prev + cherries
				}
			}
		}
	}

	ans := 0
	for c1 := 0; c1 < cols; c1++ {
		for c2 := 0; c2 < cols; c2++ {
			ans = max(ans, dp[rows-1][c1][c2])
		}
	}
	return ans
}

func main() {
	// Example: [[3,1,1],[2,5,1],[1,5,5],[2,1,1]] -> 24
	fmt.Println(cherryPickup([][]int{{3, 1, 1}, {2, 5, 1}, {1, 5, 5}, {2, 1, 1}}))
}
```

## 1467 — Probability Of A Two Boxes Having The Same Number Of Distinct Balls

```go
package main

// LeetCode #1467: Probability of a Two Boxes Having The Same Number of Distinct Balls
// https://leetcode.com/problems/probability-of-a-two-boxes-having-the-same-number-of-distinct-balls/
// Difficulty: Hard
//
// Given 2n balls of k distinct colors, each with a given count, randomly
// distribute all balls into two boxes (each gets n balls). Return the
// probability that both boxes have the same number of distinct colors.
//
// Approach: DP with combinatorics. Use DFS to distribute balls of each
// color between the two boxes, counting valid distributions.

import (
	"fmt"
)

func main() {
	// Example 1
	fmt.Println(getProbability([]int{1, 1}))
	// Example 2
	fmt.Println(getProbability([]int{2}))
	// Example 3
	fmt.Println(getProbability([]int{1, 2, 3}))
}

func getProbability(balls []int) float64 {
	n, mx := 0, 0
	for _, x := range balls {
		n += x
		mx = max(mx, x)
	}
	n >>= 1
	m := max(mx, n<<1)
	c := make([][]int, m+1)
	for i := range c {
		c[i] = make([]int, m+1)
	}
	for i := 0; i <= m; i++ {
		c[i][0] = 1
		for j := 1; j <= i; j++ {
			c[i][j] = c[i-1][j-1] + c[i-1][j]
		}
	}
	k := len(balls)
	f := make([][][]int, k)
	for i := range f {
		f[i] = make([][]int, n+1)
		for j := range f[i] {
			f[i][j] = make([]int, k<<1|1)
			for h := range f[i][j] {
				f[i][j][h] = -1
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, diff int) int {
		if i >= k {
			if j == 0 && diff == k {
				return 1
			}
			return 0
		}
		if j < 0 {
			return 0
		}
		if f[i][j][diff] != -1 {
			return f[i][j][diff]
		}
		ans := 0
		for x := 0; x <= balls[i]; x++ {
			y := 1
			if x != balls[i] {
				if x == 0 {
					y = -1
				} else {
					y = 0
				}
			}
			ans += dfs(i+1, j-x, diff+y) * c[balls[i]][x]
		}
		f[i][j][diff] = ans
		return ans
	}
	return float64(dfs(0, n, k)) / float64(c[n<<1][n])
}
```

## 1468 — Calculate Salaries

```go
package main

// LeetCode #1468: Calculate Salaries
// https://leetcode.com/problems/calculate-salaries/
// Difficulty: Medium (listed here as Hard) [Paid]
//
// Calculate salary after tax. The tax rate depends on the company's
// max salary in the same department. If max salary in the department
// is < 1000, tax is 0%. If max is between 1000 and 10000 (inclusive),
// tax is 24%. If max is > 10000, tax is 49%.
// Also, salary after tax is rounded.

import "fmt"

// Employee represents an employee record.
type Employee struct {
	ID         int
	Department string
	Salary     int
}

// calculateSalaries computes post-tax salaries for all employees.
func calculateSalaries(employees []Employee) map[int]int {
	// Find max salary per department
	deptMax := make(map[string]int)
	for _, e := range employees {
		if e.Salary > deptMax[e.Department] {
			deptMax[e.Department] = e.Salary
		}
	}

	// Calculate tax rate per department and apply to each employee
	result := make(map[int]int)
	for _, e := range employees {
		maxSal := deptMax[e.Department]
		var taxRate float64
		if maxSal < 1000 {
			taxRate = 0.0
		} else if maxSal <= 10000 {
			taxRate = 0.24
		} else {
			taxRate = 0.49
		}

		afterTax := float64(e.Salary) * (1.0 - taxRate)
		result[e.ID] = roundSalary(afterTax)
	}

	return result
}

// roundSalary rounds to nearest integer, with .5 going up.
func roundSalary(val float64) int {
	if val < 0 {
		return int(val - 0.5)
	}
	return int(val + 0.5)
}

func main() {
	employees := []Employee{
		{1, "Engineering", 8000},
		{2, "Engineering", 6000},
		{3, "Marketing", 500},
		{4, "Marketing", 700},
		{5, "Sales", 12000},
		{6, "Sales", 11000},
	}

	salaries := calculateSalaries(employees)
	fmt.Println("Salaries after tax:")
	employeeNames := map[int]string{1: "Alice", 2: "Bob", 3: "Charlie", 4: "David", 5: "Eve", 6: "Frank"}
	for _, e := range employees {
		fmt.Printf("  %s (%s): $%d -> $%d\n", employeeNames[e.ID], e.Department, e.Salary, salaries[e.ID])
	}

	// Test 2: Dept with max < 1000 (0% tax)
	employees2 := []Employee{
		{1, "Support", 800},
		{2, "Support", 900},
	}
	salaries2 := calculateSalaries(employees2)
	fmt.Println("\nTest 2 - Low salary dept (0% tax):")
	for _, e := range employees2 {
		fmt.Printf("  ID %d: $%d -> $%d (expected $%d)\n", e.ID, e.Salary, salaries2[e.ID], e.Salary)
	}

	// Test 3: Multiple departments with different rates
	employees3 := []Employee{
		{1, "DeptA", 500},    // max=500, 0% -> 500
		{2, "DeptA", 300},    // max=500, 0% -> 300
		{3, "DeptB", 5000},   // max=5000, 24% -> 3800
		{4, "DeptB", 2000},   // max=5000, 24% -> 1520
		{5, "DeptC", 20000},  // max=20000, 49% -> 10200
		{6, "DeptC", 15000},  // max=20000, 49% -> 7650
	}
	salaries3 := calculateSalaries(employees3)
	fmt.Println("\nTest 3 - Mixed departments:")
	for _, e := range employees3 {
		fmt.Printf("  ID %d (dept %s): $%d -> $%d\n", e.ID, e.Department, e.Salary, salaries3[e.ID])
	}
}
```

## 1473 — Paint House Iii

```go
package main

// LeetCode #1473: Paint House III
// https://leetcode.com/problems/paint-house-iii/
// Difficulty: Hard
//
// Approach: 3D DP
// dp[i][j][k] = min cost to paint first i houses (0-indexed), where
// house i-1 is painted color j, and there are exactly k neighborhoods.
// Transition:
//   - If houses[i-1] != 0 (already painted), only that color is allowed.
//   - If houses[i-1] == 0, try all colors 1..n.
//   - k increments when current color != previous house's color.
// Answer: min over dp[m][*][target].

import (
	"fmt"
	"math"
)

func main() {
	// Example 1
	fmt.Println(minCost([]int{0, 0, 0, 0, 0}, [][]int{{1, 10}, {10, 1}, {10, 1}, {1, 10}, {5, 1}}, 5, 2, 3))
	// Expected: 9

	// Example 2
	fmt.Println(minCost([]int{0, 2, 1, 2, 0}, [][]int{{1, 10}, {10, 1}, {10, 1}, {1, 10}, {5, 1}}, 5, 2, 3))
	// Expected: 11

	// Example 3
	fmt.Println(minCost([]int{3, 1, 2, 3}, [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, 4, 3, 3))
	// Expected: -1 (not possible)
}

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	const INF = math.MaxInt32

	// dp[i][j][k] — working with 1-indexed for i and k for simplicity
	dp := make([][][]int, m+1)
	for i := range dp {
		dp[i] = make([][]int, n+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, target+1)
			for k := range dp[i][j] {
				dp[i][j][k] = INF
			}
		}
	}

	// Base: 0 houses, 0 neighborhoods, any "last color" is 0 cost
	// We'll treat house index 0 (0 houses processed) specially.
	// Actually, let's use 0-based house index and 1-based neighborhood.
	// Simpler: dp[h][c][t] = min cost for first h houses (h from 1..m),
	// last house color c (1..n), exactly t neighborhoods (1..target).
	// Initialize first house.

	prevColor := 0
	for c := 1; c <= n; c++ {
		if houses[0] != 0 && houses[0] != c {
			continue
		}
		paintCost := 0
		if houses[0] == 0 {
			paintCost = cost[0][c-1]
		}
		dp[1][c][1] = paintCost
		prevColor = c
	}
	_ = prevColor

	// Fill DP
	for i := 2; i <= m; i++ {
		for c := 1; c <= n; c++ {
			if houses[i-1] != 0 && houses[i-1] != c {
				continue
			}
			paintCost := 0
			if houses[i-1] == 0 {
				paintCost = cost[i-1][c-1]
			}
			for t := 1; t <= target && t <= i; t++ {
				best := INF
				// Case 1: same color as previous house
				if dp[i-1][c][t] < INF {
					best = minInt(best, dp[i-1][c][t]+paintCost)
				}
				// Case 2: different color from previous house
				if t > 1 {
					for pc := 1; pc <= n; pc++ {
						if pc == c {
							continue
						}
						if dp[i-1][pc][t-1] < INF {
							best = minInt(best, dp[i-1][pc][t-1]+paintCost)
						}
					}
				}
				dp[i][c][t] = best
			}
		}
	}

	ans := INF
	for c := 1; c <= n; c++ {
		if dp[m][c][target] < ans {
			ans = dp[m][c][target]
		}
	}
	if ans == INF {
		return -1
	}
	return ans
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

## 1478 — Allocate Mailboxes

```go
package main

// LeetCode #1478: Allocate Mailboxes
// https://leetcode.com/problems/allocate-mailboxes/
// Difficulty: Hard
//
// Approach: DP + Median Cost
// Sort houses first. dp[i][j] = min distance to place j mailboxes
// among first i houses (0-indexed).
// cost[l][r] = min total distance to serve houses[l..r] with 1 mailbox
// placed at the median (optimal for minimizing sum of absolute distances).
// Transition: dp[i][j] = min over p < i of dp[p][j-1] + cost[p+1][i].

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(minDistance([]int{1, 4, 8, 10, 20}, 3))
	// Expected: 5

	fmt.Println(minDistance([]int{2, 3, 5, 12, 18}, 2))
	// Expected: 9

	fmt.Println(minDistance([]int{7, 4, 6, 1}, 1))
	// Expected: 8
}

func minDistance(houses []int, k int) int {
	sort.Ints(houses)
	n := len(houses)

	if k >= n {
		return 0
	}

	// precompute cost[i][j] = min dist for 1 mailbox serving houses[i..j]
	cost := make([][]int, n)
	for i := range cost {
		cost[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			// Median index
			mid := i + (j-i)/2
			median := houses[mid]
			total := 0
			for t := i; t <= j; t++ {
				total += absInt(houses[t] - median)
			}
			cost[i][j] = total
		}
	}

	// dp[i][j] = min distance for first i+1 houses with j+1 mailboxes
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}

	// base: 1 mailbox
	for i := 0; i < n; i++ {
		dp[i][0] = cost[0][i]
	}

	// fill dp
	for j := 1; j < k; j++ {
		for i := j; i < n; i++ {
			for p := j - 1; p < i; p++ {
				dp[i][j] = minInt(dp[i][j], dp[p][j-1]+cost[p+1][i])
			}
		}
	}

	return dp[n-1][k-1]
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

## 1479 — Sales By Day Of The Week

```go
package main

// LeetCode #1479: Sales by Day of the Week
// https://leetcode.com/problems/sales-by-day-of-the-week/
// Difficulty: Hard [Paid]
//
// Calculate total sales for each item, each day of the week.

import "fmt"
import "time"

// Sale represents a sales record.
type Sale struct {
	ItemID   int
	Date     string // "YYYY-MM-DD"
	Quantity int
	Price    float64
}

// DaySales represents sales summary for an item on a day of the week.
type DaySales struct {
	ItemID int
	Day    string // Monday, Tuesday, etc.
	Total  float64
}

// calculateDaySales computes total sales by item and day of the week.
func calculateDaySales(sales []Sale) []DaySales {
	// We need timezone for date parsing
	loc := time.UTC

	// Aggregate: itemID -> dayOfWeek -> total
	type key struct {
		itemID int
		day    string
	}
	aggregate := make(map[key]float64)

	daysOrder := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	for _, s := range sales {
		t, err := time.ParseInLocation("2006-01-02", s.Date, loc)
		if err != nil {
			continue
		}
		dayName := t.Weekday().String()
		// Normalize to "Monday", "Tuesday", etc.
		// Go's Weekday() returns Sunday=0, Monday=1, ..., Saturday=6
		aggregate[key{s.ItemID, dayName}] += float64(s.Quantity) * s.Price
	}

	// Collect results
	var results []DaySales
	for k, total := range aggregate {
		results = append(results, DaySales{k.itemID, k.day, total})
	}

	// Sort: by ItemID ASC, then by day of week order
	dayRank := make(map[string]int)
	for i, d := range daysOrder {
		dayRank[d] = i
	}
	// Also handle Go's Weekday() output: "Sunday", "Monday", etc.
	goDayRank := map[string]int{
		"Sunday":    0,
		"Monday":    1,
		"Tuesday":   2,
		"Wednesday": 3,
		"Thursday":  4,
		"Friday":    5,
		"Saturday":  6,
	}

	_ = dayRank // keep for reference

	for i := 0; i < len(results); i++ {
		for j := i + 1; j < len(results); j++ {
			swap := false
			if results[i].ItemID > results[j].ItemID {
				swap = true
			} else if results[i].ItemID == results[j].ItemID {
				if goDayRank[results[i].Day] > goDayRank[results[j].Day] {
					swap = true
				}
			}
			if swap {
				results[i], results[j] = results[j], results[i]
			}
		}
	}

	return results
}

func main() {
	sales := []Sale{
		{1, "2023-01-02", 5, 10.00},   // Monday
		{1, "2023-01-03", 3, 10.00},   // Tuesday
		{1, "2023-01-09", 2, 10.00},   // Monday
		{2, "2023-01-02", 4, 20.00},   // Monday
		{2, "2023-01-04", 1, 20.00},   // Wednesday
	}

	results := calculateDaySales(sales)
	fmt.Println("Sales by Day of Week:")
	for _, r := range results {
		fmt.Printf("  Item %d, %s: $%.2f\n", r.ItemID, r.Day, r.Total)
	}

	// Test 2: Weekend sales
	sales2 := []Sale{
		{1, "2023-01-07", 10, 5.00},   // Saturday
		{1, "2023-01-08", 8, 5.00},    // Sunday
		{2, "2023-01-07", 3, 15.00},   // Saturday
	}
	results2 := calculateDaySales(sales2)
	fmt.Println("\nTest 2 - Weekend sales:")
	for _, r := range results2 {
		fmt.Printf("  Item %d, %s: $%.2f\n", r.ItemID, r.Day, r.Total)
	}

	// Test 3: Empty
	results3 := calculateDaySales(nil)
	fmt.Printf("\nTest 3 - Empty: %d results\n", len(results3))

	// Test 4: Single sale
	sales4 := []Sale{
		{1, "2023-06-01", 2, 25.50},  // Thursday
	}
	results4 := calculateDaySales(sales4)
	fmt.Println("\nTest 4 - Single sale:")
	for _, r := range results4 {
		fmt.Printf("  Item %d, %s: $%.2f\n", r.ItemID, r.Day, r.Total)
	}
}
```

## 1483 — Kth Ancestor Of A Tree Node

```go
package main

// LeetCode #1483: Kth Ancestor of a Tree Node
// https://leetcode.com/problems/kth-ancestor-of-a-tree-node/
// Difficulty: Hard
//
// Approach: Binary Lifting (doubling)
// up[node][i] = 2^i-th ancestor of node.
// up[node][0] = parent[node]
// up[node][i] = up[up[node][i-1]][i-1]
// getKthAncestor: for each bit of k, jump up.

import "fmt"

func main() {
	// Example: n=7, parent=[-1,0,0,1,1,2,2]
	ta := Constructor(7, []int{-1, 0, 0, 1, 1, 2, 2})
	fmt.Println(ta.GetKthAncestor(3, 1)) // 1
	fmt.Println(ta.GetKthAncestor(5, 2)) // 0
	fmt.Println(ta.GetKthAncestor(6, 3)) // -1

	// Edge case
	ta2 := Constructor(1, []int{-1})
	fmt.Println(ta2.GetKthAncestor(0, 1)) // -1
}

type TreeAncestor struct {
	up [][]int // up[node][i]
}

func Constructor(n int, parent []int) TreeAncestor {
	LOG := 1
	for (1 << LOG) <= n {
		LOG++
	}
	up := make([][]int, n)
	for i := 0; i < n; i++ {
		up[i] = make([]int, LOG)
		up[i][0] = parent[i]
	}
	for j := 1; j < LOG; j++ {
		for i := 0; i < n; i++ {
			if up[i][j-1] >= 0 {
				up[i][j] = up[up[i][j-1]][j-1]
			} else {
				up[i][j] = -1
			}
		}
	}
	return TreeAncestor{up: up}
}

func (ta *TreeAncestor) GetKthAncestor(node int, k int) int {
	LOG := len(ta.up[0])
	for j := 0; j < LOG; j++ {
		if k&(1<<j) != 0 {
			node = ta.up[node][j]
			if node == -1 {
				return -1
			}
		}
	}
	return node
}
```

## 1488 — Avoid Flood In The City

```go
package main

// LeetCode #1488: Avoid Flood in The City
// https://leetcode.com/problems/avoid-flood-in-the-city/
// Difficulty: Hard
//
// rains[i] > 0 means lake rains[i] fills up.
// rains[i] == 0 means we can dry one lake.
// Return an array ans where ans[i] = -1 if rains[i] > 0,
// else ans[i] = the lake we dry (any). If flood cannot be avoided, return [].

import (
	"fmt"
	"sort"
)

// avoidFlood returns an array of actions to avoid flooding.
func avoidFlood(rains []int) []int {
	n := len(rains)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = 1 // default dry value for zero-rain days
	}

	// Map lake -> last day it rained
	lastRain := make(map[int]int)

	// Collect dry days (when rains[i] == 0)
	var dryDays []int

	for i, lake := range rains {
		if lake == 0 {
			dryDays = append(dryDays, i)
			continue
		}

		ans[i] = -1 // it's raining on this day

		if prevDay, ok := lastRain[lake]; ok {
			// This lake is already full, we need to dry it before today
			// Find a dry day after prevDay to dry this lake
			idx := sort.Search(len(dryDays), func(j int) bool {
				return dryDays[j] > prevDay
			})

			if idx >= len(dryDays) {
				// Cannot dry this lake in time -> flood unavoidable
				return []int{}
			}

			dryDay := dryDays[idx]
			ans[dryDay] = lake
			// Remove this dry day from the list
			dryDays = append(dryDays[:idx], dryDays[idx+1:]...)
		}

		lastRain[lake] = i
	}

	return ans
}

func main() {
	// Test case 1
	rains1 := []int{1, 2, 3, 4}
	result1 := avoidFlood(rains1)
	fmt.Printf("Test 1: rains=%v => %v (expected [-1,-1,-1,-1])\n", rains1, result1)

	// Test case 2
	rains2 := []int{1, 2, 0, 0, 2, 1}
	result2 := avoidFlood(rains2)
	fmt.Printf("Test 2: rains=%v => %v (expected [-1,-1,2,1,-1,-1])\n", rains2, result2)

	// Test case 3
	rains3 := []int{1, 2, 0, 1, 2}
	result3 := avoidFlood(rains3)
	fmt.Printf("Test 3: rains=%v => %v (expected [])\n", rains3, result3)

	// Test case 4
	rains4 := []int{69, 0, 0, 0, 69}
	result4 := avoidFlood(rains4)
	fmt.Printf("Test 4: rains=%v => %v\n", rains4, result4)

	// Test case 5: LeetCode example
	rains5 := []int{1, 0, 2, 0, 2, 1}
	result5 := avoidFlood(rains5)
	fmt.Printf("Test 5: rains=%v => %v\n", rains5, result5)

	// Test case 6: impossible - two floods without dry day
	rains6 := []int{1, 1}
	result6 := avoidFlood(rains6)
	fmt.Printf("Test 6: rains=%v => %v (expected [])\n", rains6, result6)
}
```

## 1489 — Find Critical And Pseudo Critical Edges In Minimum Spanning Tree

```go
package main

// LeetCode #1489: Find Critical and Pseudo-Critical Edges in Minimum Spanning Tree
// https://leetcode.com/problems/find-critical-and-pseudo-critical-edges-in-minimum-spanning-tree/
// Difficulty: Hard
//
// Approach: Kruskal per Edge
// 1. Compute MST weight normally.
// 2. For each edge i:
//    a. Forced-in: include edge i first, then run Kruskal. If weight > MST or
//       not all nodes connected → edge i is critical.
//    b. Forced-out: run Kruskal excluding edge i. If weight > MST or not all
//       connected → edge i is pseudo-critical (since excluding it makes MST worse,
//       but it's not critical if forced-in also yields MST weight).
// Return [critical, pseudoCritical].

import (
	"fmt"
	"sort"
)

func main() {
	n := 5
	edges := [][]int{
		{0, 1, 1},
		{1, 2, 1},
		{2, 3, 2},
		{0, 3, 2},
		{0, 4, 3},
		{3, 4, 3},
		{1, 4, 6},
	}
	result := findCriticalAndPseudoCriticalEdges(n, edges)
	fmt.Println(result)
	// Expected: [[0,1],[2,3,4,5]]
}

type edge struct {
	u, v, w, idx int
}

type dSU struct {
	parent []int
	rank   []int
}

func newDSU(n int) *dSU {
	p := make([]int, n)
	r := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	return &dSU{parent: p, rank: r}
}

func (d *dSU) find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.find(d.parent[x])
	}
	return d.parent[x]
}

func (d *dSU) union(x, y int) bool {
	x, y = d.find(x), d.find(y)
	if x == y {
		return false
	}
	if d.rank[x] < d.rank[y] {
		x, y = y, x
	}
	d.parent[y] = x
	if d.rank[x] == d.rank[y] {
		d.rank[x]++
	}
	return true
}

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	m := len(edges)
	elist := make([]edge, m)
	for i, e := range edges {
		elist[i] = edge{u: e[0], v: e[1], w: e[2], idx: i}
	}
	sort.Slice(elist, func(i, j int) bool {
		return elist[i].w < elist[j].w
	})

	// Helper: compute MST weight with optional forced-in and forced-out edges.
	// forcedIn == nil means no forced edge.
	// forcedOut == -1 means no exclusion.
	mstWeight := func(forcedIn *edge, forcedOut int) int {
		dsu := newDSU(n)
		weight := 0
		edgesUsed := 0

		if forcedIn != nil {
			if dsu.union(forcedIn.u, forcedIn.v) {
				weight += forcedIn.w
				edgesUsed++
			}
		}

		for _, e := range elist {
			if e.idx == forcedOut {
				continue
			}
			if forcedIn != nil && e.idx == forcedIn.idx {
				continue
			}
			if dsu.union(e.u, e.v) {
				weight += e.w
				edgesUsed++
			}
		}

		if edgesUsed != n-1 {
			return -1 // not connected
		}
		return weight
	}

	baseWeight := mstWeight(nil, -1)

	critical := make([]int, 0)
	pseudo := make([]int, 0)

	for _, e := range elist {
		// Forced out (exclude this edge)
		wWithout := mstWeight(nil, e.idx)

		// If MST weight increases or graph disconnects, this edge is critical
		if wWithout == -1 || wWithout > baseWeight {
			critical = append(critical, e.idx)
			continue
		}

		// Forced in (include this edge)
		wWith := mstWeight(&e, -1)

		// If including it still gives MST weight, it's pseudo-critical
		// (it's in some MST but not all)
		if wWith == baseWeight {
			pseudo = append(pseudo, e.idx)
		}
	}

	return [][]int{critical, pseudo}
}
```

## 1490 — Clone N Ary Tree

```go
package main

// LeetCode #1490: Clone N-ary Tree
// https://leetcode.com/problems/clone-n-ary-tree/
// Difficulty: Medium (listed here as Hard)
//
// Clone an N-ary tree. Each node has a value and a list of children.

import "fmt"

// Node represents an N-ary tree node.
type Node struct {
	Val      int
	Children []*Node
}

// cloneTree creates a deep copy of the N-ary tree.
func cloneTree(root *Node) *Node {
	if root == nil {
		return nil
	}

	// Recursively clone children
	clonedChildren := make([]*Node, len(root.Children))
	for i, child := range root.Children {
		clonedChildren[i] = cloneTree(child)
	}

	return &Node{
		Val:      root.Val,
		Children: clonedChildren,
	}
}

// Helper function to compare two trees for testing
func equalTrees(a, b *Node) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	if len(a.Children) != len(b.Children) {
		return false
	}
	for i := range a.Children {
		if !equalTrees(a.Children[i], b.Children[i]) {
			return false
		}
	}
	return true
}

// Helper to print tree (preorder)
func printTree(root *Node, indent int) {
	if root == nil {
		return
	}
	for i := 0; i < indent; i++ {
		fmt.Print("  ")
	}
	fmt.Printf("Node(%d)\n", root.Val)
	for _, child := range root.Children {
		printTree(child, indent+1)
	}
}

func main() {
	// Test case 1: single node
	root1 := &Node{Val: 1}
	clone1 := cloneTree(root1)
	fmt.Printf("Test 1 - Clone of single node: equal=%v (expected true)\n", equalTrees(root1, clone1))
	fmt.Printf("  Original pointer: %p, Clone pointer: %p (different: %v)\n",
		root1, clone1, root1 != clone1)

	// Test case 2: tree with children
	root2 := &Node{
		Val: 1,
		Children: []*Node{
			{Val: 2, Children: []*Node{
				{Val: 5},
				{Val: 6},
			}},
			{Val: 3},
			{Val: 4, Children: []*Node{
				{Val: 7},
				{Val: 8},
				{Val: 9},
			}},
		},
	}
	clone2 := cloneTree(root2)
	fmt.Printf("\nTest 2 - Clone of complex tree: equal=%v (expected true)\n", equalTrees(root2, clone2))

	// Modify original and ensure clone is unchanged
	root2.Children[0].Val = 99
	fmt.Printf("  After modifying original: equal=%v (expected false)\n", equalTrees(root2, clone2))

	fmt.Println("\n  Original tree:")
	printTree(clone2, 0)
	fmt.Println("  (clone preserved after original modification)")

	// Test case 3: nil tree
	clone3 := cloneTree(nil)
	fmt.Printf("\nTest 3 - Clone of nil: %v (expected nil)\n", clone3)

	// Test case 4: deep structural verification
	root4 := &Node{
		Val: 10,
		Children: []*Node{
			{Val: 20},
		},
	}
	clone4 := cloneTree(root4)
	fmt.Printf("\nTest 4 - Structural equality: equal=%v (expected true)\n", equalTrees(root4, clone4))
	// Verify they are truly independent
	root4.Children[0].Val = 30
	fmt.Printf("  After modifying child: equal=%v (expected false)\n", equalTrees(root4, clone4))
	fmt.Printf("  Clone child val: %d (expected 20)\n", clone4.Children[0].Val)
}
```

## 1494 — Parallel Courses Ii

```go
package main

// LeetCode #1494: Parallel Courses II
// https://leetcode.com/problems/parallel-courses-ii/
// Difficulty: Hard
//
// Approach: DP over Bitmask
// prereq[mask] = bitmask of courses that are prerequisites for courses in mask.
// Actually, we precompute pre[c] = bitmask of direct prerequisites for course c.
// dp[mask] = minimum semesters to complete courses in mask.
// For each mask, compute available = all courses whose prerequisites are satisfied
// (i.e., pre[c] & mask == pre[c] for each c not in mask).
// Then try all subsets of available with size <= k, and transition:
// dp[mask | subset] = min(dp[mask | subset], dp[mask] + 1)

import "fmt"

func main() {
	// Example 1
	fmt.Println(minNumberOfSemesters(4, [][]int{{2, 1}, {3, 1}, {1, 4}}, 2))
	// Expected: 3

	// Example 2
	fmt.Println(minNumberOfSemesters(5, [][]int{{2, 1}, {3, 1}, {4, 1}, {5, 1}}, 2))
	// Expected: 3

	// Example 3
	fmt.Println(minNumberOfSemesters(11, [][]int{}, 2))
	// Expected: 6 (11 courses, 2 per semester = ceil(11/2) = 6)
}

func minNumberOfSemesters(n int, dependencies [][]int, k int) int {
	// pre[c] = bitmask of direct prerequisites for course c (1-indexed)
	pre := make([]int, n)
	for _, dep := range dependencies {
		// dep[0] -> dep[1], meaning dep[1] has prerequisite dep[0]
		// We use 0-indexed internally
		pre[dep[1]-1] |= 1 << (dep[0] - 1)
	}

	total := 1 << n
	INF := n + 1
	dp := make([]int, total)
	for i := 1; i < total; i++ {
		dp[i] = INF
	}
	dp[0] = 0

	// Precompute required prerequisites for each mask (union of all pre[c] for c in mask)
	require := make([]int, total)
	for mask := 1; mask < total; mask++ {
		// find lowest set bit
		lsb := mask & -mask
		c := 0
		// find index of lsb
		for (1 << c) != lsb {
			c++
		}
		require[mask] = require[mask^lsb] | pre[c]
	}

	for mask := 0; mask < total; mask++ {
		if dp[mask] == INF {
			continue
		}
		// Courses that can be taken next: those whose prerequisites are satisfied
		// and that are not already taken.
		available := 0
		for c := 0; c < n; c++ {
			if mask&(1<<c) != 0 {
				continue
			}
			if pre[c]&^mask == 0 { // all prerequisites in mask
				available |= 1 << c
			}
		}

		if available == 0 {
			continue
		}

		// Try all subsets of available of size <= k
		sub := available
		for sub > 0 {
			if bitsCount(sub) <= k {
				next := mask | sub
				if dp[next] > dp[mask]+1 {
					dp[next] = dp[mask] + 1
				}
			}
			sub = (sub - 1) & available
		}
		// Also consider taking 0 courses? No, that's wasteful.
		// Actually the DP needs to consider taking available courses.
		// The loop above handles all non-empty subsets.
	}

	return dp[total-1]
}

func bitsCount(x int) int {
	cnt := 0
	for x > 0 {
		cnt++
		x &= x - 1
	}
	return cnt
}
```

## 1499 — Max Value Of Equation

```go
package main

// LeetCode #1499: Max Value of Equation
// https://leetcode.com/problems/max-value-of-equation/
// Difficulty: Hard
//
// Approach: Monotonic Deque
// For points i < j: yi + yj + |xi - xj| = (yi - xi) + (xj + yj)
// Since xi < xj, we fix j and want max of (yi - xi) for points i where
// xj - xi <= k (i.e., xi >= xj - k).
// Use a deque storing pairs (xi, yi - xi) sorted by (yi - xi) descending.
// Pop front when xi < xj - k (out of window).
// Pop back when new value is larger (keep decreasing order).

import (
	"container/list"
	"fmt"
	"math"
)

func main() {
	// Example 1
	fmt.Println(findMaxValueOfEquation([][]int{{1, 3}, {2, 0}, {5, 10}, {6, -10}}, 1))
	// Expected: 4

	// Example 2
	fmt.Println(findMaxValueOfEquation([][]int{{0, 0}, {3, 0}, {9, 2}}, 3))
	// Expected: 3
}

type pair struct {
	x, diff int // diff = y - x
}

func findMaxValueOfEquation(points [][]int, k int) int {
	ans := math.MinInt32
	dq := list.New()

	for _, p := range points {
		xj, yj := p[0], p[1]

		// Remove points out of window (xi < xj - k)
		for dq.Len() > 0 && dq.Front().Value.(pair).x < xj-k {
			dq.Remove(dq.Front())
		}

		// If deque not empty, front has max (yi - xi)
		if dq.Len() > 0 {
			bestDiff := dq.Front().Value.(pair).diff
			val := bestDiff + xj + yj
			if val > ans {
				ans = val
			}
		}

		// Insert current point, maintain decreasing order
		diff := yj - xj
		for dq.Len() > 0 && dq.Back().Value.(pair).diff <= diff {
			dq.Remove(dq.Back())
		}
		dq.PushBack(pair{x: xj, diff: diff})
	}

	return ans
}
```

## 1501 — Countries You Can Safely Invest In

```go
package main

// LeetCode #1501: Countries You Can Safely Invest In
// https://leetcode.com/problems/countries-you-can-safely-invest-in/
// Difficulty: Medium (listed here as Hard) [Paid]
//
// A country is safe to invest in if the average call duration
// of its residents is longer than the global average call duration.

import "fmt"

// Person represents a person record.
type Person struct {
	ID        int
	Name      string
	CountryID int
}

// Country represents a country record.
type Country struct {
	ID   int
	Name string
}

// Call represents a call record.
type Call struct {
	CallerID int
	CalleeID int
	Duration int // in seconds
}

// SafeCountry holds the result.
type SafeCountry struct {
	Name string
}

// findSafeCountries finds countries whose average call duration
// exceeds the global average.
func findSafeCountries(persons []Person, countries []Country, calls []Call) []SafeCountry {
	// Map country ID to country name
	countryMap := make(map[int]string)
	for _, c := range countries {
		countryMap[c.ID] = c.Name
	}

	// Map person ID to country ID
	personCountry := make(map[int]int)
	for _, p := range persons {
		personCountry[p.ID] = p.CountryID
	}

	// Global totals
	var globalTotalDuration int
	var globalTotalCalls int

	// Per-country totals
	type totals struct {
		duration int
		count    int
	}
	countryTotals := make(map[int]*totals)

	for _, c := range calls {
		globalTotalDuration += c.Duration
		globalTotalCalls++

		// Caller's country
		if countryID, ok := personCountry[c.CallerID]; ok {
			if countryTotals[countryID] == nil {
				countryTotals[countryID] = &totals{}
			}
			countryTotals[countryID].duration += c.Duration
			countryTotals[countryID].count++
		}

		// Callee's country
		if countryID, ok := personCountry[c.CalleeID]; ok {
			if countryTotals[countryID] == nil {
				countryTotals[countryID] = &totals{}
			}
			countryTotals[countryID].duration += c.Duration
			countryTotals[countryID].count++
		}
	}

	if globalTotalCalls == 0 {
		return nil
	}
	globalAvg := float64(globalTotalDuration) / float64(globalTotalCalls)

	var result []SafeCountry
	for countryID, t := range countryTotals {
		if t.count == 0 {
			continue
		}
		countryAvg := float64(t.duration) / float64(t.count)
		if countryAvg > globalAvg {
			name := countryMap[countryID]
			result = append(result, SafeCountry{Name: name})
		}
	}

	return result
}

func main() {
	persons := []Person{
		{1, "Alice", 1},
		{2, "Bob", 2},
		{3, "Charlie", 1},
		{4, "David", 3},
		{5, "Eve", 2},
	}

	countries := []Country{
		{1, "USA"},
		{2, "Canada"},
		{3, "Mexico"},
	}

	calls := []Call{
		{1, 2, 300},  // USA->Canada, 300s
		{1, 3, 200},  // USA->USA, 200s
		{2, 4, 100},  // Canada->Mexico, 100s
		{3, 5, 400},  // USA->Canada, 400s
		{4, 5, 50},   // Mexico->Canada, 50s
	}

	// Global avg = (300+200+100+400+50)/5 = 1050/5 = 210
	// USA: (Alice as caller: 300+200=500, Charlie as caller: 400, Alice as callee: none, Charlie as callee: 200) = 1100/3 = 366.67
	// Wait, let me recalculate:
	// USA calls: caller Alice(300+200=500), callee Charlie(200), caller Charlie(400) = 500+200+400=1100, count=3, avg=366.67 > 210 -> SAFE
	// Canada calls: caller Bob(100+50=150), callee Alice(300), callee Bob(50+100=150)... wait, calls are bi-directional?

	// Actually looking at the problem: A call has a caller and callee. The person's country
	// is involved when they are either the caller or callee. Duration counts for the call.
	// Each call contributes its duration to BOTH the caller's country and the callee's country.

	// Call 1: caller=1(USA,300), callee=2(Canada,300)
	// Call 2: caller=1(USA,200), callee=3(USA,200)
	// Call 3: caller=2(Canada,100), callee=4(Mexico,100)
	// Call 4: caller=3(USA,400), callee=5(Canada,400)
	// Call 5: caller=4(Mexico,50), callee=5(Canada,50)

	// Global: 300+200+100+400+50 = 1050, count=5, avg=210

	// USA: calls 1(300),2(200),4(400) = 900, count=3, avg=300
	// Canada: calls 1(300),3(100),4(400),5(50) = 850, count=4, avg=212.5
	// Mexico: calls 3(100),5(50) = 150, count=2, avg=75

	// USA avg 300 > 210 -> Safe!
	// Canada avg 212.5 > 210 -> Safe!
	// Mexico avg 75 < 210 -> Not safe

	safe := findSafeCountries(persons, countries, calls)
	fmt.Println("Safe countries for investment:")
	for _, c := range safe {
		fmt.Printf("  %s\n", c.Name)
	}

	// Test 2: All below average
	persons2 := []Person{
		{1, "Alice", 1},
		{2, "Bob", 1},
	}
	countries2 := []Country{
		{1, "TestLand"},
	}
	calls2 := []Call{
		{1, 2, 10},
	}
	safe2 := findSafeCountries(persons2, countries2, calls2)
	fmt.Printf("\nTest 2 - All below avg: %d (expected 0)\n", len(safe2))

	// Test 3: No calls
	safe3 := findSafeCountries(persons, countries, nil)
	fmt.Printf("Test 3 - No calls: %v (expected [])\n", safe3)
}
```

## 1505 — Minimum Possible Integer After At Most K Adjacent Swaps On Digits

```go
package main

// LeetCode #1505: Minimum Possible Integer After at Most K Adjacent Swaps On Digits
// https://leetcode.com/problems/minimum-possible-integer-after-at-most-k-adjacent-swaps-on-digits/
// Difficulty: Hard
//
// Approach: Fenwick Tree (BIT) + Greedy
// We process positions left-to-right. For each position, we want the smallest
// possible digit that can be moved here within k swaps.
// Use a Fenwick tree to track how many positions have been removed (shifted left).
// Maintain queues of positions for each digit 0-9.
// For each position i:
//   - Try digits 0-9 in order.
//   - For each digit with available positions, compute the cost to bring it
//     to position i: cost = position - i + (number of removed positions before it).
//   - If cost <= k, take it, update k, mark position as removed, break.

import (
	"fmt"
	"strings"
)

func main() {
	// Example 1
	fmt.Println(minInteger("4321", 4))
	// Expected: "1342"

	// Example 2
	fmt.Println(minInteger("100", 1))
	// Expected: "010"

	// Example 3
	fmt.Println(minInteger("36789", 3))
	// Expected: "36789"
}

type fenwick struct {
	tree []int
	n    int
}

func newFenwick(n int) *fenwick {
	return &fenwick{tree: make([]int, n+2), n: n}
}

func (f *fenwick) add(idx int, delta int) {
	idx++
	for idx <= f.n {
		f.tree[idx] += delta
		idx += idx & -idx
	}
}

func (f *fenwick) sum(idx int) int {
	idx++
	res := 0
	for idx > 0 {
		res += f.tree[idx]
		idx -= idx & -idx
	}
	return res
}

func (f *fenwick) rangeSum(l, r int) int {
	if r < l {
		return 0
	}
	return f.sum(r) - f.sum(l-1)
}

func minInteger(num string, k int) string {
	// queues of positions for each digit
	queues := make([][]int, 10)
	for i, ch := range num {
		d := int(ch - '0')
		queues[d] = append(queues[d], i)
	}
	// pointers for each queue
	ptr := make([]int, 10)

	n := len(num)
	bit := newFenwick(n)
	used := make([]bool, n)

	var sb strings.Builder

	for i := 0; i < n; i++ {
		// Try digits 0-9
		for d := 0; d <= 9; d++ {
			if ptr[d] >= len(queues[d]) {
				continue
			}
			pos := queues[d][ptr[d]]
			// How many positions before pos have been removed?
			removed := bit.rangeSum(0, pos-1)
			cost := pos - removed
			if cost <= k {
				// Take this digit
				k -= cost
				sb.WriteByte(byte('0' + d))
				used[pos] = true
				bit.add(pos, 1)
				ptr[d]++
				break
			}
		}
	}

	return sb.String()
}
```

## 1506 — Find Root Of N Ary Tree

```go
package main

// LeetCode #1506: Find Root of N-Ary Tree
// https://leetcode.com/problems/find-root-of-n-ary-tree/
// Difficulty: Medium (listed here as Hard)
//
// You are given a list of nodes from an N-ary tree. Each node has a value
// and a list of children. Find the root node of the tree.

import "fmt"

// Node represents an N-ary tree node.
type Node struct {
	Val      int
	Children []*Node
}

// findRoot finds the root node from a list of tree nodes.
// The root is the node that never appears as a child.
//
// Approach 1: Use indegree count. The root has indegree 0, all others have indegree >= 1.
// Approach 2: XOR all node values + all child values. The result is the root's value.
//             Then find the node with that value.
//
// We use Approach 1 for clarity and correctness (handles duplicate values).

func findRoot(tree []*Node) *Node {
	if len(tree) == 0 {
		return nil
	}
	if len(tree) == 1 {
		return tree[0]
	}

	// Count indegree: how many times each node appears as a child
	indegree := make(map[*Node]int)
	for _, node := range tree {
		for _, child := range node.Children {
			indegree[child]++
		}
	}

	// The root never appears as a child
	for _, node := range tree {
		if indegree[node] == 0 {
			return node
		}
	}

	return nil // should not happen for a valid tree
}

// findRootByXOR finds the root using XOR of all node values and child values.
// Assumes all node values are unique.
func findRootByXOR(tree []*Node) *Node {
	if len(tree) == 0 {
		return nil
	}
	if len(tree) == 1 {
		return tree[0]
	}

	var xorSum int
	valueToNode := make(map[int]*Node)

	for _, node := range tree {
		valueToNode[node.Val] = node
		xorSum ^= node.Val
		for _, child := range node.Children {
			xorSum ^= child.Val
		}
	}

	// xorSum now equals the root's value (all non-root values cancel out)
	return valueToNode[xorSum]
}

func main() {
	// Build a tree:
	//       1
	//     / | \
	//    2  3  4
	//   /
	//  5
	child5 := &Node{Val: 5}
	child2 := &Node{Val: 2, Children: []*Node{child5}}
	child3 := &Node{Val: 3}
	child4 := &Node{Val: 4}
	root1 := &Node{Val: 1, Children: []*Node{child2, child3, child4}}

	// Shuffle the list (simulates the problem input)
	tree1 := []*Node{child2, root1, child4, child3, child5}

	// Test 1: findRoot
	result1 := findRoot(tree1)
	fmt.Printf("Test 1 - findRoot: Val=%d (expected 1)\n", result1.Val)

	// Test 2: findRootByXOR
	result2 := findRootByXOR(tree1)
	fmt.Printf("Test 2 - findRootByXOR: Val=%d (expected 1)\n", result2.Val)

	// Test 3: Larger tree
	//       10
	//     /    \
	//    20    30
	//   /  \     \
	//  40  50    60
	n40 := &Node{Val: 40}
	n50 := &Node{Val: 50}
	n60 := &Node{Val: 60}
	n20 := &Node{Val: 20, Children: []*Node{n40, n50}}
	n30 := &Node{Val: 30, Children: []*Node{n60}}
	n10 := &Node{Val: 10, Children: []*Node{n20, n30}}

	tree2 := []*Node{n30, n10, n40, n60, n20, n50}
	result3 := findRoot(tree2)
	fmt.Printf("Test 3 - findRoot (larger): Val=%d (expected 10)\n", result3.Val)

	// Test 4: Single node
	tree3 := []*Node{{Val: 42}}
	result4 := findRoot(tree3)
	fmt.Printf("Test 4 - Single node: Val=%d (expected 42)\n", result4.Val)

	// Test 5: Empty
	result5 := findRoot(nil)
	fmt.Printf("Test 5 - Empty: %v (expected nil)\n", result5)

	// Test 6: Two nodes (root and child)
	c := &Node{Val: 2}
	r := &Node{Val: 1, Children: []*Node{c}}
	tree4 := []*Node{c, r}
	result6 := findRoot(tree4)
	fmt.Printf("Test 6 - Two nodes: Val=%d (expected 1)\n", result6.Val)
}
```

## 1510 — Stone Game Iv

```go
package main

// LeetCode #1510: Stone Game IV
// https://leetcode.com/problems/stone-game-iv/
// Difficulty: Hard
//
// Alice and Bob take turns removing stones. In one move, a player can remove
// k*k stones (a perfect square). Alice goes first. Return true if Alice can win
// (playing optimally), false otherwise.

import (
	"fmt"
	"math"
)

// winnerSquareGame returns true if Alice can win the stone game.
func winnerSquareGame(n int) bool {
	// dp[i] = true if the current player can win with i stones remaining
	dp := make([]bool, n+1)

	for i := 1; i <= n; i++ {
		maxSquare := int(math.Sqrt(float64(i)))
		for k := 1; k <= maxSquare; k++ {
			sq := k * k
			// If there's a move that leaves the opponent in a losing position
			if !dp[i-sq] {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}

func main() {
	// Test case 1
	n1 := 1
	result1 := winnerSquareGame(n1)
	fmt.Printf("Test 1: n=%d => %v (expected true, Alice removes 1)\n", n1, result1)

	// Test case 2
	n2 := 2
	result2 := winnerSquareGame(n2)
	fmt.Printf("Test 2: n=%d => %v (expected false, Alice removes 1, Bob removes 1)\n", n2, result2)

	// Test case 3
	n3 := 4
	result3 := winnerSquareGame(n3)
	fmt.Printf("Test 3: n=%d => %v (expected true, Alice removes 4)\n", n3, result3)

	// Test case 4
	n4 := 7
	result4 := winnerSquareGame(n4)
	fmt.Printf("Test 4: n=%d => %v (expected false)\n", n4, result4)

	// Test case 5
	n5 := 17
	result5 := winnerSquareGame(n5)
	fmt.Printf("Test 5: n=%d => %v (expected false)\n", n5, result5)

	// Test case 6: larger number
	n6 := 100
	result6 := winnerSquareGame(n6)
	fmt.Printf("Test 6: n=%d => %v\n", n6, result6)

	// Test case 7
	fmt.Println("\nFirst 20 results:")
	for n := 1; n <= 20; n++ {
		fmt.Printf("n=%d: %v\n", n, winnerSquareGame(n))
	}
}
```

## 1515 — Best Position For A Service Centre

```go
package main

import (
	"fmt"
	"math"
)

// LeetCode #1515: Best Position for a Service Centre
// https://leetcode.com/problems/best-position-for-a-service-centre/
// Difficulty: Hard
//
// Weiszfeld's algorithm (iteratively reweighted) to find the geometric median.
// The geometric median minimizes the sum of Euclidean distances.
// Because the cost function is convex, gradient descent also works reliably.

func getMinDistSum(positions [][]int) float64 {
	n := len(positions)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 0
	}

	// Convert to float64 for precision
	pts := make([][2]float64, n)
	for i, p := range positions {
		pts[i] = [2]float64{float64(p[0]), float64(p[1])}
	}

	// Start at the centroid
	x, y := 0.0, 0.0
	for _, p := range pts {
		x += p[0]
		y += p[1]
	}
	x /= float64(n)
	y /= float64(n)

	// Weiszfeld iteration:
	//   x_{k+1} = (sum_i w_i * xi) / (sum_i w_i)
	//   where w_i = 1 / dist(p_i, current_point)
	// With safeguard for when the current point coincides with a data point.

	best := math.MaxFloat64

	for iter := 0; iter < 10000; iter++ {
		numerX, numerY, denom := 0.0, 0.0, 0.0
		for _, p := range pts {
			d := math.Sqrt((x-p[0])*(x-p[0]) + (y-p[1])*(y-p[1]))
			if d < 1e-12 {
				// We are (almost) exactly on a data point – gradient is undefined.
				// Perturb slightly and let the next iteration handle it.
				continue
			}
			w := 1.0 / d
			numerX += p[0] * w
			numerY += p[1] * w
			denom += w
		}
		if denom > 0 {
			x = numerX / denom
			y = numerY / denom
		}

		// Evaluate current sum of distances
		cur := 0.0
		for _, p := range pts {
			cur += math.Sqrt((x-p[0])*(x-p[0]) + (y-p[1])*(y-p[1]))
		}

		if cur < best {
			best = cur
		}

		// Stop when change is very small
		if iter > 0 && math.Abs(cur-best) < 1e-11 {
			break
		}
	}

	return best
}

func main() {
	// Example 1:
	// Input: positions = [[0,1],[1,0],[1,2],[2,1]]
	// Output: 4.00000
	fmt.Printf("%.5f\n", getMinDistSum([][]int{{0, 1}, {1, 0}, {1, 2}, {2, 1}}))

	// Example 2:
	// Input: positions = [[1,1],[3,3]]
	// Output: 2.82843
	fmt.Printf("%.5f\n", getMinDistSum([][]int{{1, 1}, {3, 3}}))

	// Example 3:
	// Input: positions = [[1,1]]
	// Output: 0.00000
	fmt.Printf("%.5f\n", getMinDistSum([][]int{{1, 1}}))
}
```

## 1516 — Move Sub Tree Of N Ary Tree

```go
package main

import "fmt"

// LeetCode #1516: Move Sub-Tree of N-Ary Tree
// https://leetcode.com/problems/move-sub-tree-of-n-ary-tree/
// Difficulty: Hard [Paid]
//
// Given the root of an N-ary tree, a node p, and a node q,
// make q the new parent of the subtree rooted at p.
// p is not an ancestor of q (guaranteed by the problem).

// Node represents an N-ary tree node.
type Node struct {
	Val      int
	Children []*Node
}

// moveSubTree makes q the new parent of the subtree rooted at p.
// It returns the new root (which may change if the root itself is moved).
func moveSubTree(root, p, q *Node) *Node {
	if root == nil || p == nil || q == nil {
		return root
	}
	if p == root {
		// p is the root – we cannot move root under q unless we detach first.
		// Problem usually guarantees p is not root, but handle gracefully.
		removeChild(root, p)
		q.Children = append(q.Children, p)
		return root
	}

	// Find parent of p
	parentP, _ := findParent(root, p, nil)
	if parentP == nil {
		return root // p not found
	}

	// Remove p from its current parent
	removeChild(parentP, p)

	// Make q the new parent
	q.Children = append(q.Children, p)

	return root
}

// findParent traverses the tree to find parent of target.
func findParent(node, target, parent *Node) (*Node, *Node) {
	if node == nil {
		return nil, nil
	}
	if node == target {
		return parent, node
	}
	for _, child := range node.Children {
		if p, found := findParent(child, target, node); found != nil {
			return p, found
		}
	}
	return nil, nil
}

// removeChild removes child from parent's Children slice.
func removeChild(parent, child *Node) {
	for i, c := range parent.Children {
		if c == child {
			parent.Children = append(parent.Children[:i], parent.Children[i+1:]...)
			return
		}
	}
}

// Helper to build a tree from adjacency list.
// nodes[0] is the root.
func buildTree(adj [][]int) *Node {
	if len(adj) == 0 {
		return nil
	}
	nodes := make([]*Node, len(adj))
	for i := range adj {
		nodes[i] = &Node{Val: i}
	}
	for i, children := range adj {
		for _, c := range children {
			nodes[i].Children = append(nodes[i].Children, nodes[c])
		}
	}
	return nodes[0]
}

// Helper to collect tree values via preorder traversal for verification.
func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	res := []int{root.Val}
	for _, c := range root.Children {
		res = append(res, preorder(c)...)
	}
	return res
}

func main() {
	// Test case:
	// Tree: 0 -> [1, 2], 1 -> [3, 4], 2 -> [5]
	// Root: 0
	// Move subtree of node 1 under node 2 (q=2, p=1)
	// Before: 0 has children [1, 2]; 1 has children [3, 4]; 2 has child [5]
	// After:  0 has child [2]; 2 has children [5, 1]; 1 has children [3, 4]

	adj := [][]int{
		{1, 2}, // 0's children
		{3, 4}, // 1's children
		{5},    // 2's children
		{},     // 3's children
		{},     // 4's children
		{},     // 5's children
	}
	root := buildTree(adj)

	// Find nodes p=1 and q=2
	var p, q *Node
	var findNodes func(*Node)
	findNodes = func(n *Node) {
		if n == nil {
			return
		}
		if n.Val == 1 {
			p = n
		}
		if n.Val == 2 {
			q = n
		}
		for _, c := range n.Children {
			findNodes(c)
		}
	}
	findNodes(root)

	fmt.Println("Before move:")
	fmt.Println("Root preorder:", preorder(root))

	newRoot := moveSubTree(root, p, q)

	fmt.Println("After moving subtree at 1 under 2:")
	fmt.Println("Root preorder:", preorder(newRoot))
	fmt.Println("Expected: [0 2 5 1 3 4]")
}
```

## 1520 — Maximum Number Of Non Overlapping Substrings

```go
package main

// LeetCode #1520: Maximum Number of Non-Overlapping Substrings
// https://leetcode.com/problems/maximum-number-of-non-overlapping-substrings/
// Difficulty: Hard
//
// Approach: Greedy Interval
// 1. For each character, find its first and last occurrence in s.
// 2. For each character, expand its interval until all chars in the interval
//    have their full range within the interval.
// 3. Sort intervals by end ascending. Greedy pick: if start > last_end, take it.

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(maxNumOfSubstrings("adefaddaccc"))
	// Expected: ["e","f","ccc"] or similar valid order

	// Example 2
	fmt.Println(maxNumOfSubstrings("abbaccd"))
	// Expected: ["d","bb","cc"] or similar
}

type iv struct{ l, r int }

func maxNumOfSubstrings(s string) []string {
	n := len(s)

	// First and last occurrence of each char
	first := make([]int, 26)
	last := make([]int, 26)
	for i := 0; i < 26; i++ {
		first[i] = n
		last[i] = -1
	}
	for i, ch := range s {
		c := int(ch - 'a')
		if i < first[c] {
			first[c] = i
		}
		if i > last[c] {
			last[c] = i
		}
	}

	// Compute minimal interval for each character that appears
	minIntervals := make([]iv, 0)
	for c := 0; c < 26; c++ {
		if first[c] == n {
			continue
		}
		l, r := first[c], last[c]
		changed := true
		for changed {
			changed = false
			for j := l; j <= r; j++ {
				ch := int(s[j] - 'a')
				if first[ch] < l {
					l = first[ch]
					changed = true
				}
				if last[ch] > r {
					r = last[ch]
					changed = true
				}
			}
		}
		minIntervals = append(minIntervals, iv{l, r})
	}

	// Dedup by (l,r)
	seen := make(map[int]map[int]bool)
	unique := make([]iv, 0)
	for _, inv := range minIntervals {
		if seen[inv.l] == nil {
			seen[inv.l] = make(map[int]bool)
		}
		if !seen[inv.l][inv.r] {
			seen[inv.l][inv.r] = true
			unique = append(unique, inv)
		}
	}

	// Sort by end ascending
	sort.Slice(unique, func(i, j int) bool {
		return unique[i].r < unique[j].r
	})

	// Greedy pick non-overlapping substrings
	result := make([]string, 0)
	end := -1
	for _, inv := range unique {
		if inv.l > end {
			result = append(result, s[inv.l:inv.r+1])
			end = inv.r
		}
	}

	return result
}
```

## 1521 — Find A Value Of A Mysterious Function Closest To Target

```go
package main

// LeetCode #1521: Find a Value of a Mysterious Function Closest to Target
// https://leetcode.com/problems/find-a-value-of-a-mysterious-function-closest-to-target/
// Difficulty: Hard
//
// Winston has a mysterious function func(arr, l, r) that returns the
// bitwise AND of all elements in arr[l..r]. Find the minimum absolute
// difference between any func value and target.
//
// Approach: Track all possible AND values of subarrays ending at each
// position. AND values only decrease, so the set of distinct values
// is small (at most 32 per position).

import "fmt"

func main() {
	// Example 1
	fmt.Println(closestToTarget([]int{9, 12, 3, 7, 15}, 5))
	// Example 2
	fmt.Println(closestToTarget([]int{1000000, 1000000, 1000000}, 1))
	// Edge: single element
	fmt.Println(closestToTarget([]int{5}, 5))
}

func closestToTarget(arr []int, target int) int {
	ans := abs(arr[0] - target)
	pre := map[int]bool{arr[0]: true}
	for _, x := range arr {
		cur := map[int]bool{x: true}
		for y := range pre {
			cur[x&y] = true
		}
		for y := range cur {
			ans = min(ans, abs(y-target))
		}
		pre = cur
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

## 1526 — Minimum Number Of Increments On Subarrays To Form A Target Array

```go
package main

// LeetCode #1526: Minimum Number of Increments on Subarrays to Form a Target Array
// https://leetcode.com/problems/minimum-number-of-increments-on-subarrays-to-form-a-target-array/
// Difficulty: Hard
//
// Approach: One Pass (Think of diff array)
// Consider the target array. Each time we need to "raise" the level from
// previous position, that's a new operation. The answer is:
//   target[0] + sum over i > 0 of max(0, target[i] - target[i-1])
// Intuition: Imagine we start from all zeros and apply operations. Each operation
// increments a contiguous segment. The difference array gives us the number of
// times we need to start a new segment at each position.

import "fmt"

func main() {
	// Example 1
	fmt.Println(minNumberOperations([]int{1, 2, 3, 2, 1}))
	// Expected: 3

	// Example 2
	fmt.Println(minNumberOperations([]int{3, 1, 1, 2}))
	// Expected: 4

	// Example 3
	fmt.Println(minNumberOperations([]int{3, 1, 5, 4, 2}))
	// Expected: 7
}

func minNumberOperations(target []int) int {
	if len(target) == 0 {
		return 0
	}
	ans := target[0]
	for i := 1; i < len(target); i++ {
		if target[i] > target[i-1] {
			ans += target[i] - target[i-1]
		}
	}
	return ans
}
```

## 1531 — String Compression Ii

```go
package main

// LeetCode #1531: String Compression II
// https://leetcode.com/problems/string-compression-ii/
// Difficulty: Hard
//
// Approach: DP (Top-Down with memoization)
// dp[i][k] = minimum compressed length for s[i:] with at most k deletions.
// At position i with character c = s[i]:
//   - Option 1: delete s[i] (if k > 0): dp[i+1][k-1]
//   - Option 2: keep s[i], count consecutive same chars starting at i.
//     For each run length len from 1 to n-i (stopping when we either run out of
//     deletions or change character), we keep 'len' copies of c and delete the rest
//     within the run. The compressed length for this segment is:
//       1 (for the char) + (digits of len if len > 1)
//     Then add dp[i+len+deleted][k-deleted].

import (
	"fmt"
	"math"
)

func main() {
	// Example 1
	fmt.Println(getLengthOfOptimalCompression("aaabcccd", 2))
	// Expected: 4 ("aa3bccd" -> a2bccd... Let me verify)
	// Actually "aaabcccd" with 2 deletions:
	// delete a at idx 2, delete c at idx 5? Let's trust LeetCode: answer is 4.

	// Example 2
	fmt.Println(getLengthOfOptimalCompression("aabbaa", 2))
	// Expected: 2

	// Example 3
	fmt.Println(getLengthOfOptimalCompression("aaaaaaaaaaa", 0))
	// Expected: 3 ("a11" -> "a11" is 3 chars)
}

func getLengthOfOptimalCompression(s string, k int) int {
	n := len(s)
	// memo[i][k]
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, k+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dp func(i, k int) int
	dp = func(i, k int) int {
		if i == n || k >= n-i {
			return 0
		}
		if memo[i][k] != -1 {
			return memo[i][k]
		}

		best := math.MaxInt32

		// Option 1: delete s[i]
		if k > 0 {
			best = minInt(best, dp(i+1, k-1))
		}

		// Option 2: keep s[i]
		c := s[i]
		sameCount := 0
		deleted := 0
		for j := i; j < n && deleted <= k; j++ {
			if s[j] == c {
				sameCount++
			} else {
				deleted++
			}
			if deleted > k {
				break
			}
			// Compressed length for sameCount copies of c
			compLen := compLength(sameCount)
			best = minInt(best, compLen+dp(j+1, k-deleted))
		}

		memo[i][k] = best
		return best
	}

	return dp(0, k)
}

func compLength(cnt int) int {
	// compressed = 1 char + digit length if cnt > 1
	if cnt == 1 {
		return 1
	}
	if cnt < 10 {
		return 2
	}
	if cnt < 100 {
		return 3
	}
	return 4
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

## 1532 — The Most Recent Three Orders

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1532: The Most Recent Three Orders
// https://leetcode.com/problems/the-most-recent-three-orders/
// Difficulty: Hard
//
// For each customer, find the most recent 3 orders.
// Implementation simulates the SQL query using Go data structures.

// Order represents a customer order.
type Order struct {
	OrderID    int
	OrderDate  string
	CustomerID int
}

// Customer represents a customer.
type Customer struct {
	CustomerID int
	Name       string
}

type resultRow1532 struct {
	CustomerName string
	OrderID      int
	OrderDate    string
}

// mostRecentThreeOrders returns for each customer their 3 most recent orders,
// ordered by customer name ascending, then order date descending, then order ID descending.
func mostRecentThreeOrders(customers []Customer, orders []Order) []resultRow1532 {
	// Group orders by customer ID
	ordersByCustomer := make(map[int][]Order)
	for _, o := range orders {
		ordersByCustomer[o.CustomerID] = append(ordersByCustomer[o.CustomerID], o)
	}

	// For each customer, sort orders by date descending, then order ID descending
	customerMap := make(map[int]string)
	for _, c := range customers {
		customerMap[c.CustomerID] = c.Name
	}

	var results []resultRow1532

	for cid, ords := range ordersByCustomer {
		sort.Slice(ords, func(i, j int) bool {
			if ords[i].OrderDate != ords[j].OrderDate {
				return ords[i].OrderDate > ords[j].OrderDate
			}
			return ords[i].OrderID > ords[j].OrderID
		})

		// Take top 3
		limit := 3
		if len(ords) < limit {
			limit = len(ords)
		}

		for _, o := range ords[:limit] {
			results = append(results, resultRow1532{
				CustomerName: customerMap[cid],
				OrderID:      o.OrderID,
				OrderDate:    o.OrderDate,
			})
		}
	}

	// Sort results by customer name ascending, order date descending, order ID descending
	sort.Slice(results, func(i, j int) bool {
		if results[i].CustomerName != results[j].CustomerName {
			return results[i].CustomerName < results[j].CustomerName
		}
		if results[i].OrderDate != results[j].OrderDate {
			return results[i].OrderDate > results[j].OrderDate
		}
		return results[i].OrderID > results[j].OrderID
	})

	return results
}

func main() {
	customers := []Customer{
		{1, "Alice"},
		{2, "Bob"},
	}
	orders := []Order{
		{101, "2023-01-01", 1},
		{102, "2023-01-02", 1},
		{103, "2023-01-03", 1},
		{104, "2023-01-04", 1},
		{201, "2023-01-01", 2},
		{202, "2023-01-02", 2},
	}

	results := mostRecentThreeOrders(customers, orders)
	for _, r := range results {
		fmt.Printf("%s | %d | %s\n", r.CustomerName, r.OrderID, r.OrderDate)
	}
	// Expected:
	// Alice | 104 | 2023-01-04
	// Alice | 103 | 2023-01-03
	// Alice | 102 | 2023-01-02
	// Bob   | 202 | 2023-01-02
	// Bob   | 201 | 2023-01-01
}
```

## 1537 — Get The Maximum Score

```go
package main

// LeetCode #1537: Get the Maximum Score
// https://leetcode.com/problems/get-the-maximum-score/
// Difficulty: Hard
//
// Two-pointer + cumulative sum approach:
// 1. Traverse both sorted arrays with two pointers.
// 2. Keep running sums for each path.
// 3. When values equal, we can switch paths; take max of both sums.
// 4. Continue accumulating after the switch point.

import "fmt"

func main() {
	// Example: [2,4,5,8,10], [4,6,8,9] -> 30
	// Path: 2+4+6+8+10 = 30 (switch at 4, switch at 8)
	fmt.Println(maxSum([]int{2, 4, 5, 8, 10}, []int{4, 6, 8, 9}))

	// Additional tests
	fmt.Println(maxSum([]int{1, 3, 5, 7, 9}, []int{2, 4, 6, 8, 10}))
	fmt.Println(maxSum([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}))
	fmt.Println(maxSum([]int{1, 4, 5, 8, 12}, []int{2, 3, 6, 10, 11}))
	fmt.Println(maxSum([]int{1}, []int{1}))
}

const mod = 1_000_000_007

func maxSum(nums1 []int, nums2 []int) int {
	i, j := 0, 0
	sum1, sum2 := 0, 0
	n, m := len(nums1), len(nums2)

	for i < n && j < m {
		if nums1[i] < nums2[j] {
			sum1 += nums1[i]
			i++
		} else if nums1[i] > nums2[j] {
			sum2 += nums2[j]
			j++
		} else {
			// Equal — switch point: take max of both paths
			best := max(sum1, sum2) + nums1[i]
			sum1 = best
			sum2 = best
			i++
			j++
		}
	}

	for i < n {
		sum1 += nums1[i]
		i++
	}
	for j < m {
		sum2 += nums2[j]
		j++
	}

	return max(sum1, sum2) % mod
}
```

## 1542 — Find Longest Awesome Substring

```go
package main

// LeetCode #1542: Find Longest Awesome Substring
// https://leetcode.com/problems/find-longest-awesome-substring/
// Difficulty: Hard
//
// Bitmask prefix approach:
// - A substring is "awesome" if at most one digit has odd frequency.
// - Use a 10-bit mask where bit k = parity of digit k's count.
// - For each prefix position, store the first occurrence of each mask.
// - For each mask, we look for the same mask (all even) or
//   a mask differing by exactly one bit (one digit odd).

import "fmt"

func main() {
	// Example: "3242415" -> 5 ("24241" or "42415")
	fmt.Println(longestAwesome("3242415"))

	// Additional tests
	fmt.Println(longestAwesome("0"))
	fmt.Println(longestAwesome("00"))
	fmt.Println(longestAwesome("123456789"))
	fmt.Println(longestAwesome("373781"))
}

func longestAwesome(s string) int {
	// first[mask] = earliest index where this mask first appeared
	first := make(map[int]int)
	first[0] = -1 // empty prefix has mask 0

	mask := 0
	maxLen := 1

	for i, ch := range s {
		bit := 1 << (ch - '0')
		mask ^= bit

		// Check same mask (all even frequencies)
		if idx, ok := first[mask]; ok {
			maxLen = max(maxLen, i-idx)
		}

		// Check masks differing by one bit (one odd frequency)
		for d := 0; d < 10; d++ {
			neighbor := mask ^ (1 << d)
			if idx, ok := first[neighbor]; ok {
				maxLen = max(maxLen, i-idx)
			}
		}

		// Store first occurrence of this mask
		if _, ok := first[mask]; !ok {
			first[mask] = i
		}
	}

	return maxLen
}
```

## 1547 — Minimum Cost To Cut A Stick

```go
package main

// LeetCode #1547: Minimum Cost to Cut a Stick
// https://leetcode.com/problems/minimum-cost-to-cut-a-stick/
// Difficulty: Hard
//
// DP interval approach:
// 1. Add 0 and n to the cuts array, sort.
// 2. DP[i][j] = minimum cost to cut stick from cuts[i] to cuts[j].
// 3. For each interval [i,j], try every cut point k in (i,j).
// 4. DP[i][j] = min(DP[i][k] + DP[k][j] + (cuts[j]-cuts[i]))
// 5. Return DP[0][len(cuts)-1].

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// Example: n=7, cuts=[1,3,4,5] -> 16
	fmt.Println(minCost(7, []int{1, 3, 4, 5}))

	// Additional tests
	fmt.Println(minCost(9, []int{5, 6, 1, 4, 2}))
	fmt.Println(minCost(10, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	fmt.Println(minCost(3, []int{1}))
	fmt.Println(minCost(4, []int{2}))
}

func minCost(n int, cuts []int) int {
	// Add boundaries and sort
	extended := make([]int, 0, len(cuts)+2)
	extended = append(extended, 0)
	extended = append(extended, cuts...)
	extended = append(extended, n)
	sort.Ints(extended)

	m := len(extended)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, m)
	}

	// Interval DP: process by increasing length
	for length := 2; length < m; length++ {
		for i := 0; i+length < m; i++ {
			j := i + length
			dp[i][j] = math.MaxInt32
			for k := i + 1; k < j; k++ {
				cost := dp[i][k] + dp[k][j] + (extended[j] - extended[i])
				if cost < dp[i][j] {
					dp[i][j] = cost
				}
			}
		}
	}

	return dp[0][m-1]
}
```

## 1548 — The Most Similar Path In A Graph

```go
package main

// LeetCode #1548: The Most Similar Path in a Graph
// https://leetcode.com/problems/the-most-similar-path-in-a-graph/
// Difficulty: Hard [Paid]
//
// Given n cities connected by roads, each with a name, find a path of
// length equal to targetPath that minimizes the edit distance (number
// of mismatches) between city names and targetPath names.
//
// Approach: DP[i][v] = min edit distance for first i steps ending at city v.
// Reconstruct path by backtracking through DP.

import "fmt"

func main() {
	// Example 1
	fmt.Println(mostSimilar(5, [][]int{{0, 2}, {0, 3}, {1, 2}, {1, 3}, {1, 4}, {2, 4}},
		[]string{"ATL", "PEK", "LAX", "DXB", "HND"},
		[]string{"ATL", "DXB", "HND", "LAX"}))
	// Example 2
	fmt.Println(mostSimilar(4, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}},
		[]string{"BOM", "BOM", "MAA", "BOM"},
		[]string{"BOM", "MAA", "BOM"}))
	// Edge: single node
	fmt.Println(mostSimilar(1, [][]int{},
		[]string{"A"},
		[]string{"A", "B", "A"}))
}

func mostSimilar(n int, roads [][]int, names []string, targetPath []string) []int {
	m := len(targetPath)
	adj := make([][]int, n)
	for _, r := range roads {
		u, v := r[0], r[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// If no edges, handle separately
	if n == 1 {
		path := make([]int, m)
		for i := range path {
			path[i] = 0
		}
		return path
	}

	// dp[i][v] = min edit distance for first i steps (0-indexed) ending at v
	dp := make([][]int, m)
	prev := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		prev[i] = make([]int, n)
		for v := 0; v < n; v++ {
			cost := 0
			if names[v] != targetPath[i] {
				cost = 1
			}
			if i == 0 {
				dp[i][v] = cost
				prev[i][v] = -1
			} else {
				best := m + 1
				bestPrev := -1
				for _, u := range adj[v] {
					if dp[i-1][u] < best {
						best = dp[i-1][u]
						bestPrev = u
					}
				}
				dp[i][v] = best + cost
				prev[i][v] = bestPrev
			}
		}
	}

	// Find best ending city
	end := 0
	best := dp[m-1][0]
	for v := 1; v < n; v++ {
		if dp[m-1][v] < best {
			best = dp[m-1][v]
			end = v
		}
	}

	// Reconstruct path
	path := make([]int, m)
	path[m-1] = end
	for i := m - 1; i > 0; i-- {
		path[i-1] = prev[i][path[i]]
	}
	return path
}
```

## 1549 — The Most Recent Orders For Each Product

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1549: The Most Recent Orders for Each Product
// https://leetcode.com/problems/the-most-recent-orders-for-each-product/
// Difficulty: Hard
//
// For each product, find the most recent order(s) -- if multiple orders
// share the most recent date, include all of them.

// Order represents a customer order line.
type Order struct {
	OrderID   int
	OrderDate string
	ProductID int
	Quantity  int
}

// Product represents a product.
type Product struct {
	ProductID   int
	ProductName string
}

type resultRow1549 struct {
	ProductName string
	ProductID   int
	OrderID     int
	OrderDate   string
}

// mostRecentOrdersForEachProduct returns for each product its most recent order(s).
func mostRecentOrdersForEachProduct(products []Product, orders []Order) []resultRow1549 {
	// Group orders by product ID
	ordersByProduct := make(map[int][]Order)
	for _, o := range orders {
		ordersByProduct[o.ProductID] = append(ordersByProduct[o.ProductID], o)
	}

	productMap := make(map[int]string)
	for _, p := range products {
		productMap[p.ProductID] = p.ProductName
	}

	var results []resultRow1549

	for pid, ords := range ordersByProduct {
		// Find the most recent date
		maxDate := ""
		for _, o := range ords {
			if o.OrderDate > maxDate {
				maxDate = o.OrderDate
			}
		}

		// Collect all orders with that date
		for _, o := range ords {
			if o.OrderDate == maxDate {
				results = append(results, resultRow1549{
					ProductName: productMap[pid],
					ProductID:   pid,
					OrderID:     o.OrderID,
					OrderDate:   o.OrderDate,
				})
			}
		}
	}

	// Sort results by product name ascending, order ID ascending
	sort.Slice(results, func(i, j int) bool {
		if results[i].ProductName != results[j].ProductName {
			return results[i].ProductName < results[j].ProductName
		}
		return results[i].OrderID < results[j].OrderID
	})

	return results
}

func main() {
	products := []Product{
		{1, "Widget"},
		{2, "Gadget"},
	}
	orders := []Order{
		{101, "2023-01-01", 1, 5},
		{102, "2023-01-02", 1, 3},
		{103, "2023-01-02", 2, 2},
		{104, "2023-01-03", 2, 1},
	}

	results := mostRecentOrdersForEachProduct(products, orders)
	for _, r := range results {
		fmt.Printf("%s (ID %d) | Order %d | %s\n", r.ProductName, r.ProductID, r.OrderID, r.OrderDate)
	}
	// Expected:
	// Gadget (ID 2) | Order 104 | 2023-01-03
	// Widget (ID 1) | Order 102 | 2023-01-02
}
```

## 1553 — Minimum Number Of Days To Eat N Oranges

```go
package main

// LeetCode #1553: Minimum Number of Days to Eat N Oranges
// https://leetcode.com/problems/minimum-number-of-days-to-eat-n-oranges/
// Difficulty: Hard
//
// Memoized recursion:
// - If n is divisible by 2, we can eat n/2 oranges in one day (after eating n%2 oranges one by one).
// - If n is divisible by 3, we can eat 2*n/3 oranges in one day (after eating n%3 oranges one by one).
// - Otherwise, eat 1 orange.
// - Use memoization to avoid recomputation.
// - Since n can be up to 2*10^9, we use a map for memoization.

import (
	"fmt"
)

func main() {
	// Example: 10 -> 4; 6 -> 3
	fmt.Println(minDays(10))
	fmt.Println(minDays(6))

	// Additional tests
	fmt.Println(minDays(1))
	fmt.Println(minDays(2))
	fmt.Println(minDays(3))
	fmt.Println(minDays(50))
	fmt.Println(minDays(100))
}

func minDays(n int) int {
	memo := make(map[int]int)
	return dp(n, memo)
}

func dp(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}
	if val, ok := memo[n]; ok {
		return val
	}

	// Option 1: eat one at a time to make n divisible by 2, then eat half
	option1 := n%2 + 1 + dp(n/2, memo)
	// Option 2: eat one at a time to make n divisible by 3, then eat 2/3
	option2 := n%3 + 1 + dp(n/3, memo)

	result := min(option1, option2)
	memo[n] = result
	return result
}
```

## 1555 — Bank Account Summary

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1555: Bank Account Summary
// https://leetcode.com/problems/bank-account-summary/
// Difficulty: Hard
//
// Given users, transactions, and transfers, compute the net balance for each user
// and classify them as "Low Salary", "Average Salary", or "High Salary"
// based on their final balance (credit - debit + paid_by_incoming_transfers - paid_by_outgoing_transfers).
//
// Classification rules:
// - Low Salary:  balance < 20000
// - Average Salary: 20000 <= balance <= 50000
// - High Salary: balance > 50000

// User represents a bank user.
type User struct {
	UserID int
	Name   string
}

// Transaction is a credit/debit operation (credit means money IN to the user).
type Transaction struct {
	TransactionID int
	PaidBy        int
	PaidTo        int
	Amount        int
	TransactedOn  string
}

type resultRow1555 struct {
	Name    string
	Balance int
	Class   string
}

// accountSummary computes each user's balance and salary classification.
func accountSummary(users []User, transactions []Transaction) []resultRow1555 {
	balances := make(map[int]int) // userID -> net balance

	// Initialize all users with 0
	for _, u := range users {
		balances[u.UserID] = 0
	}

	// Process transactions: paid_by loses money, paid_to gains money
	for _, t := range transactions {
		balances[t.PaidBy] -= t.Amount
		balances[t.PaidTo] += t.Amount
	}

	var results []resultRow1555

	for _, u := range users {
		bal := balances[u.UserID]
		class := ""
		if bal < 20000 {
			class = "Low Salary"
		} else if bal <= 50000 {
			class = "Average Salary"
		} else {
			class = "High Salary"
		}
		results = append(results, resultRow1555{
			Name:    u.Name,
			Balance: bal,
			Class:   class,
		})
	}

	// Sort by name ascending
	sort.Slice(results, func(i, j int) bool {
		return results[i].Name < results[j].Name
	})

	return results
}

func main() {
	users := []User{
		{1, "Alice"},
		{2, "Bob"},
		{3, "Charlie"},
	}
	transactions := []Transaction{
		{1, 1, 2, 10000, "2023-01-01"},
		{2, 2, 3, 5000, "2023-01-02"},
		{3, 3, 1, 30000, "2023-01-03"},
		{4, 1, 2, 40000, "2023-01-04"},
	}

	results := accountSummary(users, transactions)
	for _, r := range results {
		fmt.Printf("%s | Balance: %d | %s\n", r.Name, r.Balance, r.Class)
	}
	// Expected:
	// Alice | Balance: -20000 | Low Salary (paid 10000+40000=50000, received 30000, net -20000)
	// Bob   | Balance: 45000  | Average Salary (paid 5000, received 10000+40000=50000, net 45000)
	// Charlie | Balance: -25000 | Low Salary (paid 30000, received 5000, net -25000)
}
```

## 1563 — Stone Game V

```go
package main

// LeetCode #1563: Stone Game V
// https://leetcode.com/problems/stone-game-v/
// Difficulty: Hard
//
// DP interval approach:
// 1. Compute prefix sums for O(1) range sum queries.
// 2. DP[i][j] = maximum score Alice can get from stones[i..j].
// 3. For each split k in [i, j-1]:
//    - leftSum = sum[i..k], rightSum = sum[k+1..j]
//    - If leftSum < rightSum: score = DP[i][k] + leftSum
//    - If leftSum > rightSum: score = DP[k+1][j] + rightSum
//    - If equal: score = max(DP[i][k], DP[k+1][j]) + leftSum
// 4. Take max over all splits.

import (
	"fmt"
)

func main() {
	// Example: [6,2,3,4,5,5] -> 18
	fmt.Println(stoneGameV([]int{6, 2, 3, 4, 5, 5}))

	// Additional tests
	fmt.Println(stoneGameV([]int{7, 7, 7, 7, 7, 7, 7}))
	fmt.Println(stoneGameV([]int{1}))
	fmt.Println(stoneGameV([]int{1, 2}))
	fmt.Println(stoneGameV([]int{3, 3, 3}))
}

func stoneGameV(stoneValue []int) int {
	n := len(stoneValue)
	if n == 1 {
		return 0
	}

	// Prefix sums
	prefix := make([]int, n+1)
	for i, v := range stoneValue {
		prefix[i+1] = prefix[i] + v
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			best := 0
			for k := i; k < j; k++ {
				leftSum := prefix[k+1] - prefix[i]
				rightSum := prefix[j+1] - prefix[k+1]

				var score int
				if leftSum < rightSum {
					score = dp[i][k] + leftSum
				} else if leftSum > rightSum {
					score = dp[k+1][j] + rightSum
				} else {
					score = max(dp[i][k], dp[k+1][j]) + leftSum
				}

				if score > best {
					best = score
				}
			}
			dp[i][j] = best
		}
	}

	return dp[0][n-1]
}
```

## 1568 — Minimum Number Of Days To Disconnect Island

```go
package main

// LeetCode #1568: Minimum Number of Days to Disconnect Island
// https://leetcode.com/problems/minimum-number-of-days-to-disconnect-island/
// Difficulty: Hard
//
// Approach:
// 1. Count islands. If not exactly 1, return 0 (already disconnected).
// 2. Try flipping each land cell to water. If it disconnects the island, return 1.
// 3. Otherwise, return 2 (worst case: remove any corner of an island).
//
// Key insight: answer is always 0, 1, or 2. If removing one cell doesn't work,
// removing two cells always works (e.g., remove the two ends of a bridge).

import "fmt"

func main() {
	// Example: grid = [[0,1,1,0],[0,1,1,0],[0,0,0,0]] -> 2
	fmt.Println(minDays([][]int{
		{0, 1, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0},
	}))

	// Example: single cell island -> 1
	fmt.Println(minDays([][]int{{1}}))

	// Example: already disconnected -> 0
	fmt.Println(minDays([][]int{
		{1, 1},
		{1, 0},
	}))

	// Example: 2x2 all land -> 2
	fmt.Println(minDays([][]int{
		{1, 1},
		{1, 1},
	}))

	// Example: line of 3 -> 1 (remove middle)
	fmt.Println(minDays([][]int{{1, 1, 1}}))
}

var dirs = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func minDays(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// Count islands
	if countIslands(grid) != 1 {
		return 0
	}

	// Try removing each land cell
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				grid[i][j] = 0
				if countIslands(grid) != 1 {
					return 1
				}
				grid[i][j] = 1
			}
		}
	}

	return 2
}

func countIslands(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				count++
				dfs(grid, visited, i, j)
			}
		}
	}
	return count
}

func dfs(grid [][]int, visited [][]bool, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return
	}
	if grid[i][j] == 0 || visited[i][j] {
		return
	}
	visited[i][j] = true
	for _, d := range dirs {
		dfs(grid, visited, i+d[0], j+d[1])
	}
}
```

## 1569 — Number Of Ways To Reorder Array To Get Same Bst

```go
package main

// LeetCode #1569: Number of Ways to Reorder Array to Get Same BST
// https://leetcode.com/problems/number-of-ways-to-reorder-array-to-get-same-bst/
// Difficulty: Hard
//
// Combinatorics + recursion:
// 1. The first element is the root.
// 2. Split remaining elements into left (< root) and right (> root).
// 3. The relative order within left and right must be preserved for BST
//    insertion, but we can interleave them arbitrarily.
// 4. Number of ways = C(len(left)+len(right), len(left)) *
//    ways(left) * ways(right) mod M.
// 5. Use Pascal's triangle or precomputed nCr for efficiency.

import (
	"fmt"
)

func main() {
	// Example: [2,1,3] -> 1 (only [2,1,3] works)
	fmt.Println(numOfWays([]int{2, 1, 3}))

	// Additional tests
	fmt.Println(numOfWays([]int{1, 2, 3}))
	fmt.Println(numOfWays([]int{3, 1, 2, 4}))
	fmt.Println(numOfWays([]int{3, 4, 5, 1, 2}))
	fmt.Println(numOfWays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}

const MOD = 1_000_000_007

func numOfWays(nums []int) int {
	n := len(nums)
	// Precompute nCr using Pascal's triangle
	comb := make([][]int, n+1)
	for i := range comb {
		comb[i] = make([]int, n+1)
		comb[i][0] = 1
		comb[i][i] = 1
		for j := 1; j < i; j++ {
			comb[i][j] = (comb[i-1][j-1] + comb[i-1][j]) % MOD
		}
	}

	// Result minus 1 (exclude original order)
	return (countWays(nums, comb) - 1 + MOD) % MOD
}

func countWays(nums []int, comb [][]int) int {
	if len(nums) <= 1 {
		return 1
	}

	root := nums[0]
	left := make([]int, 0)
	right := make([]int, 0)

	for i := 1; i < len(nums); i++ {
		if nums[i] < root {
			left = append(left, nums[i])
		} else {
			right = append(right, nums[i])
		}
	}

	// C(len(left)+len(right), len(left))
	ways := comb[len(left)+len(right)][len(left)]
	ways = (ways * countWays(left, comb)) % MOD
	ways = (ways * countWays(right, comb)) % MOD

	return ways
}
```

## 1575 — Count All Possible Routes

```go
package main

// LeetCode #1575: Count All Possible Routes
// https://leetcode.com/problems/count-all-possible-routes/
// Difficulty: Hard
//
// DP[fuel][city] approach:
// - dp[f][i] = number of ways to reach city i with exactly f fuel remaining.
// - Base: dp[fuel][start] = 1 (we start there with full fuel).
// - Transition: for each city i, for each city j != i, if fuel >= |loc[i]-loc[j]|,
//   dp[f-fuelUsed][j] += dp[f][i].
// - Result: sum of dp[any fuel][finish].
//
// Since fuel up to 200 and cities up to 100, O(fuel * n^2) works.

import "fmt"

func main() {
	// Example: locations=[2,3,6,8,4], start=1, finish=3, fuel=5 -> 4
	fmt.Println(countRoutes([]int{2, 3, 6, 8, 4}, 1, 3, 5))

	// Additional tests
	fmt.Println(countRoutes([]int{1, 2, 3}, 0, 2, 3))
	fmt.Println(countRoutes([]int{1, 2, 3}, 0, 2, 1))
	fmt.Println(countRoutes([]int{5, 2, 1}, 0, 2, 3))
	fmt.Println(countRoutes([]int{2, 3, 6, 8, 4}, 1, 3, 3))
}

const MODR = 1_000_000_007

func countRoutes(locations []int, start int, finish int, fuel int) int {
	n := len(locations)
	// dp[f][i] = number of ways to reach city i with exactly f fuel
	dp := make([][]int, fuel+1)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[fuel][start] = 1

	result := 0
	if start == finish {
		result = 1
	}

	for f := fuel; f >= 0; f-- {
		for i := 0; i < n; i++ {
			if dp[f][i] == 0 {
				continue
			}
			if i == finish && f != fuel {
				result = (result + dp[f][i]) % MODR
			}
			for j := 0; j < n; j++ {
				if i == j {
					continue
				}
				cost := abs(locations[i] - locations[j])
				if f >= cost {
					dp[f-cost][j] = (dp[f-cost][j] + dp[f][i]) % MODR
				}
			}
		}
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

## 1579 — Remove Max Number Of Edges To Keep Graph Fully Traversable

```go
package main

// LeetCode #1579: Remove Max Number of Edges to Keep Graph Fully Traversable
// https://leetcode.com/problems/remove-max-number-of-edges-to-keep-graph-fully-traversable/
// Difficulty: Hard
//
// Union-Find approach:
// - Process type-3 (both) edges first: they benefit both Alice and Bob.
// - Then process type-1 (Alice) and type-2 (Bob) edges separately.
// - Count how many edges are actually used. Total edges - used = answer.
// - If either Alice or Bob can't fully traverse, return -1.

import "fmt"

func main() {
	// Example: n=4, edges=[[3,1,2],[3,2,3],[1,1,3],[1,2,4],[1,1,2],[2,3,4]] -> 2
	fmt.Println(maxNumEdgesToRemove(4, [][]int{
		{3, 1, 2},
		{3, 2, 3},
		{1, 1, 3},
		{1, 2, 4},
		{1, 1, 2},
		{2, 3, 4},
	}))

	// Additional tests
	fmt.Println(maxNumEdgesToRemove(4, [][]int{
		{3, 1, 2},
		{3, 2, 3},
		{1, 1, 4},
		{2, 1, 4},
	}))

	fmt.Println(maxNumEdgesToRemove(2, [][]int{
		{1, 1, 2},
		{2, 1, 2},
		{3, 1, 2},
	}))

	fmt.Println(maxNumEdgesToRemove(4, [][]int{
		{3, 1, 2},
		{3, 3, 4},
		{1, 1, 3},
		{2, 2, 4},
	}))

	fmt.Println(maxNumEdgesToRemove(5, [][]int{
		{1, 1, 2},
		{2, 2, 3},
	}))
}

type unionFind struct {
	parent []int
	rank   []int
	count  int
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n+1)
	rank := make([]int, n+1)
	for i := 1; i <= n; i++ {
		parent[i] = i
	}
	return &unionFind{parent, rank, n}
}

func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *unionFind) union(x, y int) bool {
	rx, ry := uf.find(x), uf.find(y)
	if rx == ry {
		return false
	}
	if uf.rank[rx] < uf.rank[ry] {
		rx, ry = ry, rx
	}
	uf.parent[ry] = rx
	if uf.rank[rx] == uf.rank[ry] {
		uf.rank[rx]++
	}
	uf.count--
	return true
}

func maxNumEdgesToRemove(n int, edges [][]int) int {
	alice := newUnionFind(n)
	bob := newUnionFind(n)
	used := 0

	// Process type-3 edges first
	for _, e := range edges {
		if e[0] == 3 {
			connectedA := alice.union(e[1], e[2])
			connectedB := bob.union(e[1], e[2])
			if connectedA || connectedB {
				used++
			}
		}
	}

	// Process type-1 (Alice only)
	for _, e := range edges {
		if e[0] == 1 {
			if alice.union(e[1], e[2]) {
				used++
			}
		}
	}

	// Process type-2 (Bob only)
	for _, e := range edges {
		if e[0] == 2 {
			if bob.union(e[1], e[2]) {
				used++
			}
		}
	}

	// Check if both are fully traversable
	if alice.count != 1 || bob.count != 1 {
		return -1
	}

	return len(edges) - used
}
```

## 1585 — Check If String Is Transformable With Substring Sort Operations

```go
package main

import (
	"fmt"
)

// LeetCode #1585: Check If String Is Transformable With Substring Sort Operations
// https://leetcode.com/problems/check-if-string-is-transformable-with-substring-sort-operations/
// Difficulty: Hard
//
// Key insight: Sorting a substring in ascending order moves smaller digits left
// and larger digits right. Therefore a digit can only move right (not left) in
// the string relative to smaller digits.
//
// Algorithm:
// - Maintain queues of positions for each digit (0-9) in the source string s.
// - Scan t left-to-right. For each digit d:
//   1. Pop the earliest available position pos in s for digit d.
//   2. For every smaller digit sd < d, check that ALL remaining positions
//      of sd in s are AFTER pos. If any remaining sd position is before pos,
//      that smaller digit would block d's movement (since smaller digits move
//      left and would overtake d if they start before it).
// - If all checks pass, return true.

func isTransformable(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// Queue of positions for each digit in source string s
	pos := make([][]int, 10)
	for i, ch := range s {
		d := ch - '0'
		pos[d] = append(pos[d], i)
	}

	// Also track counts for quick character validation
	countS := [10]int{}
	countT := [10]int{}
	for _, ch := range s {
		countS[ch-'0']++
	}
	for _, ch := range t {
		countT[ch-'0']++
	}
	if countS != countT {
		return false
	}

	// Scan target string left to right
	for _, ch := range t {
		d := ch - '0'

		// Pop the first (earliest) occurrence of this digit in s
		if len(pos[d]) == 0 {
			return false
		}
		j := pos[d][0]
		pos[d] = pos[d][1:]

		// Check all smaller digits: any remaining position before j is a blocker.
		// A smaller digit before d in the source would, when sorted, end up before d,
		// making it impossible for d to be at the current position.
		for sd := 0; sd < int(d); sd++ {
			if len(pos[sd]) > 0 && pos[sd][0] < j {
				return false
			}
		}
	}

	return true
}

func main() {
	// Example 1:
	// Input: s = "84532", t = "34852"
	// Output: true
	fmt.Println(isTransformable("84532", "34852")) // true

	// Example 2:
	// Input: s = "34521", t = "23415"
	// Output: true
	fmt.Println(isTransformable("34521", "23415")) // true

	// Example 3:
	// Input: s = "12345", t = "12435"
	// Output: false — '4' cannot move right past '5' (4 < 5, and 5 is after 4)
	fmt.Println(isTransformable("12345", "12435")) // false

	// Additional tests:
	// Same strings
	fmt.Println(isTransformable("12345", "12345")) // true

	// "21" -> "12" works (sort entire substring)
	fmt.Println(isTransformable("21", "12")) // true

	// "12" -> "21" fails (smaller can't move right, larger can't move left)
	fmt.Println(isTransformable("12", "21")) // false

	// Length mismatch
	fmt.Println(isTransformable("123", "1234")) // false
}
```

## 1586 — Binary Search Tree Iterator Ii

```go
package main

import "fmt"

// LeetCode #1586: Binary Search Tree Iterator II
// https://leetcode.com/problems/binary-search-tree-iterator-ii/
// Difficulty: Hard
//
// BSTIteratorII provides bidirectional traversal of a BST:
// - hasNext() / next()   — forward
// - hasPrev() / prev()   — backward
//
// Approach: Perform an inorder traversal to collect all node values,
// then use a cursor to navigate. This gives O(1) for all operations
// at the cost of O(n) memory.

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BSTIteratorII is the bidirectional BST iterator.
type BSTIteratorII struct {
	vals    []int
	cursor  int // index of the last returned value; -1 means before start
}

// Constructor initializes the iterator with the root of a BST.
func Constructor(root *TreeNode) *BSTIteratorII {
	vals := make([]int, 0)
	inorder(root, &vals)
	return &BSTIteratorII{vals: vals, cursor: -1}
}

func inorder(node *TreeNode, vals *[]int) {
	if node == nil {
		return
	}
	inorder(node.Left, vals)
	*vals = append(*vals, node.Val)
	inorder(node.Right, vals)
}

// HasNext returns true if there is a next element.
func (it *BSTIteratorII) HasNext() bool {
	return it.cursor+1 < len(it.vals)
}

// Next returns the next element. It advances the cursor.
func (it *BSTIteratorII) Next() int {
	it.cursor++
	return it.vals[it.cursor]
}

// HasPrev returns true if there is a previous element.
func (it *BSTIteratorII) HasPrev() bool {
	return it.cursor > 0
}

// Prev returns the previous element. It retreats the cursor.
func (it *BSTIteratorII) Prev() int {
	it.cursor--
	return it.vals[it.cursor]
}

// BuildBST builds a BST from a level-order slice (null represented by -1).
func BuildBST(vals []int) *TreeNode {
	if len(vals) == 0 || vals[0] == -1 {
		return nil
	}
	root := &TreeNode{Val: vals[0]}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]
		if i < len(vals) && vals[i] != -1 {
			node.Left = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(vals) && vals[i] != -1 {
			node.Right = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func main() {
	// Example:
	// BST:
	//       7
	//      / \
	//     3   15
	//        /  \
	//       9   20
	//
	// Inorder: [3, 7, 9, 15, 20]

	root := BuildBST([]int{7, 3, 15, -1, -1, 9, 20})
	it := Constructor(root)

	fmt.Println("HasNext:", it.HasNext())      // true
	fmt.Println("Next:", it.Next())            // 3
	fmt.Println("HasPrev:", it.HasPrev())      // false (cursor at 0)
	fmt.Println("HasNext:", it.HasNext())      // true
	fmt.Println("Next:", it.Next())            // 7
	fmt.Println("Prev:", it.Prev())            // 3 (back to 3)
	fmt.Println("Next:", it.Next())            // 7 (forward again)
	fmt.Println("Next:", it.Next())            // 9
	fmt.Println("Next:", it.Next())            // 15
	fmt.Println("HasNext:", it.HasNext())      // true
	fmt.Println("Next:", it.Next())            // 20
	fmt.Println("HasNext:", it.HasNext())      // false
	fmt.Println("HasPrev:", it.HasPrev())      // true
	fmt.Println("Prev:", it.Prev())            // 15
	fmt.Println("Prev:", it.Prev())            // 9
	fmt.Println("Prev:", it.Prev())            // 7
	fmt.Println("Prev:", it.Prev())            // 3
	fmt.Println("HasPrev:", it.HasPrev())      // false
}
```

## 1591 — Strange Printer Ii

```go
package main

import (
	"fmt"
)

// LeetCode #1591: Strange Printer II
// https://leetcode.com/problems/strange-printer-ii/
// Difficulty: Hard
//
// The printer paints axis-aligned rectangles of a single color,
// one at a time, in some order. Later rectangles may cover earlier ones.
// Determine if the target grid can be produced.
//
// Approach:
// 1. For each color (1..60), find its bounding box (min/max row and col).
// 2. For each cell inside the bounding box of color c, if the cell's color
//    is different from c, that other color MUST have been painted AFTER c
//    (to override c's paint). This defines a directed edge: c -> other.
// 3. Build a directed graph and check if there is a cycle. If no cycle,
//    a topological order exists, meaning the grid is producible.

func isPrintable(targetGrid [][]int) bool {
	m := len(targetGrid)
	if m == 0 {
		return true
	}
	n := len(targetGrid[0])

	// Bounding boxes for each color (1-based, max 60 colors)
	type bbox struct {
		minR, maxR int
		minC, maxC int
	}
	boxes := make(map[int]*bbox)
	hasColor := make(map[int]bool)

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			color := targetGrid[r][c]
			hasColor[color] = true
			if _, ok := boxes[color]; !ok {
				boxes[color] = &bbox{minR: r, maxR: r, minC: c, maxC: c}
			} else {
				b := boxes[color]
				if r < b.minR {
					b.minR = r
				}
				if r > b.maxR {
					b.maxR = r
				}
				if c < b.minC {
					b.minC = c
				}
				if c > b.maxC {
					b.maxC = c
				}
			}
		}
	}

	// Build adjacency: if color a's bounding box contains a cell of color b (b != a),
	// then a must be painted before b (edge a -> b).
	graph := make(map[int]map[int]bool)
	for color, b := range boxes {
		if graph[color] == nil {
			graph[color] = make(map[int]bool)
		}
		for r := b.minR; r <= b.maxR; r++ {
			for c := b.minC; c <= b.maxC; c++ {
				other := targetGrid[r][c]
				if other != color {
					// color was painted first, then other painted over it
					if !graph[color][other] {
						graph[color][other] = true
					}
				}
			}
		}
	}

	// Detect cycle via DFS (topological sort / Kahn's algorithm)
	// Use three-color DFS: 0=unvisited, 1=visiting, 2=visited
	state := make(map[int]int)
	var colors []int
	for c := range hasColor {
		colors = append(colors, c)
	}

	var dfs func(int) bool
	dfs = func(node int) bool {
		state[node] = 1 // visiting
		for next := range graph[node] {
			if state[next] == 1 {
				return true // cycle found
			}
			if state[next] == 0 {
				if dfs(next) {
					return true
				}
			}
		}
		state[node] = 2 // visited
		return false
	}

	for _, c := range colors {
		if state[c] == 0 {
			if dfs(c) {
				return false
			}
		}
	}

	return true
}

func main() {
	// Example 1:
	// Input: targetGrid = [[1,1,1,1],[1,2,2,1],[1,2,2,1],[1,1,1,1]]
	// Output: true
	fmt.Println(isPrintable([][]int{
		{1, 1, 1, 1},
		{1, 2, 2, 1},
		{1, 2, 2, 1},
		{1, 1, 1, 1},
	}))

	// Example 2:
	// Input: targetGrid = [[1,1,1,1],[1,1,3,3],[1,1,3,3],[1,1,1,1]]
	// Output: true
	fmt.Println(isPrintable([][]int{
		{1, 1, 1, 1},
		{1, 1, 3, 3},
		{1, 1, 3, 3},
		{1, 1, 1, 1},
	}))

	// Example 3:
	// Input: targetGrid = [[1,1,1],[3,1,3]]
	// Output: false
	// Color 1's bounding box covers [0,0]..[1,2]. Cell (1,0)=3 and (1,2)=3
	// mean 1->3. Color 3's bounding box covers [1,0]..[1,2]. Cell (1,1)=1
	// means 3->1. Cycle: 1->3->1.
	fmt.Println(isPrintable([][]int{
		{1, 1, 1},
		{3, 1, 3},
	}))
}
```

## 1595 — Minimum Cost To Connect Two Groups Of Points

```go
package main

import (
	"fmt"
	"math"
)

// LeetCode #1595: Minimum Cost to Connect Two Groups of Points
// https://leetcode.com/problems/minimum-cost-to-connect-two-groups-of-points/
// Difficulty: Hard
//
// We have two groups of points (size1 = m, size2 = n). Each point in group 1
// must be connected to at least one point in group 2, and vice versa.
// Cost[i][j] = cost of connecting point i (group 1) to point j (group 2).
//
// DP with bitmask: dp[i][mask] = min cost to connect first i points of group 1
// to a subset of group 2 represented by mask, with the condition that all first
// i points of group 1 have at least one connection.
//
// Optimization: precompute minCost[j] for each group-2 point across all group-1
// points, to guarantee each group-2 point gets at least one connection.
//
// DP transition: dp[i][mask] = min over j (where mask has bit j set) of:
//   dp[i-1][mask without j] + cost[i-1][j]   (first connection for point i to j)
//   dp[i][mask without j] + cost[i-1][j]     (additional connection for point i)
//
// After processing all group 1 points, for each mask that covers all group 1
// points, add the minimum connection cost for any unconnected group 2 point.

func connectTwoGroups(cost [][]int) int {
	m := len(cost)
	n := len(cost[0])

	size := 1 << n
	INF := math.MaxInt32

	// dp[mask] = min cost after processing current prefix of group 1
	dp := make([]int, size)
	for i := range dp {
		dp[i] = INF
	}
	dp[0] = 0

	// Precompute minCostTo[j]: cheapest connection from ANY group-1 point to group-2 point j
	minCostTo := make([]int, n)
	for j := 0; j < n; j++ {
		minVal := math.MaxInt32
		for i := 0; i < m; i++ {
			if cost[i][j] < minVal {
				minVal = cost[i][j]
			}
		}
		minCostTo[j] = minVal
	}

	// Process each point in group 1
	for i := 0; i < m; i++ {
		ndp := make([]int, size)
		for mask := range ndp {
			ndp[mask] = INF
		}

		for mask := 0; mask < size; mask++ {
			if dp[mask] == INF {
				continue
			}
			// Try each group-2 point j
			for j := 0; j < n; j++ {
				newMask := mask | (1 << j)
				val := dp[mask] + cost[i][j]
				if val < ndp[newMask] {
					ndp[newMask] = val
				}
			}
			// Also option: connect i to j where j is already connected (stays same mask)
			// This is covered by the above loop (newMask may equal mask if j already set)
		}
		dp = ndp
	}

	// After processing all group-1 points, ensure every group-2 point is connected.
	// For each mask, if group-2 point j is NOT connected, add its min connection cost.
	answer := INF


	// Precompute extra cost to cover missing group-2 points for each mask
	extra := make([]int, size)
	for mask := 0; mask < size; mask++ {
		sum := 0
		for j := 0; j < n; j++ {
			if mask&(1<<j) == 0 {
				sum += minCostTo[j]
			}
		}
		extra[mask] = sum
	}

	for mask := 0; mask < size; mask++ {
		if dp[mask] == INF {
			continue
		}
		total := dp[mask] + extra[mask]
		if total < answer {
			answer = total
		}
	}

	return answer
}

func main() {
	// Example 1:
	// Input: cost = [[15, 96], [36, 2]]
	// Output: 17
	// Connect 0-0 (15) and 1-1 (2). Group1 both connected, group2 both connected. Total = 17.
	fmt.Println(connectTwoGroups([][]int{{15, 96}, {36, 2}}))

	// Example 2:
	// Input: cost = [[1, 3, 5], [4, 1, 1], [1, 5, 3]]
	// Output: 4
	// Connect 0-0 (1), 1-1 (1), 2-0 (1), 2-2 (1) -> total 4. All connected.
	fmt.Println(connectTwoGroups([][]int{{1, 3, 5}, {4, 1, 1}, {1, 5, 3}}))

	// Example 3:
	// Input: cost = [[2, 5, 1], [3, 4, 7], [8, 1, 2], [6, 2, 4], [3, 8, 8]]
	// Output: 10
	fmt.Println(connectTwoGroups([][]int{
		{2, 5, 1},
		{3, 4, 7},
		{8, 1, 2},
		{6, 2, 4},
		{3, 8, 8},
	}))
}
```

## 1596 — The Most Frequently Ordered Products For Each Customer

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1596: The Most Frequently Ordered Products for Each Customer
// https://leetcode.com/problems/the-most-frequently-ordered-products-for-each-customer/
// Difficulty: Hard
//
// For each customer, find the product(s) they ordered most frequently.
// If there are ties (multiple products ordered the same maximum number of times),
// include all of them.

// Order represents a customer order.
type Order struct {
	OrderID    int
	CustomerID int
	ProductID  int
}

// Product represents a product.
type Product struct {
	ProductID   int
	ProductName string
}

type resultRow1596 struct {
	CustomerID  int
	ProductName string
	ProductID   int
}

// mostFrequentProducts returns for each customer the product(s) they ordered most frequently.
func mostFrequentProducts(orders []Order, products []Product) []resultRow1596 {
	// Count orders per customer per product
	type cpKey struct {
		customerID int
		productID  int
	}
	counts := make(map[cpKey]int)
	// Also track max per customer
	maxPerCustomer := make(map[int]int)

	for _, o := range orders {
		key := cpKey{customerID: o.CustomerID, productID: o.ProductID}
		counts[key]++
		c := counts[key]
		if c > maxPerCustomer[o.CustomerID] {
			maxPerCustomer[o.CustomerID] = c
		}
	}

	productName := make(map[int]string)
	for _, p := range products {
		productName[p.ProductID] = p.ProductName
	}

	var results []resultRow1596

	for key, cnt := range counts {
		if cnt == maxPerCustomer[key.customerID] {
			results = append(results, resultRow1596{
				CustomerID:  key.customerID,
				ProductName: productName[key.productID],
				ProductID:   key.productID,
			})
		}
	}

	// Sort results by customer ID ascending, product name ascending
	sort.Slice(results, func(i, j int) bool {
		if results[i].CustomerID != results[j].CustomerID {
			return results[i].CustomerID < results[j].CustomerID
		}
		return results[i].ProductName < results[j].ProductName
	})

	return results
}

func main() {
	orders := []Order{
		{1, 1, 10},
		{2, 1, 20},
		{3, 1, 10},
		{4, 2, 30},
		{5, 2, 30},
		{6, 2, 20},
		{7, 3, 10},
		{8, 3, 20},
		{9, 3, 30},
	}
	products := []Product{
		{10, "Widget"},
		{20, "Gadget"},
		{30, "Doohickey"},
	}

	results := mostFrequentProducts(orders, products)
	for _, r := range results {
		fmt.Printf("Customer %d | %s (ID %d)\n", r.CustomerID, r.ProductName, r.ProductID)
	}
	// Expected:
	// Customer 1 | Widget (ID 10)  (ordered 2x vs Gadget 1x)
	// Customer 2 | Doohickey (ID 30)  (ordered 2x vs Gadget 1x)
	// Customer 3 | Widget, Gadget, Doohickey (each ordered 1x)
}
```

## 1597 — Build Binary Expression Tree From Infix Expression

```go
package main

import (
	"fmt"
	"strings"
)

// LeetCode #1597: Build Binary Expression Tree From Infix Expression
// https://leetcode.com/problems/build-binary-expression-tree-from-infix-expression/
// Difficulty: Hard [Paid]
//
// Build a binary expression tree from an infix expression string.
// The tree follows standard operator precedence: * and / bind tighter than + and -.
// The tree nodes are:
//
//	type Node struct {
//		Val   byte   // '0'-'9' for digits, '+', '-', '*', '/' for operators
//		Left  *Node
//		Right *Node
//	}
//
// Operators are stored in internal nodes; operands (digits) are leaves.

// Node represents a node in the binary expression tree.
type Node struct {
	Val   byte
	Left  *Node
	Right *Node
}

// expTree builds a binary expression tree from the infix expression s.
func expTree(s string) *Node {
	// Remove spaces
	s = strings.ReplaceAll(s, " ", "")

	// Shunting-yard: convert infix to postfix (RPN)
	var postfix []byte
	var ops []byte

	prec := map[byte]int{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
	}

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= '0' && ch <= '9' {
			postfix = append(postfix, ch)
		} else if ch == '(' {
			ops = append(ops, ch)
		} else if ch == ')' {
			for len(ops) > 0 && ops[len(ops)-1] != '(' {
				postfix = append(postfix, ops[len(ops)-1])
				ops = ops[:len(ops)-1]
			}
			ops = ops[:len(ops)-1] // pop '('
		} else {
			// operator
			for len(ops) > 0 && ops[len(ops)-1] != '(' &&
				prec[ops[len(ops)-1]] >= prec[ch] {
				postfix = append(postfix, ops[len(ops)-1])
				ops = ops[:len(ops)-1]
			}
			ops = append(ops, ch)
		}
	}

	for len(ops) > 0 {
		postfix = append(postfix, ops[len(ops)-1])
		ops = ops[:len(ops)-1]
	}

	// Build expression tree from postfix using a stack
	var stack []*Node
	for _, token := range postfix {
		if token >= '0' && token <= '9' {
			stack = append(stack, &Node{Val: token})
		} else {
			// Operator: pop two operands
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, &Node{
				Val:   token,
				Left:  left,
				Right: right,
			})
		}
	}

	if len(stack) != 1 {
		return nil
	}
	return stack[0]
}

// inorder returns the infix representation (with parentheses to verify structure).
func inorder(root *Node) string {
	if root == nil {
		return ""
	}
	if root.Left == nil && root.Right == nil {
		return string(root.Val)
	}
	return "(" + inorder(root.Left) + string(root.Val) + inorder(root.Right) + ")"
}

// postorder returns the postfix representation.
func postorder(root *Node) string {
	if root == nil {
		return ""
	}
	if root.Left == nil && root.Right == nil {
		return string(root.Val)
	}
	return postorder(root.Left) + postorder(root.Right) + string(root.Val)
}

func main() {
	// Example 1:
	// Input: "3*4-2*5"
	// Output: - (subtree: * and *)
	// Tree:
	//        -
	//      /   \
	//     *     *
	//    / \   / \
	//   3   4 2   5
	tree1 := expTree("3*4-2*5")
	fmt.Println("Infix:  ", inorder(tree1))
	fmt.Println("Postfix:", postorder(tree1))
	// Expected infix: ((3*4)-(2*5))
	// Expected postfix: 34*25*-

	// Example 2:
	// Input: "2-3/(5*2)+1"
	tree2 := expTree("2-3/(5*2)+1")
	fmt.Println("Infix:  ", inorder(tree2))
	fmt.Println("Postfix:", postorder(tree2))

	// Simple: "1+2"
	tree3 := expTree("1+2")
	fmt.Println("Infix:  ", inorder(tree3))   // (1+2)
	fmt.Println("Postfix:", postorder(tree3)) // 12+

	// Parentheses: "(1+2)*3"
	tree4 := expTree("(1+2)*3")
	fmt.Println("Infix:  ", inorder(tree4))   // ((1+2)*3)
	fmt.Println("Postfix:", postorder(tree4)) // 12+3*
}
```

## 1601 — Maximum Number Of Achievable Transfer Requests

```go
package main

// LeetCode #1601: Maximum Number of Achievable Transfer Requests
// https://leetcode.com/problems/maximum-number-of-achievable-transfer-requests/
// Difficulty: Hard
//
// Bitmask enumeration approach:
// - There are at most 20 requests and 20 buildings (n <= 20).
// - Try all subsets from largest to smallest (or use bitmask enumeration).
// - For each subset, simulate the transfers and check if net change is 0.
// - Return the size of the largest valid subset.

import "fmt"

func main() {
	// Example: n=5, requests=[[0,1],[1,0],[0,1],[1,2],[2,0],[3,4]] -> 5
	fmt.Println(maximumRequests(5, [][]int{
		{0, 1},
		{1, 0},
		{0, 1},
		{1, 2},
		{2, 0},
		{3, 4},
	}))

	// Additional tests
	fmt.Println(maximumRequests(3, [][]int{
		{0, 0},
		{1, 2},
		{2, 1},
	}))

	fmt.Println(maximumRequests(4, [][]int{
		{0, 1},
		{1, 0},
		{2, 3},
		{3, 2},
	}))

	fmt.Println(maximumRequests(3, [][]int{
		{0, 1},
		{1, 0},
		{1, 0},
		{0, 1},
	}))

	fmt.Println(maximumRequests(2, [][]int{
		{0, 1},
		{1, 0},
		{0, 1},
		{1, 0},
		{0, 1},
		{1, 0},
	}))
}

func maximumRequests(n int, requests [][]int) int {
	r := len(requests)
	result := 0

	// Try all subsets
	for mask := 1; mask < (1 << r); mask++ {
		size := popcount(mask)
		if size <= result {
			continue
		}
		if isValid(mask, n, requests) {
			result = size
		}
	}

	return result
}

func popcount(x int) int {
	count := 0
	for x > 0 {
		count += x & 1
		x >>= 1
	}
	return count
}

func isValid(mask int, n int, requests [][]int) bool {
	balance := make([]int, n)
	for i, req := range requests {
		if mask&(1<<i) != 0 {
			balance[req[0]]--
			balance[req[1]]++
		}
	}
	for _, b := range balance {
		if b != 0 {
			return false
		}
	}
	return true
}
```

## 1602 — Find Nearest Right Node In Binary Tree

```go
package main

import "fmt"

// LeetCode #1602: Find Nearest Right Node in Binary Tree
// https://leetcode.com/problems/find-nearest-right-node-in-binary-tree/
// Difficulty: Medium (listed in Hard section)
//
// Given the root of a binary tree and a node u, find the nearest node to the
// right of u on the same level. If u is the rightmost node on its level,
// return nil.

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// findNearestRightNode returns the nearest node to the right of u at the same level.
func findNearestRightNode(root *TreeNode, u *TreeNode) *TreeNode {
	if root == nil || u == nil {
		return nil
	}

	// BFS level-order traversal
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node == u {
				// Found u; return the next node at this level, if any
				if i+1 < levelSize {
					return queue[0] // queue[0] is the next node at this level
				}
				return nil
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return nil
}

// BuildBT builds a binary tree from a level-order slice (-1 for nil).
func BuildBT(vals []int) *TreeNode {
	if len(vals) == 0 || vals[0] == -1 {
		return nil
	}
	root := &TreeNode{Val: vals[0]}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]
		if i < len(vals) && vals[i] != -1 {
			node.Left = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(vals) && vals[i] != -1 {
			node.Right = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

// findNode finds a node with given value in the tree.
func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if left := findNode(root.Left, val); left != nil {
		return left
	}
	return findNode(root.Right, val)
}

func main() {
	// Example:
	// Tree:
	//       1
	//      / \
	//     2   3
	//    /   / \
	//   4   5   6
	//
	// Find nearest right node of 2: should be 3
	// Find nearest right node of 4: should be nil (rightmost)
	// Find nearest right node of 5: should be 6

	root := BuildBT([]int{1, 2, 3, 4, -1, 5, 6})

	node2 := findNode(root, 2)
	node4 := findNode(root, 4)
	node5 := findNode(root, 5)

	rightOf2 := findNearestRightNode(root, node2)
	rightOf4 := findNearestRightNode(root, node4)
	rightOf5 := findNearestRightNode(root, node5)

	fmt.Println("Nearest right of 2:", rightOf2) // 3
	if rightOf2 != nil {
		fmt.Println("  Val:", rightOf2.Val)
	}

	fmt.Println("Nearest right of 4:", rightOf4) // nil
	if rightOf4 != nil {
		fmt.Println("  Val:", rightOf4.Val)
	}

	fmt.Println("Nearest right of 5:", rightOf5) // 6
	if rightOf5 != nil {
		fmt.Println("  Val:", rightOf5.Val)
	}
}
```

## 1606 — Find Servers That Handled Most Number Of Requests

```go
package main

// LeetCode #1606: Find Servers That Handled Most Number of Requests
// https://leetcode.com/problems/find-servers-that-handled-most-number-of-requests/
// Difficulty: Hard

import (
	"container/heap"
	"fmt"
	"sort"
)

// --- Fenwick Tree (BIT) for ordered set of available servers ---

type BIT struct {
	tree []int
	n    int
}

func newBIT(n int) *BIT {
	b := &BIT{tree: make([]int, n+1), n: n}
	for i := 1; i <= n; i++ {
		b.add(i, 1)
	}
	return b
}

func (b *BIT) add(idx, val int) {
	for i := idx; i <= b.n; i += i & -i {
		b.tree[i] += val
	}
}

func (b *BIT) sum(idx int) int {
	s := 0
	for i := idx; i > 0; i -= i & -i {
		s += b.tree[i]
	}
	return s
}

// find smallest idx with prefix sum >= target (1-indexed)
func (b *BIT) find(target int) int {
	idx := 0
	bitMask := 1
	for bitMask <= b.n {
		bitMask <<= 1
	}
	bitMask >>= 1

	for bitMask > 0 {
		t := idx + bitMask
		if t <= b.n && b.tree[t] < target {
			target -= b.tree[t]
			idx = t
		}
		bitMask >>= 1
	}
	return idx + 1
}

// ceil returns smallest available server >= x, or wraps to first available
func (b *BIT) ceil(x int) int {
	total := b.sum(b.n)
	if total == 0 {
		return -1
	}
	before := b.sum(x)
	if before < total {
		return b.find(before + 1) - 1 // convert to 0-indexed
	}
	return b.find(1) - 1 // wrap around
}

func (b *BIT) remove(x int) {
	b.add(x+1, -1)
}

func (b *BIT) addServer(x int) {
	b.add(x+1, 1)
}

// --- Min Heap for busy servers ---

type busyItem struct {
	endTime int
	index   int
}

type busyHeap []*busyItem

func (h busyHeap) Len() int            { return len(h) }
func (h busyHeap) Less(i, j int) bool  { return h[i].endTime < h[j].endTime }
func (h busyHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *busyHeap) Push(x interface{}) { *h = append(*h, x.(*busyItem)) }
func (h *busyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h *busyHeap) Peek() *busyItem { return (*h)[0] }

// --- Main solution ---

func busiestServers(k int, arrival []int, load []int) []int {
	counts := make([]int, k)
	available := newBIT(k)
	busy := &busyHeap{}
	heap.Init(busy)

	maxCount := 0

	for i := 0; i < len(arrival); i++ {
		t := arrival[i]

		// Free completed servers
		for busy.Len() > 0 && busy.Peek().endTime <= t {
			item := heap.Pop(busy).(*busyItem)
			available.addServer(item.index)
		}

		// Find server to assign
		serverIdx := available.ceil(i % k)
		if serverIdx == -1 {
			continue // all servers busy, drop request
		}

		available.remove(serverIdx)
		heap.Push(busy, &busyItem{endTime: t + load[i], index: serverIdx})
		counts[serverIdx]++
		if counts[serverIdx] > maxCount {
			maxCount = counts[serverIdx]
		}
	}

	result := []int{}
	for i, c := range counts {
		if c == maxCount {
			result = append(result, i)
		}
	}
	sort.Ints(result)
	return result
}

func main() {
	// Test case 1: k=3, arrival=[1,2,3,4,5], load=[5,2,3,3,3] -> [1]
	k := 3
	arrival := []int{1, 2, 3, 4, 5}
	load := []int{5, 2, 3, 3, 3}
	result := busiestServers(k, arrival, load)
	fmt.Printf("k=%d arrival=%v load=%v -> %v\n", k, arrival, load, result)

	// Test case 2: k=3, arrival=[1,2,3,4,8,9,10], load=[5,2,3,3,2,1,1] -> [1]
	arrival2 := []int{1, 2, 3, 4, 8, 9, 10}
	load2 := []int{5, 2, 3, 3, 2, 1, 1}
	result2 := busiestServers(3, arrival2, load2)
	fmt.Printf("k=3 arrival=%v load=%v -> %v\n", arrival2, load2, result2)

	// Test case 3: k=1, arrival=[1], load=[1] -> [0]
	result3 := busiestServers(1, []int{1}, []int{1})
	fmt.Printf("k=1 arrival=[1] load=[1] -> %v\n", result3)
}
```

## 1610 — Maximum Number Of Visible Points

```go
package main

// LeetCode #1610: Maximum Number of Visible Points
// https://leetcode.com/problems/maximum-number-of-visible-points/
// Difficulty: Hard

import (
	"fmt"
	"math"
	"sort"
)

func visiblePoints(points [][]int, angle int, location []int) int {
	angles := []float64{}
	same := 0

	for _, p := range points {
		if p[0] == location[0] && p[1] == location[1] {
			same++
			continue
		}
		// atan2 returns angle in [-pi, pi], convert to degrees
		rad := math.Atan2(float64(p[1]-location[1]), float64(p[0]-location[0]))
		deg := rad * 180.0 / math.Pi
		angles = append(angles, deg)
	}

	sort.Float64s(angles)
	n := len(angles)

	// Duplicate with +360 for circular sliding window
	for i := 0; i < n; i++ {
		angles = append(angles, angles[i]+360.0)
	}

	maxVis := 0
	j := 0
	angleF := float64(angle)

	for i := 0; i < n; i++ {
		for j < len(angles) && angles[j] <= angles[i]+angleF {
			j++
		}
		if j-i > maxVis {
			maxVis = j - i
		}
	}

	return same + maxVis
}

func main() {
	// Test case 1: points=[[2,1],[2,2],[3,4]], angle=90, location=[1,1] -> 3
	points := [][]int{{2, 1}, {2, 2}, {3, 4}}
	result := visiblePoints(points, 90, []int{1, 1})
	fmt.Printf("points=%v angle=90 location=[1,1] -> %d (expected 3)\n", points, result)

	// Test case 2: points=[[1,1],[2,2],[3,3],[1,1]], angle=0, location=[1,1] -> 4
	// Same location points: 2 (both [1,1]).
	// Points [2,2] and [3,3] share the same angle (45 deg), both visible with angle=0.
	points2 := [][]int{{1, 1}, {2, 2}, {3, 3}, {1, 1}}
	result2 := visiblePoints(points2, 0, []int{1, 1})
	fmt.Printf("points=%v angle=0 location=[1,1] -> %d\n", points2, result2)

	// Test case 3: points=[[0,0]], angle=90, location=[1,1] -> 0
	points3 := [][]int{{0, 0}}
	result3 := visiblePoints(points3, 90, []int{1, 1})
	fmt.Printf("points=%v angle=90 location=[1,1] -> %d\n", points3, result3)
}
```

## 1611 — Minimum One Bit Operations To Make Integers Zero

```go
package main

// LeetCode #1611: Minimum One Bit Operations to Make Integers Zero
// https://leetcode.com/problems/minimum-one-bit-operations-to-make-integers-zero/
// Difficulty: Hard
//
// Solution: The minimum operations = inverse Gray code of n.
// Gray code: g(n) = n ^ (n>>1)
// Inverse Gray code (find position p such that g(p) = n):
//   ans = n ^ (n>>1) ^ (n>>2) ^ (n>>4) ^ ... (until zero)

import "fmt"

func minimumOneBitOperations(n int) int {
	ans := 0
	for n > 0 {
		ans ^= n
		n >>= 1
	}
	return ans
}

func main() {
	// Test case 1: 3 -> 2
	result := minimumOneBitOperations(3)
	fmt.Printf("minimumOneBitOperations(3) = %d (expected 2)\n", result)

	// Test case 2: 6 -> 4
	result2 := minimumOneBitOperations(6)
	fmt.Printf("minimumOneBitOperations(6) = %d (expected 4)\n", result2)

	// Test case 3: 0 -> 0
	result3 := minimumOneBitOperations(0)
	fmt.Printf("minimumOneBitOperations(0) = %d (expected 0)\n", result3)

	// Test case 4: 2 -> 3
	result4 := minimumOneBitOperations(2)
	fmt.Printf("minimumOneBitOperations(2) = %d (expected 3)\n", result4)

	// Test case 5: 9 -> 14
	result5 := minimumOneBitOperations(9)
	fmt.Printf("minimumOneBitOperations(9) = %d (expected 14)\n", result5)
}
```

## 1617 — Count Subtrees With Max Distance Between Cities

```go
package main

// LeetCode #1617: Count Subtrees With Max Distance Between Cities
// https://leetcode.com/problems/count-subtrees-with-max-distance-between-cities/
// Difficulty: Hard

import "fmt"

func countSubgraphsForEachDiameter(n int, edges [][]int) []int {
	// Floyd-Warshall for all-pairs shortest paths
	INF := 1000
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = INF
		}
		dist[i][i] = 0
	}

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		dist[u][v] = 1
		dist[v][u] = 1
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	result := make([]int, n-1) // diameters 1..n-1

	// Enumerate all non-empty subsets
	for mask := 1; mask < (1 << n); mask++ {
		// Collect nodes in this subset
		nodes := make([]int, 0, n)
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				nodes = append(nodes, i)
			}
		}
		if len(nodes) < 2 {
			continue
		}

		// Check connectivity: BFS on induced subgraph
		visited := make(map[int]bool)
		queue := []int{nodes[0]}
		visited[nodes[0]] = true
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			for _, nb := range nodes {
				if !visited[nb] && dist[cur][nb] == 1 {
					visited[nb] = true
					queue = append(queue, nb)
				}
			}
		}
		if len(visited) != len(nodes) {
			continue // not connected
		}

		// Compute diameter (max distance between any two nodes in subset)
		maxDist := 0
		for i := 0; i < len(nodes); i++ {
			for j := i + 1; j < len(nodes); j++ {
				if dist[nodes[i]][nodes[j]] > maxDist {
					maxDist = dist[nodes[i]][nodes[j]]
				}
			}
		}

		if maxDist >= 1 && maxDist <= n-1 {
			result[maxDist-1]++
		}
	}

	return result
}

func main() {
	// Test case 1: n=4, edges=[[1,2],[2,3],[2,4]] -> [3,4,0]
	n := 4
	edges := [][]int{{1, 2}, {2, 3}, {2, 4}}
	result := countSubgraphsForEachDiameter(n, edges)
	fmt.Printf("n=%d edges=%v -> %v (expected [3,4,0])\n", n, edges, result)

	// Test case 2: n=2, edges=[[1,2]] -> [1]
	n2 := 2
	edges2 := [][]int{{1, 2}}
	result2 := countSubgraphsForEachDiameter(n2, edges2)
	fmt.Printf("n=%d edges=%v -> %v (expected [1])\n", n2, edges2, result2)

	// Test case 3: n=3, edges=[[1,2],[2,3]] -> [2,1]
	n3 := 3
	edges3 := [][]int{{1, 2}, {2, 3}}
	result3 := countSubgraphsForEachDiameter(n3, edges3)
	fmt.Printf("n=%d edges=%v -> %v (expected [2,1])\n", n3, edges3, result3)
}
```

## 1622 — Fancy Sequence

```go
package main

// LeetCode #1622: Fancy Sequence
// https://leetcode.com/problems/fancy-sequence/
// Difficulty: Hard

import "fmt"

const MOD = 1000000007

// modInv computes modular inverse using Fermat's little theorem (MOD is prime)
func modInv(a int) int {
	return modPow(a, MOD-2)
}

func modPow(a, b int) int {
	res := 1
	a %= MOD
	for b > 0 {
		if b&1 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		b >>= 1
	}
	return res
}

type Fancy struct {
	arr []int // stored base values (normalized)
	mul int   // global multiplier
	add int   // global adder
}

func Constructor() Fancy {
	return Fancy{
		arr: []int{},
		mul: 1,
		add: 0,
	}
}

// Append(val): append val to sequence
func (f *Fancy) Append(val int) {
	// If mul == 0, all existing elements have value = f.add.
	// We need to normalize: set all existing stored values to 0
	// and set mul=1, add=f.add (unchanged).
	if f.mul == 0 {
		for i := range f.arr {
			f.arr[i] = 0
		}
		f.mul = 1
		// f.add stays the same
	}

	// Normalize: find stored_val such that stored_val * mul + add = val
	// stored_val = (val - add) * inv(mul) (mod MOD)
	v := (val - f.add + MOD) % MOD
	v = v * modInv(f.mul) % MOD
	f.arr = append(f.arr, v)
}

// AddAll(inc): add inc to every element
func (f *Fancy) AddAll(inc int) {
	f.add = (f.add + inc) % MOD
}

// MultAll(m): multiply every element by m
func (f *Fancy) MultAll(m int) {
	f.mul = f.mul * m % MOD
	f.add = f.add * m % MOD
}

// GetIndex(idx): return value at index, -1 if out of bounds
func (f *Fancy) GetIndex(idx int) int {
	if idx >= len(f.arr) {
		return -1
	}
	if f.mul == 0 {
		return f.add
	}
	return (f.arr[idx]*f.mul + f.add) % MOD
}

// FancySequence tests the Fancy struct operations and returns results
func FancySequence() []int {
	f := Constructor()
	results := []int{}

	f.Append(2)      // seq = [2]
	f.AddAll(3)      // seq = [5]
	f.Append(7)      // seq = [5, 7]
	f.MultAll(2)     // seq = [10, 14]
	results = append(results, f.GetIndex(0)) // 10
	f.AddAll(3)      // seq = [13, 17]
	f.Append(10)     // seq = [13, 17, 10]
	f.MultAll(2)     // seq = [26, 34, 20]
	results = append(results, f.GetIndex(0)) // 26
	results = append(results, f.GetIndex(1)) // 34
	results = append(results, f.GetIndex(2)) // 20

	return results
}

func main() {
	// Test case: sequence of operations
	f := Constructor()
	f.Append(2)
	f.AddAll(3)
	f.Append(7)
	f.MultAll(2)
	fmt.Printf("GetIndex(0) = %d (expected 10)\n", f.GetIndex(0))
	f.AddAll(3)
	f.Append(10)
	f.MultAll(2)
	fmt.Printf("GetIndex(0) = %d (expected 26)\n", f.GetIndex(0))
	fmt.Printf("GetIndex(1) = %d (expected 34)\n", f.GetIndex(1))
	fmt.Printf("GetIndex(2) = %d (expected 20)\n", f.GetIndex(2))

	// Test case: multAll(0) then append
	f2 := Constructor()
	f2.Append(5)    // [5]
	f2.MultAll(0)   // [0]
	f2.Append(3)    // [0, 3]
	f2.AddAll(2)    // [2, 5]
	fmt.Printf("After multAll(0): GetIndex(0) = %d (expected 2)\n", f2.GetIndex(0))
	fmt.Printf("After multAll(0): GetIndex(1) = %d (expected 5)\n", f2.GetIndex(1))

	// Test case: empty
	f3 := Constructor()
	fmt.Printf("Empty getIndex(0) = %d (expected -1)\n", f3.GetIndex(0))
}
```

## 1627 — Graph Connectivity With Threshold

```go
package main

// LeetCode #1627: Graph Connectivity With Threshold
// https://leetcode.com/problems/graph-connectivity-with-threshold/
// Difficulty: Hard

import "fmt"

type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{parent, rank}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	xr, yr := uf.Find(x), uf.Find(y)
	if xr == yr {
		return
	}
	if uf.rank[xr] < uf.rank[yr] {
		uf.parent[xr] = yr
	} else if uf.rank[xr] > uf.rank[yr] {
		uf.parent[yr] = xr
	} else {
		uf.parent[yr] = xr
		uf.rank[xr]++
	}
}

func (uf *UnionFind) Connected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

func areConnected(n int, threshold int, queries [][]int) []bool {
	uf := NewUnionFind(n + 1) // 1-indexed

	// For each divisor d > threshold, connect all multiples of d
	for d := threshold + 1; d <= n; d++ {
		for m := 2 * d; m <= n; m += d {
			uf.Union(d, m)
		}
	}

	result := make([]bool, len(queries))
	for i, q := range queries {
		result[i] = uf.Connected(q[0], q[1])
	}
	return result
}

func main() {
	// Test case 1: n=6, threshold=2, queries=[[1,4],[2,5],[3,6]] -> [false,false,true]
	n := 6
	threshold := 2
	queries := [][]int{{1, 4}, {2, 5}, {3, 6}}
	result := areConnected(n, threshold, queries)
	fmt.Printf("n=%d threshold=%d queries=%v -> %v (expected [false,false,true])\n",
		n, threshold, queries, result)

	// Test case 2: n=6, threshold=0, queries=[[4,5],[3,4],[3,2],[2,6],[1,3]] -> [true,false,true,false,true]
	n2 := 6
	threshold2 := 0
	queries2 := [][]int{{4, 5}, {3, 4}, {3, 2}, {2, 6}, {1, 3}}
	result2 := areConnected(n2, threshold2, queries2)
	fmt.Printf("n=%d threshold=%d queries=%v -> %v\n", n2, threshold2, queries2, result2)

	// Test case 3: n=5, threshold=1, queries=[[4,5],[4,5],[3,2],[2,3],[3,4]] -> [false,false,false,false,false]
	n3 := 5
	threshold3 := 1
	queries3 := [][]int{{4, 5}, {4, 5}, {3, 2}, {2, 3}, {3, 4}}
	result3 := areConnected(n3, threshold3, queries3)
	fmt.Printf("n=%d threshold=%d queries=%v -> %v\n", n3, threshold3, queries3, result3)
}
```

## 1632 — Rank Transform Of A Matrix

```go
package main

// LeetCode #1632: Rank Transform of a Matrix
// https://leetcode.com/problems/rank-transform-of-a-matrix/
// Difficulty: Hard

import (
	"fmt"
	"sort"
)

type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{parent, rank}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	xr, yr := uf.Find(x), uf.Find(y)
	if xr == yr {
		return
	}
	if uf.rank[xr] < uf.rank[yr] {
		uf.parent[xr] = yr
	} else if uf.rank[xr] > uf.rank[yr] {
		uf.parent[yr] = xr
	} else {
		uf.parent[yr] = xr
		uf.rank[xr]++
	}
}

func matrixRankTransform(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
	}

	// Group cells by their value
	valToCells := make(map[int][][2]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			val := matrix[i][j]
			valToCells[val] = append(valToCells[val], [2]int{i, j})
		}
	}

	// Sort unique values
	values := make([]int, 0, len(valToCells))
	for v := range valToCells {
		values = append(values, v)
	}
	sort.Ints(values)

	// Track the current max rank for each row and column
	rowMax := make([]int, m)
	colMax := make([]int, n)

	for _, val := range values {
		cells := valToCells[val]

		// Union cells that are in the same row or column (same value)
		uf := NewUnionFind(len(cells))
		rowMap := make(map[int]int) // row -> first cell index
		colMap := make(map[int]int) // col -> first cell index

		for idx, cell := range cells {
			r, c := cell[0], cell[1]
			if first, ok := rowMap[r]; ok {
				uf.Union(idx, first)
			} else {
				rowMap[r] = idx
			}
			if first, ok := colMap[c]; ok {
				uf.Union(idx, first)
			} else {
				colMap[c] = idx
			}
		}

		// Group cells by their root (connected components)
		groups := make(map[int][]int)
		for idx := range cells {
			root := uf.Find(idx)
			groups[root] = append(groups[root], idx)
		}

		// For each connected component, compute rank = max(rowMax, colMax) + 1
		for _, group := range groups {
			maxRank := 0
			for _, idx := range group {
				r, c := cells[idx][0], cells[idx][1]
				if rowMax[r] > maxRank {
					maxRank = rowMax[r]
				}
				if colMax[c] > maxRank {
					maxRank = colMax[c]
				}
			}
			rank := maxRank + 1
			for _, idx := range group {
				r, c := cells[idx][0], cells[idx][1]
				result[r][c] = rank
				rowMax[r] = rank
				colMax[c] = rank
			}
		}
	}

	return result
}

func main() {
	// Test case 1: [[1,2],[3,4]] -> [[1,2],[2,3]]
	matrix := [][]int{{1, 2}, {3, 4}}
	result := matrixRankTransform(matrix)
	fmt.Printf("matrix=%v -> %v (expected [[1,2],[2,3]])\n", matrix, result)

	// Test case 2: [[7,7],[7,7]] -> [[1,1],[1,1]]
	matrix2 := [][]int{{7, 7}, {7, 7}}
	result2 := matrixRankTransform(matrix2)
	fmt.Printf("matrix=%v -> %v (expected [[1,1],[1,1]])\n", matrix2, result2)

	// Test case 3: [[20,-21,14],[-19,4,19],[22,-47,24],[-19,4,19]]
	matrix3 := [][]int{{20, -21, 14}, {-19, 4, 19}, {22, -47, 24}, {-19, 4, 19}}
	result3 := matrixRankTransform(matrix3)
	fmt.Printf("matrix=%v -> %v\n", matrix3, result3)
}
```

## 1635 — Hopper Company Queries I

```go
package main

// LeetCode #1635: Hopper Company Queries I
// https://leetcode.com/problems/hopper-company-queries-i/
// Difficulty: Hard [Paid] (SQL)
//
// For each month of 2020, report the number of active drivers by
// month-end and the number of accepted rides that month.
//
// Approach: Process drivers and rides data in Go to simulate the
// SQL query results.

import (
	"fmt"
	"time"
)

// Driver represents a driver
type Driver struct {
	DriverID int
	JoinDate time.Time
}

// Ride represents a ride request
type Ride struct {
	RideID      int
	UserID      int
	RequestedAt time.Time
}

// AcceptedRide represents an accepted ride
type AcceptedRide struct {
	RideID        int
	DriverID      int
	RideDistance  int
	RideDuration  int
}

func main() {
	// Example
	drivers := []Driver{
		{10, parseDate("2019-12-10")},
		{8, parseDate("2020-1-13")},
		{5, parseDate("2020-2-16")},
		{7, parseDate("2020-3-8")},
		{4, parseDate("2020-5-17")},
		{1, parseDate("2020-10-24")},
		{6, parseDate("2021-1-5")},
	}
	rides := []Ride{
		{6, 82, parseDate("2019-12-9")},
		{1, 17, parseDate("2020-1-2")},
		{10, 36, parseDate("2020-1-11")},
		{11, 98, parseDate("2020-1-19")},
		{12, 44, parseDate("2020-1-29")},
		{3, 1, parseDate("2020-1-17")},
		{9, 90, parseDate("2020-2-13")},
		{2, 50, parseDate("2020-2-22")},
	}
	accepted := []AcceptedRide{
		{2, 10, 63, 38},
		{13, 10, 58, 0},
		{7, 8, 51, 23},
		{3, 5, 50, 36},
		{4, 5, 58, 19},
		{11, 7, 53, 26},
		{1, 10, 50, 42},
		{5, 7, 55, 42},
		{12, 8, 51, 23},
		{6, 7, 56, 41},
	}

	fmt.Println(hopperQueriesI(drivers, rides, accepted))
}

func parseDate(s string) time.Time {
	t, err := time.Parse("2006-1-2", s)
	if err != nil {
		t, err = time.Parse("2006-01-02", s)
		if err != nil {
			panic(err)
		}
	}
	return t
}

func hopperQueriesI(drivers []Driver, rides []Ride, accepted []AcceptedRide) [][3]int {
	// Count drivers active by end of each month in 2020
	acceptedRideSet := make(map[int]bool)
	for _, ar := range accepted {
		acceptedRideSet[ar.RideID] = true
	}

	rideMonth := make(map[int]int) // rideID -> month
	for _, r := range rides {
		if r.RequestedAt.Year() == 2020 {
			rideMonth[r.RideID] = int(r.RequestedAt.Month())
		}
	}

	activeDrivers := make([]int, 13)
	for _, d := range drivers {
		joinYear, joinMonth := d.JoinDate.Year(), d.JoinDate.Month()
		if joinYear < 2020 {
			for m := 1; m <= 12; m++ {
				activeDrivers[m]++
			}
		} else if joinYear == 2020 {
			for m := int(joinMonth); m <= 12; m++ {
				activeDrivers[m]++
			}
		}
	}

	acceptedRides := make([]int, 13)
	for rideID, month := range rideMonth {
		if acceptedRideSet[rideID] {
			acceptedRides[month]++
		}
	}

	result := make([][3]int, 12)
	for m := 1; m <= 12; m++ {
		result[m-1] = [3]int{m, activeDrivers[m], acceptedRides[m]}
	}
	return result
}
```

## 1639 — Number Of Ways To Form A Target String Given A Dictionary

```go
package main

// LeetCode #1639: Number of Ways to Form a Target String Given a Dictionary
// https://leetcode.com/problems/number-of-ways-to-form-a-target-string-given-a-dictionary/
// Difficulty: Hard

import "fmt"

const MOD1639 = 1000000007

func numWays(words []string, target string) int {
	m := len(words[0]) // number of columns
	n := len(target)   // target length

	// count[c][ch] = number of words with character ch at column c
	count := make([][26]int, m)
	for _, w := range words {
		for c, ch := range w {
			count[c][ch-'a']++
		}
	}

	// dp[j] = number of ways to form first j characters of target
	dp := make([]int, n+1)
	dp[0] = 1

	for c := 0; c < m; c++ {
		// Process right-to-left to avoid using the same column twice
		for j := n; j > 0; j-- {
			ch := target[j-1] - 'a'
			if count[c][ch] > 0 {
				dp[j] = (dp[j] + dp[j-1]*count[c][ch]) % MOD1639
			}
		}
	}

	return dp[n]
}

func main() {
	// Test case 1: words=["acca","bbbb","caca"], target="aba" -> 6
	words := []string{"acca", "bbbb", "caca"}
	target := "aba"
	result := numWays(words, target)
	fmt.Printf("words=%v target=%s -> %d (expected 6)\n", words, target, result)

	// Test case 2: words=["abba","baab"], target="bab" -> 4
	words2 := []string{"abba", "baab"}
	target2 := "bab"
	result2 := numWays(words2, target2)
	fmt.Printf("words=%v target=%s -> %d (expected 4)\n", words2, target2, result2)

	// Test case 3: words=["abcd"], target="abcd" -> 1
	words3 := []string{"abcd"}
	target3 := "abcd"
	result3 := numWays(words3, target3)
	fmt.Printf("words=%v target=%s -> %d (expected 1)\n", words3, target3, result3)

	// Test case 4: words=["abcd"], target="ac" -> 1 (col 0='a', col 2='c')
	words4 := []string{"abcd"}
	target4 := "ac"
	result4 := numWays(words4, target4)
	fmt.Printf("words=%v target=%s -> %d (expected 1)\n", words4, target4, result4)
}
```

## 1643 — Kth Smallest Instructions

```go
package main

// LeetCode #1643: Kth Smallest Instructions
// https://leetcode.com/problems/kth-smallest-instructions/
// Difficulty: Hard

import "fmt"

// nCr computes binomial coefficient using Pascal's triangle
func nCr(n, r int) int {
	if r < 0 || r > n {
		return 0
	}
	if r > n-r {
		r = n - r
	}
	res := 1
	for i := 0; i < r; i++ {
		res = res * (n - i) / (i + 1)
	}
	return res
}

func kthSmallestPath(destination []int, k int) string {
	v, h := destination[0], destination[1] // vertical (V) and horizontal (H) steps
	n := v + h                              // total steps

	result := make([]byte, n)

	for i := 0; i < n; i++ {
		if h > 0 {
			// Number of ways if we place 'H' here:
			// Remaining positions = v + h - 1, remaining V's = v
			ways := nCr(v+h-1, v)
			if k <= ways {
				result[i] = 'H'
				h--
			} else {
				result[i] = 'V'
				k -= ways
				v--
			}
		} else {
			// No H's left, must place V
			result[i] = 'V'
			v--
		}
	}

	return string(result)
}

func main() {
	// Test case 1: destination=[2,3], k=1 -> "HHHVV"
	dest := []int{2, 3}
	k := 1
	result := kthSmallestPath(dest, k)
	fmt.Printf("destination=%v k=%d -> \"%s\" (expected \"HHHVV\")\n", dest, k, result)

	// Test case 2: destination=[2,3], k=2 -> "HHVHV"
	k2 := 2
	result2 := kthSmallestPath(dest, k2)
	fmt.Printf("destination=%v k=%d -> \"%s\" (expected \"HHVHV\")\n", dest, k2, result2)

	// Test case 3: destination=[2,3], k=3 -> "HHVVH"
	k3 := 3
	result3 := kthSmallestPath(dest, k3)
	fmt.Printf("destination=%v k=%d -> \"%s\" (expected \"HHVVH\")\n", dest, k3, result3)

	// Test case 4: destination=[2,3], k=last -> "VVHHH" (last = C(5,2)=10)
	k4 := 10
	result4 := kthSmallestPath(dest, k4)
	fmt.Printf("destination=%v k=%d -> \"%s\" (expected \"VVHHH\")\n", dest, k4, result4)

	// Test case 5: destination=[1,1], k=1 -> "HV"
	dest5 := []int{1, 1}
	result5 := kthSmallestPath(dest5, 1)
	fmt.Printf("destination=%v k=1 -> \"%s\" (expected \"HV\")\n", dest5, result5)

	// Test case 6: destination=[1,1], k=2 -> "VH"
	result6 := kthSmallestPath(dest5, 2)
	fmt.Printf("destination=%v k=2 -> \"%s\" (expected \"VH\")\n", dest5, result6)
}
```

## 1644 — Lowest Common Ancestor Of A Binary Tree Ii

```go
package main

import "fmt"

// LeetCode #1644: Lowest Common Ancestor of a Binary Tree II
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree-ii/
// Difficulty: Medium (listed in Hard section)
//
// Given the root of a binary tree and two nodes p and q, return their lowest
// common ancestor. Unlike LCA I, p and q may not exist in the tree.
// If either node does not exist, return nil.

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// lowestCommonAncestor returns the LCA of p and q, or nil if either is missing.
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// First pass: check existence of both nodes
	foundP := false
	foundQ := false

	var search func(*TreeNode)
	search = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node == p {
			foundP = true
		}
		if node == q {
			foundQ = true
		}
		search(node.Left)
		search(node.Right)
	}
	search(root)

	if !foundP || !foundQ {
		return nil
	}

	// Both exist, find LCA
	var lca func(*TreeNode) *TreeNode
	lca = func(node *TreeNode) *TreeNode {
		if node == nil || node == p || node == q {
			return node
		}
		left := lca(node.Left)
		right := lca(node.Right)
		if left != nil && right != nil {
			return node
		}
		if left != nil {
			return left
		}
		return right
	}

	return lca(root)
}

// BuildBT builds a binary tree from a level-order slice (-1 for nil).
func BuildBT(vals []int) *TreeNode {
	if len(vals) == 0 || vals[0] == -1 {
		return nil
	}
	root := &TreeNode{Val: vals[0]}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]
		if i < len(vals) && vals[i] != -1 {
			node.Left = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(vals) && vals[i] != -1 {
			node.Right = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

// findNode finds a node with given value in the tree.
func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if left := findNode(root.Left, val); left != nil {
		return left
	}
	return findNode(root.Right, val)
}

func main() {
	// Example 1:
	// Tree:
	//       3
	//      / \
	//     5   1
	//    / \  / \
	//   6  2 0  8
	//     / \
	//    7   4
	//
	// LCA of 5 and 1 = 3
	// LCA of 5 and 4 = 5
	// LCA of 5 and 999 (non-existent) = nil

	root := BuildBT([]int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4})

	node5 := findNode(root, 5)
	node1 := findNode(root, 1)
	node4 := findNode(root, 4)

	fmt.Println("LCA of 5 and 1:", lowestCommonAncestor(root, node5, node1)) // 3
	if lca := lowestCommonAncestor(root, node5, node1); lca != nil {
		fmt.Println("  Val:", lca.Val)
	}

	fmt.Println("LCA of 5 and 4:", lowestCommonAncestor(root, node5, node4)) // 5
	if lca := lowestCommonAncestor(root, node5, node4); lca != nil {
		fmt.Println("  Val:", lca.Val)
	}

	// Non-existent node
	fakeNode := &TreeNode{Val: 999}
	fmt.Println("LCA of 5 and 999 (fake):", lowestCommonAncestor(root, node5, fakeNode)) // nil

	// Self LCA
	fmt.Println("LCA of 5 and 5:", lowestCommonAncestor(root, node5, node5)) // 5
	if lca := lowestCommonAncestor(root, node5, node5); lca != nil {
		fmt.Println("  Val:", lca.Val)
	}

	// Empty tree
	fmt.Println("LCA nil:", lowestCommonAncestor(nil, node5, node1)) // nil
}
```

## 1645 — Hopper Company Queries Ii

```go
package main

// LeetCode #1645: Hopper Company Queries II
// https://leetcode.com/problems/hopper-company-queries-ii/
// Difficulty: Hard [Paid] (SQL)
//
// For each month of 2020, report the percentage of working drivers
// (drivers who accepted at least one ride that month) relative to
// the number of available drivers by month-end.
//
// Approach: Process driver and ride data in Go.

import (
	"fmt"
	"time"
)

func main() {
	// Example data
	drivers := []struct {
		id       int
		joinDate string
	}{
		{10, "2019-12-10"},
		{8, "2020-1-13"},
		{5, "2020-2-16"},
		{7, "2020-3-8"},
		{4, "2020-5-17"},
		{1, "2020-10-24"},
		{6, "2021-1-5"},
	}
	rides := []struct {
		id       int
		reqDate  string
	}{
		{6, "2019-12-9"},
		{1, "2020-1-2"},
		{10, "2020-1-11"},
		{11, "2020-1-19"},
		{12, "2020-1-29"},
		{3, "2020-1-17"},
		{9, "2020-2-13"},
		{2, "2020-2-22"},
	}
	accepted := []struct {
		rideID   int
		driverID int
	}{
		{2, 10},
		{13, 10},
		{7, 8},
		{3, 5},
		{4, 5},
		{11, 7},
		{1, 10},
		{5, 7},
		{12, 8},
		{6, 7},
	}

	fmt.Println(hopperQueriesII(drivers, rides, accepted))
}

func hopperQueriesII(drivers []struct {
	id       int
	joinDate string
}, rides []struct {
	id      int
	reqDate string
}, accepted []struct {
	rideID   int
	driverID int
}) []struct {
	month             int
	workingPercentage float64
} {
	// Build a set of accepted ride IDs
	accSet := make(map[int]bool)
	for _, a := range accepted {
		accSet[a.rideID] = true
	}

	// Map ride ID -> month and month -> set of working drivers
	rideMonth := make(map[int]int)
	for _, r := range rides {
		t, _ := time.Parse("2006-1-2", r.reqDate)
		if t.Year() == 2020 {
			rideMonth[r.id] = int(t.Month())
		}
	}

	workingDrivers := make([]map[int]bool, 13)
	for i := range workingDrivers {
		workingDrivers[i] = make(map[int]bool)
	}
	for rideID, month := range rideMonth {
		if accSet[rideID] {
			for _, a := range accepted {
				if a.rideID == rideID {
					workingDrivers[month][a.driverID] = true
				}
			}
		}
	}

	// Count available drivers by end of each month
	available := make([]int, 13)
	for _, d := range drivers {
		t, _ := time.Parse("2006-1-2", d.joinDate)
		joinYear, joinMonth := t.Year(), t.Month()
		if joinYear < 2020 {
			for m := 1; m <= 12; m++ {
				available[m]++
			}
		} else if joinYear == 2020 {
			for m := int(joinMonth); m <= 12; m++ {
				available[m]++
			}
		}
	}

	result := make([]struct {
		month             int
		workingPercentage float64
	}, 12)
	for m := 1; m <= 12; m++ {
		pct := 0.0
		if available[m] > 0 {
			pct = float64(len(workingDrivers[m])) / float64(available[m]) * 100.0
		}
		result[m-1] = struct {
			month             int
			workingPercentage float64
		}{m, pct}
	}
	return result
}
```

## 1649 — Create Sorted Array Through Instructions

```go
package main

// LeetCode #1649: Create Sorted Array through Instructions
// https://leetcode.com/problems/create-sorted-array-through-instructions/
// Difficulty: Hard

import "fmt"

const MOD1649 = 1000000007

// Fenwick Tree (Binary Indexed Tree)
type BIT struct {
	tree []int
	n    int
}

func NewBIT(n int) *BIT {
	return &BIT{tree: make([]int, n+1), n: n}
}

func (b *BIT) Add(idx, val int) {
	for i := idx; i <= b.n; i += i & -i {
		b.tree[i] += val
	}
}

func (b *BIT) Sum(idx int) int {
	if idx <= 0 {
		return 0
	}
	if idx > b.n {
		idx = b.n
	}
	s := 0
	for i := idx; i > 0; i -= i & -i {
		s += b.tree[i]
	}
	return s
}

func createSortedArray(instructions []int) int {
	// Find maximum value to size the BIT
	maxVal := 0
	for _, v := range instructions {
		if v > maxVal {
			maxVal = v
		}
	}

	bit := NewBIT(maxVal)
	totalCost := 0
	totalInserted := 0

	for _, v := range instructions {
		less := bit.Sum(v - 1)               // elements < v
		total := totalInserted                // total elements so far
		greater := total - bit.Sum(v)         // elements > v

		cost := less
		if greater < cost {
			cost = greater
		}
		totalCost = (totalCost + cost) % MOD1649

		bit.Add(v, 1)
		totalInserted++
	}

	return totalCost
}

func main() {
	// Test case 1: [1,5,2,6,3] -> 3
	// 1: cost 0, 5: cost 0, 2: cost min(1,1)=1, 6: cost 0, 3: cost min(2,2)=2
	instructions := []int{1, 5, 2, 6, 3}
	result := createSortedArray(instructions)
	fmt.Printf("instructions=%v -> %d (expected 3)\n", instructions, result)

	// Test case 2: [1,2,3,4,5] -> 0 (inserted in order, no cost)
	instructions2 := []int{1, 2, 3, 4, 5}
	result2 := createSortedArray(instructions2)
	fmt.Printf("instructions=%v -> %d (expected 0)\n", instructions2, result2)

	// Test case 3: [1,2,1,2,1] -> 0
	// 1: cost 0, 2: cost 0, 1: cost min(0,1)=0, 2: cost min(1,0)=0, 1: cost min(0,2)=0
	instructions3 := []int{1, 2, 1, 2, 1}
	result3 := createSortedArray(instructions3)
	fmt.Printf("instructions=%v -> %d (expected 0)\n", instructions3, result3)

	// Test case 4: [5,4,3,2,1] -> 0 (all descending, always insert at rightmost)
	instructions4 := []int{5, 4, 3, 2, 1}
	result4 := createSortedArray(instructions4)
	fmt.Printf("instructions=%v -> %d\n", instructions4, result4)
}
```

## 1651 — Hopper Company Queries Iii

```go
package main

// LeetCode #1651: Hopper Company Queries III
// https://leetcode.com/problems/hopper-company-queries-iii/
// Difficulty: Hard [Paid] (SQL)
//
// For each 3-month window (Jan-Mar, Feb-Apr, ..., Oct-Dec) of 2020,
// compute the average ride distance and average ride duration.
//
// Approach: Aggregate ride stats by month, then compute rolling
// 3-month averages.

import (
	"fmt"
	"time"
)

func main() {
	// Example data
	rides := []struct {
		rideID   int
		reqDate  string
	}{
		{6, "2019-12-9"},
		{1, "2020-1-2"},
		{10, "2020-1-11"},
		{11, "2020-1-19"},
		{12, "2020-1-29"},
		{3, "2020-1-17"},
		{9, "2020-2-13"},
		{2, "2020-2-22"},
	}
	accepted := []struct {
		rideID  int
		dist    int
		dur     int
	}{
		{2, 63, 38},
		{13, 58, 0},
		{7, 51, 23},
		{3, 50, 36},
		{4, 58, 19},
		{11, 53, 26},
		{1, 50, 42},
		{5, 55, 42},
		{12, 51, 23},
		{6, 56, 41},
	}

	fmt.Println(hopperQueriesIII(rides, accepted))
}

func hopperQueriesIII(rides []struct {
	rideID  int
	reqDate string
}, accepted []struct {
	rideID int
	dist   int
	dur    int
}) []struct {
	month                 int
	avgRideDistance       float64
	avgRideDuration       float64
} {
	// Map ride ID -> month (only 2020)
	rideMonth := make(map[int]int)
	rideData := make(map[int]struct{ dist, dur int })
	for _, r := range rides {
		t, _ := time.Parse("2006-1-2", r.reqDate)
		if t.Year() == 2020 {
			rideMonth[r.rideID] = int(t.Month())
		}
	}
	for _, a := range accepted {
		rideData[a.rideID] = struct{ dist, dur int }{a.dist, a.dur}
	}

	// Monthly totals
	monthDist := make([]int, 13)
	monthDur := make([]int, 13)
	monthCount := make([]int, 13)

	for rideID, m := range rideMonth {
		if d, ok := rideData[rideID]; ok {
			monthDist[m] += d.dist
			monthDur[m] += d.dur
			monthCount[m]++
		}
	}

	result := make([]struct {
		month                 int
		avgRideDistance       float64
		avgRideDuration       float64
	}, 10)

	for start := 1; start <= 10; start++ {
		totalDist := 0
		totalDur := 0
		for m := start; m <= start+2; m++ {
			totalDist += monthDist[m]
			totalDur += monthDur[m]
		}
		result[start-1] = struct {
			month                 int
			avgRideDistance       float64
			avgRideDuration       float64
		}{start, float64(totalDist) / 3.0, float64(totalDur) / 3.0}
	}

	return result
}
```

## 1655 — Distribute Repeating Integers

```go
package main

import (
	"fmt"
	"math"
	"sort"
)

// LeetCode #1655: Distribute Repeating Integers
// https://leetcode.com/problems/distribute-repeating-integers/
// Difficulty: Hard
//
// We have an array nums of integers (may have duplicates) and an array quantity
// where quantity[j] is the number of items that customer j wants to order.
// Each customer must receive items of the SAME value (all identical integers).
// Determine if it's possible to satisfy all customers.
//
// Approach:
// 1. Count frequencies of each distinct number in nums.
// 2. Sort frequencies descending (larger counts first) for pruning.
// 3. Sort quantity descending (larger orders first) for better pruning.
// 4. Use DP with bitmask: can[mask] = true if the current set of frequencies
//    can satisfy the subset of customers represented by mask.
// 5. Iterate over frequencies; for each, update the DP.

func canDistribute(nums []int, quantity []int) bool {
	// Count frequencies
	freqMap := make(map[int]int)
	for _, v := range nums {
		freqMap[v]++
	}

	freqs := make([]int, 0, len(freqMap))
	for _, f := range freqMap {
		freqs = append(freqs, f)
	}

	// Sort frequencies descending for better pruning
	sort.Slice(freqs, func(i, j int) bool {
		return freqs[i] > freqs[j]
	})

	// Sort quantity descending
	sort.Slice(quantity, func(i, j int) bool {
		return quantity[i] > quantity[j]
	})

	n := len(quantity)
	size := 1 << n

	// Precompute subset sums of quantity
	subsetSum := make([]int, size)
	for mask := 1; mask < size; mask++ {
		lsb := mask & -mask
		bit := int(math.Log2(float64(lsb)))
		subsetSum[mask] = subsetSum[mask^lsb] + quantity[bit]
	}

	// dp[mask] = true if we can satisfy customer subset 'mask' with processed frequencies
	dp := make([]bool, size)
	dp[0] = true

	for _, freq := range freqs {
		// For each mask that is currently achievable, try adding this frequency
		// to cover additional customers.
		// We need to iterate backwards to avoid using the same frequency multiple times.
		for mask := size - 1; mask >= 0; mask-- {
			if !dp[mask] {
				continue
			}
			// Find the complement (customers not yet served)
			remaining := (size - 1) ^ mask
			// Try all subsets of remaining customers
			sub := remaining
			for sub > 0 {
				if subsetSum[sub] <= freq {
					dp[mask|sub] = true
				}
				sub = (sub - 1) & remaining
			}
		}

		if dp[size-1] {
			return true
		}
	}

	return dp[size-1]
}

func main() {
	// Example 1:
	// Input: nums = [1,2,3,4], quantity = [2]
	// Output: false — we have 4 distinct numbers with frequency 1 each,
	// but customer needs 2 items of the same value.
	fmt.Println(canDistribute([]int{1, 2, 3, 4}, []int{2}))

	// Example 2:
	// Input: nums = [1,2,3,3], quantity = [2]
	// Output: true — we have two 3's, give them both to the customer.
	fmt.Println(canDistribute([]int{1, 2, 3, 3}, []int{2}))

	// Example 3:
	// Input: nums = [1,1,2,2], quantity = [2,2]
	// Output: true — give 1,1 to one customer, 2,2 to the other.
	fmt.Println(canDistribute([]int{1, 1, 2, 2}, []int{2, 2}))

	// Example 4:
	// Input: nums = [1,1,2,3], quantity = [2,2]
	// Output: false — only one pair (1,1), need two pairs.
	fmt.Println(canDistribute([]int{1, 1, 2, 3}, []int{2, 2}))

	// Example 5:
	// Input: nums = [1,1,1,1,1], quantity = [2,3]
	// Output: true — give 2 of the 1's to one customer, 3 to the other.
	fmt.Println(canDistribute([]int{1, 1, 1, 1, 1}, []int{2, 3}))
}
```

## 1659 — Maximize Grid Happiness

```go
package main

import (
	"fmt"
	"math"
)

// LeetCode #1659: Maximize Grid Happiness
// https://leetcode.com/problems/maximize-grid-happiness/
// Difficulty: Hard
//
// Place introverts (I) and extroverts (E) in an m x n grid.
// - Introvert happiness = 120 (base) - 30 per neighbor (any I or E nearby)
// - Extrovert happiness = 40 (base) + 20 per neighbor (any I or E nearby)
// - Empty cell = 0
// Maximize total happiness with at most introvertsCount introverts
// and extrovertsCount extroverts.
//
// DP with state compression: process row by row. For each row, we track
// the placement of the previous row as a 3-base (ternary) number where
// 0 = empty, 1 = introvert, 2 = extrovert. n <= 5 so 3^5 = 243 states.
//
// dp[row][maskPrev][i][e] = max happiness up to row `row` when previous row
// is `maskPrev`, with i introverts and e extroverts used so far.

func getMaxGridHappiness(m int, n int, introvertsCount int, extrovertsCount int) int {
	// Total states for one row: 3^n
	states := int(math.Pow(3, float64(n)))

	// Precompute for each state:
	// - rowScore: total happiness within the row
	// - intro count, extro count
	rowScore := make([]int, states)
	introCount := make([]int, states)
	extroCount := make([]int, states)

	for s := 0; s < states; s++ {
		prev := -1
		tmp := s
		inner := 0
		introCnt := 0
		extroCnt := 0
		mask := make([]int, n)
		for pos := 0; pos < n; pos++ {
			cell := tmp % 3
			tmp /= 3
			mask[pos] = cell
			if cell == 1 {
				introCnt++
				inner += 120
			} else if cell == 2 {
				extroCnt++
				inner += 40
			}
			// Horizontal interaction with left neighbor
			if prev != -1 {
				if prev == 1 && cell == 1 {
					inner -= 60 // I-I
				} else if prev == 1 && cell == 2 {
					inner -= 10 // I-E (I loses 30, E gains 20)
				} else if prev == 2 && cell == 1 {
					inner -= 10 // E-I (same)
				} else if prev == 2 && cell == 2 {
					inner += 40 // E-E (each gains 20)
				}
			}
			prev = cell
		}
		rowScore[s] = inner
		introCount[s] = introCnt
		extroCount[s] = extroCnt

		// Precompute vertical interaction score with the same state (placeholder)
		// Will compute actual inter-row score on the fly.
	}

	// Precompute vertical interaction between two states (top row and bottom row)
	vertBetween := make([][]int, states)
	for s1 := 0; s1 < states; s1++ {
		vertBetween[s1] = make([]int, states)
		for s2 := 0; s2 < states; s2++ {
			score := 0
			tmp1, tmp2 := s1, s2
			for pos := 0; pos < n; pos++ {
				top := tmp1 % 3
				bottom := tmp2 % 3
				tmp1 /= 3
				tmp2 /= 3
				if top == 1 && bottom == 1 {
					score -= 60
				} else if top == 1 && bottom == 2 {
					score -= 10
				} else if top == 2 && bottom == 1 {
					score -= 10
				} else if top == 2 && bottom == 2 {
					score += 40
				}
			}
			vertBetween[s1][s2] = score
		}
	}

	// dp[mask][i][e] = max happiness for processed rows with previous row = mask,
	// i introverts used, e extroverts used
	dp := make([][][]int, states)
	for s := 0; s < states; s++ {
		dp[s] = make([][]int, introvertsCount+1)
		for i := 0; i <= introvertsCount; i++ {
			dp[s][i] = make([]int, extrovertsCount+1)
			for e := 0; e <= extrovertsCount; e++ {
				dp[s][i][e] = math.MinInt32
			}
		}
	}
	dp[0][0][0] = 0 // mask 0 = all empty

	for r := 0; r < m; r++ {
		ndp := make([][][]int, states)
		for s := 0; s < states; s++ {
			ndp[s] = make([][]int, introvertsCount+1)
			for i := 0; i <= introvertsCount; i++ {
				ndp[s][i] = make([]int, extrovertsCount+1)
				for e := 0; e <= extrovertsCount; e++ {
					ndp[s][i][e] = math.MinInt32
				}
			}
		}

		for prevMask := 0; prevMask < states; prevMask++ {
			for i := 0; i <= introvertsCount; i++ {
				for e := 0; e <= extrovertsCount; e++ {
					cur := dp[prevMask][i][e]
					if cur == math.MinInt32 {
						continue
					}
					// Try every possible state for the current row
					for curMask := 0; curMask < states; curMask++ {
						ni := i + introCount[curMask]
						ne := e + extroCount[curMask]
						if ni > introvertsCount || ne > extrovertsCount {
							continue
						}
						score := cur + rowScore[curMask] + vertBetween[prevMask][curMask]
						if score > ndp[curMask][ni][ne] {
							ndp[curMask][ni][ne] = score
						}
					}
				}
			}
		}
		dp = ndp
	}

	ans := 0
	for mask := 0; mask < states; mask++ {
		for i := 0; i <= introvertsCount; i++ {
			for e := 0; e <= extrovertsCount; e++ {
				if dp[mask][i][e] > ans {
					ans = dp[mask][i][e]
				}
			}
		}
	}
	return ans
}

func main() {
	// Example 1:
	// Input: m = 2, n = 3, introvertsCount = 1, extrovertsCount = 2
	// Output: 240
	fmt.Println(getMaxGridHappiness(2, 3, 1, 2))

	// Example 2:
	// Input: m = 3, n = 1, introvertsCount = 2, extrovertsCount = 1
	// Output: 260
	fmt.Println(getMaxGridHappiness(3, 1, 2, 1))

	// Example 3:
	// Input: m = 2, n = 2, introvertsCount = 4, extrovertsCount = 0
	// Output: 240
	fmt.Println(getMaxGridHappiness(2, 2, 4, 0))
}
```

## 1665 — Minimum Initial Energy To Finish Tasks

```go
package main

// LeetCode #1665: Minimum Initial Energy to Finish Tasks
// https://leetcode.com/problems/minimum-initial-energy-to-finish-tasks/
// Difficulty: Hard

import (
	"fmt"
	"sort"
)

func main() {
	tasks1 := [][]int{{1, 2}, {2, 4}, {4, 8}}
	fmt.Printf("Test 1 - Input: [[1,2],[2,4],[4,8]]\nExpected: 8\nGot: %d\n\n", minimumEffort(tasks1))

	tasks2 := [][]int{{1, 3}, {2, 4}, {10, 11}, {10, 12}, {8, 9}}
	fmt.Printf("Test 2 - Input: [[1,3],[2,4],[10,11],[10,12],[8,9]]\nExpected: 32\nGot: %d\n\n", minimumEffort(tasks2))

	tasks3 := [][]int{{1, 1}, {1, 1}}
	fmt.Printf("Test 3 - Input: [[1,1],[1,1]]\nExpected: 1\nGot: %d\n", minimumEffort(tasks3))
}

func minimumEffort(tasks [][]int) int {
	// Sort by (minimum - actual) descending: tasks with the largest energy
	// deficit (minimum required vs actual consumed) should be done first.
	sort.Slice(tasks, func(i, j int) bool {
		return (tasks[i][1] - tasks[i][0]) > (tasks[j][1] - tasks[j][0])
	})

	result := 0
	curr := 0

	for _, t := range tasks {
		actual, minimum := t[0], t[1]
		if curr < minimum {
			result += minimum - curr
			curr = minimum
		}
		curr -= actual
	}

	return result
}
```

## 1666 — Change The Root Of A Binary Tree

```go
package main

// LeetCode #1666: Change the Root of a Binary Tree
// https://leetcode.com/problems/change-the-root-of-a-binary-tree/
// Difficulty: Hard [Premium]

import "fmt"

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

func main() {
	// Build tree: root=3
	//     3
	//    / \
	//   5   1
	//  / \  / \
	// 6   2 0  8
	//    / \
	//   7   4
	root := &Node{Val: 3}
	n5 := &Node{Val: 5}
	n1 := &Node{Val: 1}
	n6 := &Node{Val: 6}
	n2 := &Node{Val: 2}
	n7 := &Node{Val: 7}
	n4 := &Node{Val: 4}
	n0 := &Node{Val: 0}
	n8 := &Node{Val: 8}

	root.Left, root.Right = n5, n1
	n5.Parent, n1.Parent = root, root
	n5.Left, n5.Right = n6, n2
	n6.Parent, n2.Parent = n5, n5
	n2.Left, n2.Right = n7, n4
	n7.Parent, n4.Parent = n2, n2
	n1.Left, n1.Right = n0, n8
	n0.Parent, n8.Parent = n1, n1

	// Change root to node 2
	newRoot := flipBinaryTree(root, n2)
	fmt.Printf("Test 1 - New root value: %d (Expected: 2)\n", newRoot.Val)
	fmt.Printf("Right child (was old parent 5): %d (Expected: 5)\n", newRoot.Right.Val)
	fmt.Printf("Left child of 5 (was old parent 3): %d (Expected: 3)\n", n5.Left.Val)
	fmt.Printf("Parent of new root is nil: %v (Expected: true)\n", newRoot.Parent == nil)
	fmt.Printf("Parent of old root is child: %d (Expected: 5)\n", root.Parent.Val)
}

func flipBinaryTree(root *Node, leaf *Node) *Node {
	// Walk from leaf up to root, reversing parent-child relationships
	return flip(leaf, nil)
}

func flip(node, newParent *Node) *Node {
	oldParent := node.Parent
	node.Parent = newParent

	// Disconnect the child pointer that now points to the new parent
	if node.Left == newParent {
		node.Left = nil
	} else if node.Right == newParent {
		node.Right = nil
	}

	if oldParent != nil {
		// Disconnect old parent's pointer to this node
		if oldParent.Left == node {
			oldParent.Left = nil
		} else if oldParent.Right == node {
			oldParent.Right = nil
		}
		flip(oldParent, node)
		// Attach old parent as a child of current node
		if node.Left == nil {
			node.Left = oldParent
		} else {
			node.Right = oldParent
		}
	}

	return node
}
```

## 1671 — Minimum Number Of Removals To Make Mountain Array

```go
package main

// LeetCode #1671: Minimum Number of Removals to Make Mountain Array
// https://leetcode.com/problems/minimum-number-of-removals-to-make-mountain-array/
// Difficulty: Hard

import "fmt"

func main() {
	nums1 := []int{1, 3, 1}
	fmt.Printf("Test 1 - Input: %v\nExpected: 0\nGot: %d\n\n", nums1, minimumMountainRemovals(nums1))

	nums2 := []int{2, 1, 1, 5, 6, 2, 3, 1}
	fmt.Printf("Test 2 - Input: %v\nExpected: 3\nGot: %d\n\n", nums2, minimumMountainRemovals(nums2))

	nums3 := []int{4, 3, 2, 1, 1, 2, 3, 1}
	fmt.Printf("Test 3 - Input: %v\nExpected: 4\nGot: %d\n", nums3, minimumMountainRemovals(nums3))
}

func minimumMountainRemovals(nums []int) int {
	n := len(nums)
	lis := make([]int, n) // longest increasing subsequence ending at i
	lds := make([]int, n) // longest decreasing subsequence starting at i

	// Compute LIS from left
	for i := 0; i < n; i++ {
		lis[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && lis[j]+1 > lis[i] {
				lis[i] = lis[j] + 1
			}
		}
	}

	// Compute LDS from right (reverse LIS)
	for i := n - 1; i >= 0; i-- {
		lds[i] = 1
		for j := n - 1; j > i; j-- {
			if nums[j] < nums[i] && lds[j]+1 > lds[i] {
				lds[i] = lds[j] + 1
			}
		}
	}

	// Find the longest bitonic subsequence (mountain)
	maxMountain := 0
	for i := 1; i < n-1; i++ {
		if lis[i] > 1 && lds[i] > 1 {
			mountainLen := lis[i] + lds[i] - 1
			if mountainLen > maxMountain {
				maxMountain = mountainLen
			}
		}
	}

	return n - maxMountain
}
```

## 1675 — Minimize Deviation In Array

```go
package main

// LeetCode #1675: Minimize Deviation in Array
// https://leetcode.com/problems/minimize-deviation-in-array/
// Difficulty: Hard
// Strategy: Max-heap. Multiply all odds by 2 (maximize), then reduce max even by /2.

import (
	"container/heap"
	"fmt"
)

// MaxHeap for ints
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func minimumDeviation(nums []int) int {
	h := &MaxHeap{}
	heap.Init(h)
	minVal := int(1 << 60)

	// Step 1: multiply all odd numbers by 2 (they can only increase),
	// track the minimum value
	for _, num := range nums {
		if num%2 == 1 {
			num *= 2
		}
		if num < minVal {
			minVal = num
		}
		heap.Push(h, num)
	}

	ans := int(1 << 60)
	for {
		maxVal := heap.Pop(h).(int)
		ans = min(ans, maxVal-minVal)
		if maxVal%2 == 1 {
			// max is odd -> can't be reduced further
			break
		}
		maxVal /= 2
		if maxVal < minVal {
			minVal = maxVal
		}
		heap.Push(h, maxVal)
	}
	return ans
}

func main() {
	// Example 1: [1,2,3,4] -> 1
	fmt.Printf("minimumDeviation([1,2,3,4]) = %d (expected 1)\n", minimumDeviation([]int{1, 2, 3, 4}))

	// Example 2: [4,1,5,20,3] -> 3
	fmt.Printf("minimumDeviation([4,1,5,20,3]) = %d (expected 3)\n", minimumDeviation([]int{4, 1, 5, 20, 3}))

	// Additional test: [3,5] -> 1
	fmt.Printf("minimumDeviation([3,5]) = %d (expected 1)\n", minimumDeviation([]int{3, 5}))

	// Single element
	fmt.Printf("minimumDeviation([10]) = %d (expected 0)\n", minimumDeviation([]int{10}))

	// All odds (correct answer is 3 not 2)
	fmt.Printf("minimumDeviation([1,3,5]) = %d (expected 3)\n", minimumDeviation([]int{1, 3, 5}))
}
```

## 1681 — Minimum Incompatibility

```go
package main

// LeetCode #1681: Minimum Incompatibility
// https://leetcode.com/problems/minimum-incompatibility/
// Difficulty: Hard
// Strategy: DP over bitmask. Precompute valid subsets of size n/k.

import (
	"fmt"
	"math"
)

func minimumIncompatibility(nums []int, k int) int {
	n := len(nums)
	subsetSize := n / k
	if subsetSize == 1 {
		return 0
	}

	// Precompute incompatibility for each valid subset of size subsetSize
	incomp := make([]int, 1<<n)
	for i := range incomp {
		incomp[i] = -1
	}
	for mask := 1; mask < 1<<n; mask++ {
		if bitsCount(mask) != subsetSize {
			continue
		}
		// Check for duplicates and find min/max
		seen := make(map[int]bool)
		minVal, maxVal := math.MaxInt32, math.MinInt32
		valid := true
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				if seen[nums[i]] {
					valid = false
					break
				}
				seen[nums[i]] = true
				if nums[i] < minVal {
					minVal = nums[i]
				}
				if nums[i] > maxVal {
					maxVal = nums[i]
				}
			}
		}
		if valid {
			incomp[mask] = maxVal - minVal
		}
	}

	// DP: dp[mask] = min incompatibility for chosen elements in mask
	dp := make([]int, 1<<n)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0

	for mask := 0; mask < 1<<n; mask++ {
		if dp[mask] == -1 {
			continue
		}
		// Remaining elements
		remaining := ((1 << n) - 1) ^ mask
		if remaining == 0 {
			continue
		}
		// Pick a valid subset from remaining
		sub := remaining
		for sub > 0 {
			if incomp[sub] != -1 {
				newMask := mask | sub
				newVal := dp[mask] + incomp[sub]
				if dp[newMask] == -1 || newVal < dp[newMask] {
					dp[newMask] = newVal
				}
			}
			sub = (sub - 1) & remaining
		}
	}

	return dp[(1<<n)-1]
}

func bitsCount(x int) int {
	count := 0
	for x > 0 {
		count += x & 1
		x >>= 1
	}
	return count
}

func main() {
	// Example 1: [1,2,1,4], k=2 -> 4
	nums1 := []int{1, 2, 1, 4}
	k1 := 2
	fmt.Printf("minimumIncompatibility(%v, %d) = %d (expected 4)\n", nums1, k1, minimumIncompatibility(nums1, k1))

	// Example 2: [6,3,8,1,3,1,2,2], k=4 -> 6
	nums2 := []int{6, 3, 8, 1, 3, 1, 2, 2}
	k2 := 4
	fmt.Printf("minimumIncompatibility(%v, %d) = %d (expected 6)\n", nums2, k2, minimumIncompatibility(nums2, k2))

	// Example 3: [5,3,3,6,3,3], k=3 -> -1
	nums3 := []int{5, 3, 3, 6, 3, 3}
	k3 := 3
	fmt.Printf("minimumIncompatibility(%v, %d) = %d (expected -1)\n", nums3, k3, minimumIncompatibility(nums3, k3))

	// Single subset size
	nums4 := []int{1, 2, 3, 4}
	k4 := 4
	fmt.Printf("minimumIncompatibility(%v, %d) = %d (expected 0)\n", nums4, k4, minimumIncompatibility(nums4, k4))
}
```

## 1687 — Delivering Boxes From Storage To Ports

```go
package main

// LeetCode #1687: Delivering Boxes from Storage to Ports
// https://leetcode.com/problems/delivering-boxes-from-storage-to-ports/
// Difficulty: Hard
// Strategy: DP with deque optimization (sliding window min).

import (
	"fmt"
)

func boxDelivering(boxes [][]int, portsCount int, maxBoxes int, maxWeight int) int {
	n := len(boxes)

	// diffTrips[i] = number of port changes between box i-1 and box i
	diffTrips := make([]int, n+2)
	for i := 1; i < n; i++ {
		if boxes[i][0] != boxes[i-1][0] {
			diffTrips[i] = 1
		}
	}
	// prefix sum of diffTrips (size n+2 to allow safe access at index n+1)
	prefDiff := make([]int, n+2)
	for i := 1; i <= n; i++ {
		prefDiff[i] = prefDiff[i-1] + diffTrips[i-1]
	}
	prefDiff[n+1] = prefDiff[n] // diffTrips[n] is always 0

	// prefix sum of weights
	prefW := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefW[i] = prefW[i-1] + boxes[i-1][1]
	}

	// DP: dp[i] = min trips to deliver first i boxes
	dp := make([]int, n+1)
	// deque stores indices j, maintaining dp[j] - prefDiff[j+1] in increasing order
	deque := make([]int, 0, n+1)
	deque = append(deque, 0)

	for i := 1; i <= n; i++ {
		// Remove boxes that exceed maxBoxes or maxWeight from front
		for len(deque) > 0 {
			j := deque[0]
			if i-j > maxBoxes || prefW[i]-prefW[j] > maxWeight {
				deque = deque[1:]
			} else {
				break
			}
		}

		// dp[i] = dp[deque[0]] + prefDiff[i] - prefDiff[deque[0]+1] + 2
		j := deque[0]
		// Each trip: 1 for storage->first port + 1 for last port->storage = 2,
		// plus port changes between boxes
		dp[i] = dp[j] + prefDiff[i] - prefDiff[j+1] + 2

		// Insert i into deque
		val := dp[i] - prefDiff[i+1]
		for len(deque) > 0 {
			last := deque[len(deque)-1]
			if dp[last]-prefDiff[last+1] >= val {
				deque = deque[:len(deque)-1]
			} else {
				break
			}
		}
		deque = append(deque, i)
	}

	return dp[n]
}

func main() {
	// Example 1
	boxes1 := [][]int{{1, 1}, {2, 1}, {1, 1}}
	portsCount1 := 2
	maxBoxes1 := 3
	maxWeight1 := 3
	fmt.Printf("boxDelivering(%v, %d, %d, %d) = %d (expected 4)\n",
		boxes1, portsCount1, maxBoxes1, maxWeight1,
		boxDelivering(boxes1, portsCount1, maxBoxes1, maxWeight1))

	// Example 2
	boxes2 := [][]int{{1, 2}, {3, 3}, {3, 1}, {3, 1}, {2, 4}}
	portsCount2 := 3
	maxBoxes2 := 3
	maxWeight2 := 6
	fmt.Printf("boxDelivering(%v, %d, %d, %d) = %d (expected 6)\n",
		boxes2, portsCount2, maxBoxes2, maxWeight2,
		boxDelivering(boxes2, portsCount2, maxBoxes2, maxWeight2))

	// Example 3
	boxes3 := [][]int{{1, 4}, {1, 2}, {2, 1}, {2, 1}, {3, 2}, {3, 4}}
	portsCount3 := 3
	maxBoxes3 := 6
	maxWeight3 := 7
	fmt.Printf("boxDelivering(%v, %d, %d, %d) = %d (expected 6)\n",
		boxes3, portsCount3, maxBoxes3, maxWeight3,
		boxDelivering(boxes3, portsCount3, maxBoxes3, maxWeight3))

	// Example 4
	boxes4 := [][]int{{2, 4}, {2, 5}, {3, 1}, {3, 2}, {3, 7}, {3, 1}, {4, 4}, {1, 3}, {5, 2}}
	portsCount4 := 5
	maxBoxes4 := 5
	maxWeight4 := 7
	fmt.Printf("boxDelivering(%v, %d, %d, %d) = %d (expected 14)\n",
		boxes4, portsCount4, maxBoxes4, maxWeight4,
		boxDelivering(boxes4, portsCount4, maxBoxes4, maxWeight4))
}
```

## 1691 — Maximum Height By Stacking Cuboids

```go
package main

// LeetCode #1691: Maximum Height by Stacking Cuboids
// https://leetcode.com/problems/maximum-height-by-stacking-cuboids/
// Difficulty: Hard
// Strategy: Sort each cuboid so height is the max dimension,
// then sort all cuboids and run LIS (DP).

import (
	"fmt"
	"sort"
)

func maxHeight(cuboids [][]int) int {
	// For each cuboid, sort dimensions so the largest is height
	for _, c := range cuboids {
		sort.Ints(c)
	}
	// Sort cuboids by dimensions (width, depth, height)
	sort.Slice(cuboids, func(i, j int) bool {
		if cuboids[i][0] != cuboids[j][0] {
			return cuboids[i][0] < cuboids[j][0]
		}
		if cuboids[i][1] != cuboids[j][1] {
			return cuboids[i][1] < cuboids[j][1]
		}
		return cuboids[i][2] < cuboids[j][2]
	})

	n := len(cuboids)
	dp := make([]int, n)
	ans := 0
	for i := 0; i < n; i++ {
		dp[i] = cuboids[i][2] // height
		for j := 0; j < i; j++ {
			// Check if cuboid j can be placed below cuboid i
			if cuboids[j][0] <= cuboids[i][0] &&
				cuboids[j][1] <= cuboids[i][1] &&
				cuboids[j][2] <= cuboids[i][2] {
				if dp[j]+cuboids[i][2] > dp[i] {
					dp[i] = dp[j] + cuboids[i][2]
				}
			}
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}

func main() {
	// Example 1: [[50,45,20],[95,37,53],[45,23,12]] -> 190
	cuboids1 := [][]int{{50, 45, 20}, {95, 37, 53}, {45, 23, 12}}
	fmt.Printf("maxHeight(%v) = %d (expected 190)\n", cuboids1, maxHeight(cuboids1))

	// Example 2: [[38,25,45],[76,35,3]] -> 76
	cuboids2 := [][]int{{38, 25, 45}, {76, 35, 3}}
	fmt.Printf("maxHeight(%v) = %d (expected 76)\n", cuboids2, maxHeight(cuboids2))

	// Example 3: [[7,11,17],[7,17,11],[11,7,17],[11,17,7],[17,7,11],[17,11,7]] -> 102
	cuboids3 := [][]int{{7, 11, 17}, {7, 17, 11}, {11, 7, 17}, {11, 17, 7}, {17, 7, 11}, {17, 11, 7}}
	fmt.Printf("maxHeight(%v) = %d (expected 102)\n", cuboids3, maxHeight(cuboids3))

	// Single cuboid
	cuboids4 := [][]int{{1, 2, 3}}
	fmt.Printf("maxHeight(%v) = %d (expected 3)\n", cuboids4, maxHeight(cuboids4))
}
```

## 1692 — Count Ways To Distribute Candies

```go
package main

// LeetCode #1692: Count Ways to Distribute Candies
// https://leetcode.com/problems/count-ways-to-distribute-candies/
// Difficulty: Hard [Paid]
//
// Given n distinct candies and k bags, count the number of ways to
// distribute all candies into exactly k non-empty bags.
// Two ways are different if at least one candy goes to a different bag.
//
// Approach: Stirling numbers of the second kind.
// dp[i][j] = ways to distribute i candies into j bags.
// Transition: dp[i][j] = j * dp[i-1][j] + dp[i-1][j-1].

import "fmt"

func main() {
	// Example 1
	fmt.Println(waysToDistribute(3, 2))
	// Example 2
	fmt.Println(waysToDistribute(4, 3))
	// Edge: k = n
	fmt.Println(waysToDistribute(3, 3))
	// Edge: k = 1
	fmt.Println(waysToDistribute(5, 1))
}

const W2DMOD = 1000000007

func waysToDistribute(n int, k int) int {
	if k > n || k == 0 {
		return 0
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= k && j <= i; j++ {
			dp[i][j] = (j*dp[i-1][j] + dp[i-1][j-1]) % W2DMOD
		}
	}
	return dp[n][k]
}
```

## 1697 — Checking Existence Of Edge Length Limited Paths

```go
package main

// LeetCode #1697: Checking Existence of Edge Length Limited Paths
// https://leetcode.com/problems/checking-existence-of-edge-length-limited-paths/
// Difficulty: Hard
// Strategy: Sort queries by limit, sort edges by weight.
// Union-Find to add edges incrementally as limits increase.

import (
	"fmt"
	"sort"
)

// Union-Find / DSU
type DSU struct {
	parent []int
	rank   []int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &DSU{parent, rank}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) Union(x, y int) {
	xr, yr := d.Find(x), d.Find(y)
	if xr == yr {
		return
	}
	if d.rank[xr] < d.rank[yr] {
		d.parent[xr] = yr
	} else if d.rank[xr] > d.rank[yr] {
		d.parent[yr] = xr
	} else {
		d.parent[yr] = xr
		d.rank[xr]++
	}
}

func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	// Sort edges by weight
	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i][2] < edgeList[j][2]
	})

	// Sort queries by limit, keeping original indices
	q := make([][4]int, len(queries)) // [limit, u, v, originalIdx]
	for i, query := range queries {
		q[i] = [4]int{query[2], query[0], query[1], i}
	}
	sort.Slice(q, func(i, j int) bool {
		return q[i][0] < q[j][0]
	})

	dsu := NewDSU(n)
	ans := make([]bool, len(queries))
	edgeIdx := 0

	for _, query := range q {
		limit, u, v, idx := query[0], query[1], query[2], query[3]
		// Add all edges with weight < limit
		for edgeIdx < len(edgeList) && edgeList[edgeIdx][2] < limit {
			dsu.Union(edgeList[edgeIdx][0], edgeList[edgeIdx][1])
			edgeIdx++
		}
		ans[idx] = dsu.Find(u) == dsu.Find(v)
	}
	return ans
}

func main() {
	// Example 1: n=3, edgeList=[[0,1,2],[1,2,4],[2,0,8],[1,0,16]], queries=[[0,1,2],[0,2,5]] -> [false,true]
	// (edge weight must be strictly less than limit)
	n1 := 3
	edgeList1 := [][]int{{0, 1, 2}, {1, 2, 4}, {2, 0, 8}, {1, 0, 16}}
	queries1 := [][]int{{0, 1, 2}, {0, 2, 5}}
	fmt.Printf("distanceLimitedPathsExist(%d, %v, %v) = %v (expected [false true])\n",
		n1, edgeList1, queries1, distanceLimitedPathsExist(n1, edgeList1, queries1))

	// Example 2: n=5, edgeList=[[0,1,10],[1,2,5],[2,3,9],[3,4,13]], queries=[[0,4,14],[1,4,13]] -> [true,false]
	n2 := 5
	edgeList2 := [][]int{{0, 1, 10}, {1, 2, 5}, {2, 3, 9}, {3, 4, 13}}
	queries2 := [][]int{{0, 4, 14}, {1, 4, 13}}
	fmt.Printf("distanceLimitedPathsExist(%d, %v, %v) = %v (expected [true false])\n",
		n2, edgeList2, queries2, distanceLimitedPathsExist(n2, edgeList2, queries2))

	// Edge case: single query
	n3 := 2
	edgeList3 := [][]int{{0, 1, 5}}
	queries3 := [][]int{{0, 1, 3}}
	fmt.Printf("distanceLimitedPathsExist(%d, %v, %v) = %v (expected [false])\n",
		n3, edgeList3, queries3, distanceLimitedPathsExist(n3, edgeList3, queries3))
}
```

## 1699 — Number Of Calls Between Two Persons

```go
package main

// LeetCode #1699: Number of Calls Between Two Persons
// https://leetcode.com/problems/number-of-calls-between-two-persons/
// Difficulty: Hard [Premium]

import "fmt"

type Call struct {
	FromId   int
	ToId     int
	Duration int
}

type CallPair struct {
	Person1       int
	Person2       int
	CallCount     int
	TotalDuration int
}

func main() {
	calls1 := []Call{
		{1, 2, 59},
		{2, 1, 11},
		{1, 3, 20},
		{3, 1, 10},
	}
	result1 := numberOfCalls(calls1)
	fmt.Println("Test 1 - Call pairs:")
	for _, r := range result1 {
		fmt.Printf("  Person1: %d, Person2: %d, CallCount: %d, TotalDuration: %d\n",
			r.Person1, r.Person2, r.CallCount, r.TotalDuration)
	}
	fmt.Println("Expected: (1,2,2,70) and (1,3,2,30)\n")

	calls2 := []Call{
		{10, 20, 30},
		{20, 10, 40},
	}
	result2 := numberOfCalls(calls2)
	fmt.Println("Test 2 - Call pairs:")
	for _, r := range result2 {
		fmt.Printf("  Person1: %d, Person2: %d, CallCount: %d, TotalDuration: %d\n",
			r.Person1, r.Person2, r.CallCount, r.TotalDuration)
	}
	fmt.Println("Expected: (10,20,2,70)")
}

func numberOfCalls(calls []Call) []CallPair {
	pairMap := make(map[[2]int]*CallPair)

	for _, c := range calls {
		p1, p2 := c.FromId, c.ToId
		if p1 > p2 {
			p1, p2 = p2, p1
		}
		key := [2]int{p1, p2}
		if pair, ok := pairMap[key]; ok {
			pair.CallCount++
			pair.TotalDuration += c.Duration
		} else {
			pairMap[key] = &CallPair{
				Person1:       p1,
				Person2:       p2,
				CallCount:     1,
				TotalDuration: c.Duration,
			}
		}
	}

	result := make([]CallPair, 0, len(pairMap))
	for _, pair := range pairMap {
		result = append(result, *pair)
	}
	return result
}
```

## 1703 — Minimum Adjacent Swaps For K Consecutive Ones

```go
package main

// LeetCode #1703: Minimum Adjacent Swaps for K Consecutive Ones
// https://leetcode.com/problems/minimum-adjacent-swaps-for-k-consecutive-ones/
// Difficulty: Hard
// Strategy: Collect positions of 1s. For k consecutive 1s,
// use median prefix sum to compute min adjacent swaps.

import (
	"fmt"
	"math"
)

func minMoves(nums []int, k int) int {
	// Collect indices of 1s
	pos := make([]int, 0)
	for i, v := range nums {
		if v == 1 {
			pos = append(pos, i)
		}
	}

	n := len(pos)
	if n < k {
		return 0
	}

	// prefix sum of positions
	pref := make([]int, n+1)
	for i := 0; i < n; i++ {
		pref[i+1] = pref[i] + pos[i]
	}

	ans := math.MaxInt32
	for i := 0; i <= n-k; i++ {
		// Window of k positions: pos[i..i+k-1]
		mid := i + k/2
		medianPos := pos[mid]

		// Number of elements on left side of median (in window)
		leftCount := mid - i
		// Number of elements on right side of median (in window)
		rightCount := i + k - 1 - mid

		// Sum of positions on left
		leftSum := pref[mid] - pref[i]
		// Sum of positions on right
		rightSum := pref[i+k] - pref[mid+1]

		// Ideal left positions: medianPos-1, medianPos-2, ...
		leftIdeal := leftCount*medianPos - leftCount*(leftCount+1)/2
		// Actual left sum
		leftActual := leftSum

		// Ideal right positions: medianPos+1, medianPos+2, ...
		rightIdeal := rightCount*medianPos + rightCount*(rightCount+1)/2
		// Actual right sum
		rightActual := rightSum

		swaps := (leftIdeal - leftActual) + (rightActual - rightIdeal)
		if swaps < ans {
			ans = swaps
		}
	}
	return ans
}

func main() {
	// Example 1: [1,0,0,1,0,1], k=2 -> 1
	nums1 := []int{1, 0, 0, 1, 0, 1}
	k1 := 2
	fmt.Printf("minMoves(%v, %d) = %d (expected 1)\n", nums1, k1, minMoves(nums1, k1))

	// Example 2: [1,0,0,0,0,0,1,1], k=3 -> 5
	nums2 := []int{1, 0, 0, 0, 0, 0, 1, 1}
	k2 := 3
	fmt.Printf("minMoves(%v, %d) = %d (expected 5)\n", nums2, k2, minMoves(nums2, k2))

	// Example 3: [1,1,0,1], k=2 -> 0
	nums3 := []int{1, 1, 0, 1}
	k3 := 2
	fmt.Printf("minMoves(%v, %d) = %d (expected 0)\n", nums3, k3, minMoves(nums3, k3))

	// Example from problem description: [1,0,0,1,0,1], k=2 -> 1 (same as example 1)
}
```

## 1707 — Maximum Xor With An Element From Array

```go
package main

// LeetCode #1707: Maximum XOR With an Element From Array
// https://leetcode.com/problems/maximum-xor-with-an-element-from-array/
// Difficulty: Hard
// Strategy: Offline Trie. Sort nums, sort queries by limit.
// Insert nums into trie as limit increases, then query max XOR.

import (
	"fmt"
	"sort"
)

// Binary Trie Node (bits 0..31 for ints up to 10^9)
type TrieNode struct {
	children [2]*TrieNode
}

type BinaryTrie struct {
	root *TrieNode
}

func NewBinaryTrie() *BinaryTrie {
	return &BinaryTrie{root: &TrieNode{}}
}

func (t *BinaryTrie) Insert(num int) {
	node := t.root
	for i := 31; i >= 0; i-- {
		bit := (num >> i) & 1
		if node.children[bit] == nil {
			node.children[bit] = &TrieNode{}
		}
		node = node.children[bit]
	}
}

func (t *BinaryTrie) QueryMaxXor(num int) int {
	// Returns max XOR value, -1 if trie is empty
	if t.root.children[0] == nil && t.root.children[1] == nil {
		return -1
	}
	node := t.root
	xor := 0
	for i := 31; i >= 0; i-- {
		bit := (num >> i) & 1
		// Try to go opposite direction for max XOR
		desired := 1 - bit
		if node.children[desired] != nil {
			xor |= (1 << i)
			node = node.children[desired]
		} else {
			node = node.children[bit]
		}
	}
	return xor
}

func maximizeXor(nums []int, queries [][]int) []int {
	// Sort nums
	sort.Ints(nums)

	// Attach original indices to queries and sort by limit
	q := make([][3]int, len(queries)) // [x, limit, originalIdx]
	for i, query := range queries {
		q[i] = [3]int{query[0], query[1], i}
	}
	sort.Slice(q, func(i, j int) bool {
		return q[i][1] < q[j][1]
	})

	trie := NewBinaryTrie()
	ans := make([]int, len(queries))
	idx := 0

	for _, query := range q {
		x, limit, origIdx := query[0], query[1], query[2]
		// Insert all nums <= limit
		for idx < len(nums) && nums[idx] <= limit {
			trie.Insert(nums[idx])
			idx++
		}
		ans[origIdx] = trie.QueryMaxXor(x)
	}
	return ans
}

func main() {
	// Example 1: nums=[0,1,2,3,4], queries=[[3,1],[1,3],[5,6]] -> [3,3,7]
	nums1 := []int{0, 1, 2, 3, 4}
	queries1 := [][]int{{3, 1}, {1, 3}, {5, 6}}
	fmt.Printf("maximizeXor(%v, %v) = %v (expected [3 3 7])\n",
		nums1, queries1, maximizeXor(nums1, queries1))

	// Example 2: nums=[5,2,4,6,6,3], queries=[[12,4],[8,1],[6,3]] -> [15,-1,5]
	nums2 := []int{5, 2, 4, 6, 6, 3}
	queries2 := [][]int{{12, 4}, {8, 1}, {6, 3}}
	fmt.Printf("maximizeXor(%v, %v) = %v (expected [15 -1 5])\n",
		nums2, queries2, maximizeXor(nums2, queries2))

	// Edge case: empty result
	nums3 := []int{10, 20}
	queries3 := [][]int{{5, 5}}
	fmt.Printf("maximizeXor(%v, %v) = %v (expected [-1])\n",
		nums3, queries3, maximizeXor(nums3, queries3))
}
```

## 1709 — Biggest Window Between Visits

```go
package main

// LeetCode #1709: Biggest Window Between Visits
// https://leetcode.com/problems/biggest-window-between-visits/
// Difficulty: Hard [Premium]

import (
	"fmt"
	"sort"
)

type Visit struct {
	UserId    int
	VisitDate string
}

type UserWindow struct {
	UserId       int
	BiggestWindow int
}

func main() {
	visits1 := []Visit{
		{1, "2020-01-01"},
		{1, "2020-01-10"},
		{1, "2020-01-28"},
		{2, "2020-01-05"},
		{2, "2020-01-20"},
	}
	result1 := biggestWindow(visits1)
	fmt.Println("Test 1:")
	for _, r := range result1 {
		fmt.Printf("  UserId: %d, BiggestWindow: %d\n", r.UserId, r.BiggestWindow)
	}

	visits2 := []Visit{
		{1, "2020-01-24"},
		{1, "2020-01-25"},
		{2, "2020-01-01"},
		{2, "2020-01-02"},
		{2, "2020-01-03"},
		{2, "2020-12-31"},
	}
	result2 := biggestWindow(visits2)
	fmt.Println("\nTest 2:")
	for _, r := range result2 {
		fmt.Printf("  UserId: %d, BiggestWindow: %d\n", r.UserId, r.BiggestWindow)
	}
}

func parseDate(s string) int {
	var year, month, day int
	fmt.Sscanf(s, "%d-%d-%d", &year, &month, &day)
	return year*365 + month*30 + day
}

func biggestWindow(visits []Visit) []UserWindow {
	userVisits := make(map[int][]int)
	userSet := make(map[int]bool)

	for _, v := range visits {
		userVisits[v.UserId] = append(userVisits[v.UserId], parseDate(v.VisitDate))
		userSet[v.UserId] = true
	}

	result := make([]UserWindow, 0)
	for uid := range userSet {
		dates := userVisits[uid]
		sort.Ints(dates)

		maxGap := 0
		for i := 1; i < len(dates); i++ {
			gap := dates[i] - dates[i-1]
			if gap > maxGap {
				maxGap = gap
			}
		}

		result = append(result, UserWindow{UserId: uid, BiggestWindow: maxGap})
	}

	return result
}
```

## 1713 — Minimum Operations To Make A Subsequence

```go
package main

// LeetCode #1713: Minimum Operations to Make a Subsequence
// https://leetcode.com/problems/minimum-operations-to-make-a-subsequence/
// Difficulty: Hard
// Strategy: Since target has distinct elements, map target values to indices.
// Then find LIS of those indices in arr.

import (
	"fmt"
	"sort"
)

func minOperations(target []int, arr []int) int {
	// Map target values to their indices
	pos := make(map[int]int)
	for i, v := range target {
		pos[v] = i
	}

	// Build list of indices in arr that also appear in target
	// This becomes the LIS problem
	indices := make([]int, 0)
	for _, v := range arr {
		if idx, ok := pos[v]; ok {
			indices = append(indices, idx)
		}
	}

	// LIS on indices using patience sorting O(n log n)
	tails := make([]int, 0)
	for _, idx := range indices {
		// Find first element >= idx in tails
		j := sort.SearchInts(tails, idx)
		if j == len(tails) {
			tails = append(tails, idx)
		} else {
			tails[j] = idx
		}
	}

	// Minimum operations = len(target) - longest common subsequence length
	return len(target) - len(tails)
}

func main() {
	// Example 1: target=[5,1,3], arr=[9,4,2,3,4] -> 2
	target1 := []int{5, 1, 3}
	arr1 := []int{9, 4, 2, 3, 4}
	fmt.Printf("minOperations(%v, %v) = %d (expected 2)\n", target1, arr1, minOperations(target1, arr1))

	// Example 2: target=[6,4,8,1,3,2], arr=[4,7,6,2,3,8,6,1] -> 3
	target2 := []int{6, 4, 8, 1, 3, 2}
	arr2 := []int{4, 7, 6, 2, 3, 8, 6, 1}
	fmt.Printf("minOperations(%v, %v) = %d (expected 3)\n", target2, arr2, minOperations(target2, arr2))

	// Example 3: target=[1,2,3], arr=[1,2,3] -> 0
	target3 := []int{1, 2, 3}
	arr3 := []int{1, 2, 3}
	fmt.Printf("minOperations(%v, %v) = %d (expected 0)\n", target3, arr3, minOperations(target3, arr3))

	// Example 4: target=[1,2,3], arr=[4,5,6] -> 3
	target4 := []int{1, 2, 3}
	arr4 := []int{4, 5, 6}
	fmt.Printf("minOperations(%v, %v) = %d (expected 3)\n", target4, arr4, minOperations(target4, arr4))
}
```

## 1714 — Sum Of Special Evenly Spaced Elements In Array

```go
package main

// LeetCode #1714: Sum Of Special Evenly-Spaced Elements In Array
// https://leetcode.com/problems/sum-of-special-evenly-spaced-elements-in-array/
// Difficulty: Hard [Premium]

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	queries := [][]int{{0, 2}, {1, 3}, {2, 1}}
	result := sumOfSpecialEvenlySpacedElements(nums, queries)
	fmt.Printf("Test 1:\n")
	fmt.Printf("Result: %v\nExpected: [25 14 30]\n\n", result)

	nums2 := []int{5, 2, 8, 3, 7, 1, 9, 4, 6}
	queries2 := [][]int{{0, 3}, {2, 2}, {1, 4}}
	result2 := sumOfSpecialEvenlySpacedElements(nums2, queries2)
	fmt.Printf("Test 2:\n")
	fmt.Printf("Result: %v\n", result2)

	nums3 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	queries3 := [][]int{{0, 1}}
	result3 := sumOfSpecialEvenlySpacedElements(nums3, queries3)
	fmt.Printf("\nTest 3 (full sum):\n")
	fmt.Printf("Result: %v\nExpected: [450]\n", result3)
}

func sumOfSpecialEvenlySpacedElements(nums []int, queries [][]int) []int {
	n := len(nums)
	sqrtN := int(math.Sqrt(float64(n)))

	// Precompute prefix sums for small step sizes (y <= sqrtN)
	// preSum[s][i] = sum of nums[i], nums[i-s], nums[i-2s], ... down to index 0
	preSum := make([][]int, sqrtN+1)
	for s := 1; s <= sqrtN; s++ {
		preSum[s] = make([]int, n)
		for i := 0; i < n; i++ {
			if i >= s {
				preSum[s][i] = preSum[s][i-s] + nums[i]
			} else {
				preSum[s][i] = nums[i]
			}
		}
	}

	ans := make([]int, len(queries))
	for idx, q := range queries {
		x, y := q[0], q[1]
		if y <= sqrtN && y > 0 {
			// Use precomputed prefix sums
			last := x + ((n-1-x)/y)*y
			sum := preSum[y][last]
			if x >= y {
				sum -= preSum[y][x-y]
			}
			ans[idx] = sum
		} else if y == 0 {
			ans[idx] = nums[x]
		} else {
			// Large step: brute force
			sum := 0
			for i := x; i < n; i += y {
				sum += nums[i]
			}
			ans[idx] = sum
		}
	}
	return ans
}
```

## 1715 — Count Apples And Oranges

```go
package main

// LeetCode #1715: Count Apples and Oranges
// https://leetcode.com/problems/count-apples-and-oranges/
// Difficulty: Hard [Premium]

import "fmt"

type Box struct {
	BoxId       int
	AppleCount  int
	OrangeCount int
}

type Chest struct {
	ChestId     int
	AppleCount  int
	OrangeCount int
}

type BoxChest struct {
	BoxId   int
	ChestId int
}

func main() {
	boxes := []Box{
		{2, 2, 3},
		{3, 1, 4},
	}
	chests := []Chest{
		{1, 5, 6},
		{2, 7, 8},
	}
	boxChests := []BoxChest{
		{2, 1},
		{3, 2},
	}
	apples, oranges := countApplesAndOranges(boxes, chests, boxChests)
	fmt.Printf("Test 1 - Apples: %d, Oranges: %d\n", apples, oranges)
	fmt.Printf("Expected: apples=15, oranges=21\n\n")

	boxes2 := []Box{
		{1, 3, 5},
	}
	chests2 := []Chest{
		{3, 2, 4},
	}
	boxChests2 := []BoxChest{
		{1, 3},
	}
	apples2, oranges2 := countApplesAndOranges(boxes2, chests2, boxChests2)
	fmt.Printf("Test 2 - Apples: %d, Oranges: %d\n", apples2, oranges2)
	fmt.Printf("Expected: apples=5, oranges=9\n")
}

func countApplesAndOranges(boxes []Box, chests []Chest, boxChests []BoxChest) (int, int) {
	chestMap := make(map[int]Chest)
	for _, c := range chests {
		chestMap[c.ChestId] = c
	}

	boxChestMap := make(map[int]int)
	for _, bc := range boxChests {
		boxChestMap[bc.BoxId] = bc.ChestId
	}

	totalApples := 0
	totalOranges := 0

	for _, b := range boxes {
		totalApples += b.AppleCount
		totalOranges += b.OrangeCount
		if chestId, ok := boxChestMap[b.BoxId]; ok {
			if chest, ok := chestMap[chestId]; ok {
				totalApples += chest.AppleCount
				totalOranges += chest.OrangeCount
			}
		}
	}

	return totalApples, totalOranges
}
```

## 1719 — Number Of Ways To Reconstruct A Tree

```go
package main

// LeetCode #1719: Number Of Ways To Reconstruct A Tree
// https://leetcode.com/problems/number-of-ways-to-reconstruct-a-tree/
// Difficulty: Hard
// Strategy: Sort nodes by degree descending, assign parents from processed neighbors.
// Validate ancestor relationships. Check multiplicity when deg equals parent deg.

import (
	"fmt"
	"sort"
)

func checkWays(pairs [][]int) int {
	// Build adjacency and degrees
	adj := make(map[int]map[int]bool)
	deg := make(map[int]int)
	nodeSet := make(map[int]bool)

	for _, p := range pairs {
		u, v := p[0], p[1]
		nodeSet[u] = true
		nodeSet[v] = true
		if adj[u] == nil {
			adj[u] = make(map[int]bool)
		}
		if adj[v] == nil {
			adj[v] = make(map[int]bool)
		}
		adj[u][v] = true
		adj[v][u] = true
		deg[u]++
		deg[v]++
	}

	n := len(nodeSet)

	// Build sorted node list
	nodes := make([]int, 0, n)
	for node := range nodeSet {
		nodes = append(nodes, node)
	}
	sort.Slice(nodes, func(i, j int) bool {
		if deg[nodes[i]] != deg[nodes[j]] {
			return deg[nodes[i]] > deg[nodes[j]]
		}
		return nodes[i] < nodes[j]
	})

	maxDeg := deg[nodes[0]]

	// If max degree < n-1: only simple paths are valid (deg <= 2 for all nodes)
	if maxDeg < n-1 {
		if maxDeg > 2 {
			return 0
		}
		leafCount := 0
		for _, v := range nodes {
			if deg[v] > 2 {
				return 0
			}
			if deg[v] == 1 {
				leafCount++
			}
		}
		if leafCount != 2 {
			return 0
		}
		// Path graph: there is exactly 1 way
		return 1
	}

	// Root = first node (highest degree). No deg=n-1 requirement.
	parent := make(map[int]int)
	parent[nodes[0]] = -1

	// Processing order for tiebreaking
	order := make(map[int]int)
	for i, v := range nodes {
		order[v] = i
	}

	// Assign parents
	for _, v := range nodes[1:] {
		bestP := -1
		bestDeg := -1
		for u := range adj[v] {
			if _, ok := parent[u]; ok { // u is processed
				if deg[u] >= deg[v] {
					if bestP == -1 || deg[u] < bestDeg || (deg[u] == bestDeg && order[u] > order[bestP]) {
						bestP = u
						bestDeg = deg[u]
					}
				}
			}
		}
		if bestP == -1 {
			return 0
		}
		parent[v] = bestP
	}

	// Validate ancestor relationships for all pairs
	isAncestor := func(anc, desc int) bool {
		for desc != -1 {
			if desc == anc {
				return true
			}
			desc = parent[desc]
		}
		return false
	}

	for _, p := range pairs {
		u, v := p[0], p[1]
		if !isAncestor(u, v) && !isAncestor(v, u) {
			return 0
		}
	}

	// Check multiplicity: if any node has same degree as parent
	// AND has another same-degree neighbor != parent
	for v, p := range parent {
		if p == -1 {
			continue
		}
		if deg[v] == deg[p] {
			for u := range adj[v] {
				if u != p && deg[u] == deg[v] {
					return 2
				}
			}
		}
	}

	return 1
}

func main() {
	// Example 1: [[1,2],[2,3]] -> 1
	fmt.Printf("checkWays(%v) = %d (expected 1)\n", [][]int{{1, 2}, {2, 3}}, checkWays([][]int{{1, 2}, {2, 3}}))

	// Example 2: [[1,2],[2,3],[1,3]] -> 2
	fmt.Printf("checkWays(%v) = %d (expected 2)\n", [][]int{{1, 2}, {2, 3}, {1, 3}}, checkWays([][]int{{1, 2}, {2, 3}, {1, 3}}))

	// Example 3: [[1,2],[2,3],[2,4],[1,5]] -> 0
	fmt.Printf("checkWays(%v) = %d (expected 0)\n", [][]int{{1, 2}, {2, 3}, {2, 4}, {1, 5}}, checkWays([][]int{{1, 2}, {2, 3}, {2, 4}, {1, 5}}))

	// Linear chain [[1,2],[2,3],[3,4]] -> 1
	fmt.Printf("checkWays(%v) = %d (expected 1)\n", [][]int{{1, 2}, {2, 3}, {3, 4}}, checkWays([][]int{{1, 2}, {2, 3}, {3, 4}}))

	// Star [[1,2],[1,3],[1,4]] -> 1
	fmt.Printf("checkWays(%v) = %d (expected 1)\n", [][]int{{1, 2}, {1, 3}, {1, 4}}, checkWays([][]int{{1, 2}, {1, 3}, {1, 4}}))
}
```

## 1721 — Swapping Nodes In A Linked List

```go
package main

// LeetCode #1721: Swapping Nodes in a Linked List
// https://leetcode.com/problems/swapping-nodes-in-a-linked-list/
// Difficulty: Medium (categorized as Hard in this repo)

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	curr := head
	for i := 1; i < len(arr); i++ {
		curr.Next = &ListNode{Val: arr[i]}
		curr = curr.Next
	}
	return head
}

func printList(head *ListNode) {
	fmt.Print("[")
	for curr := head; curr != nil; curr = curr.Next {
		if curr != head {
			fmt.Print(",")
		}
		fmt.Print(curr.Val)
	}
	fmt.Print("]")
}

func main() {
	// Test 1
	list1 := createList([]int{1, 2, 3, 4, 5})
	fmt.Printf("Test 1 - Input: [1,2,3,4,5], k=2\n")
	printList(list1)
	fmt.Println()
	result1 := swapNodes(list1, 2)
	printList(result1)
	fmt.Println(" (Expected: [1,4,3,2,5])\n")

	// Test 2
	list2 := createList([]int{7, 9, 6, 6, 7, 8, 3, 0, 9, 5})
	fmt.Printf("Test 2 - Input: [7,9,6,6,7,8,3,0,9,5], k=5\n")
	result2 := swapNodes(list2, 5)
	printList(result2)
	fmt.Println(" (Expected: [7,9,6,6,8,7,3,0,9,5])\n")

	// Test 3
	list3 := createList([]int{1})
	fmt.Printf("Test 3 - Input: [1], k=1\n")
	result3 := swapNodes(list3, 1)
	printList(result3)
	fmt.Println(" (Expected: [1])")
}

func swapNodes(head *ListNode, k int) *ListNode {
	// Find k-th node from the beginning
	first := head
	for i := 1; i < k; i++ {
		first = first.Next
	}

	// Find k-th node from the end using two-pointer technique
	second := head
	curr := first
	for curr.Next != nil {
		curr = curr.Next
		second = second.Next
	}

	// Swap values (not nodes)
	first.Val, second.Val = second.Val, first.Val
	return head
}
```

## 1723 — Find Minimum Time To Finish All Jobs

```go
package main

// LeetCode #1723: Find Minimum Time to Finish All Jobs
// https://leetcode.com/problems/find-minimum-time-to-finish-all-jobs/
// Difficulty: Hard
// Strategy: DP over bitmask. dp[mask] = min possible max time.
// Iterate w from 1 to k, each iteration adds one more worker.
// For each mask, try all subset splits: dp_prev[remaining] vs sum[sub].

import (
	"fmt"
	"math"
)

func minimumTimeRequired(jobs []int, k int) int {
	n := len(jobs)

	// Precompute sum of each subset
	sum := make([]int, 1<<n)
	for mask := 1; mask < 1<<n; mask++ {
		lsb := mask & -mask
		idx := 0
		temp := lsb
		for temp > 1 {
			temp >>= 1
			idx++
		}
		sum[mask] = sum[mask^lsb] + jobs[idx]
	}

	// dp[mask] after w workers = min possible max time for jobs in mask
	dp := make([]int, 1<<n)
	for mask := range dp {
		dp[mask] = sum[mask] // 1 worker = sum of all jobs in mask
	}

	// For 2nd through kth worker
	for w := 2; w <= k; w++ {
		next := make([]int, 1<<n)
		for mask := range next {
			next[mask] = dp[mask] // start with previous value (one fewer worker)
		}

		for mask := 0; mask < 1<<n; mask++ {
			// Try splitting mask: subset 'sub' goes to the new worker,
			// remaining = mask ^ sub goes to the existing (w-1) workers
			sub := mask
			for sub > 0 {
				// Skip full mask (sub == mask): this would mean ALL jobs go to new worker,
				// which means the previous workers do nothing — handled by initialization.
				remaining := mask ^ sub
				if dp[remaining] != math.MaxInt32 {
					candidate := max(dp[remaining], sum[sub])
					if candidate < next[mask] {
						next[mask] = candidate
					}
				}
				sub = (sub - 1) & mask
			}
		}
		dp = next
	}

	return dp[(1<<n)-1]
}

func main() {
	// LeetCode Example 1: jobs=[3,2,3], k=3 -> 3
	// (each worker gets one job: max(3,2,3)=3)
	fmt.Printf("minimumTimeRequired(%v, %d) = %d (expected 3)\n",
		[]int{3, 2, 3}, 3, minimumTimeRequired([]int{3, 2, 3}, 3))

	// LeetCode Example 2: jobs=[1,2,4,7,8], k=2 -> 11
	// split: [8,2,1]=11, [7,4]=11
	fmt.Printf("minimumTimeRequired(%v, %d) = %d (expected 11)\n",
		[]int{1, 2, 4, 7, 8}, 2, minimumTimeRequired([]int{1, 2, 4, 7, 8}, 2))

	// LeetCode Example 3: jobs=[11,2,7,4,8,10,3,1], k=3
	fmt.Printf("minimumTimeRequired(%v, %d) = %d\n",
		[]int{11, 2, 7, 4, 8, 10, 3, 1}, 3, minimumTimeRequired([]int{11, 2, 7, 4, 8, 10, 3, 1}, 3))

	// Single worker
	fmt.Printf("minimumTimeRequired(%v, %d) = %d (expected 15)\n",
		[]int{5, 5, 5}, 1, minimumTimeRequired([]int{5, 5, 5}, 1))

	// All jobs to each worker (k == n)
	fmt.Printf("minimumTimeRequired(%v, %d) = %d (expected 5)\n",
		[]int{1, 2, 3, 4, 5}, 5, minimumTimeRequired([]int{1, 2, 3, 4, 5}, 5))
}
```

## 1724 — Checking Existence Of Edge Length Limited Paths Ii

```go
package main

// LeetCode #1724: Checking Existence of Edge Length Limited Paths II
// https://leetcode.com/problems/checking-existence-of-edge-length-limited-paths-ii/
// Difficulty: Hard [Premium]

import (
	"fmt"
	"sort"
)

type DistanceLimitedPathsExist struct {
	n     int
	edges [][]int
}

func Constructor(n int, edgeList [][]int) DistanceLimitedPathsExist {
	sortedEdges := make([][]int, len(edgeList))
	copy(sortedEdges, edgeList)
	sort.Slice(sortedEdges, func(i, j int) bool {
		return sortedEdges[i][2] < sortedEdges[j][2]
	})
	return DistanceLimitedPathsExist{n: n, edges: sortedEdges}
}

func (this *DistanceLimitedPathsExist) Query(p int, q int, limit int) bool {
	parent := make([]int, this.n)
	for i := 0; i < this.n; i++ {
		parent[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(a, b int) {
		ra, rb := find(a), find(b)
		if ra != rb {
			parent[ra] = rb
		}
	}

	for _, e := range this.edges {
		if e[2] >= limit {
			break
		}
		union(e[0], e[1])
	}
	return find(p) == find(q)
}

func main() {
	// Test 1
	edges1 := [][]int{{0, 1, 2}, {1, 2, 4}, {2, 0, 8}, {1, 0, 16}}
	obj := Constructor(3, edges1)
	fmt.Println("Test 1:")
	fmt.Printf(" Query(0,2,2): %v (Expected: false)\n", obj.Query(0, 2, 2))
	fmt.Printf(" Query(0,2,5): %v (Expected: true)\n\n", obj.Query(0, 2, 5))

	// Test 2
	edges2 := [][]int{{0, 1, 10}, {1, 2, 5}, {2, 3, 3}, {0, 3, 20}}
	obj2 := Constructor(4, edges2)
	fmt.Println("Test 2:")
	fmt.Printf(" Query(0,3,15): %v (Expected: true)\n", obj2.Query(0, 3, 15))
	fmt.Printf(" Query(0,3,4):  %v (Expected: false)\n", obj2.Query(0, 3, 4))
	fmt.Printf(" Query(0,2,6):  %v (Expected: true)\n", obj2.Query(0, 2, 6))
}
```

## 1728 — Cat And Mouse Ii

```go
package main

// LeetCode #1728: Cat and Mouse II
// https://leetcode.com/problems/cat-and-mouse-ii/
// Difficulty: Hard

import "fmt"

func main() {
	grid1 := []string{
		"####F",
		"#C...#",
		"M....#",
	}
	fmt.Printf("Test 1:\nResult: %v (Expected: true)\n\n", canMouseWin(grid1, 1, 2))

	grid2 := []string{
		"M.C...F",
	}
	fmt.Printf("Test 2:\nResult: %v (Expected: true)\n\n", canMouseWin(grid2, 1, 4))

	grid3 := []string{
		"M.C...F",
	}
	fmt.Printf("Test 3:\nResult: %v (Expected: false)\n\n", canMouseWin(grid3, 1, 3))

	grid4 := []string{
		"C...#",
		"...#F",
		"....#",
		"M....",
	}
	fmt.Printf("Test 4:\nResult: %v (Expected: false)\n", canMouseWin(grid4, 2, 5))
}

func canMouseWin(grid []string, catJump int, mouseJump int) bool {
	rows := len(grid)
	cols := len(grid[0])

	// Locate positions
	var mr, mc, cr, cc, fr, fc int
	found := 0
	for i := 0; i < rows && found < 6; i++ {
		for j := 0; j < cols && found < 6; j++ {
			switch grid[i][j] {
			case 'M':
				mr, mc = i, j
				found++
			case 'C':
				cr, cc = i, j
				found++
			case 'F':
				fr, fc = i, j
				found++
			}
		}
	}

	maxMoves := rows * cols * 2
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	memo := make(map[[5]int]bool)

	var dfs func(mr, mc, cr, cc int, turn int, moves int) bool
	dfs = func(mr, mc, cr, cc int, turn int, moves int) bool {
		if moves > maxMoves {
			return false
		}
		key := [5]int{mr, mc, cr, cc, turn}
		if res, ok := memo[key]; ok {
			return res
		}

		if turn == 0 { // Mouse's turn
			// Option: stay in place
			if mr == fr && mc == fc {
				memo[key] = true
				return true
			}
			if dfs(mr, mc, cr, cc, 1, moves+1) {
				memo[key] = true
				return true
			}

			// Option: move in each direction
			for _, d := range dirs {
				for step := 1; step <= mouseJump; step++ {
					nr, nc := mr+d[0]*step, mc+d[1]*step
					if nr < 0 || nr >= rows || nc < 0 || nc >= cols || grid[nr][nc] == '#' {
						break
					}
					if nr == cr && nc == cc {
						continue
					}
					if nr == fr && nc == fc {
						memo[key] = true
						return true
					}
					if dfs(nr, nc, cr, cc, 1, moves+1) {
						memo[key] = true
						return true
					}
				}
			}
			memo[key] = false
			return false
		} else { // Cat's turn
			// Option: stay in place
			if cr == mr && cc == mc {
				memo[key] = false
				return false
			}
			if cr == fr && cc == fc {
				memo[key] = false
				return false
			}
			if !dfs(mr, mc, cr, cc, 0, moves+1) {
				memo[key] = false
				return false
			}

			// Option: move in each direction
			for _, d := range dirs {
				for step := 1; step <= catJump; step++ {
					nr, nc := cr+d[0]*step, cc+d[1]*step
					if nr < 0 || nr >= rows || nc < 0 || nc >= cols || grid[nr][nc] == '#' {
						break
					}
					if nr == mr && nc == mc {
						memo[key] = false
						return false
					}
					if nr == fr && nc == fc {
						memo[key] = false
						return false
					}
					if !dfs(mr, mc, nr, nc, 0, moves+1) {
						memo[key] = false
						return false
					}
				}
			}
			memo[key] = true
			return true
		}
	}

	return dfs(mr, mc, cr, cc, 0, 0)
}
```

## 1734 — Decode Xored Permutation

```go
package main

// LeetCode #1734: Decode XORed Permutation
// https://leetcode.com/problems/decode-xored-permutation/
// Difficulty: Medium (categorized as Hard in this repo)

import "fmt"

func main() {
	encoded1 := []int{3, 1}
	decoded1 := decode(encoded1)
	fmt.Printf("Test 1 - Input: %v\nOutput: %v\nExpected: [1,2,3]\n\n", encoded1, decoded1)

	encoded2 := []int{6, 5, 4, 6}
	decoded2 := decode(encoded2)
	fmt.Printf("Test 2 - Input: %v\nOutput: %v\nExpected: [2,4,1,5,3]\n\n", encoded2, decoded2)

	encoded3 := []int{12, 6, 2}
	decoded3 := decode(encoded3)
	fmt.Printf("Test 3 - Input: %v\nOutput: %v\n", encoded3, decoded3)
}

func decode(encoded []int) []int {
	n := len(encoded) + 1

	// XOR of all numbers from 1 to n
	allXor := 0
	for i := 1; i <= n; i++ {
		allXor ^= i
	}

	// XOR of encoded[1], encoded[3], encoded[5], ... (odd indices)
	// This gives us XOR of perm[1] ^ perm[2] ^ ... ^ perm[n-1]
	oddXor := 0
	for i := 1; i < len(encoded); i += 2 {
		oddXor ^= encoded[i]
	}

	// perm[0] = allXor ^ oddXor
	perm := make([]int, n)
	perm[0] = allXor ^ oddXor

	// Reconstruct the rest: perm[i] = perm[i-1] ^ encoded[i-1]
	for i := 1; i < n; i++ {
		perm[i] = perm[i-1] ^ encoded[i-1]
	}

	return perm
}
```

## 1735 — Count Ways To Make Array With Product

```go
package main

// LeetCode #1735: Count Ways to Make Array With Product
// https://leetcode.com/problems/count-ways-to-make-array-with-product/
// Difficulty: Hard
//
// Approach: Prime factorization + combinatorics (stars and bars).
// For each query [k, n]:
//   1. Factorize n into prime factors with counts.
//   2. For each prime with count c, distribute c identical items into k distinct
//      positions => C(c + k - 1, k - 1) ways.
//   3. Multiply results for all primes (primes are independent).

import (
	"fmt"
)

const MOD = 1_000_000_007

func powMod(a, e int) int {
	res := 1
	for e > 0 {
		if e&1 == 1 {
			res = (res * a) % MOD
		}
		a = (a * a) % MOD
		e >>= 1
	}
	return res
}

func nCr(n, r int) int {
	if r < 0 || r > n {
		return 0
	}
	if r > n-r {
		r = n - r
	}
	num, den := 1, 1
	for i := 0; i < r; i++ {
		num = (num * (n - i)) % MOD
		den = (den * (i + 1)) % MOD
	}
	return (num * powMod(den, MOD-2)) % MOD
}

func factorize(n int) map[int]int {
	factors := make(map[int]int)
	for p := 2; p*p <= n; p++ {
		for n%p == 0 {
			factors[p]++
			n /= p
		}
	}
	if n > 1 {
		factors[n]++
	}
	return factors
}

func waysToFillArray(queries [][]int) []int {
	ans := make([]int, len(queries))
	for idx, q := range queries {
		k, n := q[0], q[1]
		factors := factorize(n)
		res := 1
		for _, cnt := range factors {
			res = (res * nCr(cnt+k-1, k-1)) % MOD
		}
		ans[idx] = res
	}
	return ans
}

func main() {
	// Example test cases
	queries := [][]int{{2, 6}, {5, 1}, {73, 660}}
	result := waysToFillArray(queries)
	fmt.Println("queries=[[2,6],[5,1],[73,660]]", "→", result)
	// Expected: [4, 1, 50734910]

	// Additional tests
	fmt.Println("queries=[[1,1]] →", waysToFillArray([][]int{{1, 1}}))
	fmt.Println("queries=[[2,2]] →", waysToFillArray([][]int{{2, 2}}))
	fmt.Println("queries=[[3,4]] →", waysToFillArray([][]int{{3, 4}}))
}
```

## 1739 — Building Boxes

```go
package main

// LeetCode #1739: Building Boxes
// https://leetcode.com/problems/building-boxes/
// Difficulty: Hard
//
// You have n boxes to place on the floor. The boxes must be placed
// such that:
// - Each box is at the corner of a unit cube grid.
// - Boxes can be stacked, but each box above must be supported.
// Find the minimum number of boxes touching the floor.
//
// Approach: Greedy accumulation. Build a tetrahedral layer structure,
// counting boxes placed on the floor minimally.

import "fmt"

func main() {
	// Example 1
	fmt.Println(minimumBoxes(3))
	// Example 2
	fmt.Println(minimumBoxes(4))
	// Example 3
	fmt.Println(minimumBoxes(10))
	// Edge: n = 1
	fmt.Println(minimumBoxes(1))
}

func minimumBoxes(n int) int {
	s, k := 0, 1
	for s+k*(k+1)/2 <= n {
		s += k * (k + 1) / 2
		k++
	}
	k--
	ans := k * (k + 1) / 2
	k = 1
	for s < n {
		ans++
		s += k
		k++
	}
	return ans
}
```

## 1745 — Palindrome Partitioning Iv

```go
package main

// LeetCode #1745: Palindrome Partitioning IV
// https://leetcode.com/problems/palindrome-partitioning-iv/
// Difficulty: Hard
//
// Approach: DP palindrome table + split check.
// 1. Precompute isPal[i][j] = true if s[i..j] is palindrome.
// 2. Check pairs of split points (i, j) such that:
//    isPal[0][i-1] && isPal[i][j-1] && isPal[j][n-1] are all true.

import "fmt"

func checkPartitioning(s string) bool {
	n := len(s)
	// dp[i][j] = s[i..j] is palindrome
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	// All single chars are palindrome
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}
	// Two chars
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
		}
	}
	// Longer substrings
	for length := 3; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
			}
		}
	}

	// Try all split points: first partition ends at i-1, second ends at j-1
	for i := 1; i < n-1; i++ {
		if !dp[0][i-1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if dp[i][j-1] && dp[j][n-1] {
				return true
			}
		}
	}
	return false
}

func main() {
	// Example test cases
	fmt.Println("\"abcbdd\" →", checkPartitioning("abcbdd")) // Expected: true
	fmt.Println("\"abc\" →", checkPartitioning("abc"))       // Expected: false
	fmt.Println("\"aabb\" →", checkPartitioning("aabb"))     // Expected: true (aa|bb|or aa|b|b etc)
	fmt.Println("\"abca\" →", checkPartitioning("abca"))     // Expected: false
	fmt.Println("\"aaaa\" →", checkPartitioning("aaaa"))     // Expected: true
}
```

## 1747 — Leetflex Banned Accounts

```go
package main

// LeetCode #1747: Leetflex Banned Accounts
// https://leetcode.com/problems/leetflex-banned-accounts/
// Difficulty: Hard [Premium]

import (
	"fmt"
	"sort"
)

type LogInfo struct {
	AccountId int
	IpAddress int
	Login     int
	Logout    int
}

func main() {
	logs1 := []LogInfo{
		{1, 1, 1, 5},
		{1, 2, 2, 4},
		{2, 1, 6, 8},
	}
	result1 := leetflexBannedAccounts(logs1)
	fmt.Printf("Test 1 - Banned accounts: %v (Expected: [1])\n\n", result1)

	logs2 := []LogInfo{
		{1, 1, 1, 5},
		{1, 2, 6, 8},
		{2, 1, 1, 5},
		{2, 2, 6, 8},
		{3, 1, 1, 10},
		{3, 2, 5, 15},
	}
	result2 := leetflexBannedAccounts(logs2)
	fmt.Printf("Test 2 - Banned accounts: %v (Expected: [3])\n\n", result2)

	logs3 := []LogInfo{
		{1, 1, 1, 5},
		{1, 2, 6, 8},
		{2, 1, 1, 5},
		{2, 2, 6, 8},
	}
	result3 := leetflexBannedAccounts(logs3)
	fmt.Printf("Test 3 - Banned accounts: %v (Expected: [])\n", result3)
}

func leetflexBannedAccounts(logs []LogInfo) []int {
	// Group sessions by account
	byAccount := make(map[int][]LogInfo)
	for _, l := range logs {
		byAccount[l.AccountId] = append(byAccount[l.AccountId], l)
	}

	banned := make([]int, 0)
	for accountId, sessions := range byAccount {
		// Merge overlapping sessions for each IP, then check for overlap across IPs
		ipSessions := make(map[int][]LogInfo)
		for _, s := range sessions {
			ipSessions[s.IpAddress] = append(ipSessions[s.IpAddress], s)
		}

		// Merge per-IP sessions
		merged := make([]LogInfo, 0)
		for _, ipSess := range ipSessions {
			sort.Slice(ipSess, func(i, j int) bool {
				return ipSess[i].Login < ipSess[j].Login
			})
			mergedIP := ipSess[0]
			for i := 1; i < len(ipSess); i++ {
				if ipSess[i].Login <= mergedIP.Logout+1 {
					if ipSess[i].Logout > mergedIP.Logout {
						mergedIP.Logout = ipSess[i].Logout
					}
				} else {
					merged = append(merged, mergedIP)
					mergedIP = ipSess[i]
				}
			}
			merged = append(merged, mergedIP)
		}

		// Sort all merged sessions by login time
		sort.Slice(merged, func(i, j int) bool {
			return merged[i].Login < merged[j].Login
		})

		// Check for overlapping sessions (same account, different IPs)
		for i := 1; i < len(merged); i++ {
			if merged[i].Login <= merged[i-1].Logout {
				banned = append(banned, accountId)
				break
			}
		}
	}

	sort.Ints(banned)
	return banned
}
```

## 1751 — Maximum Number Of Events That Can Be Attended Ii

```go
package main

// LeetCode #1751: Maximum Number of Events That Can Be Attended II
// https://leetcode.com/problems/maximum-number-of-events-that-can-be-attended-ii/
// Difficulty: Hard
//
// Approach: Sort by end time + DP + Binary Search.
// Sort events by end time. For each event i, find the last event j
// that ends before event i starts (using binary search).
// dp[i][k] = max value using up to k events from first i events (1-indexed).
// Transition: skip event i or take event i + dp[prev][k-1].

import (
	"fmt"
	"sort"
)

func maxValue(events [][]int, k int) int {
	// Sort by end time
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})

	n := len(events)
	// prev[i] = index of last event that ends before events[i] starts
	prev := make([]int, n)
	for i := 0; i < n; i++ {
		start := events[i][0]
		// Binary search for rightmost event with end < start
		lo, hi := 0, i-1
		prev[i] = -1
		for lo <= hi {
			mid := (lo + hi) / 2
			if events[mid][1] < start {
				prev[i] = mid
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}

	// dp[i][j] = max value using first i events (0-indexed), at most j events
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}

	for i := 1; i <= n; i++ {
		val := events[i-1][2]
		p := prev[i-1] + 1 // 1-indexed version of prev
		for j := 1; j <= k; j++ {
			// Skip event i-1
			best := dp[i-1][j]
			// Take event i-1
			take := val + dp[p][j-1]
			if take > best {
				best = take
			}
			dp[i][j] = best
		}
	}

	return dp[n][k]
}

func main() {
	// Example test case
	events := [][]int{{1, 2, 4}, {3, 4, 3}, {2, 3, 1}}
	fmt.Println("events=[[1,2,4],[3,4,3],[2,3,1]],k=2 →", maxValue(events, 2)) // Expected: 7

	// Additional tests
	events2 := [][]int{{1, 2, 4}, {3, 4, 3}, {2, 3, 1}}
	fmt.Println("events=[[1,2,4],[3,4,3],[2,3,1]],k=1 →", maxValue(events2, 1)) // Expected: 4

	events3 := [][]int{{1, 1, 5}, {2, 2, 3}, {3, 3, 4}}
	fmt.Println("events=[[1,1,5],[2,2,3],[3,3,4]],k=3 →", maxValue(events3, 3)) // Expected: 12

	events4 := [][]int{{1, 3, 2}, {4, 5, 2}, {2, 4, 3}}
	fmt.Println("events=[[1,3,2],[4,5,2],[2,4,3]],k=2 →", maxValue(events4, 2)) // Expected: 5
}
```

## 1755 — Closest Subsequence Sum

```go
package main

// LeetCode #1755: Closest Subsequence Sum
// https://leetcode.com/problems/closest-subsequence-sum/
// Difficulty: Hard
//
// Approach: Meet-in-the-middle.
// Split array into two halves. Generate all possible subset sums for each half.
// Sort the second half. For each sum in the first half, binary search the second
// half for the value closest to (goal - sum).

import (
	"fmt"
	"sort"
)

func minAbs(x, y int) int {
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	if x < y {
		return x
	}
	return y
}

func min(vals ...int) int {
	m := vals[0]
	for _, v := range vals[1:] {
		if v < m {
			m = v
		}
	}
	return m
}

func generateSums(arr []int) []int {
	n := len(arr)
	sums := []int{0}
	for i := 0; i < n; i++ {
		curLen := len(sums)
		for j := 0; j < curLen; j++ {
			sums = append(sums, sums[j]+arr[i])
		}
	}
	return sums
}

func minAbsDifference(nums []int, goal int) int {
	n := len(nums)
	mid := n / 2

	left := generateSums(nums[:mid])
	right := generateSums(nums[mid:])

	sort.Ints(right)

	best := 1 << 60
	for _, s := range left {
		target := goal - s
		// Binary search for closest
		idx := sort.SearchInts(right, target)
		if idx < len(right) {
			diff := target - right[idx]
			if diff < 0 {
				diff = -diff
			}
			if diff < best {
				best = diff
			}
		}
		if idx > 0 {
			diff := target - right[idx-1]
			if diff < 0 {
				diff = -diff
			}
			if diff < best {
				best = diff
			}
		}
		if best == 0 {
			break
		}
	}
	return best
}

func main() {
	// Example test case
	fmt.Println("nums=[5,-7,3,5],goal=6 →", minAbsDifference([]int{5, -7, 3, 5}, 6)) // Expected: 0

	// Additional tests
	fmt.Println("nums=[1,2,3],goal=10 →", minAbsDifference([]int{1, 2, 3}, 10))       // Expected: 4 (6 vs 10)
	fmt.Println("nums=[-1,-2,-3],goal=0 →", minAbsDifference([]int{-1, -2, -3}, 0))    // Expected: 0 (sum=0 via empty)
	fmt.Println("nums=[7,-9,15,-2],goal=-5 →", minAbsDifference([]int{7, -9, 15, -2}, -5)) // Expected: 0
}
```

## 1761 — Minimum Degree Of A Connected Trio In A Graph

```go
package main

// LeetCode #1761: Minimum Degree of a Connected Trio in a Graph
// https://leetcode.com/problems/minimum-degree-of-a-connected-trio-in-a-graph/
// Difficulty: Hard
//
// Approach: Adjacency matrix + degree array.
// 1. Build degree array and adjacency matrix for the graph.
// 2. For every trio (i, j, k) with i < j < k, check if all three edges exist.
// 3. Degree of trio = degree[i] + degree[j] + degree[k] - 6 (each internal edge counted twice).
// 4. Track minimum.

import (
	"fmt"
	"math"
)

func minTrioDegree(n int, edges [][]int) int {
	deg := make([]int, n)
	adj := make([][]bool, n)
	for i := range adj {
		adj[i] = make([]bool, n)
	}

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		deg[u]++
		deg[v]++
		adj[u][v] = true
		adj[v][u] = true
	}

	minDeg := math.MaxInt32
	found := false

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if !adj[i][j] {
				continue
			}
			for k := j + 1; k < n; k++ {
				if adj[i][k] && adj[j][k] {
					found = true
					d := deg[i] + deg[j] + deg[k] - 6
					if d < minDeg {
						minDeg = d
					}
				}
			}
		}
	}

	if !found {
		return -1
	}
	return minDeg
}

func main() {
	// Example test case
	n := 6
	edges := [][]int{{1, 2}, {1, 3}, {3, 2}, {4, 1}, {5, 2}, {3, 6}}
	fmt.Println("n=6,edges=[[1,2],[1,3],[3,2],[4,1],[5,2],[3,6]] →", minTrioDegree(n, edges)) // Expected: 3

	// Additional tests
	n2 := 4
	edges2 := [][]int{{1, 2}, {1, 3}, {3, 2}}
	fmt.Println("n=4,edges=[[1,2],[1,3],[3,2]] →", minTrioDegree(n2, edges2)) // Expected: 2 (deg=2,2,2 => 2+2+2-6=0? No deg[1]=2, deg[2]=2, deg[3]=2 => 0)

	n3 := 3
	edges3 := [][]int{{1, 2}, {2, 3}, {1, 3}}
	fmt.Println("n=3,edges=[[1,2],[2,3],[1,3]] →", minTrioDegree(n3, edges3)) // Expected: 0

	n4 := 5
	edges4 := [][]int{{1, 2}, {2, 3}, {3, 1}, {1, 4}, {2, 5}}
	fmt.Println("n=5,edges=[[1,2],[2,3],[3,1],[1,4],[2,5]] →", minTrioDegree(n4, edges4)) // Expected: 2
}
```

## 1766 — Tree Of Coprimes

```go
package main

// LeetCode #1766: Tree of Coprimes
// https://leetcode.com/problems/tree-of-coprimes/
// Difficulty: Hard
//
// Approach: DFS with depth tracking per value (nums[i] ≤ 50).
// For each node, find the nearest ancestor with a coprime value.
// Since values are ≤ 50, we can precompute all coprime pairs.
// During DFS, maintain for each value (1-50) a stack of (node, depth)
// for ancestors. For each node, check all coprime values and find
// the deepest ancestor.

import (
	"fmt"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func getCoprimes(nums []int, edges [][]int) []int {
	n := len(nums)
	// Build adjacency
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// Precompute coprime pairs for values 1..50
	coprime := make([][]int, 51)
	for v := 1; v <= 50; v++ {
		for u := 1; u <= 50; u++ {
			if gcd(u, v) == 1 {
				coprime[v] = append(coprime[v], u)
			}
		}
	}

	// For each value (1..50), maintain a stack of (node, depth)
	// Use arrays of [][2]int (idx 0=node, idx 1=depth)
	valStack := make([][][2]int, 51)

	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}

	visited := make([]bool, n)

	var dfs func(u, depth int)
	dfs = func(u, depth int) {
		visited[u] = true
		val := nums[u]

		// Find best ancestor with coprime value
		bestDepth := -1
		bestNode := -1
		for _, cv := range coprime[val] {
			stack := valStack[cv]
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				if top[1] > bestDepth {
					bestDepth = top[1]
					bestNode = top[0]
				}
			}
		}
		ans[u] = bestNode

		// Push current node
		valStack[val] = append(valStack[val], [2]int{u, depth})

		// DFS children
		for _, v := range adj[u] {
			if !visited[v] {
				dfs(v, depth+1)
			}
		}

		// Pop
		valStack[val] = valStack[val][:len(valStack[val])-1]
	}

	dfs(0, 0)
	return ans
}

func main() {
	// Example test case
	nums := []int{2, 3, 3, 2}
	edges := [][]int{{0, 1}, {1, 2}, {1, 3}}
	fmt.Println("nums=[2,3,3,2],edges=[[0,1],[1,2],[1,3]] →", getCoprimes(nums, edges)) // Expected: [-1,0,0,1]

	// Additional tests
	nums2 := []int{5, 6, 10, 2, 3}
	edges2 := [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 4}}
	fmt.Println("nums=[5,6,10,2,3],edges=[[0,1],[0,2],[1,3],[2,4]] →", getCoprimes(nums2, edges2))

	nums3 := []int{1, 1, 1, 1}
	edges3 := [][]int{{0, 1}, {1, 2}, {2, 3}}
	fmt.Println("nums=[1,1,1,1],edges=[[0,1],[1,2],[2,3]] →", getCoprimes(nums3, edges3)) // Expected: [-1,0,1,2]
}
```

## 1767 — Find The Subtasks That Did Not Execute

```go
package main

// LeetCode #1767: Find the Subtasks That Did Not Execute
// https://leetcode.com/problems/find-the-subtasks-that-did-not-execute/
// Difficulty: Hard [Premium]

import (
	"fmt"
	"sort"
)

type Task struct {
	TaskId         int
	SubtasksCount  int
}

type ExecutedSubtask struct {
	TaskId    int
	SubtaskId int
}

func main() {
	tasks := []Task{
		{1, 3},
		{2, 2},
		{3, 4},
	}
	executed := []ExecutedSubtask{
		{1, 1},
		{1, 2},
		{2, 1},
		{3, 2},
		{3, 4},
	}

	result := findSubtasksThatDidNotExecute(tasks, executed)
	fmt.Println("Test 1 - Subtasks that did not execute:")
	fmt.Print("Expected: [(1,3) (2,2) (3,1) (3,3)]\nGot:      ")
	for _, r := range result {
		fmt.Printf("(%d,%d) ", r.TaskId, r.SubtaskId)
	}
	fmt.Println()

	// Test 2: empty (all executed)
	tasks2 := []Task{{1, 2}}
	executed2 := []ExecutedSubtask{{1, 1}, {1, 2}}
	result2 := findSubtasksThatDidNotExecute(tasks2, executed2)
	fmt.Println("\nTest 2 (all executed):")
	fmt.Print("Expected: []\nGot:      ")
	for _, r := range result2 {
		fmt.Printf("(%d,%d) ", r.TaskId, r.SubtaskId)
	}
	fmt.Println()
}

func findSubtasksThatDidNotExecute(tasks []Task, executed []ExecutedSubtask) []ExecutedSubtask {
	executedSet := make(map[[2]int]bool)
	for _, e := range executed {
		executedSet[[2]int{e.TaskId, e.SubtaskId}] = true
	}

	notExecuted := make([]ExecutedSubtask, 0)
	for _, t := range tasks {
		for sub := 1; sub <= t.SubtasksCount; sub++ {
			if !executedSet[[2]int{t.TaskId, sub}] {
				notExecuted = append(notExecuted, ExecutedSubtask{TaskId: t.TaskId, SubtaskId: sub})
			}
		}
	}

	sort.Slice(notExecuted, func(i, j int) bool {
		if notExecuted[i].TaskId != notExecuted[j].TaskId {
			return notExecuted[i].TaskId < notExecuted[j].TaskId
		}
		return notExecuted[i].SubtaskId < notExecuted[j].SubtaskId
	})

	return notExecuted
}
```

## 1770 — Maximum Score From Performing Multiplication Operations

```go
package main

// LeetCode #1770: Maximum Score from Performing Multiplication Operations
// https://leetcode.com/problems/maximum-score-from-performing-multiplication-operations/
// Difficulty: Hard
//
// Approach: DP[l][i] where l = number of operations done and i = left index used.
// Equivalent to: dp[l][left] = max(
//     nums[left] * mult[l] + dp[l+1][left+1],   // take from left
//     nums[right] * mult[l] + dp[l+1][left]      // take from right
// )
// where right = n - 1 - (l - left).
// Optimized: 2D DP is fine since m <= 1000.

import (
	"fmt"
)

func maximumScore(nums []int, multipliers []int) int {
	n := len(nums)
	m := len(multipliers)
	// dp[l][left] = max score using l operations with 'left' left-end picks
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for l := m - 1; l >= 0; l-- {
		for left := l; left >= 0; left-- {
			right := n - 1 - (l - left)
			takeLeft := nums[left]*multipliers[l] + dp[l+1][left+1]
			takeRight := nums[right]*multipliers[l] + dp[l+1][left]
			if takeLeft > takeRight {
				dp[l][left] = takeLeft
			} else {
				dp[l][left] = takeRight
			}
		}
	}

	return dp[0][0]
}

func main() {
	// Example test case
	fmt.Println("nums=[1,2,3],mult=[3,2,1] →", maximumScore([]int{1, 2, 3}, []int{3, 2, 1})) // Expected: 14

	// Additional tests
	fmt.Println("nums=[-5,-3,-3,-2,7,1],mult=[-10,-5,3,4,6] →",
		maximumScore([]int{-5, -3, -3, -2, 7, 1}, []int{-10, -5, 3, 4, 6}))

	fmt.Println("nums=[5],mult=[10] →", maximumScore([]int{5}, []int{10}))             // Expected: 50
	fmt.Println("nums=[1,2],mult=[1,2] →", maximumScore([]int{1, 2}, []int{1, 2}))     // Expected: 5 (1*1 + 2*2 = 5)
}
```

## 1771 — Maximize Palindrome Length From Subsequences

```go
package main

// LeetCode #1771: Maximize Palindrome Length From Subsequences
// https://leetcode.com/problems/maximize-palindrome-length-from-subsequences/
// Difficulty: Hard

import "fmt"

func main() {
	word1 := "cacb"
	word2 := "cbba"
	fmt.Printf("Test 1 - word1=%s, word2=%s\n", word1, word2)
	fmt.Printf("Result: %d (Expected: 5)\n\n", longestPalindrome(word1, word2))

	word1 = "ab"
	word2 = "ab"
	fmt.Printf("Test 2 - word1=%s, word2=%s\n", word1, word2)
	fmt.Printf("Result: %d (Expected: 3)\n\n", longestPalindrome(word1, word2))

	word1 = "aa"
	word2 = "bb"
	fmt.Printf("Test 3 - word1=%s, word2=%s\n", word1, word2)
	fmt.Printf("Result: %d (Expected: 0)\n\n", longestPalindrome(word1, word2))

	word1 = "cebdedc"
	word2 = "d"
	fmt.Printf("Test 4 - word1=%s, word2=%s\n", word1, word2)
	fmt.Printf("Result: %d (Expected: 5)\n", longestPalindrome(word1, word2))
}

func longestPalindrome(word1 string, word2 string) int {
	s := word1 + word2
	n := len(s)
	n1 := len(word1)

	// dp[i][j] = longest palindromic subsequence in s[i..j]
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	result := 0

	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
				// Must pick at least one char from each word
				if i < n1 && j >= n1 && dp[i][j] > result {
					result = dp[i][j]
				}
			} else {
				if dp[i+1][j] > dp[i][j-1] {
					dp[i][j] = dp[i+1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}

	return result
}
```

## 1776 — Car Fleet Ii

```go
package main

// LeetCode #1776: Car Fleet II
// https://leetcode.com/problems/car-fleet-ii/
// Difficulty: Hard
//
// Approach: Monotonic stack (processing from right to left).
// Each car has (position, speed). Compute collision time with the next car
// ahead. Use a stack to track the collision chain.
//
// For car i, we check the car immediately ahead (j = i+1).
// If speed[i] <= speed[j], car i will never catch car j → answer[i] = -1.
// Otherwise, compute collision time t = (pos[j] - pos[i]) / (speed[i] - speed[j]).
// If the collision happens before or when car j hits its own blocker,
// then i collides with j at time t. Otherwise, i's collision time is
// determined by its collision with car j's blocker (recursive chain).

import (
	"fmt"
)

type Car struct {
	pos, speed float64
}

func getCollisionTimes(cars [][]int) []float64 {
	n := len(cars)
	ans := make([]float64, n)
	// Stack holds indices of cars that form collision chains (from right)
	stack := make([]int, 0, n)

	for i := n - 1; i >= 0; i-- {
		pos := float64(cars[i][0])
		speed := float64(cars[i][1])

		ans[i] = -1.0

		// Remove cars that can never be caught
		for len(stack) > 0 {
			j := stack[len(stack)-1]
			// If current car is slower or equal speed, it can't catch car j
			if speed <= float64(cars[j][1]) {
				stack = stack[:len(stack)-1]
				continue
			}
			// Compute collision time with car j
			posJ := float64(cars[j][0])
			speedJ := float64(cars[j][1])
			t := (posJ - pos) / (speed - speedJ)

			// If car j has a collision and this happens after that,
			// car i will actually collide with whatever j collides with.
			if ans[j] > 0 && t >= ans[j] {
				stack = stack[:len(stack)-1]
				continue
			}

			ans[i] = t
			break
		}

		stack = append(stack, i)
	}

	return ans
}

func main() {
	// Example test case (LeetCode Example 1)
	cars := [][]int{{1, 2}, {2, 1}, {4, 3}, {7, 2}}
	fmt.Println("cars=[[1,2],[2,1],[4,3],[7,2]] →", getCollisionTimes(cars))
	// Expected: [1, -1, 3, -1]

	// Additional tests
	cars2 := [][]int{{1, 1}, {2, 2}, {3, 4}}
	fmt.Println("cars=[[1,1],[2,2],[3,4]] →", getCollisionTimes(cars2))

	// Single car should have -1
	cars3 := [][]int{{1, 1}}
	fmt.Println("cars=[[1,1]] →", getCollisionTimes(cars3))

	// User-specified test (corrected interpretation)
	cars4 := [][]int{{1, 2}, {2, 1}, {3, 3}, {5, 4}}
	fmt.Println("cars=[[1,2],[2,1],[3,3],[5,4]] →", getCollisionTimes(cars4))
}
```

## 1782 — Count Pairs Of Nodes

```go
package main

// LeetCode #1782: Count Pairs Of Nodes
// https://leetcode.com/problems/count-pairs-of-nodes/
// Difficulty: Hard
//
// Approach: Degree sort + Binary Search.
// 1. Count degree for each node.
// 2. Count edge-incident pairs for each edge (for subtracting overcount).
// 3. For each query, first find all pairs (i,j) with deg[i] + deg[j] > query
//    using two-pointer on sorted degrees.
// 4. Subtract pairs where deg[i] + deg[j] > query but (i,j) is an edge where
//    deg[i] + deg[j] - edgeCount((i,j)) <= query.
// 5. Also handle duplicate edge counts.

import (
	"fmt"
	"sort"
)

func countPairs(n int, edges [][]int, queries []int) []int {
	// Degree of each node
	deg := make([]int, n+1)
	// Edge pair counts: key = (min*100000 + max) for uniqueness
	edgeCount := make(map[int]int)
	for _, e := range edges {
		u, v := e[0], e[1]
		if u > v {
			u, v = v, u
		}
		deg[u]++
		deg[v]++
		edgeCount[u*100000+v]++
	}

	// Sorted degrees (1-indexed)
	sortedDeg := make([]int, n)
	copy(sortedDeg, deg[1:])
	sort.Ints(sortedDeg)

	ans := make([]int, len(queries))
	for qi, q := range queries {
		// Two-pointer: count pairs with deg[i] + deg[j] > q
		total := 0
		left, right := 0, n-1
		for left < right {
			if sortedDeg[left]+sortedDeg[right] > q {
				total += right - left
				right--
			} else {
				left++
			}
		}

		// Subtract edge pairs that don't satisfy when considering shared edges
		seen := make(map[int]bool)
		for _, e := range edges {
			u, v := e[0], e[1]
			if u > v {
				u, v = v, u
			}
			key := u*100000 + v
			if seen[key] {
				continue
			}
			seen[key] = true
			cnt := edgeCount[key]
			// deg[u] + deg[v] > q is needed for this edge to be counted,
			// but deg[u] + deg[v] - cnt <= q means it should be subtracted
			if deg[u]+deg[v] > q && deg[u]+deg[v]-cnt <= q {
				total--
			}
		}

		ans[qi] = total
	}

	return ans
}

func main() {
	// Example test case (LeetCode Example 1)
	n := 4
	edges := [][]int{{1, 2}, {2, 4}, {1, 3}, {2, 3}, {2, 1}}
	queries := []int{2, 3}
	fmt.Println("n=4,edges=[[1,2],[2,4],[1,3],[2,3],[2,1]],queries=[2,3] →", countPairs(n, edges, queries))
	// Expected: [6, 5]

	// Additional tests
	n2 := 5
	edges2 := [][]int{{1, 5}, {2, 5}, {3, 5}, {4, 5}}
	queries2 := []int{1, 2, 3}
	fmt.Println("n=5,edges=[[1,5],[2,5],[3,5],[4,5]],queries=[1,2,3] →", countPairs(n2, edges2, queries2))

	n3 := 2
	edges3 := [][]int{{1, 2}}
	queries3 := []int{0, 1, 2}
	fmt.Println("n=2,edges=[[1,2]],queries=[0,1,2] →", countPairs(n3, edges3, queries3))
}
```

## 1783 — Grand Slam Titles

```go
package main

// LeetCode #1783: Grand Slam Titles
// https://leetcode.com/problems/grand-slam-titles/
// Difficulty: Hard [Premium]

import (
	"fmt"
	"sort"
)

type Player struct {
	PlayerId int
	Name     string
}

type Championship struct {
	Year      int
	Wimbledon int
	FrOpen    int
	UsOpen    int
	AuOpen    int
}

type PlayerGrandSlams struct {
	PlayerId        int
	Name            string
	GrandSlamsCount int
}

func main() {
	players := []Player{
		{1, "Nadal"},
		{2, "Federer"},
		{3, "Novak"},
	}

	championships := []Championship{
		{2018, 1, 1, 1, 1},
		{2019, 1, 1, 2, 2},
		{2020, 2, 1, 1, 1},
	}

	result := grandSlamTitles(players, championships)
	fmt.Println("Test 1 - Grand Slam Titles:")
	for _, r := range result {
		fmt.Printf("  PlayerId: %d, Name: %s, GrandSlamsCount: %d\n",
			r.PlayerId, r.Name, r.GrandSlamsCount)
	}

	// Test 2: no wins
	players2 := []Player{
		{1, "Player1"},
		{2, "Player2"},
	}
	championships2 := []Championship{
		{2020, 3, 3, 3, 3},
	}
	result2 := grandSlamTitles(players2, championships2)
	fmt.Println("\nTest 2 - Player 3 won but not in players list:")
	for _, r := range result2 {
		fmt.Printf("  PlayerId: %d, Name: %s, Count: %d\n", r.PlayerId, r.Name, r.GrandSlamsCount)
	}
}

func grandSlamTitles(players []Player, championships []Championship) []PlayerGrandSlams {
	counts := make(map[int]int)
	for _, c := range championships {
		counts[c.Wimbledon]++
		counts[c.FrOpen]++
		counts[c.UsOpen]++
		counts[c.AuOpen]++
	}

	playerMap := make(map[int]string)
	for _, p := range players {
		playerMap[p.PlayerId] = p.Name
	}

	result := make([]PlayerGrandSlams, 0)
	for _, p := range players {
		if cnt, ok := counts[p.PlayerId]; ok && cnt > 0 {
			result = append(result, PlayerGrandSlams{
				PlayerId:        p.PlayerId,
				Name:            p.Name,
				GrandSlamsCount: cnt,
			})
		}
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].GrandSlamsCount != result[j].GrandSlamsCount {
			return result[i].GrandSlamsCount > result[j].GrandSlamsCount
		}
		return result[i].PlayerId < result[j].PlayerId
	})

	return result
}
```

