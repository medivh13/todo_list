package errors

import (
	http "net/http"

	constants "todo_list/src/infra/constants"
)

type HttpError struct {
	CommonError
	HttpStatusNumber int    `json:"-"`
	HttpStatusName   string `json:"type"`
}

func (err HttpError) Error() string {
	return err.ClientMessage
}

func (err CommonError) GetHttpStatus() int {
	if httpCode[err.ErrorCode] == 0 {
		return http.StatusInternalServerError
	}

	return httpCode[err.ErrorCode]
}

func (err CommonError) ToHttpError() HttpError {
	httpStatusNumber := err.GetHttpStatus()

	return HttpError{
		CommonError:      err,
		HttpStatusNumber: httpStatusNumber,
		HttpStatusName:   constants.GetHttpStatusText(httpStatusNumber),
	}
}
