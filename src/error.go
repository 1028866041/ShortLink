package src

type Error interface{
    error
    Status() int
}

type StatusError struct{
    code int
    err error
}

func (se StatusError) Error() string{
    return se.err.Error()
}

func (se StatusError) Status() int{
    return se.code
}