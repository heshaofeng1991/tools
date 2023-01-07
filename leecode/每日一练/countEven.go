package main

func countEven(num int) int {
	var result int

	for i := 2; i <= num; i++ {
		var sum int
		tmp := i

		for tmp > 0 {
			sum += tmp % 10
			tmp /= 10
		}

		if sum&1 == 0 {
			result++
		}
	}

	return result
}
