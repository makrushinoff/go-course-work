package main

import (
	"lab3/config"
	"lab3/controller"
	"lab3/database"
	_ "lab3/docs"
	"lab3/middleware"
	"lab3/repository"
	"lab3/service"
	"log"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Lab3 Rest API
// @description Rest API documentation, generated based on annotations and swag library
// @host localhost:8080
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	conf := config.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}
	db, err := config.NewConnection(&conf)
	if err != nil {
		log.Fatal(err)
	}
	dbWrapper := database.DB{DB: db}
	br := repository.BoardRepository{DB: dbWrapper}
	err = config.MigrateEntities(db)
	if err != nil {
		log.Fatal(err)
	}
	calc := service.Calculator{}
	bg := service.BoardGenerator{Calculator: calc, BoardRepository: br}
	boards, err := br.GetAllBoards()
	if err != nil {
		log.Fatal(err)
	}
	if len(boards) == 0 {
		log.Default().Println("No available boards to solve are stored yet")
		numCpu := runtime.NumCPU()
		log.Default().Println("Creating boards with ", numCpu, " threads")
		for i := 0; i < numCpu; i++ {
			go bg.GenerateBoards()
		}
 	}
	ur := repository.UserRepository{DB: dbWrapper}
	boardController := controller.BoardController{BoardService: service.BoardService{Calculator: calc, BoardRepository: br}}
	authenticationController := controller.AuthenticationController{UserService: service.UserService{UserRepository: ur}}

	authenticationMiddleware := middleware.AuthenticationMiddleware{UserRepository: ur}

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/authenticate", authenticationController.AuthenticateUser)
	router.POST("register", authenticationController.RegisterUser)
	router.GET("/boards", authenticationMiddleware.RequireAuth, boardController.GetAllBoards)
	router.GET("/boards/:boardId", authenticationMiddleware.RequireAuth, boardController.GetBoardById)
	router.POST("/calculations", authenticationMiddleware.RequireAuth, boardController.MakeCalculations)

	router.Run("localhost:8080")
}
