package pdfgenerator

import "pdf-generator-go/dto"

type GenerateRequest struct {
	Message string `json:"message" binding:"required"`
	UserId  string `json:"userId" binding:"required"`
}

type GenerateResponse struct {
	dto.BaseResponse
}
