package helper

// GetKeys returns string array pointer, which contains keys of map m
func GetKeys(m map[string]string) *[]string {
	if m == nil || len(m) < 1 {
		s := make([]string, 0)
		return &s
	}

	s := make([]string, len(m))
	i := 0
	for k := range m {
		s[i] = k
		i++
	}

	return &s
}
