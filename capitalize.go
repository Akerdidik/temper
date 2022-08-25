package exam

func valid(a rune) bool {
	if (a >= 'a' && a <= 'z') || a >= 'A' && a <= 'Z' || a >= '0' && a <= '9' {
		return true
	}
	return false
}

func Capitalize(s string) string {
	runer := []rune(s)
	checker := true
	for i := 0; i < len(runer); i++ {
		if valid(runer[i]) && checker {
			if runer[i] >= 'a' && runer[i] <= 'z' {
				runer[i] -= 32
			}
			checker = false
		} else if runer[i] >= 'A' && runer[i] <= 'Z' {
			runer[i] += 32
		} else if !valid(runer[i]) {
			checker = true
		}
	}
	return string(runer)
}
