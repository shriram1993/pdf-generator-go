package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"pdf-generator-go/internal/pdfgenerator"
)

var PdfGenerateHandler Handler = pdfGenerateHandler{pdfgenerator.Service}

type pdfGenerateHandler struct {
	service pdfgenerator.IService
}

func (pdfGenerateHandler) CreateRequest(c *gin.Context) (interface{}, error) {
	var req pdfgenerator.GenerateRequest
	if err := c.BindJSON(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func (h pdfGenerateHandler) Handle(ctx context.Context, req interface{}) (interface{}, error) {
	request, _ := req.(pdfgenerator.GenerateRequest)
	resp, err := h.service.GeneratePdf(ctx, &request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (pdfGenerateHandler) CreateResponse(resp interface{}) (interface{}, error) {
	return resp, nil
}
