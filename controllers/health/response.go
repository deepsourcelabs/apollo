package health

type Response struct {
	Results []Result `json:"results"`
}

type Result struct {
	Name  string `json:"name"`
	State string `json:"state"`
	Err   string `json:"err"`
}
