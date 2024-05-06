package types

type AsyncResponse[T any] struct {
	Data  T
	Error error
}
