package infr

type FromType[F any, T any] interface {
	From(source F) T
}

type IntoType[T any] interface {
	Into() T
}

type FI[F any, T FromType[F, T]] struct {
	Source F
	ft     T
}

func (c FI[F, T]) Into() T {
	return c.ft.From(c.Source)
}

func Into[T FromType[F, T], F any](s F) T {
	var t T
	return t.From(s)
}

type TryFromType[F, T any] interface {
	TryFrom(source F) (T, error)
}

type TryIntoType[T any] interface {
	TryInto() (T, error)
}

type TFI[F any, T TryFromType[F, T]] struct {
	Source F
	tft    T
}

func (ct TFI[F, T]) TryInto() (T, error) {
	return ct.tft.TryFrom(ct.Source)
}

func TryInto[T TryFromType[F, T], F any](s F) (T, error) {
	var t T
	return t.TryFrom(s)
}
