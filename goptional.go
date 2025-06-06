package goptional

// Maybe is a type that can be either a Just or a Nothing.
type Maybe[T any] interface {
	// IsPresent returns true if the Maybe is a Just.
	IsPresent() bool
	//Get returns the value of the Maybe if it is a Just. If the Maybe is a Nothing, Get panics.
	Get() T
}

//----

// Just is a Maybe that contains a value.
type Just[T any] struct {
	value T
}

// NewJust returns a Just containing the given value.
func NewJust[T any](value T) Maybe[T] {
	return Just[T]{value}
}

// IsPresent returns true if the Maybe is a Just.
func (j Just[T]) IsPresent() bool {
	return true
}

// Get returns the value of the Just.
func (j Just[T]) Get() T {
	return j.value
}

//----

// Nothing is a Maybe that contains no value.
type Nothing[T any] struct {
}

// NewNothing returns a Nothing.
func NewNothing[T any]() Maybe[T] {
	return Nothing[T]{}
}

// IsPresent returns false if the Maybe is a Nothing.
func (n Nothing[T]) IsPresent() bool {
	return false
}

// Get returns the value of the Maybe if it is a Just. If the Maybe is a Nothing, Get panics.
func (n Nothing[T]) Get() T {
	panic("Get() called on Nothing")
}

//----
