package board

type BoardSettings struct {
	Height int `json:"height"`
	Width  int `json:"width"`
	Bombs  int `json:"bombs"`
}
