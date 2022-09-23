package slice

func GetLatestElement[T comparable](arr []T) T {
	if len(arr) > 0 {
		return arr[len(arr)-1]
	} else {
		panic("error: len of slice must be more than 0")
	}
}
