package chatgpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

type RequestBody struct {
	Model    string              `json:"model"`
	Messages []map[string]string `json:"messages"`
}

type ChatGPTResponse struct {
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func getChatGPTRes() {
	apiKey := "$OPENAI_API_KEY"

	// 构建请求体
	requestBody := RequestBody{
		Model: "gpt-3.5-turbo",
		Messages: []map[string]string{
			{"role": "system", "content": "You are a helpful assistant."},
			{"role": "user", "content": "你好，ChatGPT!"},
		},
	}

	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Failed to create request body:", err)
		return
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		fmt.Println("Failed to create request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response:", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: received status code %d\n", resp.StatusCode)
		fmt.Println("Response body:", string(body))
		return
	}

	var chatResponse ChatGPTResponse
	err = json.Unmarshal(body, &chatResponse)
	if err != nil {
		fmt.Println("Failed to parse response:", err)
		return
	}

	if len(chatResponse.Choices) > 0 {
		fmt.Println("ChatGPT response:", chatResponse.Choices[0].Message.Content)
	} else {
		fmt.Println("No response from ChatGPT")
	}
}

func TestGetChatGPTRes(t *testing.T) {
	getChatGPTRes()
}
