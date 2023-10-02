package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	sr := []int{1,3,5,6}

	fmt.Print(searchInsert(sr, 2))

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

func longestCommonPrefix(strs []string) string {
	var longestPrefix string = ""
	var endPrefix = false

	if len(strs) > 0 {
		sort.Strings(strs)
		first := string(strs[0])
		last := string(strs[len(strs)-1])

		for i := 0; i < len(first); i++ {
			if !endPrefix && string(last[i]) == string(first[i]) {
				longestPrefix += string(last[i])
			} else {
				endPrefix = true
			}
		}
	}
	return longestPrefix
}

func isValid(str string) bool {
	var s []byte

	if str[0] == '(' || str[0] == '{' || str[0] == '[' {

		for i := 0; i < len(str); i++ {
			s = append(s, str[i])
			if len(s) > 1 {
				if s[len(s)-2] == '[' && s[len(s)-1] == ']' {
					s = s[:len(s)-2]
				} else if s[len(s)-2] == '{' && s[len(s)-1] == '}' {
					s = s[:len(s)-2]
				} else if s[len(s)-2] == '(' && s[len(s)-1] == ')' {
					s = s[:len(s)-2]
				}
			}
		}
		if len(s) == 0 {
			return true
		}
	}
	return false
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}

	return nil
}


type ListNode struct {
    Val int
    Next *ListNode
}

func searchInsert(nums []int, target int) int {


	for i :=0; i < len(nums); i++ {
		if target <= nums[i] {
			return i
		}

		if target == nums[i] {
			return i
		}
	}


	return len(nums)
}



