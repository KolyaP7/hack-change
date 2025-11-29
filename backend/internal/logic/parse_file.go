package logic

import (
	"mime/multipart"
)

func FileToDB(file *multipart.FileHeader, userId int, projectId int) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
}
