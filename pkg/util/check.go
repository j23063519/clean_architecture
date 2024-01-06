package util

func CheckMapStrStr(key []string, values map[string]string) bool {
	for _, v := range key {
		val, exist := values[v]
		if !exist {
			return false
		}
		if val == "" {
			return false
		}
	}

	return true
}
