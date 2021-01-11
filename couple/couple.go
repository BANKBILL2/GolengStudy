package couple

func split(s string) []string {
	c := []string{}
	s += "*"

	for ; len(s) > 1; s = s[2:] {
		c = append(c, s[0:2])
	}

	return c
}

func split2(str string) []string {
	words := []string{}
	str += "*"

	for len(str) > 1 {
		words, str = append(words, str[:2]), str[2:]
	}

	return words
}
