package main

// LeetCode #3528: Unit Conversion I
// https://leetcode.com/problems/unit-conversion-i/
// Difficulty: Medium
// Complexity: O(n) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", UnitConversionI(100, "cm", "m"))
	// Test case 2
	fmt.Println("Test 2:", UnitConversionI(1, "km", "m"))
	// Test case 3
	fmt.Println("Test 3:", UnitConversionI(60, "min", "hr"))
}

func UnitConversionI(value float64, fromUnit, toUnit string) float64 {
	// Base unit: meter
	toMeter := map[string]float64{
		"mm": 0.001,
		"cm": 0.01,
		"dm": 0.1,
		"m":  1.0,
		"km": 1000.0,
		"in": 0.0254,
		"ft": 0.3048,
		"yd": 0.9144,
		"mi": 1609.344,
	}

	toMeterTime := map[string]float64{
		"sec": 1.0,
		"min": 60.0,
		"hr":  3600.0,
	}

	if factor1, ok := toMeter[fromUnit]; ok {
		if factor2, ok2 := toMeter[toUnit]; ok2 {
			return value * factor1 / factor2
		}
	}

	if factor1, ok := toMeterTime[fromUnit]; ok {
		if factor2, ok2 := toMeterTime[toUnit]; ok2 {
			return value * factor1 / factor2
		}
	}

	return value
}
