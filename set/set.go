package set

type Set[T comparable] = map[T]struct{}

func New[T comparable](items ...T) Set[T] {
	set := make(Set[T], len(items))
	for _, item := range items {
		set[item] = struct{}{}
	}
	return set
}

func Add[T comparable](set Set[T], item T) {
	set[item] = struct{}{}
}

func AddAll[T comparable](set Set[T], items ...T) {
	for idx := range items {
		set[items[idx]] = struct{}{}
	}
}

func Contains[T comparable](set Set[T], item T) bool {
	_, ok := set[item]
	return ok
}

func ContainsOne[T comparable](set Set[T], items ...T) bool {
	for idx := range items {
		if _, ok := set[items[idx]]; ok {
			return true
		}
	}
	return false
}

func ContainsAll[T comparable](set Set[T], items ...T) bool {
	for idx := range items {
		if _, ok := set[items[idx]]; !ok {
			return false
		}
	}
	return true
}

func Slice[T comparable](set Set[T]) []T {
	slice := make([]T, 0, len(set))
	for item := range set {
		slice = append(slice, item)
	}
	return slice
}

func Intersect[T comparable](dst, src Set[T]) {
	if len(dst) > len(src) {
		for k := range dst {
			if _, ok := src[k]; !ok {
				delete(dst, k)
			}
		}
	} else {
		for k := range dst {
			if _, ok := src[k]; !ok {
				delete(dst, k)
			}
		}
	}
}

func Intersection[T comparable](a, b Set[T]) Set[T] {
	if len(a) > len(b) {
		a, b = b, a
	}
	result := make(Set[T], len(a))
	for k := range a {
		if _, ok := b[k]; ok {
			result[k] = struct{}{}
		}
	}
	return result
}
