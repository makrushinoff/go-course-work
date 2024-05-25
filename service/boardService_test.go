package service_test

import (
	"github.com/stretchr/testify/assert"
	"lab3/mocks"
	"lab3/model"
	"lab3/service"
	"testing"
)

var testInstance service.BoardService

func TestService_GetAllBoards(t *testing.T) {
	t.Run("Returns All Boards", func(t *testing.T) {
		calculator := mocks.NewICalculator(t)
		repository := mocks.NewIBoardRepository(t)
		testInstance = service.BoardService{
			BoardRepository: repository,
			Calculator:      calculator,
		}
		repository.On("GetAllBoards").Return([]model.BoardDto{}, nil)

		actual, _ := testInstance.GetAllBoards()

		assert.Equal(t, len(actual), 0)
	})

	t.Run("Error when return all boards", func(t *testing.T) {
		calculator := mocks.NewICalculator(t)
		repository := mocks.NewIBoardRepository(t)
		testInstance = service.BoardService{
			BoardRepository: repository,
			Calculator:      calculator,
		}
		repository.On("GetAllBoards").Return(nil, assert.AnError)

		actual, err := testInstance.GetAllBoards()

		assert.Equal(t, []model.BoardDto{}, actual)
		assert.Error(t, err)
	})
}

func TestService_GetBoardById(t *testing.T) {
	t.Run("Returns board by id", func(t *testing.T) {
		calculator := mocks.NewICalculator(t)
		repository := mocks.NewIBoardRepository(t)
		testInstance = service.BoardService{
			BoardRepository: repository,
			Calculator:      calculator,
		}

		id := "id"
		repository.On("GetBoardById", id).Return(model.BoardDto{}, nil)

		actual, _ := testInstance.GetBoardById(id)

		assert.Equal(t, model.BoardDto{}, actual)
	})

	t.Run("Error when return board by id", func(t *testing.T) {
		calculator := mocks.NewICalculator(t)
		repository := mocks.NewIBoardRepository(t)
		testInstance = service.BoardService{
			BoardRepository: repository,
			Calculator:      calculator,
		}
		id := "id"
		repository.On("GetBoardById", id).Return(model.BoardDto{}, assert.AnError)

		actual, err := testInstance.GetBoardById(id)
		assert.Equal(t, model.BoardDto{}, actual)
		assert.Error(t, err)
	})
}

func TestService_MakeCalculations(t *testing.T) {
	t.Run("Makes calculation and returns correct result", func(t *testing.T) {
		calculator := mocks.NewICalculator(t)
		repository := mocks.NewIBoardRepository(t)
		testInstance = service.BoardService{
			BoardRepository: repository,
			Calculator:      calculator,
		}
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
		boardDto := model.BoardDto{Board: boardArray}

		calculator.On("FindLoop", boardArray).Return(expected)

		actual, _ := testInstance.MakeCalculations(boardDto)

		assert.Equal(t, model.BoardWrapper{SolvedBoard: expected}, actual)
	})

	t.Run("Error when did not find loop", func(t *testing.T) {
		calculator := mocks.NewICalculator(t)
		repository := mocks.NewIBoardRepository(t)
		testInstance = service.BoardService{
			BoardRepository: repository,
			Calculator:      calculator,
		}
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

		calculator.On("FindLoop", boardArray).Return(nil)

		actual, err := testInstance.MakeCalculations(boardDto)
		assert.Equal(t, model.BoardWrapper{SolvedBoard: nil}, actual)
		assert.Error(t, err)
	})
}
