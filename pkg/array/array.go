package array

func StringInArray(input string, values []string) bool {
	for _, currVal := range values {
		if input == currVal {
			return true
		}
	}
	return false
}
