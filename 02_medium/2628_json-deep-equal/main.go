package main

// LeetCode #2628: JSON Deep Equal
// https://leetcode.com/problems/json-deep-equal/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func jsonDeepEqual(o1, o2 any) bool {
	return reflect.DeepEqual(o1, o2)
}

func deepEqualJSON(a, b string) bool {
	var v1, v2 any
	if err := json.Unmarshal([]byte(a), &v1); err != nil {
		return false
	}
	if err := json.Unmarshal([]byte(b), &v2); err != nil {
		return false
	}
	return jsonDeepEqual(v1, v2)
}

func main() {
	// Test case 1: equal objects
	fmt.Println("Test 1:", deepEqualJSON(`{"a":1,"b":2}`, `{"b":2,"a":1}`))
	// Expected: true

	// Test case 2: different types
	fmt.Println("Test 2:", deepEqualJSON(`{"a":1}`, `{"a":"1"}`))
	// Expected: false

	// Test case 3: arrays
	fmt.Println("Test 3:", deepEqualJSON(`[1,2,3]`, `[1,2,3]`))
	// Expected: true
}
