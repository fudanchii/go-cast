package infr

/// From/Into variants to represent a type transition tha is infallible

// FromType[F, T] represents a type that can transform
// a value from a given variable typed `F`.
type FromType[F, T any] interface {
	From(source F) T
}

// IntoType represents a type that implements `Into`
type IntoType[T any] interface {
	Into() T
}

// FI[F, T] represents a container that can transform a type into the target type,
// as long as the target type implements `FromType[F, T]`
type FI[F any, T FromType[F, T]] struct {
	Source F
	ft     T
}

// FI implements IntoType.
func (c FI[F, T]) Into() T {
	return c.ft.From(c.Source)
}

// Into is a shortcut for transforming F to T.
func Into[T FromType[F, T], F any](s F) T {
	var t T
	return t.From(s)
}

// From is a shortcut to create FI instance.
func From[F any, T FromType[F, T]](s F) FI[F, T] {
	return FI[F, T]{Source: s}
}

/// TryFrom/TryInto variants to represent a type transition that _is_ fallible

// TryFromType[F, T] represents a type that can transform
// a value from a given variable typed `F`,
// but also can return error if the transform fail.
type TryFromType[F, T any] interface {
	TryFrom(source F) (T, error)
}

// TryIntoType[T] represent a type that implements `TryInto`.
type TryIntoType[T any] interface {
	TryInto() (T, error)
}

// TFI[F, T] represents a container that can transform a type into the target type,
// as long as the target type implements `TryFromType[F, T]`.
type TFI[F any, T TryFromType[F, T]] struct {
	Source F
	tft    T
}

// TFI implements TryIntoType.
func (ct TFI[F, T]) TryInto() (T, error) {
	return ct.tft.TryFrom(ct.Source)
}

// TryInto is a shortcut for transforming F to T.
func TryInto[T TryFromType[F, T], F any](s F) (T, error) {
	var t T
	return t.TryFrom(s)
}

// From is a shortcut to create TFI instance.
func TryFrom[F any, T TryFromType[F, T]](s F) TFI[F, T] {
	return TFI[F, T]{Source: s}
}

func IntoSliceOf[T FromType[F, T], F any](s []F) []T {
	var sliceT []T

	for _, elt := range s {
		newT := Into[T](elt)
		sliceT = append(sliceT, newT)
	}

	return sliceT
}
