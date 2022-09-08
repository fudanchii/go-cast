package cast

type FromType[F any, T any] interface {
	From(source F) T
}

func Into[F FromType[T, F], T any](source T) F {
	var target F
	return target.From(source)
}
