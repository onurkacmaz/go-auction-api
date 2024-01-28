package utils

func InArray(key string, list []string) bool {
	for _, i := range list {
		if i == key {
			return true
		}
	}
	return false
}
