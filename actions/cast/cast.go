package cast

func To[T any](x interface{}) T {
	v, _ := x.(T)
	return v
}

func Must[T any](x interface{}) T {
	v, ok := x.(T)
	if !ok {
		panic("unexpected type")
	}
	return v
}

func As[T any](x interface{}) (T, bool) {
	v, ok := x.(T)
	return v, ok
}
