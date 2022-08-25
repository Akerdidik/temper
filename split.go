package exam

func Split(s string, sep string) []string {
	slen := len(s)
	seplen := len(sep)
	size := 0
	/* Hello World = 11, " " = 1 , indicat = 10 , size = 1
	   i = 0 , st = 0, index = 0 ,

	*/
	for i := 0; i <= slen-seplen; i++ {
		if string(s[i:i+seplen]) == sep {
			size++
		}
	}
	res := make([]string, size+1)
	i := 0
	st := 0
	dex := 0
	for ; dex <= slen-seplen; dex++ {
		if string(s[dex:dex+seplen]) == sep {
			res[i] = string(s[st:dex])
			i++
			st = dex + seplen
		}
		if dex == slen-seplen {
			res[i] = string(s[st:])
		}
	}
	return res
}
