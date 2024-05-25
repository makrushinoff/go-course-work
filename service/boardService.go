package service

import (
	"errors"
	"lab3/model"
	"lab3/repository"
	"log"
)

type (
	IBoardService interface {
		MakeCalculations(boardDto model.BoardDto) (model.BoardWrapper, error)
		GetAllBoards() ([]model.BoardDto, error)
		GetBoardById(id string) (model.BoardDto, error)
	}
	BoardService struct {
		BoardRepository repository.IBoardRepository
		Calculator      ICalculator
	}
)

func (bs BoardService) MakeCalculations(boardDto model.BoardDto) (model.BoardWrapper, error) {
	log.Println("Received request to find a solution circle loop")
	calculator := bs.Calculator
	boardArray := boardDto.Board
	solution := calculator.FindLoop(boardArray)

	if solution != nil {
		log.Println("Successfully found a solution circle loop")
		return model.BoardWrapper{SolvedBoard: solution}, nil
	}
	log.Println("Error when looking for circle loop solution on provided board")
	return model.BoardWrapper{SolvedBoard: nil}, errors.New("could not find circular loop")
}

func (bs BoardService) GetAllBoards() ([]model.BoardDto, error) {
	boards, err := bs.BoardRepository.GetAllBoards()
	if err != nil {
		return []model.BoardDto{}, err
	}
	return boards, nil
}

func (bs BoardService) GetBoardById(id string) (model.BoardDto, error) {
	board, err := bs.BoardRepository.GetBoardById(id)
	if err != nil {
		return model.BoardDto{}, err
	}
	return board, nil
}
