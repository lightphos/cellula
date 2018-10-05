package info

type Hello struct {
	text string
}

func (h Hello) Code() Result {
	return "-->hello"
}
