package math

import "fmt"

var (
	nMinValue = 1
	nMaxValue = 10000

	elementMinValue = -10000
	elementMaxValue = 10000

	nValueOutOfRangeErr       = fmt.Errorf("n must be between %d and %d", nMinValue, nMaxValue)
	elementValueOutOfRangeErr = fmt.Errorf("element value must be between %d and %d", elementMinValue, elementMaxValue)
)

// FindNthTermArithmeticSeries calculates the nth term of an arithmetic series given the first two terms.
func FindNthTermArithmeticSeries(a1, a2, n int) (int, error) {
	if n < nMinValue || n > nMaxValue {
		return -1, nValueOutOfRangeErr
	}

	if a1 < elementMinValue || a1 > elementMaxValue || a2 < elementMinValue || a2 > elementMaxValue {
		return -1, elementValueOutOfRangeErr
	}

	return a1 + (a2-a1)*(n-1), nil
}
