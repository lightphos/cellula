package info

type Result string

type Coder interface {
	Code() Result
}

func Invoke(c Coder) Result {
	return c.Code()
}
