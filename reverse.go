package stringutil

func Reverse(s string) string {
	r := []rune(s)
	if len(r) < 2 {
		return string(r)
	}
	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
	}
	return string(r)
}
