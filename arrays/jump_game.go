package arrays

/*
You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.
Return true if you can reach the last index, or false otherwise.
Example 1:
Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

Example 2:
Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.
*/

func CanJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	if len(nums) == 1 {
		return true
	}

	if nums[0] == 0 {
		return false
	}

	maxReach := 0
	for i, jump := range nums {
		if i > maxReach {
			return false // If we reach a point beyond the maximum reachable index, return false (case with 0)
		}

		if jump+i > maxReach {
			maxReach = i + jump
		}
		if maxReach >= len(nums)-1 {
			return true
		}
	}

	return false // If we finish the loop without reaching the end, return false
}
