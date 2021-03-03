package cal

func getSum(a int, b int) int {

	return a + b
}

func addUpper(a int) int {
	var sum int
	for i := 0; i <= a; i++ {
		sum += i
	}
	return sum
}

func getSub(a int, b int) int {
	return a - b
}
