package controllers

import (
	"day-2/lib/db"
	"day-2/models"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct{}

func NewUserController() UserController {
	return UserController{}
}

func (h *UserController) GetLists(c echo.Context) error {
	var users []models.User

	orm := db.NewDB()
	orm.Find(&users)

	return c.JSON(http.StatusOK, users)
}

func (h *UserController) GetDetail(c echo.Context) error {
	var user models.User

	orm := db.NewDB()
	orm.Where("id", c.Param("id")).First(&user)

	return c.JSON(http.StatusOK, user)
}

func (h *UserController) Store(c echo.Context) error {
	var user models.User

	json.NewDecoder(c.Request().Body).Decode(&user)

	orm := db.NewDB()
	result := orm.Create(&user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": result.Error})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "data stored"})
}

func (h *UserController) Update(c echo.Context) error {
	var user models.User

	json.NewDecoder(c.Request().Body).Decode(&user)

	orm := db.NewDB()
	result := orm.Where("id", c.Param("id")).Updates(&user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": result.Error})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "data updated"})
}

func (h *UserController) Delete(c echo.Context) error {
	var user models.User

	orm := db.NewDB()
	orm.Where("id", c.Param("id")).Delete(&user)

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "data deleted"})
}
