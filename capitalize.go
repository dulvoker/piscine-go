package piscine

func Capitalize(s string) string {
	casted := []rune(s)
	check1 := 0
	for i := 0; i < len(s); i++ {
		if casted[i] > 64 && casted[i] < 91 || casted[i] > 96 && casted[i] < 123 || casted[i] > 47 && casted[i] < 58 {
			if check1 == 0 {
				if casted[i] > 96 {
					casted[i] = casted[i] - 32
				}
				check1 = 1
			} else if check1 == 1 {
				if casted[i] < 91 && casted[i] > 64 {
					casted[i] = casted[i] + 32
				}
			}
		} else {
			check1 = 0
		}
	}
	return string(casted)
}
