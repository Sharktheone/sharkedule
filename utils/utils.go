package utils

import "slices"

func SliceHaveCommon[A comparable](a, b []A) bool {
	for _, v := range a {
		if slices.Contains(b, v) {
			return true
		}
	}
	return false
}
