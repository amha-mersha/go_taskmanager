package data

type TaskError struct {
	message string
}

const (
	IDNotFound       = "No item found with the specified ID."
	MalformedJSON    = "Sent a malfomed JSON."
	MismatchedFormat = "The task have a mismatched stucture."
	MissingRequireds = "There are some missing required feilds."
)

func (err TaskError) Error() string {
	return err.message
}
