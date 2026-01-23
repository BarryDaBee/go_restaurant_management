package utility

import (
	"net/http"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

type Response struct {
	Status  string      `json:"status,omitempty"`
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func BuildSuccessResponse(code int, message string, data interface{}) Response {
	return Response{
		Code:    code,
		Status:  "success",
		Message: message,
		Data:    data,
	}
}

func BuildErrorResponse(code int, message string, err interface{}) Response {
	if code == http.StatusInternalServerError {
		message = "internal server error"
		err = message
	}

	return Response{
		Code:    code,
		Status:  "error",
		Message: message,
		Error:   err,
	}
}

func ValidationResponse(err error, validate *validator.Validate) validator.ValidationErrorsTranslations {
	errs := err.(validator.ValidationErrors)
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(validate, trans)
	return errs.Translate(trans)
}
