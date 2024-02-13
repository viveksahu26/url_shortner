package errors

type UrlShortnerError struct {
	Message string
	Code    int
}

func Error(cosignError UrlShortnerError) error {
	return &UrlShortnerError{
		Message: cosignError.Message,
		Code:    cosignError.Code,
	}
}

// Assert that we implement error at build time.
var _ error = (*UrlShortnerError)(nil)

// Error implements error
func (ce *UrlShortnerError) Error() string {
	return ce.Message
}

func (ce *UrlShortnerError) ExitCode() int {
	return ce.Code
}
