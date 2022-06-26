package bussiness

import (
	"bytes"
	"context"
	"mime/multipart"
	_config "project/e-comerce/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/echo/v4"
)


func UploadFileToS3(c echo.Context, fileName string, file *multipart.FileHeader)(string, error){
	s3Config := _config.S3Config()
	s3Session,errSession := session.NewSession(s3Config)
	if errSession != nil {
        return "",errSession
    }


	src, errOpen := file.Open()
    if errOpen != nil {
        return "",errOpen
    }
    defer src.Close()

	fileBuffer := make([]byte, file.Size)
	
    uploader := s3manager.NewUploader(s3Session)
	input := &s3manager.UploadInput{
        Bucket:      aws.String("infinitysport"), 		// bucket's name
        Key:         aws.String(fileName),        		// files destination location
        Body:        bytes.NewReader(fileBuffer),       // content of the file
        ContentType: aws.String("image/jpg"),           // content type
    }
    output, errUploader := uploader.UploadWithContext(context.Background(), input)
	
	if errUploader != nil {
		return "",errUploader
	}

	return output.Location,nil
}