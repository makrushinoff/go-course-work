package model

type BoardDto struct {
	ID    int
	Board [][]int
}

type Board struct {
	ID   uint64 `gorm:"primary key;autoIncrement" json:"id"`
	Data string `json:"data"`
}

type BoardWrapper struct {
	SolvedBoard [][]int
}
