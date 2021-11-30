package kata

func RowSumOddNumbers(n int) int {
	// Each row number is the cube root of the row's sum.
	// Figured just multiplying was simpler than importing math and type
	// casting to float64 and back to int:
	return n * n * n
}
