package array_ops

func Contains[T comparable](val T, arr *[]T) bool {
	for _, rangeVal := range *arr {
		if rangeVal == val {
			return true
		}
	}
	return false
}

func Combine[T any](arr1 *[]T, arr2 *[]T) *[]T {
	combArr := make([]T, len(*arr1)+len(*arr2))
	len1 := len(*arr1)
	for i := range combArr {
		if i < len(*arr1) {
			combArr[i] = (*arr1)[i]
		} else {
			combArr[i] = (*arr2)[i-len1]
		}
	}
	return &combArr
}
