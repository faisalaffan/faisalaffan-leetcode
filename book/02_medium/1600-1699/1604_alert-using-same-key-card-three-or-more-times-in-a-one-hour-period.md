# 1604 — Alert Using Same Key Card Three Or More Times In A One Hour Period

## Deskripsi

**Soal:** [1604. Alert Using Same Key Card Three Or More Times In A One Hour Period](https://leetcode.com/problems/alert-using-same-key-card-three-or-more-times-in-a-one-hour-period/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N log N), Space: O(N)  
**Kompleksitas Ruang:** O(N)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1604: Alert Using Same Key-Card Three or More Times in a One Hour Period
// https://leetcode.com/problems/alert-using-same-key-card-three-or-more-times-in-a-one-hour-period/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(AlertNames([]string{"daniel", "daniel", "daniel", "luis", "luis", "luis", "luis"},
		[]string{"10:00", "10:40", "11:00", "09:00", "11:00", "13:00", "15:00"}))
	fmt.Println(AlertNames([]string{"alice", "alice", "alice", "bob", "bob", "bob", "bob"},
		[]string{"12:01", "12:00", "18:00", "21:00", "21:20", "21:30", "23:00"}))
	fmt.Println(AlertNames([]string{"a", "a", "a", "a", "b"},
		[]string{"00:00", "00:50", "01:00", "01:30", "00:10"}))
}

func AlertNames(keyName []string, keyTime []string) []string {
	// Time: O(N log N), Space: O(N)
	n := len(keyName)
  // Membuat map untuk pencarian O(1): key → value
	records := make(map[string][]int)

	for i := 0; i < n; i++ {
		minutes := parseTime(keyTime[i])
		records[keyName[i]] = append(records[keyName[i]], minutes)
	}

  // Membuat slice untuk menyimpan hasil
	alerted := make([]string, 0)

	for name, times := range records {
		sort.Ints(times)
		for i := 2; i < len(times); i++ {
			if times[i]-times[i-2] <= 60 {
				alerted = append(alerted, name)
				break
			}
		}
	}

	sort.Strings(alerted)
	return alerted
}

func parseTime(t string) int {
	hours := int(t[0]-'0')*10 + int(t[1]-'0')
	minutes := int(t[3]-'0')*10 + int(t[4]-'0')
	return hours*60 + minutes
}
```
