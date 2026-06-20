# 2960 — Count Tested Devices After Test Operations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func CountTestedDevicesAfterTestOperations(batteryPercentages []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n^2)  |  **Ruang:** O(1)


## 💻 Solusi Go

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
