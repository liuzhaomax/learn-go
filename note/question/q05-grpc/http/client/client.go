package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	schema "learn-go/note/question/q05-grpc/http"
	"net/http"
)

func main() {
	reqBody := schema.ReqBody{
		Data: "world",
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		panic(err)
	}
	reader := bytes.NewReader(body)
	res, err := http.Post("http://127.0.0.1:8080/hello", "application/json;charset=UTF-8", reader)
	defer res.Body.Close()
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var result schema.Result
	err = json.Unmarshal(data, &result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Msg)
}
