package arrays

/*
Given an integer array nums, rotate the array to the right by rotateStep steps, where rotateStep is non-negative.

Example 1:
Input: nums = [1,2,3,4,5,6,7], rotateStep = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]

Example 2:
Input: nums = [-1,-100,3,99], rotateStep = 2
Output: [3,99,-1,-100]
Explanation:
rotate 1 step to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]

*/

func RotateArray(nums []int, k int) {
	if len(nums) == 0 || k <= 0 {
		return
	}

	k = k % len(nums) // Handle cases where rotateStep is greater than the length of the array

	copy(nums, append(nums[len(nums)-k:], nums[:len(nums)-k]...))
	return
}
