package utils

import (
	"cmp"
	"fmt"
)

func Wrap(err *error, msg string) {
	if err != nil {
		*err = fmt.Errorf("%s: %w", msg, *err)
	}
}

func Coalesce[T any](ptrs ...*T) *T {
	for _, p := range ptrs {
		if p != nil {
			return p
		}
	}
	return nil
}

func Deref[T any](ptr *T, fallback T) T {
	if ptr != nil {
		return *ptr
	}
	return fallback
}

func Clamp[T cmp.Ordered](val, lo, hi T) T {
	return max(lo, min(val, hi))
}
