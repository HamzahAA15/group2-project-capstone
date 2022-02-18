package upload

import (
	"bytes"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadFile(filename string, path string, s *session.Session, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	size := fileHeader.Size
	buffer := make([]byte, size)
	file.Read(buffer)

	if !CheckExt(fileHeader.Filename) {
		return "", errors.New("file extension isn't equal to .png, .jpg. and .jpeg")
	}

	uploader := s3manager.NewUploader(s)

	//create a unique file name
	convName := path + "/" + filename + ".png"
	fmt.Println("unique file name: ", convName)

	result, err := uploader.Upload(&s3manager.UploadInput{
		// result, err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket: aws.String("sirclo-capstone-bucket"),
		Key:    aws.String(convName),
		// ACL:                  aws.String("public-read"), // could be private if you want it to be access by only authorized users
		Body: bytes.NewReader(buffer),
		// ContentLength:        aws.Int64(int64(size)),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("inline"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("INTELLIGENT_TIERING"),
	})
	if err != nil {
		return "", err
	}

	return result.Location, err

}

func CheckExt(filename string) bool {
	ext := string(filename[strings.LastIndex(filename, ".")+1:])

	if ext == "jpg" || ext == "jpeg" || ext == "png" {
		return true
	}
	return false
}
