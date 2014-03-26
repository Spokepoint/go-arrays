package arrays

// Contains check if 's' is in 'coll' string array
func Contains(s string, coll []string) bool {
	for _, v := range coll {
		if v == s {
			return true
		}
	}
	return false
}
