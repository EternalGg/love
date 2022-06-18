package benchmark

func Rangess(a []int) {
	var zz int
	for i := range a {
		zz += i
	}
}

func Forrr(a []int) {
	lenth := len(a)
	var zz int
	for i := 0; i < lenth; i++ {
		zz += a[i]
	}
}
