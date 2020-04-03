package todo

type insertRequest struct {
	Content string `json:"content"`
}

type response struct {
	Success bool        `json:"success"`
	Payload interface{} `json:"payload"`
}
