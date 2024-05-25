package service

import (
	"github.com/stretchr/testify/mock"
	"lab3/mocks"
	"lab3/model"
	"testing"
)

func TestBoardGenerator_GenerateBoards(t *testing.T) {
	repository := mocks.NewIBoardRepository(t)
	calculator := mocks.NewICalculator(t)
	testInstance := BoardGenerator{
		Calculator:      calculator,
		BoardRepository: repository,
	}

	expected := [][]int{
		{model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK},
		{model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE},
		{model.WHITE, model.WHITE, model.VISITED, model.VISITED, model.VISITED, model.VISITED, model.WHITE, model.BLACK, model.WHITE, model.WHITE},
		{model.WHITE, model.BLACK, model.VISITED, model.BLACK, model.BLACK, model.VISITED, model.BLACK, model.BLACK, model.WHITE, model.BLACK},
		{model.VISITED, model.VISITED, model.VISITED, model.WHITE, model.BLACK, model.VISITED, model.VISITED, model.VISITED, model.VISITED, model.WHITE},
		{model.VISITED, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.VISITED, model.BLACK},
		{model.VISITED, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.VISITED, model.VISITED, model.VISITED, model.VISITED, model.WHITE},
		{model.VISITED, model.VISITED, model.VISITED, model.VISITED, model.WHITE, model.VISITED, model.WHITE, model.BLACK, model.WHITE, model.BLACK},
		{model.WHITE, model.BLACK, model.BLACK, model.VISITED, model.BLACK, model.VISITED, model.BLACK, model.WHITE, model.BLACK, model.WHITE},
		{model.BLACK, model.WHITE, model.BLACK, model.VISITED, model.VISITED, model.VISITED, model.BLACK, model.WHITE, model.WHITE, model.BLACK},
	}

	calculator.On("FindLoop", mock.Anything).Return(expected)
	repository.On("CreateBoard", mock.Anything).Return(nil)
	testInstance.GenerateBoards()
}

func TestBoardGenerator_saveResult(t *testing.T) {
	t.Run("Save new board successfully", func(t *testing.T) {
		repository := mocks.NewIBoardRepository(t)
		calculator := mocks.NewICalculator(t)
		testInstance := BoardGenerator{
			Calculator:      calculator,
			BoardRepository: repository,
		}
		generatedBoards := []model.BoardDto{}
		boardArray := [][]int{
			{model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK},
			{model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE},
			{model.WHITE, model.WHITE, model.VISITED, model.VISITED, model.VISITED, model.VISITED, model.WHITE, model.BLACK, model.WHITE, model.WHITE},
			{model.WHITE, model.BLACK, model.VISITED, model.BLACK, model.BLACK, model.VISITED, model.BLACK, model.BLACK, model.WHITE, model.BLACK},
			{model.VISITED, model.VISITED, model.VISITED, model.WHITE, model.BLACK, model.VISITED, model.VISITED, model.VISITED, model.VISITED, model.WHITE},
			{model.VISITED, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.VISITED, model.BLACK},
			{model.VISITED, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.VISITED, model.VISITED, model.VISITED, model.VISITED, model.WHITE},
			{model.VISITED, model.VISITED, model.VISITED, model.VISITED, model.WHITE, model.VISITED, model.WHITE, model.BLACK, model.WHITE, model.BLACK},
			{model.WHITE, model.BLACK, model.BLACK, model.VISITED, model.BLACK, model.VISITED, model.BLACK, model.WHITE, model.BLACK, model.WHITE},
			{model.BLACK, model.WHITE, model.BLACK, model.VISITED, model.VISITED, model.VISITED, model.BLACK, model.WHITE, model.WHITE, model.BLACK},
		}
		boardDto := model.BoardDto{Board: boardArray}
		repository.On("CreateBoard", boardDto).Return(nil)

		testInstance.saveResult(boardArray, generatedBoards)
	})

	t.Run("No new board, since already saved", func(t *testing.T) {
		repository := mocks.NewIBoardRepository(t)
		calculator := mocks.NewICalculator(t)
		testInstance := BoardGenerator{
			Calculator:      calculator,
			BoardRepository: repository,
		}

		boardArray := [][]int{
			{model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK},
			{model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE},
			{model.WHITE, model.WHITE, model.VISITED, model.VISITED, model.VISITED, model.VISITED, model.WHITE, model.BLACK, model.WHITE, model.WHITE},
			{model.WHITE, model.BLACK, model.VISITED, model.BLACK, model.BLACK, model.VISITED, model.BLACK, model.BLACK, model.WHITE, model.BLACK},
			{model.VISITED, model.VISITED, model.VISITED, model.WHITE, model.BLACK, model.VISITED, model.VISITED, model.VISITED, model.VISITED, model.WHITE},
			{model.VISITED, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.VISITED, model.BLACK},
			{model.VISITED, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.VISITED, model.VISITED, model.VISITED, model.VISITED, model.WHITE},
			{model.VISITED, model.VISITED, model.VISITED, model.VISITED, model.WHITE, model.VISITED, model.WHITE, model.BLACK, model.WHITE, model.BLACK},
			{model.WHITE, model.BLACK, model.BLACK, model.VISITED, model.BLACK, model.VISITED, model.BLACK, model.WHITE, model.BLACK, model.WHITE},
			{model.BLACK, model.WHITE, model.BLACK, model.VISITED, model.VISITED, model.VISITED, model.BLACK, model.WHITE, model.WHITE, model.BLACK},
		}
		boardDto := model.BoardDto{Board: boardArray}
		generatedBoards := []model.BoardDto{boardDto}

		repository.On("CreateBoard", boardDto).Return(nil).Times(0)

		testInstance.saveResult(boardArray, generatedBoards)
	})
}
