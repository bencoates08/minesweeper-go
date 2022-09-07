package gameHandler

type CreateRequest struct {
	Name   string `json:"name"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
	Bombs  int    `json:"bombs"`
}
