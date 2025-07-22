package math

var (
	minNumberValue = 1
	maxNumberValue = 1000000
)

// CreateMultiplicationTable creates the multiplication table from 1 to 10 for a given number n and returns the table as an array
func CreateMultiplicationTable(n int) []int {
	if n < minNumberValue || n > maxNumberValue {
		return nil
	}

	table := make([]int, 10)
	for i := 1; i <= 10; i++ {
		table[i-1] = i * n
	}

	return table
}
