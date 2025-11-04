package results

type Result struct {
	LastOneFlag bool `json:"lastOne"`
}

func (r Result) LastOne() bool {
	return r.LastOneFlag
}
