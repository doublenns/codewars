package kata

func MakeNegative(x int) int {
	if x > 0 {
		x = 0 - x
	}
	return x
}
