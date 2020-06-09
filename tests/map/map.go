package mp

func getKeys(mp map[string]int) []string {
	var keys []string
	for k := range mp {
		keys = append(keys, k)
	}
	return keys
}
