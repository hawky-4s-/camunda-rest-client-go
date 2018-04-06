package camunda

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

type ExternalTaskClientError struct {
	err          error
	statusCode   int
	errorMessage string
}

func (e *ExternalTaskClientError) Error() string {
	return e.errorMessage
}

func newExternalTaskClientError(errorMessage string, statusCode int, err error) ExternalTaskClientError {
	return ExternalTaskClientError{
		errorMessage: errorMessage,
		statusCode:   statusCode,
		err:          err,
	}
}

// ExternalTaskNotFoundError is used when the Camunda REST API returns with 404
// TODO: incorporate TaskId or something from query into it
var ExternalTaskNotFoundError = newExternalTaskClientError("External task with id '%s' could not be found.", 404, nil)
var ExternalTaskNotAcquiredError = newExternalTaskClientError("", 0, nil)
var ExternalTaskConnectionLostError = newExternalTaskClientError("", 0, nil)
