package blueprint

import info "cellula/info"

type Blueprint struct {
}

func (b Blueprint) Code() info.Result {
	return "bp"
}
