package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/yuuulf/go-simple-api/pkg/db"
	"github.com/yuuulf/go-simple-api/pkg/user"
)

func GetUsers(c echo.Context) error {
	users, err := user.FindAll(db.Conn())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	userid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	user, err := user.FindById(db.Conn(), userid)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}
