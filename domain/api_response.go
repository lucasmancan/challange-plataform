package domain

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
