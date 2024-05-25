package service

import (
	. "lab3/model"
	"math"
)

func firstStep(board [][]int, pos Position) Position {
	// move right
	if SIZE > pos.J+2 {
		if board[pos.I][pos.J+1] != BLACK {
			return NewPosition(pos.I, pos.J+1, board[pos.I][pos.J+1])
		}
	}
	// move left
	if 0 <= pos.J-2 {
		if board[pos.I][pos.J-1] != BLACK {
			return NewPosition(pos.I, pos.J-1, board[pos.I][pos.J-1])
		}
	}
	// move up
	if 0 <= pos.I-2 {
		if board[pos.I-1][pos.J] != BLACK {
			return NewPosition(pos.I-1, pos.J, board[pos.I-1][pos.J])
		}
	}
	// move down
	if SIZE > pos.I+2 {
		if board[pos.I+1][pos.J] != BLACK {
			return NewPosition(pos.I+1, pos.J, board[pos.I+1][pos.J])
		}
	}

	return NewPosition(N, N, N)
}

func contains(vector []Position, pos Position) bool {
	for _, p := range vector {
		if p.Equals(pos) {
			return true
		}
	}
	return false
}

func findNextPositions(board [][]int, pos Position, prevPos Position, checkedPositions []Position, startPos Position) []Position {
	var positions []Position
	if prevPos.Color == BLACK {
		if pos.I-prevPos.I == 1 && pos.I < SIZE {
			positions = append(positions, NewPosition(pos.I+1, pos.J, board[pos.I+1][pos.J]))
		}
		if pos.I-prevPos.I == -1 && pos.I-1 >= 0 {
			positions = append(positions, NewPosition(pos.I-1, pos.J, board[pos.I-1][pos.J]))
		}
		if pos.J-prevPos.J == 1 && pos.J+1 < SIZE {
			positions = append(positions, NewPosition(pos.I, pos.J+1, board[pos.I][pos.J+1]))
		}
		if pos.J-prevPos.J == -1 && pos.J-1 >= 0 {
			positions = append(positions, NewPosition(pos.I, pos.J-1, board[pos.I][pos.J-1]))
		}
		for _, item := range checkedPositions {
			if contains(positions, item) {
				for i, p := range positions {
					if p.Equals(item) && !p.Equals(startPos) {
						positions = append(positions[:i], positions[i+1:]...)
						break
					}
				}
			}
		}
		return positions
	}

	if board[pos.I][pos.J] == BLACK {
		if math.Abs(float64(pos.I-prevPos.I)) == 1.0 {
			// move right
			if SIZE > pos.J+2 {
				if board[pos.I][pos.J+1] != BLACK {
					positions = append(positions, NewPosition(pos.I, pos.J+1, board[pos.I][pos.J+1]))
				}
			}
			// move left
			if pos.J-2 >= 0 {
				if board[pos.I][pos.J-1] != BLACK {
					positions = append(positions, NewPosition(pos.I, pos.J-1, board[pos.I][pos.J-1]))
				}
			}

		} else {
			// move up
			if pos.I-2 >= 0 {
				if board[pos.I-1][pos.J] != BLACK {
					positions = append(positions, NewPosition(pos.I-1, pos.J, board[pos.I-1][pos.J]))
				}
			}
			// move down
			if SIZE > pos.I+2 {
				if board[pos.I+1][pos.J] != BLACK {
					positions = append(positions, NewPosition(pos.I+1, pos.J, board[pos.I+1][pos.J]))
				}
			}

		}

	}
	if board[pos.I][pos.J] == WHITE {
		if pos.I-prevPos.I == 1 {
			if pos.I+1 < SIZE {
				if board[pos.I+1][pos.J] == BLACK {
					positions = append(positions, NewPosition(pos.I+1, pos.J, board[pos.I+1][pos.J]))
				}
			}
		}
		if pos.I-prevPos.I == -1 {
			if pos.I-1 >= 0 {
				if board[pos.I-1][pos.J] == BLACK {
					positions = append(positions, NewPosition(pos.I-1, pos.J, board[pos.I-1][pos.J]))
				}
			}
		}
		if pos.J-prevPos.J == 1 {
			if pos.J+1 < SIZE {
				if board[pos.I][pos.J+1] == BLACK {
					positions = append(positions, NewPosition(pos.I, pos.J+1, board[pos.I][pos.J+1]))
				}
			}
		}
		if pos.J-prevPos.J == -1 {
			if pos.J-1 >= 0 {
				if board[pos.I][pos.J-1] == BLACK {
					positions = append(positions, NewPosition(pos.I, pos.J-1, board[pos.I][pos.J-1]))
				}
			}
		}

	}
	for _, item := range checkedPositions {
		if contains(positions, item) {
			for i, p := range positions {
				if p.Equals(item) {
					positions = append(positions[:i], positions[i+1:]...)
					break
				}
			}
		}
	}
	return positions
}

func findBlackCells(board [][]int) []Position {
	var blacks []Position
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if board[i][j] == BLACK {
				blacks = append(blacks, NewPosition(i, j, BLACK))
			}
		}
	}
	return blacks
}

func loop(board [][]int, startPosition Position) [][]int {
	boardCopy := make([][]int, SIZE)
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			boardCopy[i] = append(boardCopy[i], board[i][j])
		}
	}
	prevPos := startPosition
	boardCopy[prevPos.I][prevPos.J] = VISITED
	var positionsToCheck []Position
	var checkedPositions []Position
	first := firstStep(board, prevPos)
	if first.Equals(NewPosition(N, N, N)) {
		return nil
	}
	positionsToCheck = append(positionsToCheck, first)
	prevPositions := make(map[Position]Position)
	prevPositions[positionsToCheck[0]] = prevPos
	checkedPositions = append(checkedPositions, startPosition)
	for len(positionsToCheck) > 0 {
		pos := positionsToCheck[len(positionsToCheck)-1]
		positionsToCheck = positionsToCheck[:len(positionsToCheck)-1]
		if pos.Equals(startPosition) {
			for i := 0; i < len(checkedPositions); i++ {
				boardCopy[checkedPositions[i].I][checkedPositions[i].J] = VISITED
			}
			return boardCopy
		}
		checkedPositions = append(checkedPositions, pos)
		nextPositions := findNextPositions(board, pos, prevPositions[pos], checkedPositions, startPosition)
		if len(nextPositions) == 0 {
			if len(positionsToCheck) > 0 {
				for checkedPositions[len(checkedPositions)-1] != prevPositions[positionsToCheck[len(positionsToCheck)-1]] {
					checkedPositions = checkedPositions[:len(checkedPositions)-1]
					if len(checkedPositions) == 0 {
						break
					}
				}
			} else {
				break
			}
		} else {
			for _, position := range nextPositions {
				positionsToCheck = append(positionsToCheck, position)
				prevPositions[position] = pos
			}
		}

	}
	return nil
}

type ICalculator interface {
	FindLoop(board [][]int) [][]int
}

type Calculator struct {
}

func (c Calculator) FindLoop(board [][]int) [][]int {
	blacks := findBlackCells(board)
	for _, black := range blacks {
		solution := loop(board, black)
		if solution != nil {
			return solution
		}
	}
	return nil
}
