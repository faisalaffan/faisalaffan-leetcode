# Easy (Mudah) — Problem 2437–2797

## 2437 — Number Of Valid Clock Times

```go
package main

// LeetCode #2437: Number of Valid Clock Times
// https://leetcode.com/problems/number-of-valid-clock-times/
// Difficulty: Easy
// Time O(1) | Space O(1)

import "fmt"

func main() {
	fmt.Println(NumberOfValidClockTimes("?5:00")) // 2
	fmt.Println(NumberOfValidClockTimes("0?:0?")) // 100
	fmt.Println(NumberOfValidClockTimes("??:??")) // 1440
}

func NumberOfValidClockTimes(time string) int {
	ways := 1

	// Hours tens digit
	if time[0] == '?' {
		if time[1] == '?' {
			ways *= 24
		} else if time[1] <= '3' {
			ways *= 3 // 0,1,2
		} else {
			ways *= 2 // 0,1
		}
	} else if time[1] == '?' {
		if time[0] == '0' || time[0] == '1' {
			ways *= 10
		} else { // time[0] == '2'
			ways *= 4
		}
	}

	// Minutes tens digit
	if time[3] == '?' {
		ways *= 6
	}
	// Minutes ones digit
	if time[4] == '?' {
		ways *= 10
	}

	return ways
}
```

## 2441 — Largest Positive Integer That Exists With Its Negative

```go
package main

// LeetCode #2441: Largest Positive Integer That Exists With Its Negative
// https://leetcode.com/problems/largest-positive-integer-that-exists-with-its-negative/
// Difficulty: Easy
// Time O(n) | Space O(n)

import "fmt"

func main() {
	fmt.Println(LargestPositiveIntegerThatExistsWithItsNegative([]int{-1, 2, -3, 3}))  // 3
	fmt.Println(LargestPositiveIntegerThatExistsWithItsNegative([]int{-1, 10, 6, 7, -7, 1})) // 7
	fmt.Println(LargestPositiveIntegerThatExistsWithItsNegative([]int{-10, 8, 6, 7, -2, -3})) // -1
}

func LargestPositiveIntegerThatExistsWithItsNegative(nums []int) int {
	seen := map[int]bool{}
	for _, n := range nums {
		seen[n] = true
	}
	best := -1
	for _, n := range nums {
		if n > 0 && seen[-n] && n > best {
			best = n
		}
	}
	return best
}
```

## 2446 — Determine If Two Events Have Conflict

```go
package main

// LeetCode #2446: Determine if Two Events Have Conflict
// https://leetcode.com/problems/determine-if-two-events-have-conflict/
// Difficulty: Easy
// Time O(1) | Space O(1)

import "fmt"

func main() {
	fmt.Println(DetermineIfTwoEventsHaveConflict([]string{"01:15", "02:00"}, []string{"02:00", "03:00"})) // true
	fmt.Println(DetermineIfTwoEventsHaveConflict([]string{"01:00", "02:00"}, []string{"01:20", "03:00"})) // true
	fmt.Println(DetermineIfTwoEventsHaveConflict([]string{"10:00", "11:00"}, []string{"14:00", "15:00"})) // false
}

func timeToMin(t string) int {
	return int(t[0]-'0')*600 + int(t[1]-'0')*60 + int(t[3]-'0')*10 + int(t[4]-'0')
}

func DetermineIfTwoEventsHaveConflict(event1 []string, event2 []string) bool {
	s1, e1 := timeToMin(event1[0]), timeToMin(event1[1])
	s2, e2 := timeToMin(event2[0]), timeToMin(event2[1])
	return !(e1 < s2 || e2 < s1)
}
```

## 2451 — Odd String Difference

```go
package main

// LeetCode #2451: Odd String Difference
// https://leetcode.com/problems/odd-string-difference/
// Difficulty: Easy
// Time O(n * m) | Space O(n)

import "fmt"

func main() {
	fmt.Println(OddStringDifference([]string{"adc", "wzy", "abc"})) // "abc"
	fmt.Println(OddStringDifference([]string{"aaa", "bob", "ccc", "ddd"})) // "bob"
}

func differenceArray(s string) string {
	diff := make([]byte, len(s)-1)
	for i := 1; i < len(s); i++ {
		diff[i-1] = s[i] - s[i-1]
	}
	return string(diff)
}

func OddStringDifference(words []string) string {
	diff0 := differenceArray(words[0])
	diff1 := differenceArray(words[1])

	if string(diff0) != string(diff1) {
		diff2 := differenceArray(words[2])
		if string(diff0) == string(diff2) {
			return words[1]
		}
		return words[0]
	}

	for i := 2; i < len(words); i++ {
		if string(differenceArray(words[i])) != string(diff0) {
			return words[i]
		}
	}
	return ""
}
```

## 2455 — Average Value Of Even Numbers That Are Divisible By Three

```go
package main

// LeetCode #2455: Average Value of Even Numbers That Are Divisible by Three
// https://leetcode.com/problems/average-value-of-even-numbers-that-are-divisible-by-three/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(AverageValueOfEvenNumbersThatAreDivisibleByThree([]int{1, 3, 6, 10, 12, 15})) // 9
	fmt.Println(AverageValueOfEvenNumbersThatAreDivisibleByThree([]int{1, 2, 4, 7, 10}))       // 0
}

func AverageValueOfEvenNumbersThatAreDivisibleByThree(nums []int) int {
	sum, count := 0, 0
	for _, n := range nums {
		if n%6 == 0 {
			sum += n
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return sum / count
}
```

## 2460 — Apply Operations To An Array

```go
package main

// LeetCode #2460: Apply Operations to an Array
// https://leetcode.com/problems/apply-operations-to-an-array/
// Difficulty: Easy
// Time O(n) | Space O(n)

import "fmt"

func main() {
	fmt.Println(ApplyOperationsToAnArray([]int{1, 2, 2, 1, 1, 0})) // [1,4,2,0,0,0]
	fmt.Println(ApplyOperationsToAnArray([]int{0, 1}))              // [1,0]
}

func ApplyOperationsToAnArray(nums []int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}
	}

	res := make([]int, n)
	idx := 0
	for _, v := range nums {
		if v != 0 {
			res[idx] = v
			idx++
		}
	}
	return res
}
```

## 2465 — Number Of Distinct Averages

```go
package main

// LeetCode #2465: Number of Distinct Averages
// https://leetcode.com/problems/number-of-distinct-averages/
// Difficulty: Easy
// Time O(n log n) | Space O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(NumberOfDistinctAverages([]int{4, 1, 4, 0, 3, 5})) // 2
	fmt.Println(NumberOfDistinctAverages([]int{1, 100}))            // 1
}

func NumberOfDistinctAverages(nums []int) int {
	sort.Ints(nums)
	seen := map[int]bool{}
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		seen[nums[i]+nums[j]] = true
	}
	return len(seen)
}
```

## 2469 — Convert The Temperature

```go
package main

// LeetCode #2469: Convert the Temperature
// https://leetcode.com/problems/convert-the-temperature/
// Difficulty: Easy
// Time O(1) | Space O(1)

import "fmt"

func main() {
	fmt.Println(ConvertTheTemperature(36.50)) // [309.65, 97.7]
	fmt.Println(ConvertTheTemperature(122.11)) // [395.26, 251.798]
}

func ConvertTheTemperature(celsius float64) []float64 {
	return []float64{celsius + 273.15, celsius*1.8 + 32.0}
}
```

## 2475 — Number Of Unequal Triplets In Array

```go
package main

// LeetCode #2475: Number of Unequal Triplets in Array
// https://leetcode.com/problems/number-of-unequal-triplets-in-array/
// Difficulty: Easy
// Time O(n^3) | Space O(1)

import "fmt"

func main() {
	fmt.Println(NumberOfUnequalTripletsInArray([]int{4, 4, 2, 4, 3})) // 3
	fmt.Println(NumberOfUnequalTripletsInArray([]int{1, 1, 1, 1, 1}))  // 0
}

func NumberOfUnequalTripletsInArray(nums []int) int {
	n := len(nums)
	count := 0
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i] != nums[j] && nums[i] != nums[k] && nums[j] != nums[k] {
					count++
				}
			}
		}
	}
	return count
}
```

## 2480 — Form A Chemical Bond

```go
package main

// LeetCode #2480: Form a Chemical Bond
// https://leetcode.com/problems/form-a-chemical-bond/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(FormAChemicalBond())
}

func FormAChemicalBond() any {
	// TODO: implement
	return nil
}
```

## 2481 — Minimum Cuts To Divide A Circle

```go
package main

// LeetCode #2481: Minimum Cuts to Divide a Circle
// https://leetcode.com/problems/minimum-cuts-to-divide-a-circle/
// Difficulty: Easy
// Time O(1) | Space O(1)

import "fmt"

func main() {
	fmt.Println(MinimumCutsToDivideACircle(4)) // 2
	fmt.Println(MinimumCutsToDivideACircle(3)) // 3
	fmt.Println(MinimumCutsToDivideACircle(1)) // 0
}

func MinimumCutsToDivideACircle(n int) int {
	if n == 1 {
		return 0
	}
	if n%2 == 0 {
		return n / 2
	}
	return n
}
```

## 2485 — Find The Pivot Integer

```go
package main

// LeetCode #2485: Find the Pivot Integer
// https://leetcode.com/problems/find-the-pivot-integer/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(FindThePivotInteger(8)) // 6
	fmt.Println(FindThePivotInteger(1)) // 1
	fmt.Println(FindThePivotInteger(4)) // -1
}

func FindThePivotInteger(n int) int {
	total := n * (n + 1) / 2
	sum := 0
	for x := 1; x <= n; x++ {
		sum += x
		if sum == total-sum+x {
			return x
		}
	}
	return -1
}
```

## 2490 — Circular Sentence

```go
package main

// LeetCode #2490: Circular Sentence
// https://leetcode.com/problems/circular-sentence/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CircularSentence("leetcode exercises sound delightful")) // true
	fmt.Println(CircularSentence("eetcode"))                            // true
	fmt.Println(CircularSentence("Leetcode is cool"))                   // false
}

func CircularSentence(sentence string) bool {
	if sentence[0] != sentence[len(sentence)-1] {
		return false
	}
	for i := 0; i < len(sentence); i++ {
		if sentence[i] == ' ' && sentence[i-1] != sentence[i+1] {
			return false
		}
	}
	return true
}
```

## 2496 — Maximum Value Of A String In An Array

```go
package main

// LeetCode #2496: Maximum Value of a String in an Array
// https://leetcode.com/problems/maximum-value-of-a-string-in-an-array/
// Difficulty: Easy
// Time O(n * m) | Space O(1)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(MaximumValueOfAStringInAnArray([]string{"alic3", "bob", "3", "4", "00000"})) // 5
	fmt.Println(MaximumValueOfAStringInAnArray([]string{"1", "01", "001", "0001"}))           // 1
}

func MaximumValueOfAStringInAnArray(strs []string) int {
	maxVal := 0
	for _, s := range strs {
		val := 0
		if isNumeric(s) {
			val, _ = strconv.Atoi(s)
		} else {
			val = len(s)
		}
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

func isNumeric(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return len(s) > 0
}
```

## 2500 — Delete Greatest Value In Each Row

```go
package main

// LeetCode #2500: Delete Greatest Value in Each Row
// https://leetcode.com/problems/delelete-greatest-value-in-each-row/
// Difficulty: Easy
// Time O(n * m log m) | Space O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(DeleteGreatestValueInEachRow([][]int{{1, 2, 4}, {3, 3, 1}})) // 8
	fmt.Println(DeleteGreatestValueInEachRow([][]int{{10}}))                 // 10
}

func DeleteGreatestValueInEachRow(grid [][]int) int {
	for i := range grid {
		sort.Ints(grid[i])
	}

	m := len(grid[0])
	sum := 0
	for col := m - 1; col >= 0; col-- {
		maxVal := 0
		for row := 0; row < len(grid); row++ {
			if grid[row][col] > maxVal {
				maxVal = grid[row][col]
			}
		}
		sum += maxVal
	}
	return sum
}
```

## 2504 — Concatenate The Name And The Profession

```go
package main

// LeetCode #2504: Concatenate the Name and the Profession
// https://leetcode.com/problems/concatenate-the-name-and-the-profession/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(ConcatenateTheNameAndTheProfession())
}

func ConcatenateTheNameAndTheProfession() any {
	// TODO: implement
	return nil
}
```

## 2506 — Count Pairs Of Similar Strings

```go
package main

// LeetCode #2506: Count Pairs Of Similar Strings
// https://leetcode.com/problems/count-pairs-of-similar-strings/
// Difficulty: Easy
// Time O(n * m) | Space O(n)

import "fmt"

func main() {
	fmt.Println(CountPairsOfSimilarStrings([]string{"aba", "aabb", "abcd", "bac", "aabc"})) // 2
	fmt.Println(CountPairsOfSimilarStrings([]string{"aabb", "ab", "ba"}))                    // 3
}

func charMask(s string) int {
	mask := 0
	for i := 0; i < len(s); i++ {
		mask |= 1 << (s[i] - 'a')
	}
	return mask
}

func CountPairsOfSimilarStrings(words []string) int {
	count := 0
	for i := 0; i < len(words); i++ {
		maskI := charMask(words[i])
		for j := i + 1; j < len(words); j++ {
			if maskI == charMask(words[j]) {
				count++
			}
		}
	}
	return count
}
```

## 2511 — Maximum Enemy Forts That Can Be Captured

```go
package main

// LeetCode #2511: Maximum Enemy Forts That Can Be Captured
// https://leetcode.com/problems/maximum-enemy-forts-that-can-be-captured/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(MaximumEnemyFortsThatCanBeCaptured([]int{1, 0, 0, -1, 0, 0, 0, 0, 1})) // 4
	fmt.Println(MaximumEnemyFortsThatCanBeCaptured([]int{0, 0, 1, -1}))                 // 0
}

func MaximumEnemyFortsThatCanBeCaptured(forts []int) int {
	maxCap := 0
	for i := 0; i < len(forts); i++ {
		if forts[i] == 1 {
			// Move right
			for j := i + 1; j < len(forts); j++ {
				if forts[j] == -1 {
					if j-i-1 > maxCap {
						maxCap = j - i - 1
					}
					break
				} else if forts[j] == 1 {
					break
				}
			}
		} else if forts[i] == -1 {
			// Move right (capturing from -1 to 1)
			for j := i + 1; j < len(forts); j++ {
				if forts[j] == 1 {
					if j-i-1 > maxCap {
						maxCap = j - i - 1
					}
					break
				} else if forts[j] == -1 {
					break
				}
			}
		}
	}
	return maxCap
}
```

## 2515 — Shortest Distance To Target String In A Circular Array

```go
package main

// LeetCode #2515: Shortest Distance to Target String in a Circular Array
// https://leetcode.com/problems/shortest-distance-to-target-string-in-a-circular-array/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(ShortestDistanceToTargetStringInACircularArray([]string{"hello", "i", "am", "leetcode", "hello"}, "hello", 1)) // 1
	fmt.Println(ShortestDistanceToTargetStringInACircularArray([]string{"a", "b", "leetcode"}, "leetcode", 0))                // 1
}

func ShortestDistanceToTargetStringInACircularArray(words []string, target string, startIndex int) int {
	n := len(words)
	minDist := n

	for i, w := range words {
		if w == target {
			dist := abs(i - startIndex)
			if dist > n-dist {
				dist = n - dist
			}
			if dist < minDist {
				minDist = dist
			}
		}
	}
	return minDist
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

## 2520 — Count The Digits That Divide A Number

```go
package main

// LeetCode #2520: Count the Digits That Divide a Number
// https://leetcode.com/problems/count-the-digits-that-divide-a-number/
// Difficulty: Easy
// Time O(log n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CountTheDigitsThatDivideANumber(7))    // 1
	fmt.Println(CountTheDigitsThatDivideANumber(121))  // 2
	fmt.Println(CountTheDigitsThatDivideANumber(1248)) // 4
}

func CountTheDigitsThatDivideANumber(num int) int {
	count := 0
	n := num
	for n > 0 {
		digit := n % 10
		if digit != 0 && num%digit == 0 {
			count++
		}
		n /= 10
	}
	return count
}
```

## 2525 — Categorize Box According To Criteria

```go
package main

// LeetCode #2525: Categorize Box According to Criteria
// https://leetcode.com/problems/categorize-box-according-to-criteria/
// Difficulty: Easy
// Time O(1) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CategorizeBoxAccordingToCriteria(1000, 35, 700, 300)) // "Heavy"
	fmt.Println(CategorizeBoxAccordingToCriteria(200, 50, 800, 50))   // "Neither"
}

func CategorizeBoxAccordingToCriteria(length int, width int, height int, mass int) string {
	volume := length * width * height
	isBulky := length >= 10000 || width >= 10000 || height >= 10000 || volume >= 1000000000
	isHeavy := mass >= 100

	if isBulky && isHeavy {
		return "Both"
	}
	if isBulky {
		return "Bulky"
	}
	if isHeavy {
		return "Heavy"
	}
	return "Neither"
}
```

## 2529 — Maximum Count Of Positive Integer And Negative Integer

```go
package main

// LeetCode #2529: Maximum Count of Positive Integer and Negative Integer
// https://leetcode.com/problems/maximum-count-of-positive-integer-and-negative-integer/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(MaximumCountOfPositiveIntegerAndNegativeInteger([]int{-2, -1, -1, 1, 2, 3})) // 3
	fmt.Println(MaximumCountOfPositiveIntegerAndNegativeInteger([]int{-3, -2, -1, 0, 0, 1, 2})) // 3
}

func MaximumCountOfPositiveIntegerAndNegativeInteger(nums []int) int {
	pos, neg := 0, 0
	for _, n := range nums {
		if n > 0 {
			pos++
		} else if n < 0 {
			neg++
		}
	}
	if pos > neg {
		return pos
	}
	return neg
}
```

## 2535 — Difference Between Element Sum And Digit Sum Of An Array

```go
package main

// LeetCode #2535: Difference Between Element Sum and Digit Sum of an Array
// https://leetcode.com/problems/difference-between-element-sum-and-digit-sum-of-an-array/
// Difficulty: Easy
// Time O(n log m) | Space O(1)

import "fmt"

func main() {
	fmt.Println(DifferenceBetweenElementSumAndDigitSumOfAnArray([]int{1, 15, 6, 3})) // 9
	fmt.Println(DifferenceBetweenElementSumAndDigitSumOfAnArray([]int{1, 2, 3, 4}))  // 0
}

func DifferenceBetweenElementSumAndDigitSumOfAnArray(nums []int) int {
	elementSum := 0
	digitSum := 0
	for _, n := range nums {
		elementSum += n
		for n > 0 {
			digitSum += n % 10
			n /= 10
		}
	}
	result := elementSum - digitSum
	if result < 0 {
		return -result
	}
	return result
}
```

## 2540 — Minimum Common Value

```go
package main

// LeetCode #2540: Minimum Common Value
// https://leetcode.com/problems/minimum-common-value/
// Difficulty: Easy
// Time O(n + m) | Space O(1)

import "fmt"

func main() {
	fmt.Println(MinimumCommonValue([]int{1, 2, 3}, []int{2, 4}))       // 2
	fmt.Println(MinimumCommonValue([]int{1, 2, 3, 6}, []int{2, 3, 4, 5})) // 2
}

func MinimumCommonValue(nums1 []int, nums2 []int) int {
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			return nums1[i]
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return -1
}
```

## 2544 — Alternating Digit Sum

```go
package main

// LeetCode #2544: Alternating Digit Sum
// https://leetcode.com/problems/alternating-digit-sum/
// Difficulty: Easy
// Time O(log n) | Space O(log n)

import "fmt"

func main() {
	fmt.Println(AlternatingDigitSum(521)) // 4
	fmt.Println(AlternatingDigitSum(111)) // 1
	fmt.Println(AlternatingDigitSum(886996)) // 0
}

func AlternatingDigitSum(n int) int {
	digits := []int{}
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	// Reverse to get original order
	sum := 0
	sign := 1
	for i := len(digits) - 1; i >= 0; i-- {
		sum += digits[i] * sign
		sign = -sign
	}
	return sum
}
```

## 2549 — Count Distinct Numbers On Board

```go
package main

// LeetCode #2549: Count Distinct Numbers on Board
// https://leetcode.com/problems/count-distinct-numbers-on-board/
// Difficulty: Easy
// Time O(1) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CountDistinctNumbersOnBoard(5)) // 4
	fmt.Println(CountDistinctNumbersOnBoard(2)) // 1
	fmt.Println(CountDistinctNumbersOnBoard(1)) // 1
}

func CountDistinctNumbersOnBoard(n int) int {
	if n == 1 {
		return 1
	}
	return n - 1
}
```

## 2553 — Separate The Digits In An Array

```go
package main

// LeetCode #2553: Separate the Digits in an Array
// https://leetcode.com/problems/separate-the-digits-in-an-array/
// Difficulty: Easy
// Time O(n log m) | Space O(n log m)

import "fmt"

func main() {
	fmt.Println(SeparateTheDigitsInAnArray([]int{13, 25, 83, 77})) // [1,3,2,5,8,3,7,7]
	fmt.Println(SeparateTheDigitsInAnArray([]int{7, 1, 3, 9}))     // [7,1,3,9]
}

func SeparateTheDigitsInAnArray(nums []int) []int {
	res := []int{}
	for _, n := range nums {
		digits := []int{}
		for n > 0 {
			digits = append(digits, n%10)
			n /= 10
		}
		for i := len(digits) - 1; i >= 0; i-- {
			res = append(res, digits[i])
		}
	}
	return res
}
```

## 2558 — Take Gifts From The Richest Pile

```go
package main

// LeetCode #2558: Take Gifts From the Richest Pile
// https://leetcode.com/problems/take-gifts-from-the-richest-pile/
// Difficulty: Easy
// Time O(k * n) | Space O(1)

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(TakeGiftsFromTheRichestPile([]int{25, 64, 9, 4, 100}, 4)) // 29
	fmt.Println(TakeGiftsFromTheRichestPile([]int{1, 1, 1, 1}, 4))         // 4
}

func TakeGiftsFromTheRichestPile(gifts []int, k int) int64 {
	for t := 0; t < k; t++ {
		maxIdx := 0
		for i := 1; i < len(gifts); i++ {
			if gifts[i] > gifts[maxIdx] {
				maxIdx = i
			}
		}
		gifts[maxIdx] = int(math.Sqrt(float64(gifts[maxIdx])))
	}
	sum := int64(0)
	for _, g := range gifts {
		sum += int64(g)
	}
	return sum
}
```

## 2562 — Find The Array Concatenation Value

```go
package main

// LeetCode #2562: Find the Array Concatenation Value
// https://leetcode.com/problems/find-the-array-concatenation-value/
// Difficulty: Easy
// Time O(n) | Space O(1)

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(FindTheArrayConcatenationValue([]int{7, 52, 2, 4})) // 596
	fmt.Println(FindTheArrayConcatenationValue([]int{5, 14, 13, 8, 12})) // 673
}

func FindTheArrayConcatenationValue(nums []int) int64 {
	sum := int64(0)
	i, j := 0, len(nums)-1
	for i < j {
		concat := int64(nums[i])
		n := nums[j]
		digits := 0
		if n == 0 {
			digits = 1
		} else {
			digits = int(math.Log10(float64(n))) + 1
		}
		concat = concat * pow10(digits) + int64(n)
		sum += concat
		i++
		j--
	}
	if i == j {
		sum += int64(nums[i])
	}
	return sum
}

func pow10(n int) int64 {
	p := int64(1)
	for i := 0; i < n; i++ {
		p *= 10
	}
	return p
}
```

## 2566 — Maximum Difference By Remapping A Digit

```go
package main

// LeetCode #2566: Maximum Difference by Remapping a Digit
// https://leetcode.com/problems/maximum-difference-by-remapping-a-digit/
// Difficulty: Easy
// Time O(log n) | Space O(log n)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(MaximumDifferenceByRemappingADigit(11891)) // 99009
	fmt.Println(MaximumDifferenceByRemappingADigit(90))    // 99
}

func MaximumDifferenceByRemappingADigit(num int) int {
	s := strconv.Itoa(num)

	// Find max: replace first non-9 digit with 9
	maxStr := []byte(s)
	for i := 0; i < len(maxStr); i++ {
		if maxStr[i] != '9' {
			replaceWith := maxStr[i]
			for j := i; j < len(maxStr); j++ {
				if maxStr[j] == replaceWith {
					maxStr[j] = '9'
				}
			}
			break
		}
	}
	maxVal, _ := strconv.Atoi(string(maxStr))

	// Find min: replace first non-0 digit (or non-1) with 0
	minStr := []byte(s)
	for i := 0; i < len(minStr); i++ {
		if minStr[i] != '0' && minStr[i] != '1' {
			replaceWith := minStr[i]
			for j := i; j < len(minStr); j++ {
				if minStr[j] == replaceWith {
					if i == 0 {
						minStr[j] = '1'
					} else {
						minStr[j] = '0'
					}
				}
			}
			break
		}
	}
	minVal, _ := strconv.Atoi(string(minStr))

	return maxVal - minVal
}
```

## 2570 — Merge Two 2D Arrays By Summing Values

```go
package main

// LeetCode #2570: Merge Two 2D Arrays by Summing Values
// https://leetcode.com/problems/merge-two-2d-arrays-by-summing-values/
// Difficulty: Easy
// Time O(n + m) | Space O(n + m)

import "fmt"

func main() {
	fmt.Println(MergeTwoTwoDArraysBySummingValues([][]int{{1, 2}, {2, 3}, {4, 5}}, [][]int{{1, 4}, {3, 2}, {4, 1}})) // [[1,6],[2,3],[3,2],[4,6]]
	fmt.Println(MergeTwoTwoDArraysBySummingValues([][]int{{2, 4}, {3, 6}, {5, 5}}, [][]int{{1, 3}, {4, 3}}))          // [[1,3],[2,4],[3,6],[4,3],[5,5]]
}

func MergeTwoTwoDArraysBySummingValues(nums1 [][]int, nums2 [][]int) [][]int {
	res := [][]int{}
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i][0] == nums2[j][0] {
			res = append(res, []int{nums1[i][0], nums1[i][1] + nums2[j][1]})
			i++
			j++
		} else if nums1[i][0] < nums2[j][0] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}
	for i < len(nums1) {
		res = append(res, nums1[i])
		i++
	}
	for j < len(nums2) {
		res = append(res, nums2[j])
		j++
	}
	return res
}
```

## 2574 — Left And Right Sum Differences

```go
package main

// LeetCode #2574: Left and Right Sum Differences
// https://leetcode.com/problems/left-and-right-sum-differences/
// Difficulty: Easy
// Time O(n) | Space O(n)

import "fmt"

func main() {
	fmt.Println(LeftAndRightSumDifferences([]int{10, 4, 8, 3})) // [15,1,11,22]
	fmt.Println(LeftAndRightSumDifferences([]int{1}))            // [0]
}

func LeftAndRightSumDifferences(nums []int) []int {
	n := len(nums)
	total := 0
	for _, v := range nums {
		total += v
	}
	res := make([]int, n)
	leftSum := 0
	for i, v := range nums {
		rightSum := total - leftSum - v
		diff := leftSum - rightSum
		if diff < 0 {
			diff = -diff
		}
		res[i] = diff
		leftSum += v
	}
	return res
}
```

## 2578 — Split With Minimum Sum

```go
package main

// LeetCode #2578: Split With Minimum Sum
// https://leetcode.com/problems/split-with-minimum-sum/
// Difficulty: Easy
// Time O(log n) | Space O(log n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(SplitWithMinimumSum(4325)) // 59
	fmt.Println(SplitWithMinimumSum(687))  // 75
}

func SplitWithMinimumSum(num int) int {
	digits := []int{}
	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}
	sort.Ints(digits)

	num1, num2 := 0, 0
	for i := 0; i < len(digits); i++ {
		if i%2 == 0 {
			num1 = num1*10 + digits[i]
		} else {
			num2 = num2*10 + digits[i]
		}
	}
	return num1 + num2
}
```

## 2582 — Pass The Pillow

```go
package main

// LeetCode #2582: Pass the Pillow
// https://leetcode.com/problems/pass-the-pillow/
// Difficulty: Easy
// Time O(1) | Space O(1)

import "fmt"

func main() {
	fmt.Println(PassThePillow(4, 5)) // 2
	fmt.Println(PassThePillow(3, 2)) // 3
}

func PassThePillow(n int, time int) int {
	cycle := 2 * (n - 1)
	t := time % cycle
	if t < n {
		return t + 1
	}
	return n - (t - n + 1)
}
```

## 2586 — Count The Number Of Vowel Strings In Range

```go
package main

// LeetCode #2586: Count the Number of Vowel Strings in Range
// https://leetcode.com/problems/count-the-number-of-vowel-strings-in-range/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CountTheNumberOfVowelStringsInRange([]string{"are", "amy", "u"}, 0, 2)) // 2
	fmt.Println(CountTheNumberOfVowelStringsInRange([]string{"hey", "aeo", "mu", "ooo", "artro"}, 1, 4)) // 3
}

func isVowel(ch byte) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
}

func CountTheNumberOfVowelStringsInRange(words []string, left int, right int) int {
	count := 0
	for i := left; i <= right; i++ {
		if len(words[i]) > 0 && isVowel(words[i][0]) && isVowel(words[i][len(words[i])-1]) {
			count++
		}
	}
	return count
}
```

## 2591 — Distribute Money To Maximum Children

```go
package main

// LeetCode #2591: Distribute Money to Maximum Children
// https://leetcode.com/problems/distribute-money-to-maximum-children/
// Difficulty: Easy
// Time O(1) | Space O(1)

import "fmt"

func main() {
	fmt.Println(DistributeMoneyToMaximumChildren(20, 3)) // 1
	fmt.Println(DistributeMoneyToMaximumChildren(16, 2)) // 2
}

func DistributeMoneyToMaximumChildren(money int, children int) int {
	if money < children {
		return -1
	}

	// Give each child 1 dollar first
	money -= children

	// Now we have 7-dollar increments (to make 8) for as many children as possible
	count := money / 7
	money %= 7

	// If count > children, we over-assigned
	if count > children {
		return children - 1
	}

	// If we have 3 children left and money = 3, we can't give it optimally
	remaining := children - count
	if remaining == 0 && money > 0 {
		count--
	} else if remaining == 1 && money == 3 {
		count--
	}

	return count
}
```

## 2595 — Number Of Even And Odd Bits

```go
package main

// LeetCode #2595: Number of Even and Odd Bits
// https://leetcode.com/problems/number-of-even-and-odd-bits/
// Difficulty: Easy
// Time O(log n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(NumberOfEvenAndOddBits(17)) // [2,0]
	fmt.Println(NumberOfEvenAndOddBits(2))  // [0,1]
}

func NumberOfEvenAndOddBits(n int) []int {
	even, odd := 0, 0
	idx := 0
	for n > 0 {
		if n&1 == 1 {
			if idx%2 == 0 {
				even++
			} else {
				odd++
			}
		}
		n >>= 1
		idx++
	}
	return []int{even, odd}
}
```

## 2600 — K Items With The Maximum Sum

```go
package main

// LeetCode #2600: K Items With the Maximum Sum
// https://leetcode.com/problems/k-items-with-the-maximum-sum/
// Difficulty: Easy
// Time O(1) | Space O(1)

import "fmt"

func main() {
	fmt.Println(KItemsWithTheMaximumSum(3, 2, 0, 2)) // 2
	fmt.Println(KItemsWithTheMaximumSum(3, 2, 0, 4)) // 3
}

func KItemsWithTheMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	if k <= numOnes {
		return k
	}
	k -= numOnes
	if k <= numZeros {
		return numOnes
	}
	k -= numZeros
	return numOnes - k
}
```

## 2605 — Form Smallest Number From Two Digit Arrays

```go
package main

// LeetCode #2605: Form Smallest Number From Two Digit Arrays
// https://leetcode.com/problems/form-smallest-number-from-two-digit-arrays/
// Difficulty: Easy
// Time O(n + m) | Space O(1)

import "fmt"

func main() {
	fmt.Println(FormSmallestNumberFromTwoDigitArrays([]int{4, 1, 3}, []int{5, 7}))       // 15
	fmt.Println(FormSmallestNumberFromTwoDigitArrays([]int{3, 5, 2, 6}, []int{3, 1, 7})) // 3
}

func FormSmallestNumberFromTwoDigitArrays(nums1 []int, nums2 []int) int {
	seen := [10]bool{}
	for _, n := range nums1 {
		seen[n] = true
	}

	common := 10
	for _, n := range nums2 {
		if seen[n] && n < common {
			common = n
		}
	}
	if common < 10 {
		return common
	}

	min1, min2 := 10, 10
	for _, n := range nums1 {
		if n < min1 {
			min1 = n
		}
	}
	for _, n := range nums2 {
		if n < min2 {
			min2 = n
		}
	}
	if min1 < min2 {
		return min1*10 + min2
	}
	return min2*10 + min1
}
```

## 2609 — Find The Longest Balanced Substring Of A Binary String

```go
package main

// LeetCode #2609: Find the Longest Balanced Substring of a Binary String
// https://leetcode.com/problems/find-the-longest-balanced-substring-of-a-binary-string/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(FindTheLongestBalancedSubstringOfABinaryString("01000111")) // 6
	fmt.Println(FindTheLongestBalancedSubstringOfABinaryString("00111"))    // 4
	fmt.Println(FindTheLongestBalancedSubstringOfABinaryString("111"))      // 0
}

func FindTheLongestBalancedSubstringOfABinaryString(s string) int {
	maxLen := 0
	i := 0
	n := len(s)

	for i < n {
		zeros, ones := 0, 0
		for i < n && s[i] == '0' {
			zeros++
			i++
		}
		for i < n && s[i] == '1' {
			ones++
			i++
		}
		pairLen := min(zeros, ones) * 2
		if pairLen > maxLen {
			maxLen = pairLen
		}
	}
	return maxLen
}
```

## 2614 — Prime In Diagonal

```go
package main

// LeetCode #2614: Prime In Diagonal
// https://leetcode.com/problems/prime-in-diagonal/
// Difficulty: Easy
// Time: O(n * sqrt(max(nums))) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(diagonalPrime([][]int{{1, 2, 3}, {5, 6, 7}, {9, 10, 11}}))
	fmt.Println(diagonalPrime([][]int{{1, 2, 3}, {5, 17, 7}, {9, 11, 10}}))
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func diagonalPrime(nums [][]int) int {
	n := len(nums)
	ans := 0
	for i := 0; i < n; i++ {
		if isPrime(nums[i][i]) && nums[i][i] > ans {
			ans = nums[i][i]
		}
		if isPrime(nums[i][n-i-1]) && nums[i][n-i-1] > ans {
			ans = nums[i][n-i-1]
		}
	}
	return ans
}
```

## 2619 — Array Prototype Last

```go
package main

// LeetCode #2619: Array Prototype Last
// https://leetcode.com/problems/array-prototype-last/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript problem, adapted to Go. Returns last element of a slice.

import "fmt"

func main() {
	fmt.Println(arrayPrototypeLast([]int{1, 2, 3}))
	fmt.Println(arrayPrototypeLast([]int{}))
}

func arrayPrototypeLast(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	return arr[len(arr)-1]
}
```

## 2620 — Counter

```go
package main

// LeetCode #2620: Counter
// https://leetcode.com/problems/counter/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript closure problem, adapted to Go. Returns a counter function.

import "fmt"

func main() {
	counter := counter(10)
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}

func counter(n int) func() int {
	return func() int {
		n++
		return n - 1
	}
}
```

## 2621 — Sleep

```go
package main

// LeetCode #2621: Sleep
// https://leetcode.com/problems/sleep/
// Difficulty: Easy
// Time: O(millis) | Space: O(1)
// Note: JavaScript async problem, adapted to Go.

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	sleep(100)
	fmt.Println("Slept for", time.Since(start).Milliseconds(), "ms")
}

func sleep(millis int) {
	time.Sleep(time.Duration(millis) * time.Millisecond)
}
```

## 2626 — Array Reduce Transformation

```go
package main

// LeetCode #2626: Array Reduce Transformation
// https://leetcode.com/problems/array-reduce-transformation/
// Difficulty: Easy
// Time: O(n) | Space: O(1)
// Note: JavaScript problem, adapted to Go. Reduces slice using a function.

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	sum := func(acc, curr int) int { return acc + curr }
	fmt.Println(arrayReduceTransformation(nums, sum, 0))

	nums2 := []int{1, 2, 3, 4}
	product := func(acc, curr int) int { return acc * curr }
	fmt.Println(arrayReduceTransformation(nums2, product, 1))
}

func arrayReduceTransformation(nums []int, fn func(int, int) int, init int) int {
	result := init
	for _, num := range nums {
		result = fn(result, num)
	}
	return result
}
```

## 2629 — Function Composition

```go
package main

// LeetCode #2629: Function Composition
// https://leetcode.com/problems/function-composition/
// Difficulty: Easy
// Time: O(n) | Space: O(1)
// Note: JavaScript problem, adapted to Go. Composes functions right-to-left.

import "fmt"

func main() {
	double := func(x int) int { return x * 2 }
	addOne := func(x int) int { return x + 1 }
	square := func(x int) int { return x * x }

	fns := []func(int) int{square, double, addOne}
	composed := functionComposition(fns)
	fmt.Println(composed(5)) // (5+1)*2 squared = 144
}

func functionComposition(functions []func(int) int) func(int) int {
	return func(x int) int {
		result := x
		for i := len(functions) - 1; i >= 0; i-- {
			result = functions[i](result)
		}
		return result
	}
}
```

## 2634 — Filter Elements From Array

```go
package main

// LeetCode #2634: Filter Elements from Array
// https://leetcode.com/problems/filter-elements-from-array/
// Difficulty: Easy
// Time: O(n) | Space: O(n)
// Note: JavaScript problem, adapted to Go. Filters slice using a predicate.

import "fmt"

func main() {
	nums := []int{0, 10, 20, 30}
	greaterThan10 := func(n int, i int) bool { return n > 10 }
	fmt.Println(FilterElementsFromArray(nums, greaterThan10))

	nums2 := []int{1, 2, 3}
	firstIndex := func(n int, i int) bool { return i == 0 }
	fmt.Println(FilterElementsFromArray(nums2, firstIndex))
}

func FilterElementsFromArray(arr []int, fn func(int, int) bool) []int {
	result := []int{}
	for i, v := range arr {
		if fn(v, i) {
			result = append(result, v)
		}
	}
	return result
}
```

## 2635 — Apply Transform Over Each Element In Array

```go
package main

// LeetCode #2635: Apply Transform Over Each Element in Array
// https://leetcode.com/problems/apply-transform-over-each-element-in-array/
// Difficulty: Easy
// Time: O(n) | Space: O(n)
// Note: JavaScript problem, adapted to Go. Maps a function over a slice.

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	double := func(n int, i int) int { return n * 2 }
	fmt.Println(ApplyTransformOverEachElementInArray(nums, double))

	nums2 := []int{1, 2, 3}
	timesIndex := func(n int, i int) int { return n * i }
	fmt.Println(ApplyTransformOverEachElementInArray(nums2, timesIndex))
}

func ApplyTransformOverEachElementInArray(arr []int, fn func(int, int) int) []int {
	result := make([]int, len(arr))
	for i, v := range arr {
		result[i] = fn(v, i)
	}
	return result
}
```

## 2639 — Find The Width Of Columns Of A Grid

```go
package main

// LeetCode #2639: Find the Width of Columns of a Grid
// https://leetcode.com/problems/find-the-width-of-columns-of-a-grid/
// Difficulty: Easy
// Time: O(m * n) | Space: O(n)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(FindTheWidthOfColumnsOfAGrid([][]int{{1}, {22}, {333}}))
	fmt.Println(FindTheWidthOfColumnsOfAGrid([][]int{{-15, 1, 3}, {15, 7, 12}, {5, 6, -2}}))
}

func FindTheWidthOfColumnsOfAGrid(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
	ans := make([]int, n)
	for j := 0; j < n; j++ {
		maxLen := 0
		for i := 0; i < m; i++ {
			width := len(strconv.Itoa(grid[i][j]))
			if width > maxLen {
				maxLen = width
			}
		}
		ans[j] = maxLen
	}
	return ans
}
```

## 2643 — Row With Maximum Ones

```go
package main

// LeetCode #2643: Row With Maximum Ones
// https://leetcode.com/problems/row-with-maximum-ones/
// Difficulty: Easy
// Time: O(m * n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(RowWithMaximumOnes([][]int{{0, 1}, {1, 0}}))
	fmt.Println(RowWithMaximumOnes([][]int{{0, 0, 0}, {0, 1, 1}}))
}

func RowWithMaximumOnes(mat [][]int) []int {
	maxRow, maxCount := 0, 0
	for i := 0; i < len(mat); i++ {
		count := 0
		for _, val := range mat[i] {
			if val == 1 {
				count++
			}
		}
		if count > maxCount {
			maxCount = count
			maxRow = i
		}
	}
	return []int{maxRow, maxCount}
}
```

## 2644 — Find The Maximum Divisibility Score

```go
package main

// LeetCode #2644: Find the Maximum Divisibility Score
// https://leetcode.com/problems/find-the-maximum-divisibility-score/
// Difficulty: Easy
// Time: O(|nums| * |divisors|) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(FindTheMaximumDivisibilityScore([]int{2, 3, 4, 5, 6}, []int{2, 3, 4}))
	fmt.Println(FindTheMaximumDivisibilityScore([]int{4, 7, 9, 3, 9}, []int{5, 2, 3}))
}

func FindTheMaximumDivisibilityScore(nums []int, divisors []int) int {
	ans := divisors[0]
	maxScore := 0
	for _, d := range divisors {
		score := 0
		for _, n := range nums {
			if n%d == 0 {
				score++
			}
		}
		if score > maxScore || (score == maxScore && d < ans) {
			maxScore = score
			ans = d
		}
	}
	return ans
}
```

## 2648 — Generate Fibonacci Sequence

```go
package main

// LeetCode #2648: Generate Fibonacci Sequence
// https://leetcode.com/problems/generate-fibonacci-sequence/
// Difficulty: Easy
// Time: O(n) | Space: O(1)
// Note: JavaScript generator problem, adapted to Go.

import "fmt"

func main() {
	fib := GenerateFibonacciSequence()
	for i := 0; i < 5; i++ {
		fmt.Println(fib())
	}
}

func GenerateFibonacciSequence() func() int {
	a, b := 0, 1
	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}
```

## 2651 — Calculate Delayed Arrival Time

```go
package main

// LeetCode #2651: Calculate Delayed Arrival Time
// https://leetcode.com/problems/calculate-delayed-arrival-time/
// Difficulty: Easy
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(CalculateDelayedArrivalTime(15, 5))
	fmt.Println(CalculateDelayedArrivalTime(23, 10))
}

func CalculateDelayedArrivalTime(arrivalTime int, delayedTime int) int {
	return (arrivalTime + delayedTime) % 24
}
```

## 2652 — Sum Multiples

```go
package main

// LeetCode #2652: Sum Multiples
// https://leetcode.com/problems/sum-multiples/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(SumMultiples(7))
	fmt.Println(SumMultiples(10))
}

func SumMultiples(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 || i%7 == 0 {
			sum += i
		}
	}
	return sum
}
```

## 2656 — Maximum Sum With Exactly K Elements

```go
package main

// LeetCode #2656: Maximum Sum With Exactly K Elements
// https://leetcode.com/problems/maximum-sum-with-exactly-k-elements/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(MaximumSumWithExactlyKElements([]int{1, 2, 3, 4, 5}, 3))
	fmt.Println(MaximumSumWithExactlyKElements([]int{5, 5, 5}, 2))
}

func MaximumSumWithExactlyKElements(nums []int, k int) int {
	maxVal := 0
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}
	// Sum of arithmetic series: maxVal + (maxVal+1) + ... + (maxVal+k-1)
	return k * (2*maxVal + k - 1) / 2
}
```

## 2660 — Determine The Winner Of A Bowling Game

```go
package main

// LeetCode #2660: Determine the Winner of a Bowling Game
// https://leetcode.com/problems/determine-the-winner-of-a-bowling-game/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(DetermineTheWinnerOfABowlingGame([]int{4, 10, 7, 9}, []int{6, 5, 2, 3}))
	fmt.Println(DetermineTheWinnerOfABowlingGame([]int{3, 5, 7, 6}, []int{8, 10, 10, 2}))
}

func DetermineTheWinnerOfABowlingGame(player1 []int, player2 []int) int {
	score1 := computeScore(player1)
	score2 := computeScore(player2)
	if score1 > score2 {
		return 1
	} else if score2 > score1 {
		return 2
	}
	return 0
}

func computeScore(pins []int) int {
	score := 0
	for i, v := range pins {
		score += v
		if (i >= 1 && pins[i-1] == 10) || (i >= 2 && pins[i-2] == 10) {
			score += v
		}
	}
	return score
}
```

## 2665 — Counter Ii

```go
package main

// LeetCode #2665: Counter II
// https://leetcode.com/problems/counter-ii/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript problem, adapted to Go. Returns an object with increment/decrement/reset.

import "fmt"

func main() {
	counter := CounterIi(5)
	fmt.Println(counter.increment())
	fmt.Println(counter.reset())
	fmt.Println(counter.decrement())
}

type Counter struct {
	init  int
	value int
}

func (c *Counter) increment() int {
	c.value++
	return c.value
}

func (c *Counter) decrement() int {
	c.value--
	return c.value
}

func (c *Counter) reset() int {
	c.value = c.init
	return c.value
}

func CounterIi(init int) *Counter {
	return &Counter{init: init, value: init}
}
```

## 2666 — Allow One Function Call

```go
package main

// LeetCode #2666: Allow One Function Call
// https://leetcode.com/problems/allow-one-function-call/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript problem, adapted to Go. Ensures fn is called at most once.

import "fmt"

func main() {
	fn := func(x int) int { return x * 2 }
	onceFn := AllowOneFunctionCall(fn)
	fmt.Println(onceFn(5))
	fmt.Println(onceFn(10))
}

func AllowOneFunctionCall(fn func(int) int) func(int) int {
	called := false
	return func(x int) int {
		if called {
			return 0
		}
		called = true
		return fn(x)
	}
}
```

## 2667 — Create Hello World Function

```go
package main

// LeetCode #2667: Create Hello World Function
// https://leetcode.com/problems/create-hello-world-function/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript problem, adapted to Go. Returns a function that always returns "Hello World".

import "fmt"

func main() {
	f := CreateHelloWorldFunction()
	fmt.Println(f())
	fmt.Println(f())
}

func CreateHelloWorldFunction() func() string {
	return func() string {
		return "Hello World"
	}
}
```

## 2668 — Find Latest Salaries

```go
package main

// LeetCode #2668: Find Latest Salaries
// https://leetcode.com/problems/find-latest-salaries/
// Difficulty: Easy [Paid]
// Time: O(n log n) | Space: O(n)
// Note: SQL/JS problem adapted to Go. Find latest salary per department.

import (
	"fmt"
	"sort"
)

func main() {
	salaries := []EmployeeSalary{
		{1, "Alice", 50000, "Eng", "2023-01-01"},
		{1, "Alice", 55000, "Eng", "2023-06-01"},
		{2, "Bob", 60000, "Sales", "2023-01-01"},
	}
	fmt.Println(FindLatestSalaries(salaries))
}

type EmployeeSalary struct {
	EmpID  int
	Name   string
	Salary int
	Dept   string
	Date   string
}

func FindLatestSalaries(salaries []EmployeeSalary) []EmployeeSalary {
	if len(salaries) == 0 {
		return nil
	}

	sort.Slice(salaries, func(i, j int) bool {
		if salaries[i].EmpID != salaries[j].EmpID {
			return salaries[i].EmpID < salaries[j].EmpID
		}
		return salaries[i].Date > salaries[j].Date
	})

	result := []EmployeeSalary{}
	seen := map[int]bool{}
	for _, s := range salaries {
		if !seen[s.EmpID] {
			seen[s.EmpID] = true
			result = append(result, s)
		}
	}
	return result
}
```

## 2669 — Count Artist Occurrences On Spotify Ranking List

```go
package main

// LeetCode #2669: Count Artist Occurrences On Spotify Ranking List
// https://leetcode.com/problems/count-artist-occurrences-on-spotify-ranking-list/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(n)
// Note: SQL/JS problem adapted to Go.

import "fmt"

func main() {
	songs := []struct {
		SongID int
		Artist string
	}{
		{1, "Drake"},
		{2, "Taylor Swift"},
		{3, "Drake"},
		{4, "Ed Sheeran"},
	}
	fmt.Println(CountArtistOccurrencesOnSpotifyRankingList(songs))
}

func CountArtistOccurrencesOnSpotifyRankingList(songs []struct {
	SongID int
	Artist string
}) map[string]int {
	counts := map[string]int{}
	for _, s := range songs {
		counts[s.Artist]++
	}
	return counts
}
```

## 2670 — Find The Distinct Difference Array

```go
package main

// LeetCode #2670: Find the Distinct Difference Array
// https://leetcode.com/problems/find-the-distinct-difference-array/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(FindTheDistinctDifferenceArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(FindTheDistinctDifferenceArray([]int{3, 2, 3, 4, 2}))
}

func FindTheDistinctDifferenceArray(nums []int) []int {
	n := len(nums)
	suffixDistinct := make([]int, n+1)
	seen := map[int]bool{}

	for i := n - 1; i >= 0; i-- {
		suffixDistinct[i] = suffixDistinct[i+1]
		if !seen[nums[i]] {
			seen[nums[i]] = true
			suffixDistinct[i]++
		}
	}

	seen = map[int]bool{}
	prefixDistinct := 0
	ans := make([]int, n)

	for i := 0; i < n; i++ {
		if !seen[nums[i]] {
			seen[nums[i]] = true
			prefixDistinct++
		}
		ans[i] = prefixDistinct - suffixDistinct[i+1]
	}

	return ans
}
```

## 2677 — Chunk Array

```go
package main

// LeetCode #2677: Chunk Array
// https://leetcode.com/problems/chunk-array/
// Difficulty: Easy
// Time: O(n) | Space: O(n)
// Note: JavaScript problem, adapted to Go. Splits array into chunks of given size.

import "fmt"

func main() {
	fmt.Println(ChunkArray([]int{1, 2, 3, 4, 5}, 1))
	fmt.Println(ChunkArray([]int{1, 9, 6, 3, 2}, 3))
}

func ChunkArray(arr []int, size int) [][]int {
	var result [][]int
	for i := 0; i < len(arr); i += size {
		end := i + size
		if end > len(arr) {
			end = len(arr)
		}
		result = append(result, arr[i:end])
	}
	return result
}
```

## 2678 — Number Of Senior Citizens

```go
package main

// LeetCode #2678: Number of Senior Citizens
// https://leetcode.com/problems/number-of-senior-citizens/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(NumberOfSeniorCitizens([]string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}))
	fmt.Println(NumberOfSeniorCitizens([]string{"1313579440F2036", "2921522980M5644"}))
}

func NumberOfSeniorCitizens(details []string) int {
	count := 0
	for _, d := range details {
		age, _ := strconv.Atoi(d[11:13])
		if age > 60 {
			count++
		}
	}
	return count
}
```

## 2682 — Find The Losers Of The Circular Game

```go
package main

// LeetCode #2682: Find the Losers of the Circular Game
// https://leetcode.com/problems/find-the-losers-of-the-circular-game/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(FindTheLosersOfTheCircularGame(5, 2))
	fmt.Println(FindTheLosersOfTheCircularGame(4, 4))
}

func FindTheLosersOfTheCircularGame(n int, k int) []int {
	visited := make([]bool, n)
	i := 0
	step := k
	for !visited[i] {
		visited[i] = true
		i = (i + step) % n
		step += k
	}

	result := []int{}
	for i := 0; i < n; i++ {
		if !visited[i] {
			result = append(result, i+1) // 1-indexed
		}
	}
	return result
}
```

## 2687 — Bikes Last Time Used

```go
package main

// LeetCode #2687: Bikes Last Time Used
// https://leetcode.com/problems/bikes-last-time-used/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(n)
// Note: SQL/JS problem adapted to Go. Find last time each bike was used.

import "fmt"

func main() {
	rides := []struct {
		BikeID int
		Time   int
	}{
		{1, 100},
		{2, 150},
		{1, 200},
		{3, 50},
	}
	fmt.Println(BikesLastTimeUsed(rides))
}

func BikesLastTimeUsed(rides []struct {
	BikeID int
	Time   int
}) map[int]int {
	lastUsed := map[int]int{}
	for _, r := range rides {
		if r.Time > lastUsed[r.BikeID] {
			lastUsed[r.BikeID] = r.Time
		}
	}
	return lastUsed
}
```

## 2689 — Extract Kth Character From The Rope Tree

```go
package main

// LeetCode #2689: Extract Kth Character From The Rope Tree
// https://leetcode.com/problems/extract-kth-character-from-the-rope-tree/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(h)
// Note: A rope tree is a binary tree where leaves contain substrings and internal nodes concatenate children.

import "fmt"

func main() {
	// Rope tree representing "abc" + "defg" = "abcdefg"
	root := &RopeNode{left: &RopeNode{val: "abc"}, right: &RopeNode{val: "defg"}}
	fmt.Println(string(ExtractKthCharacterFromTheRopeTree(root, 5))) // 'e'
}

type RopeNode struct {
	val   string
	left  *RopeNode
	right *RopeNode
}

func ExtractKthCharacterFromTheRopeTree(root *RopeNode, k int) byte {
	var dfs func(*RopeNode) string
	dfs = func(node *RopeNode) string {
		if node == nil {
			return ""
		}
		if node.left == nil && node.right == nil {
			return node.val
		}
		return dfs(node.left) + dfs(node.right)
	}
	return dfs(root)[k-1]
}
```

## 2690 — Infinite Method Object

```go
package main

// LeetCode #2690: Infinite Method Object
// https://leetcode.com/problems/infinite-method-object/
// Difficulty: Easy [Paid]
// Time: O(1) | Space: O(1)
// Note: JavaScript Proxy problem, adapted to Go. Returns an object that returns "methodName" for any method.

import "fmt"

func main() {
	obj := InfiniteMethodObject()
	fmt.Println(obj("abc"))
	fmt.Println(obj("xyz"))
}

func InfiniteMethodObject() func(string) string {
	return func(methodName string) string {
		return methodName
	}
}
```

## 2695 — Array Wrapper

```go
package main

// LeetCode #2695: Array Wrapper
// https://leetcode.com/problems/array-wrapper/
// Difficulty: Easy
// Time: O(n) | Space: O(n)
// Note: JavaScript problem, adapted to Go. Wraps an array with string conversion and addition.

import (
	"fmt"
	"strings"
)

func main() {
	w1 := ArrayWrapper([]int{1, 2})
	w2 := ArrayWrapper([]int{3, 4})
	fmt.Println(ArrayWrapperAdd(w1, w2))
	fmt.Println(ArrayWrapperString(w1))
}

type ArrayWrapperVal struct {
	nums []int
}

func ArrayWrapper(nums []int) *ArrayWrapperVal {
	return &ArrayWrapperVal{nums: nums}
}

func ArrayWrapperAdd(a, b *ArrayWrapperVal) int {
	sum := 0
	for _, v := range a.nums {
		sum += v
	}
	for _, v := range b.nums {
		sum += v
	}
	return sum
}

func ArrayWrapperString(w *ArrayWrapperVal) string {
	strs := make([]string, len(w.nums))
	for i, v := range w.nums {
		strs[i] = fmt.Sprintf("%d", v)
	}
	return "[" + strings.Join(strs, ",") + "]"
}
```

## 2696 — Minimum String Length After Removing Substrings

```go
package main

// LeetCode #2696: Minimum String Length After Removing Substrings
// https://leetcode.com/problems/minimum-string-length-after-removing-substrings/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(MinimumStringLengthAfterRemovingSubstrings("ABFCACDB"))
	fmt.Println(MinimumStringLengthAfterRemovingSubstrings("ACBBD"))
}

func MinimumStringLengthAfterRemovingSubstrings(s string) int {
	stack := []rune{}
	for _, c := range s {
		if len(stack) > 0 &&
			((stack[len(stack)-1] == 'A' && c == 'B') ||
				(stack[len(stack)-1] == 'C' && c == 'D')) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}
	return len(stack)
}
```

## 2697 — Lexicographically Smallest Palindrome

```go
package main

// LeetCode #2697: Lexicographically Smallest Palindrome
// https://leetcode.com/problems/lexicographically-smallest-palindrome/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(LexicographicallySmallestPalindrome("egcfe"))
	fmt.Println(LexicographicallySmallestPalindrome("abcd"))
}

func LexicographicallySmallestPalindrome(s string) string {
	runes := []rune(s)
	i, j := 0, len(runes)-1
	for i < j {
		if runes[i] < runes[j] {
			runes[j] = runes[i]
		} else {
			runes[i] = runes[j]
		}
		i++
		j--
	}
	return string(runes)
}
```

## 2703 — Return Length Of Arguments Passed

```go
package main

// LeetCode #2703: Return Length of Arguments Passed
// https://leetcode.com/problems/return-length-of-arguments-passed/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript problem, adapted to Go. Returns number of arguments.

import "fmt"

func main() {
	fmt.Println(ReturnLengthOfArgumentsPassed(1, 2, 3))
	fmt.Println(ReturnLengthOfArgumentsPassed("a", "b"))
}

func ReturnLengthOfArgumentsPassed(args ...interface{}) int {
	return len(args)
}
```

## 2704 — To Be Or Not To Be

```go
package main

// LeetCode #2704: To Be Or Not To Be
// https://leetcode.com/problems/to-be-or-not-to-be/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript problem, adapted to Go. Expect-type assertion.

import (
	"fmt"
	"reflect"
)

func main() {
	{
		expect := ToBeOrNotToBe(5)
		result, err := expect.toBe(5)
		fmt.Println(result, err)
	}
	{
		expect := ToBeOrNotToBe(5)
		result, err := expect.notToBe(5)
		fmt.Println(result, err)
	}
}

type Expect struct {
	val interface{}
}

func (e *Expect) toBe(expected interface{}) (bool, string) {
	if reflect.DeepEqual(e.val, expected) {
		return true, ""
	}
	return false, "Not Equal"
}

func (e *Expect) notToBe(expected interface{}) (bool, string) {
	if !reflect.DeepEqual(e.val, expected) {
		return true, ""
	}
	return false, "Equal"
}

func ToBeOrNotToBe(val interface{}) *Expect {
	return &Expect{val: val}
}
```

## 2706 — Buy Two Chocolates

```go
package main

// LeetCode #2706: Buy Two Chocolates
// https://leetcode.com/problems/buy-two-chocolates/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(BuyTwoChocolates([]int{1, 2, 2}, 3))
	fmt.Println(BuyTwoChocolates([]int{3, 2, 3}, 3))
}

func BuyTwoChocolates(prices []int, money int) int {
	sort.Ints(prices)
	if len(prices) >= 2 && prices[0]+prices[1] <= money {
		return money - prices[0] - prices[1]
	}
	return money
}
```

## 2710 — Remove Trailing Zeros From A String

```go
package main

// LeetCode #2710: Remove Trailing Zeros From a String
// https://leetcode.com/problems/remove-trailing-zeros-from-a-string/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(RemoveTrailingZerosFromAString("51230100"))
	fmt.Println(RemoveTrailingZerosFromAString("123"))
}

func RemoveTrailingZerosFromAString(num string) string {
	return strings.TrimRight(num, "0")
}
```

## 2715 — Timeout Cancellation

```go
package main

// LeetCode #2715: Timeout Cancellation
// https://leetcode.com/problems/timeout-cancellation/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript problem, adapted to Go. Returns a cancel function.

import (
	"fmt"
	"time"
)

func main() {
	cancel := TimeoutCancellation(func() { fmt.Println("executed") }, 100)
	time.Sleep(50 * time.Millisecond)
	cancel()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("done")
}

func TimeoutCancellation(fn func(), delay int) func() {
	timer := time.AfterFunc(time.Duration(delay)*time.Millisecond, fn)
	return func() {
		timer.Stop()
	}
}
```

## 2716 — Minimize String Length

```go
package main

// LeetCode #2716: Minimize String Length
// https://leetcode.com/problems/minimize-string-length/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(MinimizeStringLength("aaabc"))
	fmt.Println(MinimizeStringLength("cbbd"))
}

func MinimizeStringLength(s string) int {
	seen := map[rune]bool{}
	for _, c := range s {
		seen[c] = true
	}
	return len(seen)
}
```

## 2717 — Semi Ordered Permutation

```go
package main

// LeetCode #2717: Semi-Ordered Permutation
// https://leetcode.com/problems/semi-ordered-permutation/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(SemiOrderedPermutation([]int{2, 1, 4, 3}))
	fmt.Println(SemiOrderedPermutation([]int{2, 4, 1, 3}))
}

func SemiOrderedPermutation(nums []int) int {
	n := len(nums)
	pos1, posN := 0, 0
	for i, v := range nums {
		if v == 1 {
			pos1 = i
		}
		if v == n {
			posN = i
		}
	}

	swaps := pos1 + (n - 1 - posN)
	if pos1 > posN {
		swaps--
	}
	return swaps
}
```

## 2723 — Add Two Promises

```go
package main

// LeetCode #2723: Add Two Promises
// https://leetcode.com/problems/add-two-promises/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript Promise problem, adapted to Go. Adds two integer results asynchronously.

import (
	"fmt"
	"time"
)

func main() {
	p1 := func() int { time.Sleep(50 * time.Millisecond); return 10 }
	p2 := func() int { time.Sleep(100 * time.Millisecond); return 20 }
	result := AddTwoPromises(p1, p2)
	fmt.Println(result)
}

func AddTwoPromises(promise1 func() int, promise2 func() int) int {
	result1 := make(chan int)
	result2 := make(chan int)

	go func() { result1 <- promise1() }()
	go func() { result2 <- promise2() }()

	return <-result1 + <-result2
}
```

## 2724 — Sort By

```go
package main

// LeetCode #2724: Sort By
// https://leetcode.com/problems/sort-by/
// Difficulty: Easy
// Time: O(n log n) | Space: O(1)
// Note: JavaScript problem, adapted to Go.

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fn := func(n int) int { return n % 2 }
	fmt.Println(SortBy(arr, fn))

	arr2 := []int{1, 2, 3, 4, 5}
	fn2 := func(n int) int { return n }
	fmt.Println(SortBy(arr2, fn2))
}

func SortBy(arr []int, fn func(int) int) []int {
	sort.SliceStable(arr, func(i, j int) bool {
		return fn(arr[i]) < fn(arr[j])
	})
	return arr
}
```

## 2725 — Interval Cancellation

```go
package main

// LeetCode #2725: Interval Cancellation
// https://leetcode.com/problems/interval-cancellation/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript setInterval problem, adapted to Go.

import (
	"fmt"
	"time"
)

func main() {
	count := 0
	cancel := IntervalCancellation(func() {
		count++
		fmt.Println("tick", count)
	}, 50)
	time.Sleep(200 * time.Millisecond)
	cancel()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("total ticks:", count)
}

func IntervalCancellation(fn func(), delay int) func() {
	ticker := time.NewTicker(time.Duration(delay) * time.Millisecond)
	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				fn()
			case <-done:
				ticker.Stop()
				return
			}
		}
	}()

	return func() {
		close(done)
	}
}
```

## 2726 — Calculator With Method Chaining

```go
package main

// LeetCode #2726: Calculator with Method Chaining
// https://leetcode.com/problems/calculator-with-method-chaining/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript problem, adapted to Go. Calculator with method chaining.

import "fmt"

func main() {
	result := CalculatorWithMethodChaining(10).add(5).subtract(7).getResult()
	fmt.Println(result)
}

type Calculator struct {
	value int
}

func (c *Calculator) add(val int) *Calculator {
	c.value += val
	return c
}

func (c *Calculator) subtract(val int) *Calculator {
	c.value -= val
	return c
}

func (c *Calculator) multiply(val int) *Calculator {
	c.value *= val
	return c
}

func (c *Calculator) divide(val int) *Calculator {
	c.value /= val
	return c
}

func (c *Calculator) power(val int) *Calculator {
	result := 1
	for i := 0; i < val; i++ {
		result *= c.value
	}
	c.value = result
	return c
}

func (c *Calculator) getResult() int {
	return c.value
}

func CalculatorWithMethodChaining(initialValue int) *Calculator {
	return &Calculator{value: initialValue}
}
```

## 2727 — Is Object Empty

```go
package main

// LeetCode #2727: Is Object Empty
// https://leetcode.com/problems/is-object-empty/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript problem, adapted to Go. Checks if array/slice is empty.

import "fmt"

func main() {
	fmt.Println(IsObjectEmpty([]int{}))
	fmt.Println(IsObjectEmpty([]int{1, 2}))
}

func IsObjectEmpty(obj interface{}) bool {
	switch v := obj.(type) {
	case []int:
		return len(v) == 0
	case []string:
		return len(v) == 0
	default:
		return false
	}
}
```

## 2728 — Count Houses In A Circular Street

```go
package main

// LeetCode #2728: Count Houses in a Circular Street
// https://leetcode.com/problems/count-houses-in-a-circular-street/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(1)
// Note: Interactive problem adapted to Go. Uses street interface.

import "fmt"

func main() {
	// Simulate: street with 3 houses, open first door
	street := &StreetSim{houses: []bool{true, false, false}, idx: 0}
	fmt.Println(CountHousesInACircularStreet(street))
}

type Street interface {
	OpenDoor()
	CloseDoor()
	IsDoorOpen() bool
	MoveRight()
	MoveLeft()
}

type StreetSim struct {
	houses []bool
	idx    int
}

func (s *StreetSim) OpenDoor()    { s.houses[s.idx] = true }
func (s *StreetSim) CloseDoor()   { s.houses[s.idx] = false }
func (s *StreetSim) IsDoorOpen() bool { return s.houses[s.idx] }
func (s *StreetSim) MoveRight()   { s.idx = (s.idx + 1) % len(s.houses) }
func (s *StreetSim) MoveLeft()    { s.idx = (s.idx - 1 + len(s.houses)) % len(s.houses) }

func CountHousesInACircularStreet(street Street) int {
	// Open current door as marker
	street.OpenDoor()

	count := 0
	for {
		street.MoveRight()
		count++
		if street.IsDoorOpen() {
			break
		}
	}

	street.CloseDoor()
	return count
}
```

## 2729 — Check If The Number Is Fascinating

```go
package main

// LeetCode #2729: Check if The Number is Fascinating
// https://leetcode.com/problems/check-if-the-number-is-fascinating/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(CheckIfTheNumberIsFascinating(192))
	fmt.Println(CheckIfTheNumberIsFascinating(100))
}

func CheckIfTheNumberIsFascinating(n int) bool {
	concat := strconv.Itoa(n) + strconv.Itoa(n*2) + strconv.Itoa(n*3)
	if len(concat) != 9 {
		return false
	}

	digits := []byte(concat)
	sort.Slice(digits, func(i, j int) bool { return digits[i] < digits[j] })
	return string(digits) == "123456789"
}
```

## 2733 — Neither Minimum Nor Maximum

```go
package main

// LeetCode #2733: Neither Minimum nor Maximum
// https://leetcode.com/problems/neither-minimum-nor-maximum/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(NeitherMinimumNorMaximum([]int{3, 2, 1, 4}))
	fmt.Println(NeitherMinimumNorMaximum([]int{1, 2}))
}

func NeitherMinimumNorMaximum(nums []int) int {
	if len(nums) < 3 {
		return -1
	}

	minVal, maxVal := nums[0], nums[0]
	for _, v := range nums {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}

	for _, v := range nums {
		if v != minVal && v != maxVal {
			return v
		}
	}

	return -1
}
```

## 2739 — Total Distance Traveled

```go
package main

// LeetCode #2739: Total Distance Traveled
// https://leetcode.com/problems/total-distance-traveled/
// Difficulty: Easy
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(TotalDistanceTraveled(5, 10))
	fmt.Println(TotalDistanceTraveled(1, 2))
}

func TotalDistanceTraveled(mainTank int, additionalTank int) int {
	total := 0
	for mainTank > 0 {
		if mainTank >= 5 {
			mainTank -= 5
			total += 50
			if additionalTank > 0 {
				additionalTank--
				mainTank++
			}
		} else {
			total += mainTank * 10
			mainTank = 0
		}
	}
	return total
}
```

## 2744 — Find Maximum Number Of String Pairs

```go
package main

// LeetCode #2744: Find Maximum Number of String Pairs
// https://leetcode.com/problems/find-maximum-number-of-string-pairs/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(FindMaximumNumberOfStringPairs([]string{"cd", "ac", "dc", "ca", "zz"}))
	fmt.Println(FindMaximumNumberOfStringPairs([]string{"ab", "ba", "cc"}))
}

func FindMaximumNumberOfStringPairs(words []string) int {
	seen := map[string]bool{}
	count := 0
	for _, w := range words {
		rev := reverse(w)
		if seen[rev] {
			count++
		}
		seen[w] = true
	}
	return count
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
```

## 2748 — Number Of Beautiful Pairs

```go
package main

// LeetCode #2748: Number of Beautiful Pairs
// https://leetcode.com/problems/number-of-beautiful-pairs/
// Difficulty: Easy
// Time: O(n^2) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(NumberOfBeautifulPairs([]int{2, 5, 1, 4}))
	fmt.Println(NumberOfBeautifulPairs([]int{11, 21, 12}))
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func firstDigit(n int) int {
	for n >= 10 {
		n /= 10
	}
	return n
}

func lastDigit(n int) int {
	return n % 10
}

func NumberOfBeautifulPairs(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if gcd(firstDigit(nums[i]), lastDigit(nums[j])) == 1 {
				count++
			}
		}
	}
	return count
}
```

## 2758 — Next Day

```go
package main

// LeetCode #2758: Next Day
// https://leetcode.com/problems/next-day/
// Difficulty: Easy [Paid]
// Time: O(1) | Space: O(1)
// Note: JS Date problem adapted to Go.

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(NextDay("2024-03-01"))
	fmt.Println(NextDay("2024-12-31"))
}

func NextDay(date string) string {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return ""
	}
	return t.AddDate(0, 0, 1).Format("2006-01-02")
}
```

## 2760 — Longest Even Odd Subarray With Threshold

```go
package main

// LeetCode #2760: Longest Even Odd Subarray With Threshold
// https://leetcode.com/problems/longest-even-odd-subarray-with-threshold/
// Difficulty: Easy
// Time: O(n^2) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(LongestEvenOddSubarrayWithThreshold([]int{3, 2, 5, 4}, 5))
	fmt.Println(LongestEvenOddSubarrayWithThreshold([]int{4, 5, 2, 1}, 4))
}

func LongestEvenOddSubarrayWithThreshold(nums []int, threshold int) int {
	maxLen := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 != 0 || nums[i] > threshold {
			continue
		}
		length := 1
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > threshold {
				break
			}
			if nums[j]%2 == nums[j-1]%2 {
				break
			}
			length++
		}
		if length > maxLen {
			maxLen = length
		}
	}
	return maxLen
}
```

## 2765 — Longest Alternating Subarray

```go
package main

// LeetCode #2765: Longest Alternating Subarray
// https://leetcode.com/problems/longest-alternating-subarray/
// Difficulty: Easy
// Time: O(n^2) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(LongestAlternatingSubarray([]int{2, 3, 4, 3, 4}))
	fmt.Println(LongestAlternatingSubarray([]int{4, 5, 6}))
}

func LongestAlternatingSubarray(nums []int) int {
	n := len(nums)
	ans := -1
	for i := 0; i < n-1; i++ {
		if nums[i+1]-nums[i] != 1 {
			continue
		}
		length := 2
		expected := -1 // next diff should be -1
		for j := i + 2; j < n; j++ {
			diff := nums[j] - nums[j-1]
			if diff != expected {
				break
			}
			length++
			expected = -expected
		}
		if length > ans {
			ans = length
		}
	}
	return ans
}
```

## 2769 — Find The Maximum Achievable Number

```go
package main

// LeetCode #2769: Find the Maximum Achievable Number
// https://leetcode.com/problems/find-the-maximum-achievable-number/
// Difficulty: Easy
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(FindTheMaximumAchievableNumber(4, 1))
	fmt.Println(FindTheMaximumAchievableNumber(3, 2))
}

func FindTheMaximumAchievableNumber(num int, t int) int {
	return num + 2*t
}
```

## 2774 — Array Upper Bound

```go
package main

// LeetCode #2774: Array Upper Bound
// https://leetcode.com/problems/array-upper-bound/
// Difficulty: Easy [Paid]
// Time: O(log n) | Space: O(1)
// Note: JS problem, adapted to Go. Returns upper bound of target in sorted array.

import "fmt"

func main() {
	fmt.Println(ArrayUpperBound([]int{1, 2, 2, 2, 3}, 2))
	fmt.Println(ArrayUpperBound([]int{1, 3, 5}, 4))
}

func ArrayUpperBound(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left - 1
}
```

## 2778 — Sum Of Squares Of Special Elements

```go
package main

// LeetCode #2778: Sum of Squares of Special Elements
// https://leetcode.com/problems/sum-of-squares-of-special-elements/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(SumOfSquaresOfSpecialElements([]int{1, 2, 3, 4}))
	fmt.Println(SumOfSquaresOfSpecialElements([]int{2, 7, 1, 19, 18, 3}))
}

func SumOfSquaresOfSpecialElements(nums []int) int {
	n := len(nums)
	sum := 0
	for i, v := range nums {
		if n%(i+1) == 0 {
			sum += v * v
		}
	}
	return sum
}
```

## 2784 — Check If Array Is Good

```go
package main

// LeetCode #2784: Check if Array is Good
// https://leetcode.com/problems/check-if-array-is-good/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(CheckIfArrayIsGood([]int{2, 1, 3}))
	fmt.Println(CheckIfArrayIsGood([]int{3, 4, 4, 1, 2, 1}))
}

func CheckIfArrayIsGood(nums []int) bool {
	n := len(nums) - 1
	counts := make([]int, n+1)
	for _, v := range nums {
		if v > n {
			return false
		}
		counts[v]++
	}
	if counts[n] != 2 {
		return false
	}
	for i := 1; i < n; i++ {
		if counts[i] != 1 {
			return false
		}
	}
	return true
}
```

## 2788 — Split Strings By Separator

```go
package main

// LeetCode #2788: Split Strings by Separator
// https://leetcode.com/problems/split-strings-by-separator/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(SplitStringsBySeparator([]string{"one.two.three", "four.five", "six"}, '.'))
	fmt.Println(SplitStringsBySeparator([]string{"$easy$", "$problem$"}, '$'))
}

func SplitStringsBySeparator(words []string, separator byte) []string {
	result := []string{}
	for _, w := range words {
		parts := strings.Split(w, string(separator))
		for _, p := range parts {
			if p != "" {
				result = append(result, p)
			}
		}
	}
	return result
}
```

## 2794 — Create Object From Two Arrays

```go
package main

// LeetCode #2794: Create Object from Two Arrays
// https://leetcode.com/problems/create-object-from-two-arrays/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(n)
// Note: JS problem, adapted to Go. Creates a map from keys and values arrays.

import "fmt"

func main() {
	fmt.Println(CreateObjectFromTwoArrays([]string{"a", "b", "c"}, []int{1, 2, 3}))
	fmt.Println(CreateObjectFromTwoArrays([]string{"x"}, []int{10}))
}

func CreateObjectFromTwoArrays(keys []string, values []int) map[string]int {
	result := make(map[string]int, len(keys))
	for i, k := range keys {
		if i < len(values) {
			result[k] = values[i]
		}
	}
	return result
}
```

## 2796 — Repeat String

```go
package main

// LeetCode #2796: Repeat String
// https://leetcode.com/problems/repeat-string/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(n)
// Note: JS problem, adapted to Go. Repeats string n times.

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(RepeatString("abc", 3))
	fmt.Println(RepeatString("x", 5))
}

func RepeatString(s string, n int) string {
	return strings.Repeat(s, n)
}
```

## 2797 — Partial Function With Placeholders

```go
package main

// LeetCode #2797: Partial Function with Placeholders
// https://leetcode.com/problems/partial-function-with-placeholders/
// Difficulty: Easy [Paid]
// Time: O(1) | Space: O(1)
// Note: JS problem, adapted to Go. Partially applies arguments.

import "fmt"

func main() {
	add := func(a, b, c int) int { return a + b + c }
	partial := PartialFunctionWithPlaceholders(add, 1, nil, 3)
	fmt.Println(partial(2))
}

func PartialFunctionWithPlaceholders(fn func(int, int, int) int, args ...interface{}) func(int) int {
	return func(x int) int {
		realArgs := [3]int{}
		argIdx := 0
		for i, a := range args {
			if a == nil {
				realArgs[i] = x
			} else {
				realArgs[i] = a.(int)
				argIdx++
			}
		}
		return fn(realArgs[0], realArgs[1], realArgs[2])
	}
}
```

