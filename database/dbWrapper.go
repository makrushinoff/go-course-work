package database

import (
	"gorm.io/gorm"
	"lab3/model"
)

type IDB interface {
	FindBoardById(id string) (model.Board, error)
	FindUserByLogin(login string) (model.User, error)
	CreateUser(user model.User) error
	CreateBoard(board model.Board) error
	FindAllBoards() ([]model.Board, error)
}

type DB struct {
	DB *gorm.DB
}

func (d DB) FindBoardById(id string) (model.Board, error) {
	var board model.Board
	err := d.DB.Where("id = ?", id).First(&board).Error
	return board, err
}

func (d DB) FindUserByLogin(login string) (model.User, error) {
	var user model.User
	err := d.DB.Where("login = ?", login).First(&user).Error
	return user, err
}

func (d DB) CreateUser(user model.User) error {
	return d.DB.Create(&user).Error
}

func (d DB) CreateBoard(board model.Board) error {
	return d.DB.Create(&board).Error
}

func (d DB) FindAllBoards() ([]model.Board, error) {
	var boards []model.Board
	err := d.DB.Find(&boards).Error
	return boards, err
}
