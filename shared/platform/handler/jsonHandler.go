package handler

import (
	"io/ioutil"
	"net/http"

	"github.com/bperezgo/backoffice-businesses/shared/platform/handlertypes"
	"github.com/bperezgo/backoffice-businesses/shared/platform/logger"
	"github.com/gin-gonic/gin"
)

type JsonHandler struct {
	logger *logger.Logger
}

func NewJsonHandler() *JsonHandler {
	return &JsonHandler{
		logger: logger.GetLogger(),
	}
}

func (h *JsonHandler) Adapt(handler Handler) func(c *gin.Context) {
	return func(c *gin.Context) {
		request := handler.GetEmptyRequest()

		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			h.logger.Error(logger.LogInput{
				Action: "JsonHandler.Adapt",
				State:  logger.FAILED,
				Error: &logger.Error{
					Message: err.Error(),
				},
			})
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request body",
			})
			return
		}

		request.Body = body
		headers := handlertypes.Headers{}

		if err := c.ShouldBindHeader(&headers); err != nil {
			h.logger.Error(logger.LogInput{
				Action: "JsonHandler.Adapt",
				State:  logger.FAILED,
				Error: &logger.Error{
					Message: err.Error(),
				},
			})
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request body",
			})
			return
		}

		request.Headers = headers
		response := handler.Function(c, request)

		c.JSON(response.HttpStatus, response.Body)
	}
}
