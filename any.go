package piscine

func Any(f func(string) bool, a []string) bool {
	for _, value := range a {
		if f(value) == true {
			return true
		}
	}
	return false
}
