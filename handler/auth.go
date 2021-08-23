package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tocoteron/toco-auth/auth"
	"github.com/tocoteron/toco-auth/model"
)

func SignUp(c echo.Context) error {
	type Response struct {
		Token string `json:"token"`
	}

	serverSetting, ok := c.Get("server_setting").(*model.ServerSetting)
	if !ok {
		return c.String(
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
		)
	}

	user := &model.User{}
	if err := c.Bind(user); err != nil {
		return err
	}

	token, err := auth.GenerateToken(serverSetting, user)
	if err != nil {
		return c.String(
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
		)
	}

	c.Logger().Infof("[Sign up] User: %s, Token: %s", user.ID, token)

	response := Response{
		Token: token,
	}

	return c.JSON(http.StatusCreated, response)
}

func SignIn(c echo.Context) error {
	type Response struct {
		Token string `json:"token"`
	}

	serverSetting, ok := c.Get("server_setting").(*model.ServerSetting)
	if !ok {
		c.Logger().Info("UNKO1")
		return c.String(
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
		)
	}

	user := &model.User{}
	if err := c.Bind(user); err != nil {
		c.Logger().Info("UNKO2")
		return c.String(
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
		)
	}

	token, err := auth.GenerateToken(serverSetting, user)
	if err != nil {
		c.Logger().Info("UNKO3")
		return c.String(
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
		)
	}

	c.Logger().Infof("[Sign in] User: %s, Token: %s", user.ID, token)

	response := Response{
		Token: token,
	}

	return c.JSON(http.StatusOK, response)
}
