package controllers

import (
	"day-2/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BookController struct{}

func NewBookController() BookController {
	return BookController{}
}

func (h *BookController) GetLists(c echo.Context) error {
	books := make([]models.Book, 0)

	books = append(books, models.Book{
		Id:     1,
		Title:  "Sample BOOK",
		Author: "Mr Jack",
	})
	books = append(books, models.Book{
		Id:     2,
		Title:  "Dream Book",
		Author: "Mr Pettruci",
	})

	res := map[string]interface{}{
		"data": books,
	}

	return c.JSON(http.StatusOK, res)
}

func (h *BookController) GetDetail(c echo.Context) error {
	books := models.Book{
		Id:     1,
		Title:  "Sample BOOK",
		Author: "Mr Jack",
	}

	res := map[string]interface{}{
		"data": books,
	}

	return c.JSON(http.StatusOK, res)
}

func (h *BookController) Store(c echo.Context) error {
	books := models.Book{
		Id:     1,
		Title:  "Sample BOOK",
		Author: "Mr Jack",
	}

	res := map[string]interface{}{
		"message": "books stored",
		"data":    books,
	}

	return c.JSON(http.StatusOK, res)
}

func (h *BookController) Update(c echo.Context) error {
	books := models.Book{
		Id:     1,
		Title:  "Sample BOOK",
		Author: "Mr Jack",
	}

	res := map[string]interface{}{
		"message": "data updated",
		"data":    books,
	}

	return c.JSON(http.StatusOK, res)
}

func (h *BookController) Delete(c echo.Context) error {
	res := map[string]interface{}{
		"message": "data deleted",
	}

	return c.JSON(http.StatusOK, res)
}
