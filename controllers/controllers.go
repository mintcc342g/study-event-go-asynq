package controllers

import (
	"github.com/labstack/echo/v4"
)

type (
	// ResponseBody ...
	ResponseBody struct {
		StatusCode int         `json:"resultCode" example:"000"`
		ResultMsg  string      `json:"resultMsg" example:"Request OK"`
		ResultData interface{} `json:"resultData,omitempty"`
	}
)

func response(c echo.Context, code int, resMsg string, result ...interface{}) error {
	res := ResponseBody{
		StatusCode: code,
		ResultMsg:  resMsg,
	}
	if result != nil {
		res.ResultData = result[0]
	}

	return c.JSON(code, res)
}
