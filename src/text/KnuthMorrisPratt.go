package text

// Try to find needle inside of a haystack usign Rabin-Karp algorithm
func KnuthMorrisPratt(needle string, haystack string) int {
	M := len(needle)
	N := len(haystack)
	T := make([]int, 0)
	Tsize := 0
	i := 0
	k := 0
	m := 0

	for i=0; i<N; i++ {
		if haystack[i] == needle[0] {
			T = append(T, i)
		}
	}
	Tsize = len(T)

	for i=0; i<N-M && k<Tsize; i++ {
		m = T[k]

		if needle[i] != haystack[m+i] {
			i = 0
			k++

		} else if i == M-1 {
			return m
		}
	}

	return -1
}
