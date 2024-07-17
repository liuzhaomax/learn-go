package SMS

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

// 帐号密码见password.log
const twilioAccount = ""
const twilioPassword = ""

func TestTwilio(t *testing.T) {
	// Twilio API
	apiURL := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", twilioAccount)
	// 数据
	data := url.Values{}
	data.Set("To", fmt.Sprintf("+86%s", "15802246228"))
	data.Set("From", "+16184321798")
	data.Set("Body", "你好，赶紧付费啊123")
	// 创建请求
	req, err := http.NewRequest("POST", apiURL, strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(twilioAccount, twilioPassword)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Response Status Code:", resp.StatusCode)
	}

	// 读取结果 测试用
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response Body:", string(body))
}
