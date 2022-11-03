package allocs

func f1() []int64 {
	var x []int64
	size := 10
	for i := 0; i < 3; i++ {
		x = make([]int64, size)
	}
	return x
}
