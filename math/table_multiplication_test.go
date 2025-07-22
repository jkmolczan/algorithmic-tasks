package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateMultiplicationTable(t *testing.T) {
	tests := map[string]struct {
		n        int
		expected []int
	}{
		"should return multiplication of 1": {
			n:        1,
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		"should return multiplication of 9": {
			n:        9,
			expected: []int{9, 18, 27, 36, 45, 54, 63, 72, 81, 90},
		},
		"should return nil for n < 1": {
			n:        -1,
			expected: nil,
		},
		"should return nil for n > 1000000": {
			n:        1000001,
			expected: nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := CreateMultiplicationTable(test.n)

			assert.Equal(t, test.expected, result)
		})
	}
}
