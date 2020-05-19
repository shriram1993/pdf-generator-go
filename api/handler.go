package api

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"pdf-generator-go/dto"
)

type Handler interface {
	CreateRequest(c *gin.Context) (interface{}, error)
	Handle(ctx context.Context, req interface{}) (interface{}, error)
	CreateResponse(resp interface{}) (interface{}, error)
}

func CommonHandler(handler Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := handler.CreateRequest(c)
		if err != nil {
			bytes, _ := json.Marshal(req)
			c.Data(http.StatusBadRequest, c.ContentType(), bytes)
			return
		}

		response, err := handler.Handle(c, req)

		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.BaseResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, response)
	}
}

func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, dto.BaseResponse{
		Success: true,
		Message: "Working",
	})
	return
}
