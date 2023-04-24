package array_ops

func Map[F any, T any](arr *[]F, trans func(x F) T) *[]T {
	res := make([]T, len(*arr))
	for i, elem := range *arr {
		res[i] = trans(elem)
	}
	return &res
}
