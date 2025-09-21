package slice

func Filter[T any](items []T, filter func(T) bool) []T {
	filtered := make([]T, 0, len(items))
	for _, item := range items {
		if filter(item) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

func Map[T, R any](items []T, mapper func(it T) R) []R {
	mapped := make([]R, len(items))
	for i, item := range items {
		mapped[i] = mapper(item)
	}
	return mapped
}

func Reduce[T, R any](items []T, init R, reducer func(R, T) R) R {
	result := init
	for _, item := range items {
		result = reducer(result, item)
	}
	return result
}
