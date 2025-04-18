package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PerformanceData struct {
	Language string  `json:"language"`
	Test     string  `json:"test"`
	Time     float64 `json:"time"`
}

func main() {
	resp, err := http.Get("https://api.langgraph.com/performance?languages=Python,Go&test=arithmetic")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	var data []PerformanceData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	for _, record := range data {
		fmt.Printf("Language: %s, Test: %s, Time: %.2f ms\n", record.Language, record.Test, record.Time)
	}
}
