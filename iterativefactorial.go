package piscine

func IterativeFactorial(nb int) int {
	if nb > 0 {
		result := 1
		for i := 1; i <= nb; i++ {
			result *= i
			if result > 2147483647 {
				return 0
			}
		}
		return result

	} else if nb == 1 || nb == 0 {
		return 1
	} else if nb < 0 {
		return 0
	} else {
		return 0
	}
}
