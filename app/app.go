package app

import "net/http"

type Response struct {
	Message string
	Data    any
}

type ResponseAble interface {
	Response | ResponseOK | ResponseBadRequest | ResponseInternalServerError
}

type ResponseOK Response
type ResponseBadRequest Response
type ResponseInternalServerError Response

func ResponseStatus(response any) (int, any) {
	statusCode := http.StatusOK
	switch response.(type) {
	case ResponseInternalServerError:
		statusCode = http.StatusInternalServerError
	case ResponseBadRequest:
		statusCode = http.StatusBadRequest
	case Response:
		statusCode = http.StatusOK
	}
	return statusCode, response
}
