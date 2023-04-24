package array_ops

func Filter[T any](arr *[]T, trans func(x T) bool) *[]T {
	res := make([]T, 0, len(*arr))
	for _, elem := range *arr {
		if trans(elem) {
			res = append(res, elem)
		}
	}
	return &res
}
