package controller

import (
	"lab3/model"
	"lab3/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BoardController struct {
	BoardService service.IBoardService
}

// MakeCalculations	godoc
// @Summary 		find circle loop
// @Description 	receives a 2D array from request body and search on it a circle loop, depends on default conditions
// @Param 			calculations body model.BoardDto true "Submit"
// @Produce			application/json
// @Tags 			calculations
// @Success 		200 {object} model.BoardDto
// @Router			/calculations [post]
func (bc BoardController) MakeCalculations(context *gin.Context) {
	var boardDto model.BoardDto
	if err := context.BindJSON(&boardDto); err == nil {
		calculations, err := bc.BoardService.MakeCalculations(boardDto)
		if err != nil {
			context.IndentedJSON(http.StatusExpectationFailed, gin.H{"error": err.Error()})
		}
		context.IndentedJSON(http.StatusOK, calculations)
	}
}

// MakeCalculations	godoc
// @Summary 		get all possible boards
// @Description 	retrieves all saved board entities from database
// @Produce			application/json
// @Tags 			boards
// @Success 		200 {object} []model.BoardDto
// @Router			/boards [get]
func (bc BoardController) GetAllBoards(context *gin.Context) {
	boards, err := bc.BoardService.GetAllBoards()
	if err != nil {
		context.IndentedJSON(http.StatusExpectationFailed, gin.H{"error": err.Error()})
	}
	context.IndentedJSON(http.StatusOK, boards)
}

// MakeCalculations	godoc
// @Summary 		get possible board by id
// @Description 	respond board dto by provided id
// @Param			boardId path string true  "use id"
// @Produce			application/json
// @Tags 			boards
// @Success 		200 {object} model.BoardDto
// @Router			/boards/{boardId} [get]
func (bc BoardController) GetBoardById(context *gin.Context) {
	boardById, err := bc.BoardService.GetBoardById(context.Param("boardId"))
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	context.IndentedJSON(http.StatusOK, boardById)
}
