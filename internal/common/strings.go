package common

func Contains(f string, s []string) bool {

	for _, v := range s {
		if f == v {
			return true
		}
	}

	return false
}
