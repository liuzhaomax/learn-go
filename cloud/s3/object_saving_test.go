package s3

import (
	"fmt"
	env "github.com/alibabacloud-go/darabonba-env/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"testing"
	"time"
)

var (
	bucketName = "prismer-official"
	region     = "us-east-1"
)

func uploadFile(c *gin.Context) {
	// 获取ak
	akId := env.GetEnv(tea.String("AWS_ACCESS_KEY_ID"))
	akSecret := env.GetEnv(tea.String("AWS_ACCESS_KEY_SECRET"))
	// 限制上传文件大小为 50MB
	form, err := c.MultipartForm()
	if err != nil {
		return
	}
	files := form.File["file"]
	maxFileSize := 50 << 20

	for _, fileHeader := range files {
		// 检查文件大小
		if fileHeader.Size > int64(maxFileSize) {
			c.String(http.StatusBadRequest, "File %s is too large: %d bytes", fileHeader.Filename, fileHeader.Size)
			return
		}
		// 获取文件后缀
		ext := fileHeader.Filename
		// 生成对象键，包含路径
		objectKey := fmt.Sprintf("uploads_test/%d/%02d/%s/%s", time.Now().Year(), time.Now().Month(), "userId", ext)

		// 打开文件
		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error opening the file"})
			return
		}
		defer file.Close()

		// 初始化 AWS 会话
		awsSession, err := session.NewSession(&aws.Config{
			Region: aws.String(region),
			Credentials: credentials.NewStaticCredentials(
				*akId,     // AWS Access Key ID
				*akSecret, // 的AWS Secret Access Key
				"",        // 通常为空
			),
		})
		if err != nil {
			log.Fatalf("Failed to create session: %v", err)
		}
		service := s3.New(awsSession)

		// 上传文件到 S3
		_, err = service.PutObject(&s3.PutObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(objectKey),
			Body:   file,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file to S3"})
			return
		}

		fileURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucketName, region, objectKey)
		c.JSON(http.StatusOK, gin.H{"message": "Successfully uploaded", "fileURL": fileURL})
	}
}

func TestSaveObject(t *testing.T) {
	r := gin.Default()
	r.POST("/upload", uploadFile)

	fmt.Println("Server is running on :8080")
	log.Fatal(r.Run(":8080"))
}
