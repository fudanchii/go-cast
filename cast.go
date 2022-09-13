package cast

type FromType[F any, T any] interface {
	From(source F) T
}

type C[F any, T FromType[F, T]] struct {
	Source F
	ft     T
}

func (c C[F, T]) Into() T {
	return c.ft.From(c.Source)
}

func Into[T FromType[F, T], F any](s F) T {
	var t T
	return t.From(s)
}

type TryFromType[F, T any] interface {
	TryFrom(source F) (T, error)
}

type CT[F any, T TryFromType[F, T]] struct {
	Source F
	tft    T
}

func (ct CT[F, T]) TryInto() (T, error) {
	return ct.tft.TryFrom(ct.Source)
}

func TryInto[T TryFromType[F, T], F any](s F) (T, error) {
	var t T
	return t.TryFrom(s)
}
