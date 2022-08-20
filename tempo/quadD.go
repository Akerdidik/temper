package quad

func QuadD(x, y int) string {
	compare := ""
	if x < 0 || y < 0 {
		return ""
	} else {
		for h := 0; h < y; h++ {
			for w := 0; w < x; w++ {
				if w == 0 && (h == 0 || h == y-1) {
					compare += "A"
				} else if w > 0 && w < x-1 && (h == 0 || h == y-1) || h < y-1 && h > 0 && (w == 0 || w == x-1) {
					compare += "B"
				} else if (h == y-1 || h == 0) && (w == 0 || w == x-1) {
					compare += "C"
				} else {
					compare += " "
				}
			}
			compare += "\n"
		}
	}
	return compare
}
