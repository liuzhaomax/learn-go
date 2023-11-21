package _map

import (
	"encoding/json"
	"fmt"
	"os"
)

// 无法遍历interface{}，需要将其值导入新的确定类型的切片，再进行遍历
func DoMap() {
	var mmm map[string]interface{}
	jsonBody, _ := os.ReadFile("data.json")
	err := json.Unmarshal(jsonBody, &mmm)
	if err != nil {
		panic(err)
		return
	}
	var sliceNudges []map[string]interface{}
	for _, val := range mmm["nudges"].([]interface{}) {
		sliceNudges = append(sliceNudges, val.(map[string]interface{}))
	}
	for i := 0; i < len(sliceNudges); i++ {
		sliceNudges[i]["tip_key"] = 777
	}
	fmt.Println(mmm) // mmm的值也跟着变

	var data map[string]interface{}
	for i := 0; i < len(sliceNudges); i++ {
		str1 := fmt.Sprintf("%s", sliceNudges[i]["data"])
		err = json.Unmarshal([]byte(str1), &data)
		if err != nil {
			panic(err)
		}
		fmt.Println(data)
	}

	//var cstSh, _ = time.LoadLocation("Asia/Shanghai")
	//fmt.Println(time.Now().In(cstSh).Format("2006-01-02T15:04:05.000-0700"))
}
