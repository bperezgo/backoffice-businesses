package handler

import (
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
		method := handler.GetMethod()
		request := handler.GetEmptyRequest()
		if method == POST {
			if err := c.ShouldBindJSON(&request.Body); err != nil {
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
			response := handler.Function(request)

			c.JSON(http.StatusOK, response.Body)

			return
		}

		if method == GET {
			// TODO: Get query params
			response := handler.Function(request)

			c.JSON(http.StatusOK, response.Body)
			return
		}
	}
}
