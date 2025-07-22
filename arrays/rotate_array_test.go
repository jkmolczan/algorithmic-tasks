package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RotateArray(t *testing.T) {

	tests := map[string]struct {
		arr        []int
		rotateStep int
		expected   []int
	}{
		"should do nothing for empty array": {
			arr:        []int{},
			rotateStep: 0,
			expected:   []int{},
		},
		"should do nothing for non-empty array and rotation step set to 0": {
			arr:        []int{-1, -100, 3, 99},
			rotateStep: 0,
			expected:   []int{-1, -100, 3, 99},
		},
		"should do nothing for non-empty array and rotation step set to -1": {
			arr:        []int{-1, -100, 3, 99},
			rotateStep: -1,
			expected:   []int{-1, -100, 3, 99},
		},
		"should rotate array [-1,-100,3,99] by 2 steps": {
			arr:        []int{-1, -100, 3, 99},
			rotateStep: 2,
			expected:   []int{3, 99, -1, -100},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			RotateArray(test.arr, test.rotateStep)

			assert.Equal(t, test.expected, test.arr)
		})
	}
}
