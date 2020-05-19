package api

import (
	"encoding/json"
	"fmt"
	"gi_visa/utils/errorutil"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Request struct {
	UserId string `form:"userId" binding:"required"`
}

func GetPdfHandler(c *gin.Context) {
	var req Request
	if err := c.Bind(&req); err != nil {
		bytes, _ := json.Marshal(req)
		c.Data(http.StatusBadRequest, c.ContentType(), bytes)
		return
	}

	files, _ := ioutil.ReadDir(fmt.Sprintf("storage/%s", req.UserId))

	if len(files) > 0 {
		file := files[len(files)-1]
		allBytes, _ := ioutil.ReadFile(fmt.Sprintf("storage/%s/%s", req.UserId, file.Name()))
		fileMime := http.DetectContentType(allBytes)
		c.Writer.Header().Set("Content-Type", fileMime)
		c.Writer.Header().Set("Content-Length", strconv.Itoa(len(allBytes)))
		c.Writer.Write(allBytes)
		return
	}
	bytes, _ := json.Marshal(errorutil.NewFunctionalError("File not available"))
	c.Data(http.StatusBadRequest, c.ContentType(), bytes)
	return
}
