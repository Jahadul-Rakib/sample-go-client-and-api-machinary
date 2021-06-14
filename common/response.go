package common

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type ResponseDTO struct {
	Status  string      `json:"status" xml:"status"`
	Message string      `json:"message" xml:"message"`
	Data    interface{} `json:"data" xml:"data"`
}
type ErrorResponseDTO struct {
	Status  string      `json:"status" xml:"status"`
	Error   interface{} `json:"error" xml:"error"`
	Message string      `json:"message" xml:"message"`
}

func SuccessResponse(context echo.Context, message string, data interface{}) error {
	if context.Request().Header.Get(echo.HeaderContentType) == echo.MIMEApplicationXMLCharsetUTF8 {
		context.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationXMLCharsetUTF8)
		return context.XML(http.StatusOK, ResponseDTO{
			Status:  "success",
			Message: message,
			Data:    data,
		})
	} else {
		context.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		return context.JSON(http.StatusOK, ResponseDTO{
			Status:  "success",
			Message: message,
			Data:    data,
		})
	}
}

func ErrorResponse(context echo.Context, error interface{}, message string) error {
	if context.Request().Header.Get(echo.HeaderContentType) == echo.MIMEApplicationXMLCharsetUTF8 {
		context.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationXMLCharsetUTF8)
		return context.XML(http.StatusBadRequest, ErrorResponseDTO{
			Status:  "error",
			Error:   error,
			Message: message,
		})
	} else {
		context.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		return context.JSON(http.StatusBadRequest, ErrorResponseDTO{
			Status:  "error",
			Error:   error,
			Message: message,
		})
	}
}
