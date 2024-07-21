package internal

type OptionalVal[TVal any] struct {
	Value    TVal
	Assigned bool
}
