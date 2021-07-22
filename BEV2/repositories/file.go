package repositories

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
)

func UploadFile(files []*multipart.FileHeader) ([]string, error) {
	var filename []string

	for _, file := range files {
		f, _ := file.Open()
		defer f.Close()
		tempFile, err := ioutil.TempFile("public", "image-*.png")
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		defer tempFile.Close()
		filebytes, err := ioutil.ReadAll(f)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		tempFile.Write(filebytes)

		filename = append(filename, tempFile.Name())
	}
	return filename, nil
}
