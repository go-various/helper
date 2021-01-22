package errutil

// InternalError represents an error generated internally,
// presumably not due to invalid user input
type InternalError struct {
	Err string
}

func (e InternalError) Error() string {
	return e.Err
}
type ExternalError struct {
	Err string
}

func (e ExternalError) Error() string {
	return e.Err
}