package common

import "strconv"

func BinaryToDecimal(s []string) ([]int, error) {

	var a []int

	for _, v := range s {
		if len(v) > 0 {
			i, err := strconv.ParseInt(v, 2, 64)

			if err != nil {
				return a, err
			}

			a = append(a, int(i))
		}
	}

	return a, nil

}

// ToInt converts a slice of strings
// int a slice of Ints
func ToInt(s []string) ([]int, error) {

	var a []int

	for _, n := range s {

		if len(n) > 0 {
			i, err := strconv.Atoi(n)

			if err != nil {
				return a, err
			}

			a = append(a, i)
		}

	}

	return a, nil
}
