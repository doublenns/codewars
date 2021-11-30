package kata

func Pyramid(n int) [][]int {
	res := [][]int{}
	for i := 0; i < n; i++ {
		row := []int{}
		for j := 0; j <= i; j++ {
			row = append(row, 1)
		}
		res = append(res, row)
	}
	return res
}

func PyramidOptimized(n int) [][]int {
	row := [][]int{}
	cell := []int{}

	for i := 0; i < n; i++ {
		cell = append(cell, 1)
		row = append(row, cell)
	}

	return row
}
