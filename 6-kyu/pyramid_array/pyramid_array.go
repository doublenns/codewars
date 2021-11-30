package kata

func Pyramid(n int) [][]int {
	res := [][]int{}
	for i := 0; i < n; i++ {
		brick := []int{}
		for j := 0; j <= i; j++ {
			brick = append(brick, 1)
		}
		res = append(res, brick)
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
