package service

import (
	"lab3/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstStep(t *testing.T) {
	board := [][]int{
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
	t.Run("Move right", func(t *testing.T) {
		position := model.Position{0, 0, model.BLACK}

		var actual model.Position = firstStep(board, position)

		assert.Equal(t, model.Position{0, 1, model.WHITE}, actual)
	})
	t.Run("Move left", func(t *testing.T) {
		position := model.Position{5, 9, model.BLACK}

		var actual model.Position = firstStep(board, position)

		assert.Equal(t, model.Position{5, 8, model.WHITE}, actual)
	})
	t.Run("Move up", func(t *testing.T) {
		position := model.Position{8, 1, model.BLACK}

		var actual model.Position = firstStep(board, position)

		assert.Equal(t, model.Position{7, 1, model.WHITE}, actual)
	})
	t.Run("Move down", func(t *testing.T) {
		position := model.Position{7, 9, model.BLACK}

		var actual model.Position = firstStep(board, position)

		assert.Equal(t, model.Position{8, 9, model.WHITE}, actual)
	})
	t.Run("Can not move", func(t *testing.T) {
		position := model.Position{6, 9, model.BLACK}

		var actual model.Position = firstStep(board, position)

		assert.Equal(t, model.Position{model.N, model.N, model.N}, actual)
	})
}

func TestContains(t *testing.T) {
	positions := []model.Position{{0, 0, model.BLACK}, {1, 0, model.BLACK}, {2, 0, model.BLACK}}

	t.Run("Does contain", func(t *testing.T) {
		position := model.Position{0, 0, model.BLACK}

		actual := contains(positions, position)

		assert.True(t, actual)
	})
	t.Run("Does not contain", func(t *testing.T) {
		position := model.Position{0, 2, model.BLACK}

		actual := contains(positions, position)

		assert.False(t, actual)
	})
}

func TestFindLoop(t *testing.T) {
	t.Run("Successfully", func(t *testing.T) {
		board := [][]int{
			{model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK},
			{model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE},
			{model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.WHITE},
			{model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.BLACK},
			{model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE},
			{model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.WHITE, model.BLACK},
			{model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE},
			{model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.BLACK},
			{model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE},
			{model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK},
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
		c := Calculator{}
		actual := c.FindLoop(board)

		assert.Equal(t, expected, actual)
	})

	t.Run("Failed", func(t *testing.T) {
		board := [][]int{
			{model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.BLACK, model.BLACK, model.WHITE, model.WHITE},
			{model.BLACK, model.WHITE, model.WHITE, model.WHITE, model.BLACK, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK},
			{model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.WHITE, model.BLACK},
			{model.BLACK, model.WHITE, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.BLACK, model.WHITE},
			{model.BLACK, model.BLACK, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.WHITE, model.WHITE, model.WHITE},
			{model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.BLACK, model.WHITE},
			{model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.WHITE, model.BLACK, model.BLACK, model.BLACK, model.BLACK, model.BLACK},
			{model.WHITE, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.BLACK, model.BLACK, model.WHITE},
			{model.BLACK, model.BLACK, model.BLACK, model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.WHITE, model.WHITE},
			{model.WHITE, model.WHITE, model.BLACK, model.WHITE, model.WHITE, model.WHITE, model.BLACK, model.BLACK, model.BLACK, model.WHITE},
		}
		c := Calculator{}
		actual := c.FindLoop(board)

		assert.Nil(t, actual)
	})

}
