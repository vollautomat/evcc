package util

// Addr returns the address of the given value.
func Addr[T any](v T) *T {
	return &v
}
