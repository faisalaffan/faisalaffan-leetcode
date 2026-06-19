package main

// LeetCode #2622: Cache With Time Limit
// https://leetcode.com/problems/cache-with-time-limit/
// Difficulty: Medium
// Time: O(1) per operation average | Space: O(n)

import (
	"fmt"
	"time"
)

type TimedCache struct {
	cache    map[int]timedValue
}

type timedValue struct {
	value     int
	expiresAt time.Time
}

func NewTimedCache() TimedCache {
	return TimedCache{cache: make(map[int]timedValue)}
}

func (tc *TimedCache) Set(key int, value int, duration time.Duration) bool {
	now := time.Now()
	if existing, ok := tc.cache[key]; ok {
		if existing.expiresAt.After(now) {
			tc.cache[key] = timedValue{value, now.Add(duration)}
			return false // Key exists and hasn't expired
		}
	}
	tc.cache[key] = timedValue{value, now.Add(duration)}
	return true // Key was inserted (didn't exist or was expired)
}

func (tc *TimedCache) Get(key int) (int, bool) {
	now := time.Now()
	if existing, ok := tc.cache[key]; ok {
		if existing.expiresAt.After(now) {
			return existing.value, true
		}
		delete(tc.cache, key)
	}
	return 0, false
}

func (tc *TimedCache) Count() int {
	now := time.Now()
	count := 0
	for key, tv := range tc.cache {
		if tv.expiresAt.After(now) {
			count++
		} else {
			delete(tc.cache, key)
		}
	}
	return count
}

func main() {
	cache := NewTimedCache()

	// Test case 1: Set and Get
	cache.Set(1, 100, 500*time.Millisecond)
	val, ok := cache.Get(1)
	fmt.Println("Test 1:", val, ok)
	// Expected: 100 true

	// Test case 2: Count
	fmt.Println("Test 2:", cache.Count())
	// Expected: 1

	// Test case 3: Overwrite existing
	cache.Set(1, 200, 500*time.Millisecond)
	val, ok = cache.Get(1)
	fmt.Println("Test 3:", val, ok)
	// Expected: 200 true
}
