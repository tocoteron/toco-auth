package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/tocoteron/toco-auth/model"
)

func ServerSettingProvider(setting *model.ServerSetting) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("server_setting", setting)
			return next(c)
		}
	}
}
