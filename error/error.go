package error

type responseError struct {
	originalErr error
	contextInfo errorContext
}

type errorContext struct {
	Field   string
	Message string
}

// Error returns the message of a responseError
func (e responseError) Error() string {
	return e.originalErr.Error()
}

// //New register the new error
// func New(err error, msg string) *responseError {
// 	return responseError{
// 		originalError: err,
// 		contextInfo: errorContext{
// 			Message: msg
// 		}
// 	}
// }
