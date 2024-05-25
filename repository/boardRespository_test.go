package repository

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"lab3/mocks"
	"lab3/model"
	"testing"
)

func TestBoardRepository_CreateBoard(t *testing.T) {
	t.Run("Creates board successfully", func(t *testing.T) {
		db := mocks.NewIDB(t)
		testInstance := BoardRepository{DB: db}

		boardArray := [][]int{
			{model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK},
			{model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE},
			{model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.WHITE},
			{model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.BLACK},
			{model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE},
			{model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.WHITE, model.BLACK},
			{model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.BLACK},
			{model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.WHITE, model.BLACK, model.BLACK, model.BLACK},
			{model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE},
			{model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK}}
		boardDto := model.BoardDto{Board: boardArray}
		boardData, _ := json.Marshal(boardArray)
		boardEntity := model.Board{Data: string(boardData)}

		db.On("CreateBoard", boardEntity).Return(nil)

		actualErr := testInstance.CreateBoard(boardDto)

		assert.Nil(t, actualErr)
	})
	t.Run("error on create board", func(t *testing.T) {
		db := mocks.NewIDB(t)
		testInstance := BoardRepository{DB: db}

		boardArray := [][]int{
			{model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK},
			{model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE},
			{model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.WHITE},
			{model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.BLACK},
			{model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE},
			{model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.WHITE, model.BLACK},
			{model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.BLACK},
			{model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.WHITE, model.BLACK, model.BLACK, model.BLACK},
			{model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE},
			{model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK}}
		boardDto := model.BoardDto{Board: boardArray}
		boardData, _ := json.Marshal(boardArray)
		boardEntity := model.Board{Data: string(boardData)}

		db.On("CreateBoard", boardEntity).Return(assert.AnError)
		actualErr := testInstance.CreateBoard(boardDto)

		assert.Error(t, actualErr)
	})
}

func TestBoardRepository_GetAllBoards(t *testing.T) {
	t.Run("Get all boards successfully", func(t *testing.T) {
		db := mocks.NewIDB(t)
		testInstance := BoardRepository{DB: db}

		boardArray := [][]int{
			{model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK},
			{model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE},
			{model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.WHITE},
			{model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.BLACK},
			{model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE},
			{model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.WHITE, model.BLACK},
			{model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.BLACK},
			{model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.WHITE, model.BLACK, model.BLACK, model.BLACK},
			{model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE},
			{model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK}}
		boardDto := model.BoardDto{Board: boardArray}
		boardData, _ := json.Marshal(boardArray)
		boardEntity := model.Board{Data: string(boardData)}
		boardEntities := []model.Board{boardEntity}

		db.On("FindAllBoards").Return(boardEntities, nil)

		actual, actualErr := testInstance.GetAllBoards()

		assert.Nil(t, actualErr)
		assert.Equal(t, []model.BoardDto{boardDto}, actual)
	})
	t.Run("error when get boards from DB", func(t *testing.T) {
		db := mocks.NewIDB(t)
		testInstance := BoardRepository{DB: db}

		boardArray := [][]int{
			{model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK},
			{model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE},
			{model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.WHITE},
			{model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.BLACK},
			{model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE},
			{model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.WHITE, model.BLACK},
			{model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.BLACK},
			{model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.WHITE, model.BLACK, model.BLACK, model.BLACK},
			{model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE},
			{model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK}}
		boardData, _ := json.Marshal(boardArray)
		boardEntity := model.Board{Data: string(boardData)}
		boardEntities := []model.Board{boardEntity}

		db.On("FindAllBoards").Return(boardEntities, assert.AnError)

		actual, actualErr := testInstance.GetAllBoards()

		assert.Error(t, actualErr)
		assert.Equal(t, []model.BoardDto{}, actual)
	})
}

func TestBoardRepository_GetBoardById(t *testing.T) {
	t.Run("Get board by id successfully", func(t *testing.T) {
		db := mocks.NewIDB(t)
		testInstance := BoardRepository{DB: db}

		boardArray := [][]int{
			{model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK},
			{model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE},
			{model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.WHITE},
			{model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.BLACK},
			{model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE},
			{model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.WHITE, model.BLACK},
			{model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.BLACK},
			{model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.WHITE, model.BLACK, model.BLACK, model.BLACK},
			{model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE},
			{model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK}}
		boardDto := model.BoardDto{Board: boardArray, ID: 1}
		boardData, _ := json.Marshal(boardArray)
		boardEntity := model.Board{Data: string(boardData), ID: 1}

		db.On("FindBoardById", "1").Return(boardEntity, nil)

		actual, actualErr := testInstance.GetBoardById("1")

		assert.Nil(t, actualErr)
		assert.Equal(t, boardDto, actual)
	})
	t.Run("error when get board from DB", func(t *testing.T) {
		db := mocks.NewIDB(t)
		testInstance := BoardRepository{DB: db}

		boardArray := [][]int{
			{model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK},
			{model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE},
			{model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.WHITE},
			{model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.BLACK},
			{model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE},
			{model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.WHITE, model.BLACK},
			{model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.BLACK},
			{model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.WHITE, model.BLACK, model.BLACK, model.BLACK},
			{model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE},
			{model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK}}
		boardData, _ := json.Marshal(boardArray)
		boardEntity := model.Board{Data: string(boardData), ID: 1}

		db.On("FindBoardById", "1").Return(boardEntity, assert.AnError)

		actual, actualErr := testInstance.GetBoardById("1")

		assert.Error(t, actualErr)
		assert.Equal(t, model.BoardDto{}, actual)
	})
}
