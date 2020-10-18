package swap

func Swap(x int, y int) (int, int) {
	return y, x
}

func SwapPointers(x *int, y *int) {
	*x, *y = *y, *x
}
