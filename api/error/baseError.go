package error

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func ResponseBadRequestError(c echo.Context, msg string) error {
	return c.JSON(400, BaseResponse{
		Success: 0,
		Message: msg,
	})
}

func ResponseUnprocessableEntityError(c echo.Context, msg string) error {
	return c.JSON(422, BaseResponse{
		Success: 0,
		Message: msg,
	})
}

func ResponseInternalServerError(c echo.Context, opts ...string) error {
	message := "Internal Server Error"
	if len(opts) > 0 {
		message = opts[0]
	}

	return c.JSON(500, BaseResponse{
		Success: 0,
		Message: message,
	})
}

func ResponseUnauthorized(c echo.Context, extra ...string) error {
	if extra[0] != "" {
		return c.JSON(401, BaseResponse{
			Data:    extra[0],
			Success: 0,
			Message: "Unauthorized",
		})
	}
	return c.JSON(401, BaseResponse{
		Success: 0,
		Message: "Unauthorized",
	})
}

// func ResponseForbidden(c echo.Context) error {
// 	return c.JSON(403, BaseResponse{
// 		Success: 0,
// 		Message: "Forbidden",
// 	})
// }
func ResponseForbidden(c echo.Context, msg string) error {
	return c.JSON(403, BaseResponse{
		Success: 0,
		Message: msg,
	})
}

func ResponseValidationErrors(c echo.Context, e validator.ValidationErrors) error {
	var errs []ValidationErrorsResponse
	var errMsg string
	var errCode int

	for _, err := range e {
		//getting errors of viper
		switch err.ActualTag() {
		case "dbUnique":
			errCode = viper.GetInt(`http_response.common.bad_request.already_exists.code`)
			errMsg = viper.GetString(`http_response.common.bad_request.already_exists.msg`)
		default:
			errMsg = "Failed on: " + err.ActualTag()
			errCode = 0
		}

		errs = append(errs, ValidationErrorsResponse{
			Field: err.Field(),
			Msg:   errMsg,
			Code:  errCode,
		})
	}
	return c.JSON(400, BaseResponse{
		Data:    errs,
		Success: 0,
		Message: "validation failed",
	})
}

//ResponseValidationMappedErrors for responses stardard
func ResponseValidationMappedErrors(c echo.Context, stardardResponse StandardMappedResponse) error {
	return c.JSON(stardardResponse.StatusCode, stardardResponse)
}

//DefaltBadRequestFailValidateResponse ...
func DefaltBadRequestFailValidateResponse(validateErrors []ValidationErrorsResponse) (stardardResponse StandardMappedResponse) {
	stardardResponse.StatusCode = viper.GetInt(`http_response.common.bad_request.status_code`)
	stardardResponse.Code = viper.GetInt(`http_response.common.bad_request.validate_failed.code`)
	stardardResponse.Message = viper.GetString(`http_response.common.bad_request.validate_failed.msg`)
	stardardResponse.Success = false
	stardardResponse.Data = validateErrors
	return
}

//MakeStardardResponse important: errorOnYml is a referring string error on errors.yml
func MakeStardardResponse(errorOnYml string, data interface{}) (StandardMappedResponse, error) {
	stardardResponse := StandardMappedResponse{}
	stardardResponse.StatusCode = viper.GetInt(`http_response.common.bad_request.status_code`)
	stardardResponse.Code = viper.GetInt(errorOnYml + `.code`)
	stardardResponse.Message = viper.GetString(errorOnYml + `.msg`)
	stardardResponse.Success = false
	stardardResponse.Data = data
	return stardardResponse, errors.New(stardardResponse.Message)
}

type ErrorCode int

const (
	InternalServerError      = ErrorCode(500)
	BadRequestError          = ErrorCode(400)
	UnauthorizedError        = ErrorCode(401)
	NotFoundError            = ErrorCode(404)
	UnprocessableEntityError = ErrorCode(422)
	Forbidden                = ErrorCode(403)
)

type AppError struct {
	Error   error
	Message string
	Code    ErrorCode
}

func (appError *AppError) FormattedMessage() string {
	message := appError.Message

	if message == "" {
		defaultMessage := map[ErrorCode]string{
			InternalServerError: "Internal Server Error",
			UnauthorizedError:   "Unauthorized",
		}

		if val, ok := defaultMessage[appError.Code]; ok {
			message = val
		}
	}

	return message
}

func (appError *AppError) Response() BaseResponse {
	return BaseResponse{
		Success: 0,
		Message: appError.FormattedMessage(),
	}
}

// func (appError *AppError) Json(c echo.Context) error {
// 	return c.JSON(int(appError.Code), appError.Response())
// }
