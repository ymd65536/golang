package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// MainPage ページを表示する
func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	}
}
