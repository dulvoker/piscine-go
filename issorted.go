package piscine

func f(a, b int) int {
	return a - b
}

func IsSorted(f func(a, b int) int, a []int) bool {
	for i, each := range a {
		if i == len(a)-2 {
			break
		} else if f(each, a[i+1]) != f(a[i+1], a[i+2]) {
			return false
		}
	}
	return true
}
