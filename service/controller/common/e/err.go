package e

type ValidtionError struct {
	Msg string
}

func(e *ValidtionError) Error() string{
	return e.Msg
}

type DBError struct {
	Msg string
}

func (e *DBError) Error() string {
	return e.Msg
}

type LimitError struct {
	Msg string
}

func (e *LimitError) Error() string {
	return e.Msg
}