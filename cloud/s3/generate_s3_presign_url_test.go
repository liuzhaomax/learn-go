package s3

import (
	"bytes"
	"encoding/json"
	"fmt"
	env "github.com/alibabacloud-go/darabonba-env/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

type DifyRequest struct {
	Inputs       map[string]string `json:"inputs"`
	ResponseMode string            `json:"response_mode"`
	User         string            `json:"user"`
}

type DifyResponse struct {
	WorkflowID string `json:"workflow_run_id"`
	TaskID     string `json:"task_id"`
	Data       struct {
		Outputs struct {
			Text string `json:"text"`
		} `json:"outputs"`
	} `json:"data"`
}

func generatePresignedURL(c *gin.Context) {
	// 获取文件名
	fileName := c.Query("filename")
	if fileName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Filename is required"})
		return
	}

	// 获取ak
	akId := env.GetEnv(tea.String("AWS_ACCESS_KEY_ID"))
	akSecret := env.GetEnv(tea.String("AWS_ACCESS_KEY_SECRET"))
	// 初始化 AWS 会话
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(
			*akId,     // AWS Access Key ID
			*akSecret, // AWS Secret Access Key
			"",        // 通常为空
		),
	})
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}

	// 创建S3服务客户端
	svc := s3.New(awsSession)

	// 生成预签名的请求
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String("prismer-official"),
		Key:    aws.String(fileName),
	})

	// 获取预签名URL, 设置有效期为1小时
	urlStr, err := req.Presign(1 * time.Hour) // 设置为1小时
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to sign request: %s", err)})
		return
	}

	// 将urlStr发给dify，返回dify结果
	// sendUrlToDify(c, urlStr)

	// 将text发送给dify，返回dify结果
	sendTextToDify(c, urlStr)

	// 从s3拉取文件并返回
	// download(c, urlStr, fileName)

	// 返回预签名URL
	// c.JSON(http.StatusOK, gin.H{"url": urlStr})
}

func sendTextToDify(c *gin.Context, urlStr string) {
	// 从s3拉取文件并转化为text字符串
	text, err := downloadAndConvertToTxt(urlStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to download from s3: %s", err)})
		return
	}
	fmt.Printf("Bytes: %d\n", len(text))
	fmt.Println(text)

	// TODO 发给dify
	apiKey := "app-TvvOVkZbhnGgPQPRq6hTEtvm" // 请替换为你的实际api_key
	requestBody := DifyRequest{
		Inputs: map[string]string{
			"text": text,
		},
		ResponseMode: "blocking",
		User:         "abc-123",
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	client := &http.Client{}
	reqDify, err := http.NewRequest("POST", "https://prismer.dev/v1/workflows/run", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	reqDify.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	reqDify.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(reqDify)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	var difyResponse DifyResponse
	err = json.NewDecoder(resp.Body).Decode(&difyResponse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"text": difyResponse.Data.Outputs.Text})
}

func TestGenerateS3PrisignUrl(t *testing.T) {
	r := gin.Default()
	r.GET("/presign", generatePresignedURL)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %s", err)
	}
}
