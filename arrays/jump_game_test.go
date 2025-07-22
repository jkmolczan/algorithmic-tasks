package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_JumpGame(t *testing.T) {
	tests := map[string]struct {
		arr            []int
		expectedOutput bool
	}{
		"should reach the last index for [0] (only one element in array)": {
			arr:            []int{0},
			expectedOutput: true,
		},
		"should reach the last index for [1] (only one element in array)": {
			arr:            []int{1},
			expectedOutput: true,
		},
		"should not reach the last index for [] (empty array)": {
			arr:            []int{},
			expectedOutput: false,
		},
		"should not reach the last index for [0, 1]": {
			arr:            []int{0, 1},
			expectedOutput: false,
		},
		"should reach the last index for [1, 0]": {
			arr:            []int{1, 0},
			expectedOutput: true,
		},
		"should reach the last index for [1,1,1,1,1,1]": {
			arr:            []int{1, 1, 1, 1, 1, 1},
			expectedOutput: true,
		},
		"should reach the last index for [2, 0, 1, 1, 1, 1]": {
			arr:            []int{2, 0, 1, 1, 1, 1},
			expectedOutput: true,
		},
		"should reach last index for [1,1,2,2,0,1,1]": {
			arr:            []int{1, 1, 2, 2, 0, 1, 1},
			expectedOutput: true,
		},
		"should reach the last index for [2,3,1,1,4]": {
			arr:            []int{2, 3, 1, 1, 4},
			expectedOutput: true,
		},
		"should not reach the last index for [3,2,1,0,4]": {
			arr:            []int{3, 2, 1, 0, 4},
			expectedOutput: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := CanJump(test.arr)

			assert.Equal(t, test.expectedOutput, result)
		})
	}
}
