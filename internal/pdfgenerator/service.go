package pdfgenerator

import (
	"context"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"github.com/sirupsen/logrus"
	"os"
	"pdf-generator-go/dto"
	"time"
)

var Service IService = service{}

type IService interface {
	GeneratePdf(ctx context.Context, request *GenerateRequest) (*GenerateResponse, error)
}

type service struct {
}

func (service) GeneratePdf(ctx context.Context, request *GenerateRequest) (*GenerateResponse, error) {
	go generatePdf(request)
	return &GenerateResponse{
		BaseResponse: dto.BaseResponse{
			Success: true,
			Message: "PDF generating in background",
		},
	}, nil
}

func generatePdf(request *GenerateRequest) error {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	os.MkdirAll(fmt.Sprintf("storage/%s", request.UserId), os.ModePerm)

	filePath := fmt.Sprintf("storage/%s/%v.pdf", request.UserId, timestamp)
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, fmt.Sprintf("file create on timestamp %v with message: %s", timestamp, request.Message))
	if err := pdf.OutputFileAndClose(filePath); err != nil {
		logrus.WithField("error", err).Error()
	}
	return nil
}
