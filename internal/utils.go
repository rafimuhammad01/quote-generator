package internal

import "strings"

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func RemoveInSliceString(arrStr []string, selector string) []string {
	var newArrStr []string
	for _, v := range arrStr {
		trimmedString := strings.TrimSpace(v)
		if trimmedString != selector {
			newArrStr = append(newArrStr, trimmedString)
		}
	}

	return newArrStr
}
