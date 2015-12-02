package text

// Hash string and return sum of its ascii characters function
func Hash(s string) int {
	sum := 0

	for _, c := range s {
		sum += int(c)
	}

	return sum
}

// N function for fast rehashing
func NHash(hash int, first byte, last byte) int {
	return hash - int(first) + int(last)
}

// Try to find needle inside of a haystack usign Rabin-Karp algorithm
func RabinKarp(needle string, haystack string) int {
	M := len(needle)
	N := len(haystack)
	Sx := Hash(needle)
	Sy := Hash(haystack[0:M])

	for i:=0; i<N-M; i++ {
		if Sx == Sy && needle == haystack[i:i+M] {
			return i
		}

		Sy = NHash(Sy, haystack[i], haystack[i+M])
	}

	return -1
}
