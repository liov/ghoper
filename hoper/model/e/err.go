package e

type ValidtionError struct {
	Msg string
}

func (e *ValidtionError) Error() string {
	return e.Msg
}
