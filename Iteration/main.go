package iteration

func Repeat(character string, num int) string {
	var repeated string
	for i := 0; i < num; i++ {
		repeated += character
	}
	return repeated
}

func clone(s string) string {
	return s
}

func Compare(x, y string) int {
	if x == y {
		return 0
	}
	if x < y {
		return -1
	}
	return +1
}
