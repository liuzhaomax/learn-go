package zhipu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

func TestVisitZhipu(t *testing.T) {
	url := "https://open.bigmodel.cn/api/paas/v4/chat/completions"
	method := "POST"
	apiKey := os.Getenv("ZHIPU_API_KEY")

	payload := map[string]interface{}{
		"model": "glm-4-airx",
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": "请用golang写一个hello world",
			},
		},
	}

	requestBody, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling payload:", err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}

	req.Header.Add("Authorization", "Bearer "+apiKey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", "acw_tc=2f624a2e17219843029878411e38f259e30d7f7b43d8aafc4c689322faae1a")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response:", string(body))
}
