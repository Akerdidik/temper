package quad

func QuadA(x, y int) string {
	compare := ""
	if x < 0 || y < 0 {
		return ""
	} else {
		for h := 0; h < y; h++ {
			for w := 0; w < x; w++ {
				if h == 0 && w == 0 || h == 0 && w == x-1 || h == y-1 && (w == 0 || w == x-1) {
					compare += "o"
				} else if w > 0 && w < x-1 && (h == 0 || h == y-1) {
					compare += "-"
				} else if w > 0 && w < x-1 && (h > 0 || h < y-1) {
					compare += " "
				} else {
					compare += "|"
				}
			}
			compare += "\n"
		}
	}
	return compare
}
