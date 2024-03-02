package main

import "strings"

func singleNumber(nums []int) int {
	res := 0
	for n := range nums {
		res ^= nums[n]
	}
	return res
}

// 27
func removeElement(nums []int, val int) int {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}

	return index
}

// 26

func removeDuplicates(nums []int) int {
	var last int = nums[0]
	var index int = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != last {
			nums[index] = nums[i]
			last = nums[i]
			index++
		}
	}

	return index
}

// 35
func searchInsert(nums []int, target int) int {
	var left, right int = 0, len(nums) - 1
	for left <= right {
		med := (left + right) / 2
		if nums[med] == target {
			return med
		} else if nums[med] < target {
			left = med + 1
		} else {
			right = med - 1
		}
	}
	return left
}

// 88
func merge(nums1 []int, m int, nums2 []int, n int) {
	var right = m + n - 1
	m--
	n--
	for m >= 0 && n >= 0 {
		if nums1[m] >= nums2[n] {
			nums1[right] = nums1[m]
			m--
		} else {
			nums1[right] = nums2[n]
			n--
		}
		right--
	}
	for n >= 0 {
		nums1[right] = nums2[n]
		n--
		right--
	}
}

// 3042
func countPrefixSuffixPairs(words []string) int {
	res := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if isMatch(words[i], words[j]) {
				res++
			}
		}
	}
	return res
}

func isMatch(word1, word2 string) bool {
	return strings.HasPrefix(word2, word1) && strings.HasSuffix(word2, word1)
}

// 3038
func maxOperations(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	res := 1
	sum := nums[0] + nums[1]
	for i := 2; i < len(nums)-1; i += 2 {
		if sum == nums[i]+nums[i+1] {
			res++
		} else {
			break
		}
	}
	return res
}

// 3010
func minimumCost(nums []int) int {
	res := 150

	left, right := 1, len(nums)-1

	for left < right {
		if res > nums[left]+nums[right] {
			res = nums[left] + nums[right]
		}
		if nums[left] < nums[right] {
			right--
		} else {
			left++
		}
	}

	return res + nums[0]
}

// 3024
func triangleType(nums []int) string {
	if isTriangle(nums) {
		return "none"
	}
	if nums[0] == nums[1] && nums[1] == nums[2] {
		return "equilateral"
	}

	if nums[0] == nums[1] || nums[0] == nums[2] {
		return "isosceles"
	}
	return "scalene"
}

func isTriangle(nums []int) bool {
	if nums[0] >= nums[1]+nums[2] || nums[1] >= nums[0]+nums[2] || nums[2] >= nums[0]+nums[1] {
		return false
	}
	return true
}
