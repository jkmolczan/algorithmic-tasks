package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindNthTermArithmeticSeries(t *testing.T) {

	tests := map[string]struct {
		a1, a2, n     int
		expected      int
		expectedError error
	}{
		"should return 4th value for a1=2, a2=3, series 2,3,4,5,6...": {
			a1:       2,
			a2:       3,
			n:        4,
			expected: 5,
		},
		"should return 10th value for a1 = 1, a2 = 3, series 1,3,5,7,9,11,13,15,17,19,21...": {
			a1:       1,
			a2:       3,
			n:        10,
			expected: 19,
		},
		"should return 1st value for a1 = 1, a2 = 3, series 1,3,5,7,9,11,13,15,17,19,21...": {
			a1:       1,
			a2:       3,
			n:        1,
			expected: 1,
		},
		"should return 2nd value for a1 = 1, a2 = 3, series 1,3,5,7,9,11,13,15,17,19,21...": {
			a1:       1,
			a2:       3,
			n:        2,
			expected: 3,
		},
		"should return error for n < 1": {
			a1:            1,
			a2:            3,
			n:             0,
			expected:      -1,
			expectedError: nValueOutOfRangeErr,
		},
		"should return error for n > 10000": {
			a1:            1,
			a2:            3,
			n:             10001,
			expected:      -1,
			expectedError: nValueOutOfRangeErr,
		},
		"should return error for a1 < -10000": {
			a1:            -10001,
			a2:            3,
			n:             1,
			expected:      -1,
			expectedError: elementValueOutOfRangeErr,
		},
		"should return error for a1 > 10000": {
			a1:            10001,
			a2:            3,
			n:             1,
			expected:      -1,
			expectedError: elementValueOutOfRangeErr,
		},
		"should return error for a2 < -10000": {
			a1:            1,
			a2:            -10001,
			n:             1,
			expected:      -1,
			expectedError: elementValueOutOfRangeErr,
		},
		"should return error for a2 > 10000": {
			a1:            1,
			a2:            10001,
			n:             1,
			expected:      -1,
			expectedError: elementValueOutOfRangeErr,
		},
		"should return proper value for a1 == a2": {
			a1:       5,
			a2:       5,
			n:        6,
			expected: 5,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := FindNthTermArithmeticSeries(test.a1, test.a2, test.n)

			assert.Equal(t, test.expected, result)
			if test.expectedError != nil {
				assert.EqualError(t, err, test.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
