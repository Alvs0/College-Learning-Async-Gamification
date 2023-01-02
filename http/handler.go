package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"log"
)

const (
	NOT_FOUND_CODE             = 404
	INTERNAL_SERVER_ERROR_CODE = 500
)

type ErrorWrapper struct {
	error   error
	message string
	code    int
}

type HandlerFunc func(req *http.Request, res http.ResponseWriter) (result interface{}, err *ErrorWrapper)

type HttpHandler struct {
	handlerFunc HandlerFunc
}

func NewHandler(handlerFunc HandlerFunc) HttpHandler {
	return HttpHandler{
		handlerFunc: handlerFunc,
	}
}

func WrapError(code int, err error) *ErrorWrapper {
	return &ErrorWrapper{
		error:   err,
		message: err.Error(),
		code:    code,
	}
}

func (ths HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	result, errWithCode := ths.handlerFunc(req, res)
	if errWithCode != nil {
		encodeJsonErrorResponse(res, errWithCode.code, errWithCode.message)
		return
	}

	resultByte, err := json.Marshal(result)
	if err != nil {
		encodeJsonErrorResponse(res, INTERNAL_SERVER_ERROR_CODE, err.Error())
		return
	}

	if result != nil {
		res.Header().Set("Content-Type", "application/json")
		_, err = res.Write(resultByte)
		if err != nil {
			encodeJsonErrorResponse(res, INTERNAL_SERVER_ERROR_CODE, err.Error())
			return
		}
	}
}

func encodeJsonErrorResponse(respWriter http.ResponseWriter, statusCode int, errorMessage string) {
	type errorResponse struct {
		Message string `json:"message"`
	}

	respJSON, err := json.Marshal(errorResponse{Message: errorMessage})
	if err != nil {
		respJSON = []byte("{'message': 'Internal Server Error'}")
		log.Println(fmt.Sprintf("[encodeJsonErrorResponse] error marshaling response with error: %v ", err.Error()))
	}

	respWriter.Header().Set("Content-Type", "application/json")
	respWriter.WriteHeader(statusCode)
	_, err = respWriter.Write(respJSON)
	if err != nil {
		log.Println(fmt.Sprintf("[encodeJsonErrorResponse] error write response with error: %v", err.Error()))
	}
}
