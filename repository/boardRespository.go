package repository

import (
	"encoding/json"
	"lab3/database"
	"lab3/model"
)

type (
	IBoardRepository interface {
		CreateBoard(board model.BoardDto) error
		GetBoardById(id string) (model.BoardDto, error)
		GetAllBoards() ([]model.BoardDto, error)
	}
	BoardRepository struct {
		DB database.IDB
	}
)

func (br BoardRepository) CreateBoard(board model.BoardDto) error {
	boardEntity := model.Board{}
	out, _ := json.Marshal(board.Board)
	boardEntity.Data = string(out)
	createErr := br.DB.CreateBoard(boardEntity)
	if createErr != nil {
		return createErr
	}

	return nil
}

func (br BoardRepository) GetBoardById(id string) (model.BoardDto, error) {
	entity, err := br.DB.FindBoardById(id)
	if err != nil {
		return model.BoardDto{}, err
	}
	dataBytes := []byte(entity.Data)
	var boardArray [][]int
	json.Unmarshal(dataBytes, &boardArray)
	boardDto := model.BoardDto{Board: boardArray, ID: int(entity.ID)}
	return boardDto, nil
}

func (br BoardRepository) GetAllBoards() ([]model.BoardDto, error) {
	entities, err := br.DB.FindAllBoards()
	if err != nil {
		return []model.BoardDto{}, err
	}
	boardDtos := []model.BoardDto{}
	for _, boardEntity := range entities {
		dataBytes := []byte(boardEntity.Data)
		var boardArray [][]int
		json.Unmarshal(dataBytes, &boardArray)
		boardDtos = append(boardDtos, model.BoardDto{Board: boardArray, ID: int(boardEntity.ID)})
	}
	return boardDtos, nil
}
