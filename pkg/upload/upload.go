package upload

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"path/filepath"
)

type FileBag struct {
	File multipart.FileHeader
	Name string
	Dest string
}

func (fb FileBag) Upload(c *gin.Context) (string, error) {
	extension := filepath.Ext(fb.File.Filename)
	newFileName := fb.Name + extension
	dest := fb.Dest + newFileName

	err := c.SaveUploadedFile(&fb.File, dest)

	return dest, err
}
