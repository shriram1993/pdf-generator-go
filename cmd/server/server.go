package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"pdf-generator-go/api"
	"strconv"
)

func init() {

}
func main() {
	httpServer()
	shutdownChan := make(chan os.Signal)
	signal.Notify(shutdownChan, os.Interrupt)
	<-shutdownChan
}

func httpServer() {
	router := gin.New()
	gin.SetMode(gin.DebugMode)
	router.GET("/health", api.HealthHandler)
	router.POST("/generatePdf", api.CommonHandler(api.PdfGenerateHandler))
	router.GET("/getPdf", api.GetPdfHandler)
	go func() {
		port := 8080
		logrus.Infof("Starting Web server")
		if err := router.Run(":" + strconv.Itoa(port)); err != nil {
			logrus.Errorf("Error in starting web server on port %d. Error: %v", port, err)
			os.Exit(1)
		}
	}()
}
