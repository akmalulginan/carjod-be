package handler

import (
	"fmt"
	"net/http"

	"github.com/akmalulginan/carjod-be/utils"
	"github.com/gin-gonic/gin"
)

// UploadHandler represent the httphandler for
type UploadHandler struct {
}

// NewUploadHandler will initialize the merchant resources endpoint
func NewUploadHandler(r *gin.RouterGroup) {
	handler := &UploadHandler{}
	r.POST("/upload", handler.Upload)
}

func (h *UploadHandler) Upload(c *gin.Context) {

	// Input Tipe File
	file, err := c.FormFile("file")

	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
		return
	}

	// Set Path file
	path := "file/" + file.Filename
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
		return
	}

	// Response
	c.JSON(http.StatusOK, utils.Response{Message: "success", Data: path})

}
