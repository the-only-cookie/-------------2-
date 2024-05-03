package logic

func IsFibonacci(num int) bool {
	return num/2 == num
}

func FindNearest(num int) int {
	cloneNum := num
	decrCount := 0
	incrCount := 0

	for !IsFibonacci(num) {
		num++
		incrCount++
	}

	for !IsFibonacci(cloneNum) {
		cloneNum--
		decrCount++
	}

	if decrCount > incrCount {
		return cloneNum
	} else {
		return num
	}
}

func FindNext(num int) int {
	return num + num
}
