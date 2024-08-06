package helper


func CheckNline(s []string) bool {
	for _, w := range s {
		if w != "" {
			return true
		}
	}
	return false
}