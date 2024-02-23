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

func AppendIfMissing[T comparable](slice []*T, s *T) []*T {
	for _, ele := range slice {
		if *ele == *s {
			return slice
		}
	}
	return append(slice, s)
}

//func AppendMultipleIfMissing(slice []*string, s []*string) []*string {
//	for _, ele := range s {
//		slice = AppendIfMissing(slice, ele)
//	}
//	return slice
//}

func AppendSliceIfMissing[T comparable](slice []*T, s ...T) []*T {
	for _, ele := range s {
		slice = AppendIfMissingPtr(slice, ele)
	}
	return slice
}

func AppendIfMissingPtr[T comparable](slice []*T, s T) []*T {
	for _, ele := range slice {
		if *ele == s {
			return slice
		}
	}
	return append(slice, &s)
}
