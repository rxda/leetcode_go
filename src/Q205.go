package src

func isIsomorphic(s string, t string) bool {
	if s == t{
		return true
	}

	diffMap1 := make(map[byte]byte, len(s))
	diffMap2 := make(map[byte]byte, len(s))
	for i := 0; i < len(s); i++ {
		if val, ok := diffMap1[s[i]];!ok{
			diffMap1[s[i]] = t[i]
		}else{
			if val != t[i]{
				return false
			}
		}
		if val, ok := diffMap2[t[i]];!ok{
			diffMap2[t[i]] = s[i]
		}else{
			if val != s[i]{
				return false
			}
		}
	}
	return true
}

