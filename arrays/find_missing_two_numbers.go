package arrays

// Given an array containing all the numbers from 1 to n except two numbers, find the two missing elements.
// For example, given the array [4, 2, 3], the missing elements are 1 and 5.

func FindMissingTwoNumbers(arr []int) (int, int) {
	var missingElements []int
	for i := 1; i <= len(arr)+2; i++ {
		if !ifArrContainsElement(arr, i) {
			missingElements = append(missingElements, i)
		}
	}

	return missingElements[0], missingElements[1]
}

func ifArrContainsElement(arr []int, num int) bool {
	for _, value := range arr {
		if value == num {
			return true
		}
	}
	return false
}
