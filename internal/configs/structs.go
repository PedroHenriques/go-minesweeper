package configs

type SizeOption struct {
	NumMines int
	NumRows  int
	NumCols  int
}

type Configs struct {
	SizeOptions map[string]SizeOption `json:"SizeOptions"`
}
