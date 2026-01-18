package server

type chatRequest struct {
	Title string `json:"title"`
}

type messageRequest struct {
	Text string `json:"text"`
}

type chatResponse struct {
	Chat     interface{} `json:"chat"`
	Messages interface{} `json:"messages"`
}
