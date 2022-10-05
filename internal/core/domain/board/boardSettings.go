package board

type BoardSettings struct {
	Height int `json:"height"`
	Width  int `json:"width"`
	Bombs  int `json:"bombs"`
}

func NewBoardSettings(height int, width int, bombs int) BoardSettings {
	return BoardSettings{Height: height, Width: width, Bombs: bombs}
}
