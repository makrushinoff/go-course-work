package service

import (
	"lab3/model"
	"lab3/repository"
	"log"
	"math/rand"
)

type (
	IBoardGenerator interface {
		GenerateBoards()
	}
	BoardGenerator struct {
		Calculator      ICalculator
		BoardRepository repository.IBoardRepository
	}
)

func (bg *BoardGenerator) createRandomBoard(rows, cols int) [][]int {
	board := make([][]int, rows)
	for i := range board {
		board[i] = make([]int, cols)
	}

	for i := range board {
		rows := board[i]
		for j := range rows {
			randomNumber := rand.Intn(2)
			if randomNumber == 0 {
				board[i][j] = 0
			} else {
				board[i][j] = 1
			}
		}
	}

	return board
}

func (bg *BoardGenerator) GenerateBoards() {
	boardSize := model.SIZE
	generatedBoards := []model.BoardDto{}
	for i := 0; i < 2000; i++ {
		var boardArray [][]int
		needMoreTry := true
		calculator := bg.Calculator
		for needMoreTry {
			boardArray = bg.createRandomBoard(boardSize, boardSize)
			solution := calculator.FindLoop(boardArray)
			if solution != nil {
				needMoreTry = false
				break
			}
		}
		bg.saveResult(boardArray, generatedBoards)
	}

	log.Println("Workable boards are successfully generated")
}

func (bg *BoardGenerator) saveResult(boardArray [][]int, generatedBoards []model.BoardDto) {
	result := model.BoardDto{Board: boardArray}
	if !checkBoardAlreadyGenerated(generatedBoards, result) {
		generatedBoards = append(generatedBoards, result)
		err := bg.BoardRepository.CreateBoard(result)
		if err != nil {
			return
		}
	}
}

func checkBoardAlreadyGenerated(generatedBoards []model.BoardDto, boardToAdd model.BoardDto) bool {
	for _, generatedBoard := range generatedBoards {
		for i := 0; i < model.SIZE; i++ {
			for j := 0; j < model.SIZE; j++ {
				if generatedBoard.Board[i][j] != boardToAdd.Board[i][j] {
					return true
				}
			}
		}
	}
	return false
}
