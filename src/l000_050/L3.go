package main

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
