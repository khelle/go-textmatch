package text

//
func initLast(s string) map[int]int {
	var c int
	var last map[int]int
	var ok bool

	last = make(map[int]int)

	for i:=0; i<len(s); i++ {
		c = int(s[i])

		if _, ok = last[c]; !ok || last[c] < i {
			last[c] = i
		}
	}

	return last
}

//
func findLast(c byte, last map[int]int) int {

	if val, ok := last[int(c)]; ok {
		return val
	}

	return -1
}

// Try to find needle inside of a haystack usign Boyer-Moore algorithm
func BoyerMoore(needle string, haystack string) int {
	M := len(needle)
	N := len(haystack)

	i := M-1
	j := M-1
	min := 0
	last := initLast(needle)
	cnt := 0

	for ; i<= N-1; {
		if needle[j] == haystack[i] {
			if j == 0 {
				return i
			} else {
				i--;
				j--;
			}
		} else {
			if j<1+findLast(haystack[i], last) {
				min=j
			} else {
				min=1+findLast(haystack[i], last)
			}
			i = i+M-min
			j = M-1
		}
	}

	return cnt
}
