package board

type BoardSettings struct {
	Height int   `json:"height"`
	Width  int   `json:"width"`
	Bombs  int   `json:"bombs"`
	Seed   int64 `json:"seed"`
}

func NewBoardSettings(height int, width int, bombs int, seed int64) BoardSettings {
	return BoardSettings{Height: height, Width: width, Bombs: bombs, Seed: seed}
}
