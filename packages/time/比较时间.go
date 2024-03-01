package time

import (
	"fmt"
	"time"
)

func HasPassedGivenHours(timeStr3339 string, hours int) (bool, error) {
	now := time.Now()
	createdAt, err := time.Parse(time.RFC3339, timeStr3339)
	if err != nil {
		fmt.Println("解析时间字符串时出错：", err)
		return false, err
	}
	duration := now.Sub(createdAt).Hours()
	if int(duration) > hours {
		return true, nil
	}
	return false, nil
}
