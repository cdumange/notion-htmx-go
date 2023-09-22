package common

// Error a custom error type.
type Error string

// Error returns the error message.
func (e Error) Error() string { return string(e) }
