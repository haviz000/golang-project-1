package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/haviz000/golang-project-1/controller"
	"github.com/haviz000/golang-project-1/repository"
	"github.com/haviz000/golang-project-1/service"
	"gorm.io/gorm"
)

func BookRoute(router *gin.Engine, db *gorm.DB) {
	bookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepository)
	bookController := controller.NewBookController(bookService)

	router.POST("/books", bookController.CreateBook)
	router.GET("/books", bookController.GetAllBook)
	router.GET("/books/:book_id", bookController.GetBookByID)
	router.PUT("/books/:book_id", bookController.UpdateBook)
	router.DELETE("/books/:book_id", bookController.DeleteBook)
}
