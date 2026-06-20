# Easy (Mudah) — Problem ��3173

## 2798 — Number Of Employees Who Met The Target

```go
package main

// LeetCode #2798: Number of Employees Who Met the Target
// https://leetcode.com/problems/number-of-employees-who-met-the-target/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(NumberOfEmployeesWhoMetTheTarget([]int{0, 1, 2, 3, 4}, 2))
	fmt.Println(NumberOfEmployeesWhoMetTheTarget([]int{5, 1, 4, 2, 2}, 6))
}

func NumberOfEmployeesWhoMetTheTarget(hours []int, target int) int {
	count := 0
	for _, h := range hours {
		if h >= target {
			count++
		}
	}
	return count
}
```

## 2803 — Factorial Generator

```go
package main

// LeetCode #2803: Factorial Generator
// https://leetcode.com/problems/factorial-generator/
// Difficulty: Easy [Paid]
// Time: O(1) per call | Space: O(1)
// Note: JS problem, adapted to Go. Generator that yields factorials.

import "fmt"

func main() {
	gen := FactorialGenerator()
	for i := 0; i < 5; i++ {
		fmt.Println(gen())
	}
}

func FactorialGenerator() func() int {
	n := 0
	curr := 1
	return func() int {
		if n == 0 {
			n++
			return 1
		}
		curr *= n
		n++
		return curr
	}
}
```

## 2804 — Array Prototype Foreach

```go
package main

// LeetCode #2804: Array Prototype ForEach
// https://leetcode.com/problems/array-prototype-foreach/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(1)
// Note: JS problem, adapted to Go. Applies callback to each element.

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	sum := 0
	ArrayPrototypeForeach(nums, func(v int) {
		sum += v
	})
	fmt.Println(sum)
}

func ArrayPrototypeForeach(arr []int, callback func(int)) {
	for _, v := range arr {
		callback(v)
	}
}
```

## 2806 — Account Balance After Rounded Purchase

```go
package main

// LeetCode #2806: Account Balance After Rounded Purchase
// https://leetcode.com/problems/account-balance-after-rounded-purchase/
// Difficulty: Easy
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(AccountBalanceAfterRoundedPurchase(9))
	fmt.Println(AccountBalanceAfterRoundedPurchase(15))
}

func AccountBalanceAfterRoundedPurchase(purchaseAmount int) int {
	return 100 - ((purchaseAmount+5)/10)*10
}
```

## 2810 — Faulty Keyboard

```go
package main

// LeetCode #2810: Faulty Keyboard
// https://leetcode.com/problems/faulty-keyboard/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import (
	"fmt"
)

func main() {
	fmt.Println(FaultyKeyboard("string"))
	fmt.Println(FaultyKeyboard("poiinter"))
}

func FaultyKeyboard(s string) string {
	result := []rune{}
	for _, ch := range s {
		if ch == 'i' {
			// Reverse the current result
			for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
				result[i], result[j] = result[j], result[i]
			}
		} else {
			result = append(result, ch)
		}
	}
	return string(result)
}
```

## 2815 — Max Pair Sum In An Array

```go
package main

// LeetCode #2815: Max Pair Sum in an Array
// https://leetcode.com/problems/max-pair-sum-in-an-array/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(MaxPairSumInAnArray([]int{51, 71, 17, 24, 42}))
	fmt.Println(MaxPairSumInAnArray([]int{1, 2, 3, 4}))
}

func maxDigit(n int) int {
	maxD := 0
	for n > 0 {
		d := n % 10
		if d > maxD {
			maxD = d
		}
		n /= 10
	}
	return maxD
}

func MaxPairSumInAnArray(nums []int) int {
	maxVal := make([]int, 10) // digits 0-9
	ans := -1
	for _, n := range nums {
		md := maxDigit(n)
		if maxVal[md] > 0 {
			sum := maxVal[md] + n
			if sum > ans {
				ans = sum
			}
		}
		if n > maxVal[md] {
			maxVal[md] = n
		}
	}
	return ans
}
```

## 2822 — Inversion Of Object

```go
package main

// LeetCode #2822: Inversion of Object
// https://leetcode.com/problems/inversion-of-object/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(n)
// Note: JS problem, adapted to Go. Swaps keys and values of a map.

import "fmt"

func main() {
	input := map[string]int{"a": 1, "b": 2, "c": 1}
	fmt.Println(InversionOfObject(input))
}

func InversionOfObject(obj map[string]int) map[int]string {
	result := make(map[int]string, len(obj))
	for k, v := range obj {
		result[v] = k
	}
	return result
}
```

## 2824 — Count Pairs Whose Sum Is Less Than Target

```go
package main

// LeetCode #2824: Count Pairs Whose Sum is Less than Target
// https://leetcode.com/problems/count-pairs-whose-sum-is-less-than-target/
// Difficulty: Easy
// Time: O(n^2) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(CountPairsWhoseSumIsLessThanTarget([]int{-1, 1, 2, 3, 1}, 2))
	fmt.Println(CountPairsWhoseSumIsLessThanTarget([]int{-6, 2, 5, -2, -7, -1, 3}, -2))
}

func CountPairsWhoseSumIsLessThanTarget(nums []int, target int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] < target {
				count++
			}
		}
	}
	return count
}
```

## 2828 — Check If A String Is An Acronym Of Words

```go
package main

// LeetCode #2828: Check if a String Is an Acronym of Words
// https://leetcode.com/problems/check-if-a-string-is-an-acronym-of-words/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(CheckIfAStringIsAnAcronymOfWords([]string{"alice", "bob", "charlie"}, "abc"))
	fmt.Println(CheckIfAStringIsAnAcronymOfWords([]string{"an", "apple"}, "a"))
}

func CheckIfAStringIsAnAcronymOfWords(words []string, s string) bool {
	if len(words) != len(s) {
		return false
	}
	for i, w := range words {
		if w[0] != s[i] {
			return false
		}
	}
	return true
}
```

## 2833 — Furthest Point From Origin

```go
package main

// LeetCode #2833: Furthest Point From Origin
// https://leetcode.com/problems/furthest-point-from-origin/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(FurthestPointFromOrigin("L_RL__R"))
	fmt.Println(FurthestPointFromOrigin("_R__LL_"))
}

func FurthestPointFromOrigin(moves string) int {
	countL, countR, countUnderscore := 0, 0, 0
	for _, c := range moves {
		switch c {
		case 'L':
			countL++
		case 'R':
			countR++
		case '_':
			countUnderscore++
		}
	}
	diff := countL - countR
	if diff < 0 {
		diff = -diff
	}
	return diff + countUnderscore
}
```

## 2837 — Total Traveled Distance

```go
package main

// LeetCode #2837: Total Traveled Distance
// https://leetcode.com/problems/total-traveled-distance/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(n)
// Note: SQL problem, adapted to Go. Computes total travel distance per user.

import "fmt"

func main() {
	rides := []struct {
		UserID   int
		Distance int
	}{
		{1, 10},
		{1, 15},
		{2, 20},
		{3, 0},
	}
	fmt.Println(TotalTraveledDistance(rides))
}

func TotalTraveledDistance(rides []struct {
	UserID   int
	Distance int
}) map[int]int {
	total := map[int]int{}
	for _, r := range rides {
		total[r.UserID] += r.Distance
	}
	return total
}
```

## 2839 — Check If Strings Can Be Made Equal With Operations I

```go
package main

// LeetCode #2839: Check if Strings Can be Made Equal With Operations I
// https://leetcode.com/problems/check-if-strings-can-be-made-equal-with-operations-i/
// Difficulty: Easy
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(CheckIfStringsCanBeMadeEqualWithOperationsI("abcd", "cdab"))
	fmt.Println(CheckIfStringsCanBeMadeEqualWithOperationsI("abcd", "dacb"))
}

func CheckIfStringsCanBeMadeEqualWithOperationsI(s1 string, s2 string) bool {
	// Can swap characters at even indices (0<->2) and odd indices (1<->3)
	// Check that multiset of chars at same parity positions match
	return (s1[0] == s2[0] || s1[0] == s2[2]) &&
		(s1[2] == s2[2] || s1[2] == s2[0]) &&
		(s1[1] == s2[1] || s1[1] == s2[3]) &&
		(s1[3] == s2[3] || s1[3] == s2[1])
}
```

## 2843 — Count Symmetric Integers

```go
package main

// LeetCode #2843: Count Symmetric Integers
// https://leetcode.com/problems/count-symmetric-integers/
// Difficulty: Easy
// Time: O(high - low) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(CountSymmetricIntegers(1, 100))
	fmt.Println(CountSymmetricIntegers(1200, 1230))
}

func CountSymmetricIntegers(low int, high int) int {
	count := 0
	for n := low; n <= high; n++ {
		if isSymmetric(n) {
			count++
		}
	}
	return count
}

func isSymmetric(n int) bool {
	s := fmt.Sprintf("%d", n)
	if len(s)%2 == 1 {
		return false
	}
	half := len(s) / 2
	sum1, sum2 := 0, 0
	for i := 0; i < half; i++ {
		sum1 += int(s[i] - '0')
		sum2 += int(s[i+half] - '0')
	}
	return sum1 == sum2
}
```

## 2848 — Points That Intersect With Cars

```go
package main

// LeetCode #2848: Points That Intersect With Cars
// https://leetcode.com/problems/points-that-intersect-with-cars/
// Difficulty: Easy
// Time: O(n * range) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(PointsThatIntersectWithCars([][]int{{3, 6}, {1, 5}, {4, 7}}))
	fmt.Println(PointsThatIntersectWithCars([][]int{{1, 3}, {5, 8}}))
}

func PointsThatIntersectWithCars(nums [][]int) int {
	points := make([]bool, 101)
	for _, car := range nums {
		for p := car[0]; p <= car[1]; p++ {
			points[p] = true
		}
	}
	count := 0
	for _, v := range points {
		if v {
			count++
		}
	}
	return count
}
```

## 2853 — Highest Salaries Difference

```go
package main

// LeetCode #2853: Highest Salaries Difference
// https://leetcode.com/problems/highest-salaries-difference/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(1)
// Note: SQL problem, adapted to Go. Difference between highest salaries in two departments.

import "fmt"

func main() {
	salaries := []struct {
		Name       string
		Salary     int
		Department string
	}{
		{"Alice", 100000, "Engineering"},
		{"Bob", 90000, "Engineering"},
		{"Charlie", 80000, "Marketing"},
		{"David", 95000, "Marketing"},
	}
	fmt.Println(HighestSalariesDifference(salaries))
}

func HighestSalariesDifference(salaries []struct {
	Name       string
	Salary     int
	Department string
}) int {
	maxEng, maxMkt := 0, 0
	for _, s := range salaries {
		if s.Department == "Engineering" && s.Salary > maxEng {
			maxEng = s.Salary
		} else if s.Department == "Marketing" && s.Salary > maxMkt {
			maxMkt = s.Salary
		}
	}
	if maxEng > maxMkt {
		return maxEng - maxMkt
	}
	return maxMkt - maxEng
}
```

## 2855 — Minimum Right Shifts To Sort The Array

```go
package main

// LeetCode #2855: Minimum Right Shifts to Sort the Array
// https://leetcode.com/problems/minimum-right-shifts-to-sort-the-array/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(MinimumRightShiftsToSortTheArray([]int{3, 4, 5, 1, 2}))
	fmt.Println(MinimumRightShiftsToSortTheArray([]int{1, 3, 5}))
}

func MinimumRightShiftsToSortTheArray(nums []int) int {
	n := len(nums)
	descentIdx := -1
	for i := 0; i < n-1; i++ {
		if nums[i] > nums[i+1] {
			if descentIdx != -1 {
				return -1
			}
			descentIdx = i
		}
	}
	if descentIdx == -1 {
		return 0
	}
	if nums[n-1] > nums[0] {
		return -1
	}
	return n - descentIdx - 1
}
```

## 2859 — Sum Of Values At Indices With K Set Bits

```go
package main

// LeetCode #2859: Sum of Values at Indices With K Set Bits
// https://leetcode.com/problems/sum-of-values-at-indices-with-k-set-bits/
// Difficulty: Easy
// Time: O(n * log n) | Space: O(1)

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(SumOfValuesAtIndicesWithKSetBits([]int{5, 10, 1, 5, 2}, 1))
	fmt.Println(SumOfValuesAtIndicesWithKSetBits([]int{4, 3, 2, 1}, 2))
}

func SumOfValuesAtIndicesWithKSetBits(nums []int, k int) int {
	sum := 0
	for i, v := range nums {
		if bits.OnesCount(uint(i)) == k {
			sum += v
		}
	}
	return sum
}
```

## 2864 — Maximum Odd Binary Number

```go
package main

// LeetCode #2864: Maximum Odd Binary Number
// https://leetcode.com/problems/maximum-odd-binary-number/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(MaximumOddBinaryNumber("010"))
	fmt.Println(MaximumOddBinaryNumber("0101"))
}

func MaximumOddBinaryNumber(s string) string {
	ones := strings.Count(s, "1")
	zeros := len(s) - ones
	return strings.Repeat("1", ones-1) + strings.Repeat("0", zeros) + "1"
}
```

## 2869 — Minimum Operations To Collect Elements

```go
package main

// LeetCode #2869: Minimum Operations to Collect Elements
// https://leetcode.com/problems/minimum-operations-to-collect-elements/
// Difficulty: Easy
// Time: O(n) | Space: O(k)

import "fmt"

func main() {
	fmt.Println(MinimumOperationsToCollectElements([]int{3, 1, 5, 4, 2}, 2))
	fmt.Println(MinimumOperationsToCollectElements([]int{3, 1, 5, 4, 2}, 5))
}

func MinimumOperationsToCollectElements(nums []int, k int) int {
	seen := make([]bool, k+1)
	collected := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] >= 1 && nums[i] <= k && !seen[nums[i]] {
			seen[nums[i]] = true
			collected++
		}
		if collected == k {
			return len(nums) - i
		}
	}
	return len(nums)
}
```

## 2873 — Maximum Value Of An Ordered Triplet I

```go
package main

// LeetCode #2873: Maximum Value of an Ordered Triplet I
// https://leetcode.com/problems/maximum-value-of-an-ordered-triplet-i/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(MaximumValueOfAnOrderedTripletI([]int{12, 6, 1, 2, 7}))
	fmt.Println(MaximumValueOfAnOrderedTripletI([]int{1, 10, 3, 2, 5}))
}

func MaximumValueOfAnOrderedTripletI(nums []int) int64 {
	n := len(nums)
	if n < 3 {
		return 0
	}

	maxNum := int64(nums[0])
	maxDiff := int64(nums[0] - nums[1])
	ans := int64(0)

	for i := 2; i < n; i++ {
		val := maxDiff * int64(nums[i])
		if val > ans {
			ans = val
		}
		if int64(nums[i-1]) > maxNum {
			maxNum = int64(nums[i-1])
		}
		if maxNum-int64(nums[i]) > maxDiff {
			maxDiff = maxNum - int64(nums[i])
		}
	}

	return ans
}
```

## 2877 — Create A Dataframe From List

```go
package main

// LeetCode #2877: Create a DataFrame from List
// https://leetcode.com/problems/create-a-dataframe-from-list/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we implement equivalent logic using a struct slice.

import "fmt"

func main() {
	// LeetCode name: createDataFrameFromList
	fmt.Println(CreateADataframeFromList([][]int{{1, 15}, {2, 11}, {3, 11}, {4, 20}}))
	fmt.Println(CreateADataframeFromList([][]int{{5, 25}, {6, 30}}))
}

// type DataFrame represents the solution output.
type DataFrame []struct {
	StudentID int
	Age       int
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: createDataFrameFromList
func CreateADataframeFromList(studentData [][]int) DataFrame {
	df := make(DataFrame, len(studentData))
	for i, row := range studentData {
		df[i] = struct {
			StudentID int
			Age       int
		}{StudentID: row[0], Age: row[1]}
	}
	return df
}
```

## 2878 — Get The Size Of A Dataframe

```go
package main

// LeetCode #2878: Get the Size of a DataFrame
// https://leetcode.com/problems/get-the-size-of-a-dataframe/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we return the dimensions [rows, cols].

import "fmt"

func main() {
	// LeetCode name: getDataframeSize
	fmt.Println(GetTheSizeOfADataframe([][]int{{1, 15}, {2, 11}, {3, 11}, {4, 20}})) // [4 2]
	fmt.Println(GetTheSizeOfADataframe([][]int{{5, 25}}))                             // [1 2]
}

// Time: O(1) | Space: O(1)
// LeetCode submission name: getDataframeSize
func GetTheSizeOfADataframe(df [][]int) []int {
	if len(df) == 0 {
		return []int{0, 0}
	}
	return []int{len(df), len(df[0])}
}
```

## 2879 — Display The First Three Rows

```go
package main

// LeetCode #2879: Display the First Three Rows
// https://leetcode.com/problems/display-the-first-three-rows/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we return the first 3 rows of the dataframe.

import "fmt"

func main() {
	// LeetCode name: selectFirstRows
	fmt.Println(DisplayTheFirstThreeRows([][]int{{1, 15}, {2, 11}, {3, 11}, {4, 20}}))
	// [[1 15] [2 11] [3 11]]

	fmt.Println(DisplayTheFirstThreeRows([][]int{{1, 15}}))
	// [[1 15]]
}

// Time: O(n) where n = min(3, rows) | Space: O(1)
// LeetCode submission name: selectFirstRows
func DisplayTheFirstThreeRows(df [][]int) [][]int {
	if len(df) > 3 {
		return df[:3]
	}
	return df
}
```

## 2880 — Select Data

```go
package main

// LeetCode #2880: Select Data
// https://leetcode.com/problems/select-data/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we select rows by student_id and return the age values.

import "fmt"

func main() {
	// LeetCode name: selectData
	fmt.Println(SelectData([][]int{{101, 15}, {102, 11}, {103, 11}}, 101)) // [15]
	fmt.Println(SelectData([][]int{{101, 20}, {102, 22}, {101, 21}}, 101)) // [20, 21]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: selectData
func SelectData(df [][]int, studentID int) []int {
	ages := []int{}
	for _, row := range df {
		if row[0] == studentID {
			ages = append(ages, row[1])
		}
	}
	return ages
}
```

## 2881 — Create A New Column

```go
package main

// LeetCode #2881: Create a New Column
// https://leetcode.com/problems/create-a-new-column/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we add a "grade" column computed from existing data.

import "fmt"

func main() {
	// LeetCode name: createBonusColumn
	fmt.Println(CreateANewColumn([][]int{{101, 15}, {102, 11}, {103, 20}}))
	// [[101 15 30] [102 11 22] [103 20 40]]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: createBonusColumn
func CreateANewColumn(df [][]int) [][]int {
	result := make([][]int, len(df))
	for i, row := range df {
		// bonus = salary * 2
		result[i] = []int{row[0], row[1], row[1] * 2}
	}
	return result
}
```

## 2882 — Drop Duplicate Rows

```go
package main

// LeetCode #2882: Drop Duplicate Rows
// https://leetcode.com/problems/drop-duplicate-rows/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we remove rows with duplicate emails.

import "fmt"

func main() {
	// LeetCode name: dropDuplicateEmails
	fmt.Println(DropDuplicateRows([][]string{{"1", "a@b.com"}, {"2", "c@d.com"}, {"3", "a@b.com"}}))
	// [[1 a@b.com] [2 c@d.com]]

	fmt.Println(DropDuplicateRows([][]string{{"1", "x@y.com"}, {"2", "x@y.com"}, {"3", "x@y.com"}}))
	// [[1 x@y.com]]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: dropDuplicateEmails
func DropDuplicateRows(df [][]string) [][]string {
	seen := make(map[string]bool)
	result := [][]string{}
	for _, row := range df {
		email := row[1]
		if !seen[email] {
			seen[email] = true
			result = append(result, row)
		}
	}
	return result
}
```

## 2883 — Drop Missing Data

```go
package main

// LeetCode #2883: Drop Missing Data
// https://leetcode.com/problems/drop-missing-data/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we remove rows where the name column is empty.

import "fmt"

func main() {
	// LeetCode name: dropMissingData
	fmt.Println(DropMissingData([][]string{{"1", "Alice", "15"}, {"2", "", "11"}, {"3", "Bob", "12"}}))
	// [[1 Alice 15] [3 Bob 12]]

	fmt.Println(DropMissingData([][]string{{"1", "", "10"}}))
	// []
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: dropMissingData
func DropMissingData(df [][]string) [][]string {
	result := [][]string{}
	for _, row := range df {
		if row[1] != "" {
			result = append(result, row)
		}
	}
	return result
}
```

## 2884 — Modify Columns

```go
package main

// LeetCode #2884: Modify Columns
// https://leetcode.com/problems/modify-columns/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we modify a column by multiplying salary by 2.

import "fmt"

func main() {
	// LeetCode name: modifySalaryColumn
	fmt.Println(ModifyColumns([][]int{{1, 100}, {2, 200}, {3, 300}}))
	// [[1 200] [2 400] [3 600]]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: modifySalaryColumn
func ModifyColumns(df [][]int) [][]int {
	result := make([][]int, len(df))
	for i, row := range df {
		result[i] = []int{row[0], row[1] * 2}
	}
	return result
}
```

## 2885 — Rename Columns

```go
package main

// LeetCode #2885: Rename Columns
// https://leetcode.com/problems/rename-columns/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we rename columns: id->student_id, first->first_name, last->last_name, age->age_in_years.

import "fmt"

func main() {
	// LeetCode name: renameColumns
	// Input: [id, first, last, age]
	fmt.Println(RenameColumns([][]string{{"1", "John", "Doe", "15"}, {"2", "Jane", "Smith", "20"}}))
	// [[1 John Doe 15] [2 Jane Smith 20]]
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: renameColumns
func RenameColumns(df [][]string) [][]string {
	// Column renaming is a metadata operation; data itself doesn't change.
	return df
}
```

## 2886 — Change Data Type

```go
package main

// LeetCode #2886: Change Data Type
// https://leetcode.com/problems/change-data-type/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we convert the grade column from float64 to int (truncation).

import "fmt"

func main() {
	// LeetCode name: changeDataType
	// Input: [student_id, grade (float)]
	fmt.Println(ChangeDataType([][]float64{{1, 3.5}, {2, 4.2}, {3, 2.8}}))
	// [[1 3] [2 4] [3 2]]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: changeDataType
func ChangeDataType(df [][]float64) [][]int {
	result := make([][]int, len(df))
	for i, row := range df {
		result[i] = []int{int(row[0]), int(row[1])}
	}
	return result
}
```

## 2887 — Fill Missing Data

```go
package main

// LeetCode #2887: Fill Missing Data
// https://leetcode.com/problems/fill-missing-data/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we fill missing (zero/empty) quantity values with 0.

import "fmt"

func main() {
	// LeetCode name: fillMissingValues
	// Input: [name, quantity]. Empty string means missing.
	fmt.Println(FillMissingData([][]string{{"A", "10"}, {"B", ""}, {"C", "5"}, {"D", ""}}))
	// [[A 10] [B 0] [C 5] [D 0]]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: fillMissingValues
func FillMissingData(df [][]string) [][]string {
	result := make([][]string, len(df))
	for i, row := range df {
		if row[1] == "" {
			result[i] = []string{row[0], "0"}
		} else {
			result[i] = row
		}
	}
	return result
}
```

## 2888 — Reshape Data Concatenate

```go
package main

// LeetCode #2888: Reshape Data: Concatenate
// https://leetcode.com/problems/reshape-data-concatenate/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we concatenate two dataframes vertically.

import "fmt"

func main() {
	// LeetCode name: concatenateDataFrames
	fmt.Println(ReshapeDataConcatenate([][]int{{1, 15}, {2, 11}}, [][]int{{3, 12}, {4, 14}}))
	// [[1 15] [2 11] [3 12] [4 14]]
}

// Time: O(n+m) | Space: O(n+m)
// LeetCode submission name: concatenateDataFrames
func ReshapeDataConcatenate(df1, df2 [][]int) [][]int {
	result := make([][]int, 0, len(df1)+len(df2))
	result = append(result, df1...)
	result = append(result, df2...)
	return result
}
```

## 2889 — Reshape Data Pivot

```go
package main

// LeetCode #2889: Reshape Data: Pivot
// https://leetcode.com/problems/reshape-data-pivot/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we pivot data from [month, city, temperature] to [month, city1, city2, ...].

import "fmt"

func main() {
	// LeetCode name: pivotTable
	// Input: [month, city, temperature]
	data := [][]string{
		{"January", "London", "5"},
		{"January", "Paris", "7"},
		{"February", "London", "6"},
		{"February", "Paris", "8"},
	}
	fmt.Println(ReshapeDataPivot(data))
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: pivotTable
func ReshapeDataPivot(data [][]string) map[string]map[string]string {
	result := make(map[string]map[string]string)
	for _, row := range data {
		month, city, temp := row[0], row[1], row[2]
		if result[month] == nil {
			result[month] = make(map[string]string)
		}
		result[month][city] = temp
	}
	return result
}
```

## 2890 — Reshape Data Melt

```go
package main

// LeetCode #2890: Reshape Data: Melt
// https://leetcode.com/problems/reshape-data-melt/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we melt from wide format [product, store1, store2] to
// long format [product, store, price].

import "fmt"

func main() {
	// LeetCode name: meltTable
	// Input: [product_id, store1_price, store2_price]
	fmt.Println(ReshapeDataMelt([][]string{{"P1", "100", "200"}, {"P2", "150", "250"}}))
	// [[P1 store1 100] [P1 store2 200] [P2 store1 150] [P2 store2 250]]
}

// Time: O(n * m) where m is number of store columns | Space: O(n * m)
// LeetCode submission name: meltTable
func ReshapeDataMelt(data [][]string) [][]string {
	result := [][]string{}
	for _, row := range data {
		productID := row[0]
		for j := 1; j < len(row); j++ {
			store := fmt.Sprintf("store%d", j)
			result = append(result, []string{productID, store, row[j]})
		}
	}
	return result
}
```

## 2891 — Method Chaining

```go
package main

// LeetCode #2891: Method Chaining
// https://leetcode.com/problems/method-chaining/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we filter animals with weight > 100, sort by weight,
// and rename the weight column.

import (
	"fmt"
	"sort"
)

func main() {
	// LeetCode name: methodChaining
	// Input: [name, species, age, weight]
	animals := [][]int{{1, 1, 5, 50}, {2, 1, 3, 120}, {3, 2, 4, 150}, {4, 1, 2, 80}}
	fmt.Println(MethodChaining(animals))
	// [[2 1 3 120] [3 2 4 150]]
}

// Time: O(n log n) | Space: O(n)
// LeetCode submission name: methodChaining
func MethodChaining(animals [][]int) [][]int {
	// Filter: weight > 100 (column index 3)
	filtered := [][]int{}
	for _, a := range animals {
		if a[3] > 100 {
			filtered = append(filtered, a)
		}
	}
	// Sort by weight ascending
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i][3] < filtered[j][3]
	})
	return filtered
}
```

## 2894 — Divisible And Non Divisible Sums Difference

```go
package main

// LeetCode #2894: Divisible and Non-divisible Sums Difference
// https://leetcode.com/problems/divisible-and-non-divisible-sums-difference/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: differenceOfSums
	fmt.Println(DivisibleAndNonDivisibleSumsDifference(10, 3)) // 19
	fmt.Println(DivisibleAndNonDivisibleSumsDifference(5, 6))  // 15
	fmt.Println(DivisibleAndNonDivisibleSumsDifference(5, 1))  // -15
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: differenceOfSums
func DivisibleAndNonDivisibleSumsDifference(n int, m int) int {
	num1 := 0 // sum of numbers not divisible by m
	num2 := 0 // sum of numbers divisible by m
	for i := 1; i <= n; i++ {
		if i%m == 0 {
			num2 += i
		} else {
			num1 += i
		}
	}
	return num1 - num2
}
```

## 2899 — Last Visited Integers

```go
package main

// LeetCode #2899: Last Visited Integers
// https://leetcode.com/problems/last-visited-integers/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: lastVisitedIntegers
	fmt.Println(LastVisitedIntegers([]int{1, 2, -1, -1, -1})) // [2, 1, -1]
	fmt.Println(LastVisitedIntegers([]int{1, -1, 2, -1, -1})) // [1, 2, 1]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: lastVisitedIntegers
func LastVisitedIntegers(nums []int) []int {
	seen := []int{}
	result := []int{}
	k := 0

	for _, num := range nums {
		if num != -1 {
			seen = append(seen, num)
			k = 0
		} else {
			k++
			if k <= len(seen) {
				result = append(result, seen[len(seen)-k])
			} else {
				result = append(result, -1)
			}
		}
	}
	return result
}
```

## 2900 — Longest Unequal Adjacent Groups Subsequence I

```go
package main

// LeetCode #2900: Longest Unequal Adjacent Groups Subsequence I
// https://leetcode.com/problems/longest-unequal-adjacent-groups-subsequence-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: getWordsInLongestSubsequence
	fmt.Println(LongestUnequalAdjacentGroupsSubsequenceI(3, []string{"e", "a", "b"}, []int{0, 0, 1})) // ["e","b"]
	fmt.Println(LongestUnequalAdjacentGroupsSubsequenceI(4, []string{"a", "b", "c", "d"}, []int{1, 0, 1, 0})) // ["a","b","c","d"]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: getWordsInLongestSubsequence
func LongestUnequalAdjacentGroupsSubsequenceI(n int, words []string, groups []int) []string {
	result := []string{words[0]}
	for i := 1; i < n; i++ {
		if groups[i] != groups[i-1] {
			result = append(result, words[i])
		}
	}
	return result
}
```

## 2903 — Find Indices With Index And Value Difference I

```go
package main

// LeetCode #2903: Find Indices With Index and Value Difference I
// https://leetcode.com/problems/find-indices-with-index-and-value-difference-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: findIndices
	fmt.Println(FindIndicesWithIndexAndValueDifferenceI([]int{5, 1, 4, 1}, 2, 4)) // [0,3]
	fmt.Println(FindIndicesWithIndexAndValueDifferenceI([]int{2, 1}, 0, 0))       // [0,0]
	fmt.Println(FindIndicesWithIndexAndValueDifferenceI([]int{1, 2, 3}, 2, 4))    // [-1,-1]
}

// Time: O(n^2) | Space: O(1)
// LeetCode submission name: findIndices
func FindIndicesWithIndexAndValueDifferenceI(nums []int, indexDifference int, valueDifference int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if abs(i-j) >= indexDifference && abs(nums[i]-nums[j]) >= valueDifference {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

## 2908 — Minimum Sum Of Mountain Triplets I

```go
package main

// LeetCode #2908: Minimum Sum of Mountain Triplets I
// https://leetcode.com/problems/minimum-sum-of-mountain-triplets-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: minimumSum
	fmt.Println(MinimumSumOfMountainTripletsI([]int{8, 6, 1, 5, 3})) // 9
	fmt.Println(MinimumSumOfMountainTripletsI([]int{5, 4, 8, 7, 10, 2})) // 13
	fmt.Println(MinimumSumOfMountainTripletsI([]int{6, 5, 4, 3, 4, 5})) // -1
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: minimumSum
func MinimumSumOfMountainTripletsI(nums []int) int {
	n := len(nums)
	if n < 3 {
		return -1
	}

	// prefix[i] = min value in nums[0..i]
	prefix := make([]int, n)
	prefix[0] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] < prefix[i-1] {
			prefix[i] = nums[i]
		} else {
			prefix[i] = prefix[i-1]
		}
	}

	// suffix[i] = min value in nums[i..n-1]
	suffix := make([]int, n)
	suffix[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if nums[i] < suffix[i+1] {
			suffix[i] = nums[i]
		} else {
			suffix[i] = suffix[i+1]
		}
	}

	minSum := -1
	for j := 1; j < n-1; j++ {
		if prefix[j-1] < nums[j] && suffix[j+1] < nums[j] {
			sum := prefix[j-1] + nums[j] + suffix[j+1]
			if minSum == -1 || sum < minSum {
				minSum = sum
			}
		}
	}
	return minSum
}
```

## 2913 — Subarrays Distinct Element Sum Of Squares I

```go
package main

// LeetCode #2913: Subarrays Distinct Element Sum of Squares I
// https://leetcode.com/problems/subarrays-distinct-element-sum-of-squares-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: sumCounts
	fmt.Println(SubarraysDistinctElementSumOfSquaresI([]int{1, 2, 1})) // 15
	fmt.Println(SubarraysDistinctElementSumOfSquaresI([]int{2, 2}))    // 3
}

// Time: O(n^2) | Space: O(n)
// LeetCode submission name: sumCounts
func SubarraysDistinctElementSumOfSquaresI(nums []int) int {
	n := len(nums)
	total := 0
	for i := 0; i < n; i++ {
		seen := make(map[int]bool)
		distinct := 0
		for j := i; j < n; j++ {
			if !seen[nums[j]] {
				seen[nums[j]] = true
				distinct++
			}
			total += distinct * distinct
		}
	}
	return total
}
```

## 2917 — Find The K Or Of An Array

```go
package main

// LeetCode #2917: Find the K-or of an Array
// https://leetcode.com/problems/find-the-k-or-of-an-array/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: findKOr
	fmt.Println(FindTheKOrOfAnArray([]int{7, 12, 9, 8, 9, 15}, 4)) // 9
	fmt.Println(FindTheKOrOfAnArray([]int{2, 12, 1, 11, 4, 5}, 6)) // 0
	fmt.Println(FindTheKOrOfAnArray([]int{10, 8, 5, 9, 11, 6, 8}, 1)) // 15
}

// Time: O(n * 32) | Space: O(1)
// LeetCode submission name: findKOr
func FindTheKOrOfAnArray(nums []int, k int) int {
	result := 0
	for bit := 0; bit < 32; bit++ {
		count := 0
		for _, num := range nums {
			if num&(1<<bit) != 0 {
				count++
			}
		}
		if count >= k {
			result |= (1 << bit)
		}
	}
	return result
}
```

## 2923 — Find Champion I

```go
package main

// LeetCode #2923: Find Champion I
// https://leetcode.com/problems/find-champion-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: findChampion
	fmt.Println(FindChampionI([][]int{{0, 1}, {0, 0}}))          // 0
	fmt.Println(FindChampionI([][]int{{0, 0, 1}, {1, 0, 1}, {0, 0, 0}})) // 1
}

// Time: O(n^2) | Space: O(1)
// LeetCode submission name: findChampion
func FindChampionI(grid [][]int) int {
	n := len(grid)
	for i := 0; i < n; i++ {
		isChampion := true
		for j := 0; j < n; j++ {
			if i != j && grid[i][j] != 1 {
				isChampion = false
				break
			}
		}
		if isChampion {
			return i
		}
	}
	return -1
}
```

## 2928 — Distribute Candies Among Children I

```go
package main

// LeetCode #2928: Distribute Candies Among Children I
// https://leetcode.com/problems/distribute-candies-among-children-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: distributeCandies
	fmt.Println(DistributeCandiesAmongChildrenI(5, 2)) // 3
	fmt.Println(DistributeCandiesAmongChildrenI(3, 3)) // 10
}

// Time: O(limit^2) | Space: O(1)
// LeetCode submission name: distributeCandies
func DistributeCandiesAmongChildrenI(n int, limit int) int {
	ways := 0
	for a := 0; a <= limit && a <= n; a++ {
		for b := 0; b <= limit && a+b <= n; b++ {
			c := n - a - b
			if c <= limit {
				ways++
			}
		}
	}
	return ways
}
```

## 2932 — Maximum Strong Pair Xor I

```go
package main

// LeetCode #2932: Maximum Strong Pair XOR I
// https://leetcode.com/problems/maximum-strong-pair-xor-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: maximumStrongPairXor
	fmt.Println(MaximumStrongPairXorI([]int{1, 2, 3, 4, 5})) // 7
	fmt.Println(MaximumStrongPairXorI([]int{10, 100}))        // 0
	fmt.Println(MaximumStrongPairXorI([]int{5, 6, 25, 30}))   // 7
}

// Time: O(n^2) | Space: O(1)
// LeetCode submission name: maximumStrongPairXor
func MaximumStrongPairXorI(nums []int) int {
	n := len(nums)
	maxXor := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			x, y := nums[i], nums[j]
			if abs(x-y) <= min(x, y) {
				if x^y > maxXor {
					maxXor = x ^ y
				}
			}
		}
	}
	return maxXor
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

## 2937 — Make Three Strings Equal

```go
package main

// LeetCode #2937: Make Three Strings Equal
// https://leetcode.com/problems/make-three-strings-equal/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: findMinimumOperations
	fmt.Println(MakeThreeStringsEqual("abc", "abb", "ab")) // 2
	fmt.Println(MakeThreeStringsEqual("dac", "bac", "cac")) // -1
	fmt.Println(MakeThreeStringsEqual("a", "a", "a"))       // 0
}

// Time: O(min(len(s1), len(s2), len(s3))) | Space: O(1)
// LeetCode submission name: findMinimumOperations
func MakeThreeStringsEqual(s1 string, s2 string, s3 string) int {
	// Find longest common prefix length
	i := 0
	for i < len(s1) && i < len(s2) && i < len(s3) {
		if s1[i] != s2[i] || s1[i] != s3[i] {
			break
		}
		i++
	}
	if i == 0 {
		return -1
	}
	return (len(s1) - i) + (len(s2) - i) + (len(s3) - i)
}
```

## 2942 — Find Words Containing Character

```go
package main

// LeetCode #2942: Find Words Containing Character
// https://leetcode.com/problems/find-words-containing-character/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: findWordsContaining
	fmt.Println(FindWordsContainingCharacter([]string{"leet", "code"}, 'e')) // [0, 1]
	fmt.Println(FindWordsContainingCharacter([]string{"abc", "bcd", "aaaa", "cbc"}, 'a')) // [0, 2]
	fmt.Println(FindWordsContainingCharacter([]string{"abc", "bcd", "aaaa", "cbc"}, 'z')) // []
}

// Time: O(n * m) where m is max word length | Space: O(1) excluding output
// LeetCode submission name: findWordsContaining
func FindWordsContainingCharacter(words []string, x byte) []int {
	result := []int{}
	for i, word := range words {
		for j := 0; j < len(word); j++ {
			if word[j] == x {
				result = append(result, i)
				break
			}
		}
	}
	return result
}
```

## 2946 — Matrix Similarity After Cyclic Shifts

```go
package main

// LeetCode #2946: Matrix Similarity After Cyclic Shifts
// https://leetcode.com/problems/matrix-similarity-after-cyclic-shifts/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: areSimilar
	fmt.Println(MatrixSimilarityAfterCyclicShifts([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 4)) // false
	fmt.Println(MatrixSimilarityAfterCyclicShifts([][]int{{1, 2, 1, 2}, {5, 5, 5, 5}, {6, 3, 6, 3}}, 2)) // true
	fmt.Println(MatrixSimilarityAfterCyclicShifts([][]int{{2, 2}, {2, 2}}, 3)) // true
}

// Time: O(n * m) | Space: O(1)
// LeetCode submission name: areSimilar
func MatrixSimilarityAfterCyclicShifts(mat [][]int, k int) bool {
	for _, row := range mat {
		m := len(row)
		if m == 0 {
			continue
		}
		shift := k % m
		if shift == 0 {
			continue
		}
		for j := 0; j < m; j++ {
			// After left shift by k, row[(j + k) % m] moves to position j
			if row[j] != row[(j+shift)%m] {
				return false
			}
		}
	}
	return true
}
```

## 2951 — Find The Peaks

```go
package main

// LeetCode #2951: Find the Peaks
// https://leetcode.com/problems/find-the-peaks/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: findPeaks
	fmt.Println(FindThePeaks([]int{2, 4, 4}))    // []
	fmt.Println(FindThePeaks([]int{1, 4, 3, 8, 5})) // [1, 3]
}

// Time: O(n) | Space: O(1) excluding output
// LeetCode submission name: findPeaks
func FindThePeaks(mountain []int) []int {
	result := []int{}
	for i := 1; i < len(mountain)-1; i++ {
		if mountain[i] > mountain[i-1] && mountain[i] > mountain[i+1] {
			result = append(result, i)
		}
	}
	return result
}
```

## 2956 — Find Common Elements Between Two Arrays

```go
package main

// LeetCode #2956: Find Common Elements Between Two Arrays
// https://leetcode.com/problems/find-common-elements-between-two-arrays/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: findIntersectionValues
	fmt.Println(FindCommonElementsBetweenTwoArrays([]int{4, 3, 2, 3, 1}, []int{2, 2, 5, 2, 3, 6})) // [3, 4]
	fmt.Println(FindCommonElementsBetweenTwoArrays([]int{3, 4, 2, 3}, []int{1, 5})) // [0, 0]
}

// Time: O(n + m) | Space: O(n + m)
// LeetCode submission name: findIntersectionValues
func FindCommonElementsBetweenTwoArrays(nums1 []int, nums2 []int) []int {
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)
	for _, v := range nums1 {
		set1[v] = true
	}
	for _, v := range nums2 {
		set2[v] = true
	}
	count1 := 0
	for _, v := range nums1 {
		if set2[v] {
			count1++
		}
	}
	count2 := 0
	for _, v := range nums2 {
		if set1[v] {
			count2++
		}
	}
	return []int{count1, count2}
}
```

## 2960 — Count Tested Devices After Test Operations

```go
package main

// LeetCode #2960: Count Tested Devices After Test Operations
// https://leetcode.com/problems/count-tested-devices-after-test-operations/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: countTestedDevices
	fmt.Println(CountTestedDevicesAfterTestOperations([]int{1, 1, 2, 1, 3})) // 3
	fmt.Println(CountTestedDevicesAfterTestOperations([]int{0, 1, 2}))       // 2
}

// Time: O(n^2) | Space: O(1)
// LeetCode submission name: countTestedDevices
func CountTestedDevicesAfterTestOperations(batteryPercentages []int) int {
	n := len(batteryPercentages)
	count := 0
	for i := 0; i < n; i++ {
		if batteryPercentages[i] > 0 {
			count++
			for j := i + 1; j < n; j++ {
				if batteryPercentages[j] > 0 {
					batteryPercentages[j]--
				}
			}
		}
	}
	return count
}
```

## 2965 — Find Missing And Repeated Values

```go
package main

// LeetCode #2965: Find Missing and Repeated Values
// https://leetcode.com/problems/find-missing-and-repeated-values/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: findMissingAndRepeatedValues
	fmt.Println(FindMissingAndRepeatedValues([][]int{{1, 3}, {2, 2}})) // [2, 4]
	fmt.Println(FindMissingAndRepeatedValues([][]int{{9, 1, 7}, {8, 9, 2}, {3, 4, 6}})) // [9, 5]
}

// Time: O(n^2) | Space: O(n^2)
// LeetCode submission name: findMissingAndRepeatedValues
func FindMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	total := n * n
	seen := make(map[int]int)
	repeated := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			val := grid[i][j]
			seen[val]++
			if seen[val] == 2 {
				repeated = val
			}
		}
	}
	for i := 1; i <= total; i++ {
		if seen[i] == 0 {
			return []int{repeated, i}
		}
	}
	return []int{repeated, 0}
}
```

## 2970 — Count The Number Of Incremovable Subarrays I

```go
package main

// LeetCode #2970: Count the Number of Incremovable Subarrays I
// https://leetcode.com/problems/count-the-number-of-incremovable-subarrays-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: incremovableSubarrayCount
	fmt.Println(CountTheNumberOfIncremovableSubarraysI([]int{1, 2, 3, 4})) // 10
	fmt.Println(CountTheNumberOfIncremovableSubarraysI([]int{6, 5, 7, 8})) // 7
	fmt.Println(CountTheNumberOfIncremovableSubarraysI([]int{8, 7, 6, 6})) // 3
}

// Time: O(n^3) | Space: O(n)
// LeetCode submission name: incremovableSubarrayCount
func CountTheNumberOfIncremovableSubarraysI(nums []int) int {
	n := len(nums)
	count := 0

	for l := 0; l < n; l++ {
		for r := l; r < n; r++ {
			// Check if array without nums[l..r] is strictly increasing
			if isStrictlyIncreasing(nums, l, r) {
				count++
			}
		}
	}
	return count
}

func isStrictlyIncreasing(nums []int, l, r int) bool {
	prev := -1
	for i := 0; i < len(nums); i++ {
		if i >= l && i <= r {
			continue
		}
		if nums[i] <= prev {
			return false
		}
		prev = nums[i]
	}
	return true
}
```

## 2974 — Minimum Number Game

```go
package main

// LeetCode #2974: Minimum Number Game
// https://leetcode.com/problems/minimum-number-game/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	// LeetCode name: numberGame
	fmt.Println(MinimumNumberGame([]int{5, 4, 2, 3})) // [3, 2, 5, 4]
	fmt.Println(MinimumNumberGame([]int{2, 5}))       // [5, 2]
}

// Time: O(n log n) | Space: O(n)
// LeetCode submission name: numberGame
func MinimumNumberGame(nums []int) []int {
	sort.Ints(nums)
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i += 2 {
		result[i] = nums[i+1]
		result[i+1] = nums[i]
	}
	return result
}
```

## 2980 — Check If Bitwise Or Has Trailing Zeros

```go
package main

// LeetCode #2980: Check if Bitwise OR Has Trailing Zeros
// https://leetcode.com/problems/check-if-bitwise-or-has-trailing-zeros/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: hasTrailingZeros
	fmt.Println(CheckIfBitwiseOrHasTrailingZeros([]int{1, 2, 3, 4, 5})) // true
	fmt.Println(CheckIfBitwiseOrHasTrailingZeros([]int{2, 4, 8, 16}))   // true
	fmt.Println(CheckIfBitwiseOrHasTrailingZeros([]int{1, 3, 5, 7, 9})) // false
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: hasTrailingZeros
func CheckIfBitwiseOrHasTrailingZeros(nums []int) bool {
	evenCount := 0
	for _, num := range nums {
		if num%2 == 0 {
			evenCount++
			if evenCount >= 2 {
				return true
			}
		}
	}
	return false
}
```

## 2985 — Calculate Compressed Mean

```go
package main

// LeetCode #2985: Calculate Compressed Mean
// https://leetcode.com/problems/calculate-compressed-mean/
// Difficulty: Easy [Paid - Database]
//
// Note: This is a SQL problem on LeetCode. In Go, we implement the
// equivalent weighted average calculation.

import "fmt"

func main() {
	// LeetCode name: calculateCompressedMean
	// Input: item_count, order_occurrences pairs
	fmt.Println(CalculateCompressedMean([][]int{{1, 500}, {2, 1000}, {3, 800}, {4, 200}}))
	fmt.Println(CalculateCompressedMean([][]int{{1, 5}, {2, 10}, {3, 5}})) // 2
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: calculateCompressedMean
// data[i] = [item_count, order_occurrences]
func CalculateCompressedMean(data [][]int) float64 {
	var totalItems, totalOrders float64
	for _, row := range data {
		totalItems += float64(row[0] * row[1])
		totalOrders += float64(row[1])
	}
	result := totalItems / totalOrders
	// Round to 2 decimal places
	return float64(int(result*100+0.5)) / 100
}
```

## 2987 — Find Expensive Cities

```go
package main

// LeetCode #2987: Find Expensive Cities
// https://leetcode.com/problems/find-expensive-cities/
// Difficulty: Easy [Paid - Database]
//
// Note: This is a SQL problem on LeetCode. In Go, we implement the
// equivalent logic: find cities with avg price > overall avg price.

import (
	"fmt"
	"sort"
)

func main() {
	// LeetCode name: findExpensiveCities
	// Input: list of [city, price] pairs
	fmt.Println(FindExpensiveCities([][]string{
		{"New York", "1000"},
		{"New York", "1200"},
		{"Los Angeles", "800"},
		{"Los Angeles", "600"},
		{"Chicago", "500"},
	})) // [New York]

	fmt.Println(FindExpensiveCities([][]string{
		{"A", "200"},
		{"B", "100"},
		{"B", "100"},
	})) // [A B]
}

// Time: O(n log n) | Space: O(n)
// LeetCode submission name: findExpensiveCities
func FindExpensiveCities(listings [][]string) []string {
	// Calculate overall average
	var totalPrice float64
	for _, row := range listings {
		price := parsePrice(row[1])
		totalPrice += price
	}
	avgPrice := totalPrice / float64(len(listings))

	// Calculate per-city average
	citySum := make(map[string]float64)
	cityCount := make(map[string]int)
	for _, row := range listings {
		city := row[0]
		price := parsePrice(row[1])
		citySum[city] += price
		cityCount[city]++
	}

	result := []string{}
	for city, sum := range citySum {
		if sum/float64(cityCount[city]) > avgPrice {
			result = append(result, city)
		}
	}
	sort.Strings(result)
	return result
}

func parsePrice(s string) float64 {
	// Simple string to float conversion
	var result float64
	var neg bool
	for i := 0; i < len(s); i++ {
		if s[i] == '-' {
			neg = true
		} else if s[i] >= '0' && s[i] <= '9' {
			result = result*10 + float64(s[i]-'0')
		} else if s[i] == '.' {
			// Decimal part
			div := 10.0
			for j := i + 1; j < len(s); j++ {
				if s[j] >= '0' && s[j] <= '9' {
					result += float64(s[j]-'0') / div
					div *= 10
				}
			}
			break
		}
	}
	if neg {
		return -result
	}
	return result
}
```

## 2990 — Loan Types

```go
package main

// LeetCode #2990: Loan Types
// https://leetcode.com/problems/loan-types/
// Difficulty: Easy [Paid - Database]
//
// Note: This is a SQL problem on LeetCode. In Go, we implement the
// equivalent logic: find users who have both 'Refinance' and 'Mortgage' loans.

import (
	"fmt"
	"sort"
)

func main() {
	// LeetCode name: loanTypes
	fmt.Println(LoanTypes([][]string{
		{"101", "Mortgage"},
		{"101", "Refinance"},
		{"102", "Mortgage"},
		{"103", "Refinance"},
		{"103", "Student"},
	})) // [101]

	fmt.Println(LoanTypes([][]string{
		{"1", "Refinance"},
		{"1", "Mortgage"},
		{"2", "Refinance"},
		{"2", "Student"},
		{"3", "Mortgage"},
	})) // [1]
}

// Time: O(n log n) | Space: O(n)
// LeetCode submission name: loanTypes
func LoanTypes(loans [][]string) []string {
	userLoans := make(map[string]map[string]bool)
	for _, row := range loans {
		userID := row[0]
		loanType := row[1]
		if userLoans[userID] == nil {
			userLoans[userID] = make(map[string]bool)
		}
		userLoans[userID][loanType] = true
	}

	result := []string{}
	for userID, types := range userLoans {
		if types["Refinance"] && types["Mortgage"] {
			result = append(result, userID)
		}
	}
	sort.Strings(result)
	return result
}
```

## 2996 — Smallest Missing Integer Greater Than Sequential Prefix Sum

```go
package main

// LeetCode #2996: Smallest Missing Integer Greater Than Sequential Prefix Sum
// https://leetcode.com/problems/smallest-missing-integer-greater-than-sequential-prefix-sum/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: missingInteger
	fmt.Println(SmallestMissingIntegerGreaterThanSequentialPrefixSum([]int{1, 2, 3, 2, 5}))        // 6
	fmt.Println(SmallestMissingIntegerGreaterThanSequentialPrefixSum([]int{3, 4, 5, 1, 12, 14, 13})) // 15
	fmt.Println(SmallestMissingIntegerGreaterThanSequentialPrefixSum([]int{4, 5, 6, 7, 8, 9, 10, 11})) // 60
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: missingInteger
func SmallestMissingIntegerGreaterThanSequentialPrefixSum(nums []int) int {
	// Find longest sequential prefix (each element = prev + 1)
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			sum += nums[i]
		} else {
			break
		}
	}

	// Find smallest missing >= sum
	seen := make(map[int]bool)
	for _, v := range nums {
		seen[v] = true
	}
	for seen[sum] {
		sum++
	}
	return sum
}
```

## 3000 — Maximum Area Of Longest Diagonal Rectangle

```go
package main

// LeetCode #3000: Maximum Area of Longest Diagonal Rectangle
// https://leetcode.com/problems/maximum-area-of-longest-diagonal-rectangle/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: areaOfMaxDiagonal
	fmt.Println(MaximumAreaOfLongestDiagonalRectangle([][]int{{9, 3}, {8, 6}})) // 48
	fmt.Println(MaximumAreaOfLongestDiagonalRectangle([][]int{{3, 4}, {4, 3}})) // 12
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: areaOfMaxDiagonal
func MaximumAreaOfLongestDiagonalRectangle(dimensions [][]int) int {
	maxDiag := 0
	maxArea := 0
	for _, dim := range dimensions {
		l, w := dim[0], dim[1]
		diagSq := l*l + w*w
		area := l * w
		if diagSq > maxDiag || (diagSq == maxDiag && area > maxArea) {
			maxDiag = diagSq
			maxArea = area
		}
	}
	return maxArea
}
```

## 3005 — Count Elements With Maximum Frequency

```go
package main

// LeetCode #3005: Count Elements With Maximum Frequency
// https://leetcode.com/problems/count-elements-with-maximum-frequency/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: maxFrequencyElements
	fmt.Println(CountElementsWithMaximumFrequency([]int{1, 2, 2, 3, 1, 4})) // 4
	fmt.Println(CountElementsWithMaximumFrequency([]int{1, 2, 3, 4, 5}))    // 5
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: maxFrequencyElements
func CountElementsWithMaximumFrequency(nums []int) int {
	freq := make(map[int]int)
	maxFreq := 0
	for _, v := range nums {
		freq[v]++
		if freq[v] > maxFreq {
			maxFreq = freq[v]
		}
	}
	total := 0
	for _, f := range freq {
		if f == maxFreq {
			total += f
		}
	}
	return total
}
```

## 3010 — Divide An Array Into Subarrays With Minimum Cost I

```go
package main

// LeetCode #3010: Divide an Array Into Subarrays With Minimum Cost I
// https://leetcode.com/problems/divide-an-array-into-subarrays-with-minimum-cost-i/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	// LeetCode name: minimumCost
	fmt.Println(DivideAnArrayIntoSubarraysWithMinimumCostI([]int{1, 2, 3, 12})) // 6
	fmt.Println(DivideAnArrayIntoSubarraysWithMinimumCostI([]int{5, 4, 3, 2, 1})) // 8
}

// Time: O(n log n) | Space: O(1)
// LeetCode submission name: minimumCost
// Cost = nums[0] + sum of two smallest elements from nums[1:]
func DivideAnArrayIntoSubarraysWithMinimumCostI(nums []int) int {
	// First subarray starts at nums[0], so nums[0] is always included.
	// For remaining subarrays, we pick the two smallest elements.
	rest := nums[1:]
	sort.Ints(rest)
	return nums[0] + rest[0] + rest[1]
}
```

## 3014 — Minimum Number Of Pushes To Type Word I

```go
package main

// LeetCode #3014: Minimum Number of Pushes to Type Word I
// https://leetcode.com/problems/minimum-number-of-pushes-to-type-word-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: minimumPushes
	fmt.Println(MinimumNumberOfPushesToTypeWordI("abcde")) // 5
	fmt.Println(MinimumNumberOfPushesToTypeWordI("xycdefghij")) // 12
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: minimumPushes
// Each key can hold up to 4 distinct letters. First 8 letters cost 1 push each,
// next 8 cost 2 pushes, etc.
func MinimumNumberOfPushesToTypeWordI(word string) int {
	n := len(word)
	pushes := 0
	// First 8 distinct chars: 1 push each
	// Next 8: 2 pushes each, etc.
	for i := 0; i < n; i++ {
		pushes += (i / 8) + 1
	}
	return pushes
}
```

## 3019 — Number Of Changing Keys

```go
package main

// LeetCode #3019: Number of Changing Keys
// https://leetcode.com/problems/number-of-changing-keys/
// Difficulty: Easy

import "fmt"
import "unicode"

func main() {
	// LeetCode name: countKeyChanges
	fmt.Println(NumberOfChangingKeys("aAbBcC")) // 2
	fmt.Println(NumberOfChangingKeys("AaAaAaaA")) // 0
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: countKeyChanges
func NumberOfChangingKeys(s string) int {
	count := 0
	for i := 1; i < len(s); i++ {
		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[i-1])) {
			count++
		}
	}
	return count
}
```

## 3024 — Type Of Triangle

```go
package main

// LeetCode #3024: Type of Triangle
// https://leetcode.com/problems/type-of-triangle/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	// LeetCode name: triangleType
	fmt.Println(TypeOfTriangle([]int{3, 3, 3})) // equilateral
	fmt.Println(TypeOfTriangle([]int{3, 4, 5})) // scalene
	fmt.Println(TypeOfTriangle([]int{3, 3, 5})) // isosceles
	fmt.Println(TypeOfTriangle([]int{1, 2, 3})) // none
}

// Time: O(1) | Space: O(1)
// LeetCode submission name: triangleType
func TypeOfTriangle(nums []int) string {
	sort.Ints(nums)
	a, b, c := nums[0], nums[1], nums[2]

	// Check if valid triangle
	if a+b <= c {
		return "none"
	}

	if a == b && b == c {
		return "equilateral"
	}
	if a == b || b == c || a == c {
		return "isosceles"
	}
	return "scalene"
}
```

## 3028 — Ant On The Boundary

```go
package main

// LeetCode #3028: Ant on the Boundary
// https://leetcode.com/problems/ant-on-the-boundary/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: returnToBoundaryCount
	fmt.Println(AntOnTheBoundary([]int{2, 3, -5})) // 1
	fmt.Println(AntOnTheBoundary([]int{3, 2, -3, -2})) // 1
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: returnToBoundaryCount
func AntOnTheBoundary(nums []int) int {
	pos := 0
	count := 0
	for _, v := range nums {
		pos += v
		if pos == 0 {
			count++
		}
	}
	return count
}
```

## 3032 — Count Numbers With Unique Digits Ii

```go
package main

// LeetCode #3032: Count Numbers With Unique Digits II
// https://leetcode.com/problems/count-numbers-with-unique-digits-ii/
// Difficulty: Easy [Paid]
//
// Note: This is a premium problem. In Go, we implement the equivalent logic.

import "fmt"

func main() {
	// LeetCode name: numberCount
	fmt.Println(CountNumbersWithUniqueDigitsIi(1, 20))  // 19
	fmt.Println(CountNumbersWithUniqueDigitsIi(9, 19))  // 10
}

// Time: O((b-a) * log(b)) | Space: O(1)
// LeetCode submission name: numberCount
func CountNumbersWithUniqueDigitsIi(a int, b int) int {
	count := 0
	for i := a; i <= b; i++ {
		if hasUniqueDigits(i) {
			count++
		}
	}
	return count
}

func hasUniqueDigits(n int) bool {
	seen := make(map[int]bool)
	for n > 0 {
		digit := n % 10
		if seen[digit] {
			return false
		}
		seen[digit] = true
		n /= 10
	}
	return true
}
```

## 3033 — Modify The Matrix

```go
package main

// LeetCode #3033: Modify the Matrix
// https://leetcode.com/problems/modify-the-matrix/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: modifiedMatrix
	fmt.Println(ModifyTheMatrix([][]int{{1, 2, -1}, {4, -1, 6}, {7, 8, 9}}))
	// [[1 2 9] [4 8 6] [7 8 9]]
}

// Time: O(n * m) | Space: O(1) (modifies in place)
// LeetCode submission name: modifiedMatrix
func ModifyTheMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return matrix
	}
	m, n := len(matrix), len(matrix[0])

	// Find max in each column
	colMax := make([]int, n)
	for j := 0; j < n; j++ {
		maxVal := -1
		for i := 0; i < m; i++ {
			if matrix[i][j] > maxVal {
				maxVal = matrix[i][j]
			}
		}
		colMax[j] = maxVal
	}

	// Replace -1 values
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == -1 {
				matrix[i][j] = colMax[j]
			}
		}
	}
	return matrix
}
```

## 3038 — Maximum Number Of Operations With The Same Score I

```go
package main

// LeetCode #3038: Maximum Number of Operations With the Same Score I
// https://leetcode.com/problems/maximum-number-of-operations-with-the-same-score-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: maxOperations
	fmt.Println(MaximumNumberOfOperationsWithTheSameScoreI([]int{3, 2, 1, 4, 5})) // 2
	fmt.Println(MaximumNumberOfOperationsWithTheSameScoreI([]int{3, 2, 6, 1, 4})) // 1
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: maxOperations
func MaximumNumberOfOperationsWithTheSameScoreI(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	score := nums[0] + nums[1]
	count := 1
	for i := 2; i+1 < len(nums); i += 2 {
		if nums[i]+nums[i+1] == score {
			count++
		} else {
			break
		}
	}
	return count
}
```

## 3042 — Count Prefix And Suffix Pairs I

```go
package main

// LeetCode #3042: Count Prefix and Suffix Pairs I
// https://leetcode.com/problems/count-prefix-and-suffix-pairs-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: countPrefixSuffixPairs
	fmt.Println(CountPrefixAndSuffixPairsI([]string{"a", "aba", "ababa", "aa"})) // 4
	fmt.Println(CountPrefixAndSuffixPairsI([]string{"pa", "papa", "ma", "mama"})) // 2
}

// Time: O(n^2 * m) where m is max word length | Space: O(1)
// LeetCode submission name: countPrefixSuffixPairs
func CountPrefixAndSuffixPairsI(words []string) int {
	n := len(words)
	count := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isPrefixAndSuffix(words[i], words[j]) {
				count++
			}
		}
	}
	return count
}

func isPrefixAndSuffix(a, b string) bool {
	if len(a) > len(b) {
		return false
	}
	// Check prefix
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	// Check suffix
	for i := 0; i < len(a); i++ {
		if a[i] != b[len(b)-len(a)+i] {
			return false
		}
	}
	return true
}
```

## 3046 — Split The Array

```go
package main

// LeetCode #3046: Split the Array
// https://leetcode.com/problems/split-the-array/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: isPossibleToSplit
	fmt.Println(SplitTheArray([]int{1, 1, 2, 2, 3, 4})) // true
	fmt.Println(SplitTheArray([]int{1, 1, 1, 1}))       // false
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: isPossibleToSplit
// Each number can appear at most twice (once in each half of the split)
func SplitTheArray(nums []int) bool {
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
		if freq[v] > 2 {
			return false
		}
	}
	return true
}
```

## 3051 — Find Candidates For Data Scientist Position

```go
package main

// LeetCode #3051: Find Candidates for Data Scientist Position
// https://leetcode.com/problems/find-candidates-for-data-scientist-position/
// Difficulty: Easy [Paid - Database]
//
// Note: This is a SQL problem on LeetCode. In Go, we implement the
// equivalent logic: find candidates with Python, Tableau, and PostgreSQL skills.

import (
	"fmt"
	"sort"
)

func main() {
	// LeetCode name: findCandidates
	// Input: [candidate_id, skill]
	candidates := [][]string{
		{"101", "Python"},
		{"101", "Tableau"},
		{"101", "PostgreSQL"},
		{"102", "Python"},
		{"102", "Tableau"},
		{"103", "Python"},
		{"103", "PostgreSQL"},
	}
	fmt.Println(FindCandidatesForDataScientistPosition(candidates))
	// [101]
}

// Time: O(n log n) | Space: O(n)
// LeetCode submission name: findCandidates
func FindCandidatesForDataScientistPosition(candidates [][]string) []int {
	skills := make(map[int]map[string]bool)

	for _, row := range candidates {
		id := parseInt(row[0])
		skill := row[1]
		if skills[id] == nil {
			skills[id] = make(map[string]bool)
		}
		skills[id][skill] = true
	}

	result := []int{}
	for id, s := range skills {
		if s["Python"] && s["Tableau"] && s["PostgreSQL"] {
			result = append(result, id)
		}
	}
	sort.Ints(result)
	return result
}

func parseInt(s string) int {
	n := 0
	for i := 0; i < len(s); i++ {
		n = n*10 + int(s[i]-'0')
	}
	return n
}
```

## 3053 — Classifying Triangles By Lengths

```go
package main

// LeetCode #3053: Classifying Triangles by Lengths
// https://leetcode.com/problems/classifying-triangles-by-lengths/
// Difficulty: Easy [Paid - Database]
//
// Note: This is a SQL problem on LeetCode. In Go, we implement the
// equivalent logic: classify triangles by side lengths.

import (
	"fmt"
	"sort"
)

func main() {
	// LeetCode name: triangleType
	// Input: [side_a, side_b, side_c]
	fmt.Println(ClassifyingTrianglesByLengths([][]int{{3, 3, 3}, {3, 4, 5}, {3, 3, 5}, {1, 2, 3}}))
	// [Equilateral Scalene Isosceles None]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: triangleType
func ClassifyingTrianglesByLengths(triangles [][]int) []string {
	result := make([]string, len(triangles))
	for i, t := range triangles {
		sides := []int{t[0], t[1], t[2]}
		sort.Ints(sides)
		a, b, c := sides[0], sides[1], sides[2]

		if a+b <= c {
			result[i] = "None"
		} else if a == b && b == c {
			result[i] = "Equilateral"
		} else if a == b || b == c || a == c {
			result[i] = "Isosceles"
		} else {
			result[i] = "Scalene"
		}
	}
	return result
}
```

## 3059 — Find All Unique Email Domains

```go
package main

// LeetCode #3059: Find All Unique Email Domains
// https://leetcode.com/problems/find-all-unique-email-domains/
// Difficulty: Easy [Paid - Database]
//
// Note: This is a SQL problem on LeetCode. In Go, we implement the
// equivalent logic: extract and count unique email domains.

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// LeetCode name: findUniqueDomains
	emails := []string{"alice@leetcode.com", "bob@leetcode.com", "charlie@gmail.com"}
	fmt.Println(FindAllUniqueEmailDomains(emails))
	// [gmail.com leetcode.com]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: findUniqueDomains
func FindAllUniqueEmailDomains(emails []string) []string {
	domains := make(map[string]bool)
	for _, email := range emails {
		parts := strings.Split(email, "@")
		if len(parts) == 2 {
			domains[parts[1]] = true
		}
	}
	result := make([]string, 0, len(domains))
	for d := range domains {
		result = append(result, d)
	}
	sort.Strings(result)
	return result
}
```

## 3062 — Winner Of The Linked List Game

```go
package main

// LeetCode #3062: Winner of the Linked List Game
// https://leetcode.com/problems/winner-of-the-linked-list-game/
// Difficulty: Easy [Paid]
//
// Note: This is a premium problem. We implement the equivalent logic.
// The game: pairs of consecutive nodes. Even score increases if first > second,
// odd score increases if second > first. Return "Even" or "Odd".

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// LeetCode name: gameResult
	// Create: 2 -> 1 -> 4 -> 5 -> 2 -> 9
	head := &ListNode{2, &ListNode{1, &ListNode{4, &ListNode{5, &ListNode{2, &ListNode{9, nil}}}}}}
	fmt.Println(WinnerOfTheLinkedListGame(head)) // Odd

	head2 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	fmt.Println(WinnerOfTheLinkedListGame(head2)) // Odd
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: gameResult
func WinnerOfTheLinkedListGame(head *ListNode) string {
	evenScore := 0
	oddScore := 0
	curr := head
	for curr != nil && curr.Next != nil {
		first, second := curr.Val, curr.Next.Val
		if first > second {
			evenScore++
		} else if second > first {
			oddScore++
		}
		curr = curr.Next.Next
	}
	if evenScore > oddScore {
		return "Even"
	}
	if oddScore > evenScore {
		return "Odd"
	}
	return "Tie"
}
```

## 3063 — Linked List Frequency

```go
package main

// LeetCode #3063: Linked List Frequency
// https://leetcode.com/problems/linked-list-frequency/
// Difficulty: Easy [Paid]

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// LeetCode name: frequencies
	head := &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{3, nil}}}}}}
	fmt.Println(LinkedListFrequency(head)) // map[1:1 2:2 3:3]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: frequencies
func LinkedListFrequency(head *ListNode) map[int]int {
	freq := make(map[int]int)
	curr := head
	for curr != nil {
		freq[curr.Val]++
		curr = curr.Next
	}
	return freq
}
```

## 3065 — Minimum Operations To Exceed Threshold Value I

```go
package main

// LeetCode #3065: Minimum Operations to Exceed Threshold Value I
// https://leetcode.com/problems/minimum-operations-to-exceed-threshold-value-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: minOperations
	fmt.Println(MinimumOperationsToExceedThresholdValueI([]int{2, 11, 10, 1, 3}, 10)) // 3
	fmt.Println(MinimumOperationsToExceedThresholdValueI([]int{1, 1, 2, 4, 9}, 9))    // 4
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: minOperations
func MinimumOperationsToExceedThresholdValueI(nums []int, k int) int {
	count := 0
	for _, v := range nums {
		if v < k {
			count++
		}
	}
	return count
}
```

## 3069 — Distribute Elements Into Two Arrays I

```go
package main

// LeetCode #3069: Distribute Elements Into Two Arrays I
// https://leetcode.com/problems/distribute-elements-into-two-arrays-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: resultArray
	fmt.Println(DistributeElementsIntoTwoArraysI([]int{2, 1, 3, 4})) // [2, 3, 4, 1]
	fmt.Println(DistributeElementsIntoTwoArraysI([]int{5, 4, 3, 8})) // [5, 3, 4, 8]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: resultArray
func DistributeElementsIntoTwoArraysI(nums []int) []int {
	arr1 := []int{nums[0]}
	arr2 := []int{nums[1]}
	for i := 2; i < len(nums); i++ {
		if arr1[len(arr1)-1] > arr2[len(arr2)-1] {
			arr1 = append(arr1, nums[i])
		} else {
			arr2 = append(arr2, nums[i])
		}
	}
	return append(arr1, arr2...)
}
```

## 3074 — Apple Redistribution Into Boxes

```go
package main

// LeetCode #3074: Apple Redistribution into Boxes
// https://leetcode.com/problems/apple-redistribution-into-boxes/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	// LeetCode name: minimumBoxes
	fmt.Println(AppleRedistributionIntoBoxes([]int{1, 3, 2}, []int{4, 3, 1, 5, 2})) // 2
	fmt.Println(AppleRedistributionIntoBoxes([]int{5, 5, 5}, []int{2, 4, 2, 7}))    // 4
}

// Time: O(n log n) | Space: O(1)
// LeetCode submission name: minimumBoxes
func AppleRedistributionIntoBoxes(apples []int, capacity []int) int {
	totalApples := 0
	for _, a := range apples {
		totalApples += a
	}
	sort.Sort(sort.Reverse(sort.IntSlice(capacity)))
	boxes := 0
	for _, c := range capacity {
		boxes++
		totalApples -= c
		if totalApples <= 0 {
			return boxes
		}
	}
	return boxes
}
```

## 3079 — Find The Sum Of Encrypted Integers

```go
package main

// LeetCode #3079: Find the Sum of Encrypted Integers
// https://leetcode.com/problems/find-the-sum-of-encrypted-integers/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: sumOfEncryptedInt
	fmt.Println(FindTheSumOfEncryptedIntegers([]int{10, 21, 31})) // 66
	fmt.Println(FindTheSumOfEncryptedIntegers([]int{1, 2, 3}))    // 6
}

// Time: O(n * d) where d is number of digits | Space: O(1)
// LeetCode submission name: sumOfEncryptedInt
func FindTheSumOfEncryptedIntegers(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += encrypt(v)
	}
	return sum
}

func encrypt(n int) int {
	maxDigit := 0
	digits := 0
	for n > 0 {
		d := n % 10
		if d > maxDigit {
			maxDigit = d
		}
		digits++
		n /= 10
	}
	result := 0
	for i := 0; i < digits; i++ {
		result = result*10 + maxDigit
	}
	return result
}
```

## 3083 — Existence Of A Substring In A String And Its Reverse

```go
package main

// LeetCode #3083: Existence of a Substring in a String and Its Reverse
// https://leetcode.com/problems/existence-of-a-substring-in-a-string-and-its-reverse/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: isSubstringPresent
	fmt.Println(ExistenceOfASubstringInAStringAndItsReverse("leetcode")) // true
	fmt.Println(ExistenceOfASubstringInAStringAndItsReverse("abcba"))   // true
	fmt.Println(ExistenceOfASubstringInAStringAndItsReverse("abcd"))    // false
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: isSubstringPresent
func ExistenceOfASubstringInAStringAndItsReverse(s string) bool {
	// Build set of all substrings of length 2
	substrings := make(map[string]bool)
	for i := 0; i < len(s)-1; i++ {
		substrings[s[i:i+2]] = true
	}

	// Check reverse for any of those substrings
	for i := len(s) - 1; i > 0; i-- {
		if substrings[string(s[i])+string(s[i-1])] {
			return true
		}
	}
	return false
}
```

## 3090 — Maximum Length Substring With Two Occurrences

```go
package main

// LeetCode #3090: Maximum Length Substring With Two Occurrences
// https://leetcode.com/problems/maximum-length-substring-with-two-occurrences/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: maximumLengthSubstring
	fmt.Println(MaximumLengthSubstringWithTwoOccurrences("bcbbbcba")) // 4
	fmt.Println(MaximumLengthSubstringWithTwoOccurrences("aaaa"))      // 2
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: maximumLengthSubstring
func MaximumLengthSubstringWithTwoOccurrences(s string) int {
	left := 0
	freq := make(map[byte]int)
	maxLen := 0

	for right := 0; right < len(s); right++ {
		freq[s[right]]++
		for freq[s[right]] > 2 {
			freq[s[left]]--
			left++
		}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}
```

## 3095 — Shortest Subarray With Or At Least K I

```go
package main

// LeetCode #3095: Shortest Subarray With OR at Least K I
// https://leetcode.com/problems/shortest-subarray-with-or-at-least-k-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: minimumSubarrayLength
	fmt.Println(ShortestSubarrayWithOrAtLeastKI([]int{1, 2, 3}, 2)) // 1
	fmt.Println(ShortestSubarrayWithOrAtLeastKI([]int{2, 1, 8}, 10)) // 3
	fmt.Println(ShortestSubarrayWithOrAtLeastKI([]int{1, 2}, 10))    // -1
}

// Time: O(n^2) | Space: O(1)
// LeetCode submission name: minimumSubarrayLength
func ShortestSubarrayWithOrAtLeastKI(nums []int, k int) int {
	n := len(nums)
	minLen := n + 1
	for i := 0; i < n; i++ {
		orVal := 0
		for j := i; j < n; j++ {
			orVal |= nums[j]
			if orVal >= k {
				if j-i+1 < minLen {
					minLen = j - i + 1
				}
				break
			}
		}
	}
	if minLen > n {
		return -1
	}
	return minLen
}
```

## 3099 — Harshad Number

```go
package main

// LeetCode #3099: Harshad Number
// https://leetcode.com/problems/harshad-number/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: sumOfTheDigitsOfHarshadNumber
	fmt.Println(HarshadNumber(18)) // 9
	fmt.Println(HarshadNumber(23)) // -1
}

// Time: O(log n) | Space: O(1)
// LeetCode submission name: sumOfTheDigitsOfHarshadNumber
func HarshadNumber(x int) int {
	sum := 0
	n := x
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	if x%sum == 0 {
		return sum
	}
	return -1
}
```

## 3105 — Longest Strictly Increasing Or Strictly Decreasing Subarray

```go
package main

// LeetCode #3105: Longest Strictly Increasing or Strictly Decreasing Subarray
// https://leetcode.com/problems/longest-strictly-increasing-or-strictly-decreasing-subarray/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: longestMonotonicSubarray
	fmt.Println(LongestStrictlyIncreasingOrStrictlyDecreasingSubarray([]int{1, 4, 3, 3, 2})) // 2
	fmt.Println(LongestStrictlyIncreasingOrStrictlyDecreasingSubarray([]int{3, 3, 3, 3}))  // 1
	fmt.Println(LongestStrictlyIncreasingOrStrictlyDecreasingSubarray([]int{3, 2, 1}))     // 3
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: longestMonotonicSubarray
func LongestStrictlyIncreasingOrStrictlyDecreasingSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	inc := 1
	dec := 1
	maxLen := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			inc++
			dec = 1
		} else if nums[i] < nums[i-1] {
			dec++
			inc = 1
		} else {
			inc = 1
			dec = 1
		}
		if inc > maxLen {
			maxLen = inc
		}
		if dec > maxLen {
			maxLen = dec
		}
	}
	return maxLen
}
```

## 3110 — Score Of A String

```go
package main

// LeetCode #3110: Score of a String
// https://leetcode.com/problems/score-of-a-string/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: scoreOfString
	fmt.Println(ScoreOfAString("hello")) // 13
	fmt.Println(ScoreOfAString("zaz"))   // 50
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: scoreOfString
func ScoreOfAString(s string) int {
	score := 0
	for i := 1; i < len(s); i++ {
		diff := int(s[i]) - int(s[i-1])
		if diff < 0 {
			diff = -diff
		}
		score += diff
	}
	return score
}
```

## 3114 — Latest Time You Can Obtain After Replacing Characters

```go
package main

// LeetCode #3114: Latest Time You Can Obtain After Replacing Characters
// https://leetcode.com/problems/latest-time-you-can-obtain-after-replacing-characters/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: findLatestTime
	fmt.Println(LatestTimeYouCanObtainAfterReplacingCharacters("1?:?4")) // 11:54
	fmt.Println(LatestTimeYouCanObtainAfterReplacingCharacters("0?:5?")) // 09:59
}

// Time: O(1) | Space: O(1)
// LeetCode submission name: findLatestTime
func LatestTimeYouCanObtainAfterReplacingCharacters(s string) string {
	time := []byte(s)

	// Replace hours
	if time[0] == '?' {
		if time[1] == '?' || time[1] <= '1' {
			time[0] = '1'
		} else {
			time[0] = '0'
		}
	}
	if time[1] == '?' {
		if time[0] == '1' {
			time[1] = '1'
		} else {
			time[1] = '9'
		}
	}

	// Replace minutes
	if time[3] == '?' {
		time[3] = '5'
	}
	if time[4] == '?' {
		time[4] = '9'
	}

	return string(time)
}
```

## 3120 — Count The Number Of Special Characters I

```go
package main

// LeetCode #3120: Count the Number of Special Characters I
// https://leetcode.com/problems/count-the-number-of-special-characters-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: numberOfSpecialChars
	fmt.Println(CountTheNumberOfSpecialCharactersI("aaAbcBC")) // 3
	fmt.Println(CountTheNumberOfSpecialCharactersI("abcd"))    // 0
	fmt.Println(CountTheNumberOfSpecialCharactersI("abAB"))   // 2
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: numberOfSpecialChars
func CountTheNumberOfSpecialCharactersI(word string) int {
	lower := make(map[byte]bool)
	upper := make(map[byte]bool)
	for i := 0; i < len(word); i++ {
		c := word[i]
		if c >= 'a' && c <= 'z' {
			lower[c] = true
		} else if c >= 'A' && c <= 'Z' {
			upper[c] = true
		}
	}
	count := 0
	for c := byte('a'); c <= 'z'; c++ {
		if lower[c] && upper[c-'a'+'A'] {
			count++
		}
	}
	return count
}
```

## 3127 — Make A Square With The Same Color

```go
package main

// LeetCode #3127: Make a Square with the Same Color
// https://leetcode.com/problems/make-a-square-with-the-same-color/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: canMakeSquare
	grid1 := [][]byte{{'B', 'W', 'B'}, {'B', 'W', 'W'}, {'B', 'W', 'B'}}
	fmt.Println(MakeASquareWithTheSameColor(grid1)) // true

	grid2 := [][]byte{{'B', 'W', 'B'}, {'W', 'B', 'W'}, {'B', 'W', 'B'}}
	fmt.Println(MakeASquareWithTheSameColor(grid2)) // false

	grid3 := [][]byte{{'B', 'W', 'B'}, {'B', 'W', 'W'}, {'B', 'W', 'W'}}
	fmt.Println(MakeASquareWithTheSameColor(grid3)) // true
}

// Time: O(1) | Space: O(1)
// LeetCode submission name: canMakeSquare
func MakeASquareWithTheSameColor(grid [][]byte) bool {
	// Check all 2x2 subgrids
	dirs := [][2]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			bCount := 0
			for _, d := range dirs {
				if grid[i+d[0]][j+d[1]] == 'B' {
					bCount++
				}
			}
			// If at least 3 cells have the same color, changing 1 makes all 4
			if bCount >= 3 || bCount <= 1 {
				return true
			}
		}
	}
	return false
}
```

## 3131 — Find The Integer Added To Array I

```go
package main

// LeetCode #3131: Find the Integer Added to Array I
// https://leetcode.com/problems/find-the-integer-added-to-array-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: addedInteger
	fmt.Println(FindTheIntegerAddedToArrayI([]int{2, 6, 4}, []int{9, 7, 5})) // 3
	fmt.Println(FindTheIntegerAddedToArrayI([]int{10}, []int{5}))             // -5
}

// Time: O(n log n) | Space: O(1)
// LeetCode submission name: addedInteger
func FindTheIntegerAddedToArrayI(nums1 []int, nums2 []int) int {
	min1, min2 := nums1[0], nums2[0]
	for _, v := range nums1 {
		if v < min1 {
			min1 = v
		}
	}
	for _, v := range nums2 {
		if v < min2 {
			min2 = v
		}
	}
	return min2 - min1
}
```

## 3136 — Valid Word

```go
package main

// LeetCode #3136: Valid Word
// https://leetcode.com/problems/valid-word/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: isValid
	fmt.Println(ValidWord("Hello123")) // true
	fmt.Println(ValidWord("hi"))       // false
	fmt.Println(ValidWord("AAab"))     // false (no vowel)
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: isValid
func ValidWord(word string) bool {
	if len(word) < 3 {
		return false
	}
	hasVowel := false
	hasConsonant := false
	for i := 0; i < len(word); i++ {
		c := word[i]
		if c >= '0' && c <= '9' {
			continue
		}
		if c >= 'A' && c <= 'Z' {
			c = c - 'A' + 'a'
		}
		if c < 'a' || c > 'z' {
			return false
		}
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			hasVowel = true
		} else {
			hasConsonant = true
		}
	}
	return hasVowel && hasConsonant
}
```

## 3142 — Check If Grid Satisfies Conditions

```go
package main

// LeetCode #3142: Check if Grid Satisfies Conditions
// https://leetcode.com/problems/check-if-grid-satisfies-conditions/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: satisfiesConditions
	fmt.Println(CheckIfGridSatisfiesConditions([][]int{{1, 0, 2}, {1, 0, 2}})) // true
	fmt.Println(CheckIfGridSatisfiesConditions([][]int{{1, 1, 1}, {0, 0, 0}})) // false
}

// Time: O(n * m) | Space: O(1)
// LeetCode submission name: satisfiesConditions
func CheckIfGridSatisfiesConditions(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// Below must be equal
			if i+1 < m && grid[i][j] != grid[i+1][j] {
				return false
			}
			// Right must be different
			if j+1 < n && grid[i][j] == grid[i][j+1] {
				return false
			}
		}
	}
	return true
}
```

## 3146 — Permutation Difference Between Two Strings

```go
package main

// LeetCode #3146: Permutation Difference between Two Strings
// https://leetcode.com/problems/permutation-difference-between-two-strings/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: findPermutationDifference
	fmt.Println(PermutationDifferenceBetweenTwoStrings("abc", "bac")) // 2
	fmt.Println(PermutationDifferenceBetweenTwoStrings("abcde", "edcba")) // 12
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: findPermutationDifference
func PermutationDifferenceBetweenTwoStrings(s string, t string) int {
	pos := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		pos[t[i]] = i
	}
	diff := 0
	for i := 0; i < len(s); i++ {
		d := i - pos[s[i]]
		if d < 0 {
			d = -d
		}
		diff += d
	}
	return diff
}
```

## 3150 — Invalid Tweets Ii

```go
package main

// LeetCode #3150: Invalid Tweets II
// https://leetcode.com/problems/invalid-tweets-ii/
// Difficulty: Easy [Paid - Database]
//
// Note: This is a SQL problem on LeetCode. In Go, we implement the
// equivalent logic: find tweet IDs where content length > 140 or
// content contains invalid characters.

import (
	"fmt"
	"strings"
)

func main() {
	// LeetCode name: invalidTweets
	tweets := map[int]string{
		1: "Hello world!",
		2: strings.Repeat("a", 150),
		3: "Valid tweet content",
	}
	fmt.Println(InvalidTweetsIi(tweets))
	// [2]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: invalidTweets
func InvalidTweetsIi(tweets map[int]string) []int {
	result := []int{}
	for id, content := range tweets {
		if len(content) > 140 {
			result = append(result, id)
		}
	}
	return result
}
```

## 3151 — Special Array I

```go
package main

// LeetCode #3151: Special Array I
// https://leetcode.com/problems/special-array-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: isArraySpecial
	fmt.Println(SpecialArrayI([]int{1}))       // true
	fmt.Println(SpecialArrayI([]int{2, 1, 4})) // true
	fmt.Println(SpecialArrayI([]int{4, 3, 1, 6})) // false
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: isArraySpecial
func SpecialArrayI(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if (nums[i]%2 == 0) == (nums[i-1]%2 == 0) {
			return false
		}
	}
	return true
}
```

## 3158 — Find The Xor Of Numbers Which Appear Twice

```go
package main

// LeetCode #3158: Find the XOR of Numbers Which Appear Twice
// https://leetcode.com/problems/find-the-xor-of-numbers-which-appear-twice/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: duplicateNumbersXOR
	fmt.Println(FindTheXorOfNumbersWhichAppearTwice([]int{1, 2, 2, 1})) // 3
	fmt.Println(FindTheXorOfNumbersWhichAppearTwice([]int{1, 2, 3}))    // 0
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: duplicateNumbersXOR
func FindTheXorOfNumbersWhichAppearTwice(nums []int) int {
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}
	xor := 0
	for v, f := range freq {
		if f == 2 {
			xor ^= v
		}
	}
	return xor
}
```

## 3162 — Find The Number Of Good Pairs I

```go
package main

// LeetCode #3162: Find the Number of Good Pairs I
// https://leetcode.com/problems/find-the-number-of-good-pairs-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: numberOfPairs
	fmt.Println(FindTheNumberOfGoodPairsI([]int{1, 2, 4, 12}, []int{2, 4}, 3)) // 2
	fmt.Println(FindTheNumberOfGoodPairsI([]int{1, 2, 3}, []int{1, 2, 3}, 1))  // 5
}

// Time: O(n * m) | Space: O(1)
// LeetCode submission name: numberOfPairs
func FindTheNumberOfGoodPairsI(nums1 []int, nums2 []int, k int) int {
	count := 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i]%(nums2[j]*k) == 0 {
				count++
			}
		}
	}
	return count
}
```

## 3168 — Minimum Number Of Chairs In A Waiting Room

```go
package main

// LeetCode #3168: Minimum Number of Chairs in a Waiting Room
// https://leetcode.com/problems/minimum-number-of-chairs-in-a-waiting-room/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: minimumChairs
	fmt.Println(MinimumNumberOfChairsInAWaitingRoom("EEEE"))   // 4
	fmt.Println(MinimumNumberOfChairsInAWaitingRoom("ELELEL")) // 1
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: minimumChairs
func MinimumNumberOfChairsInAWaitingRoom(s string) int {
	current := 0
	maxChairs := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'E' {
			current++
			if current > maxChairs {
				maxChairs = current
			}
		} else {
			current--
		}
	}
	return maxChairs
}
```

## 3172 — Second Day Verification

```go
package main

// LeetCode #3172: Second Day Verification
// https://leetcode.com/problems/second-day-verification/
// Difficulty: Easy [Paid - Database]
//
// Note: This is a SQL problem on LeetCode. In Go, we implement the
// equivalent logic: determine which user IDs had their verification
// completed on the second day.

import (
	"fmt"
	"sort"
)

func main() {
	// LeetCode name: secondDayVerify
	// Input: [user_id, day, action] where action is 'submitted' or 'verified'
	actions := [][]string{
		{"1", "1", "submitted"},
		{"1", "2", "verified"},
		{"2", "1", "submitted"},
		{"3", "1", "submitted"},
		{"3", "3", "verified"},
	}
	fmt.Println(SecondDayVerification(actions))
	// [1]
}

// Time: O(n log n) | Space: O(n)
// LeetCode submission name: secondDayVerify
func SecondDayVerification(actions [][]string) []int {
	submitted := make(map[int]int)
	verified := make(map[int]int)
	for _, row := range actions {
		id := parseInt(row[0])
		day := parseInt(row[1])
		action := row[2]
		if action == "submitted" {
			submitted[id] = day
		} else if action == "verified" {
			verified[id] = day
		}
	}
	result := []int{}
	for id, submitDay := range submitted {
		if verifyDay, ok := verified[id]; ok && verifyDay == submitDay+1 {
			result = append(result, id)
		}
	}
	sort.Ints(result)
	return result
}

func parseInt(s string) int {
	n := 0
	for i := 0; i < len(s); i++ {
		n = n*10 + int(s[i]-'0')
	}
	return n
}
```

## 3173 — Bitwise Or Of Adjacent Elements

```go
package main

// LeetCode #3173: Bitwise OR of Adjacent Elements
// https://leetcode.com/problems/bitwise-or-of-adjacent-elements/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	// LeetCode name: orArray
	fmt.Println(BitwiseOrOfAdjacentElements([]int{1, 2, 3, 4})) // [3, 3, 7]
	fmt.Println(BitwiseOrOfAdjacentElements([]int{5, 1, 6}))     // [5, 7]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: orArray
func BitwiseOrOfAdjacentElements(nums []int) []int {
	result := make([]int, len(nums)-1)
	for i := 0; i < len(nums)-1; i++ {
		result[i] = nums[i] | nums[i+1]
	}
	return result
}
```

