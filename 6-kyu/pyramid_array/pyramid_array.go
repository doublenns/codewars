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
