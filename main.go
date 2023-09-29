package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(romanToInt("MCMXCIV"))

}
func isMonotonic(nums []int) bool {
	increasing := true
	decreasing := true
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			decreasing = false
		} else if nums[i] < nums[i-1] {
			increasing = false
		}
		if !increasing && !decreasing {
			return false
		}
	}
	return true
}

func twoSum(nums []int, target int) []int {
	t := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		j, ok := t[target-nums[i]]
		if ok {
			result := []int{j, i}
			return result
		}
		t[nums[i]] = i
	}
	result := []int{-1, -1}
	return result
}

func romanToInt(s string) int {
	sum := 0
	var m = make(map[string]int)
	m["I"] = 1
	m["V"] = 5
	m["X"] = 10
	m["L"] = 50
	m["C"] = 100
	m["D"] = 500
	m["M"] = 1000
	s = strings.Replace(s, "IV", "IIII", 2)
	s = strings.Replace(s, "IX", "VIIII", 2)
	s = strings.Replace(s, "XL", "XXXX", 2)
	s = strings.Replace(s, "XC", "LXXXX", 2)
	s = strings.Replace(s, "CD", "CCCC", 2)
	s = strings.Replace(s, "CM", "DCCCC", 2)

	for i := 0; i < len(s); i++ {
		sum = sum + m[string([]rune(s)[i])]
	}

	return sum
}
