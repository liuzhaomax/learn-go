package SMS

import (
	env "github.com/alibabacloud-go/darabonba-env/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	string_ "github.com/alibabacloud-go/darabonba-string/client"
	time "github.com/alibabacloud-go/darabonba-time/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	console "github.com/alibabacloud-go/tea-console/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"testing"
)

// 使用AK&SK初始化账号Client
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi.Client, _err error) {
	config := &openapi.Config{}
	config.AccessKeyId = accessKeyId
	config.AccessKeySecret = accessKeySecret
	_result = &dysmsapi.Client{}
	_result, _err = dysmsapi.NewClient(config)
	return _result, _err
}

func SendSMSCode(args []*string) (_err error) {
	client, _err := CreateClient(env.GetEnv(tea.String("ACCESS_KEY_ID")), env.GetEnv(tea.String("ACCESS_KEY_SECRET")))
	if _err != nil {
		return _err
	}

	// 1.发送短信
	sendReq := &dysmsapi.SendSmsRequest{
		PhoneNumbers:  args[0],
		SignName:      args[1],
		TemplateCode:  args[2],
		TemplateParam: args[3],
	}
	sendResp, _err := client.SendSms(sendReq)
	if _err != nil {
		return _err
	}

	code := sendResp.Body.Code
	if !tea.BoolValue(util.EqualString(code, tea.String("OK"))) {
		console.Log(tea.String("错误信息: " + tea.StringValue(sendResp.Body.Message)))
		return _err
	}

	bizId := sendResp.Body.BizId
	// 2. 等待 10 秒后查询结果
	_err = util.Sleep(tea.Int(10000))
	if _err != nil {
		return _err
	}
	// 3.查询结果
	phoneNums := string_.Split(args[0], tea.String(","), tea.Int(-1))
	for _, phoneNum := range phoneNums {
		queryReq := &dysmsapi.QuerySendDetailsRequest{
			PhoneNumber: util.AssertAsString(phoneNum),
			BizId:       bizId,
			SendDate:    time.Format(tea.String("yyyyMMdd")),
			PageSize:    tea.Int64(10),
			CurrentPage: tea.Int64(1),
		}
		queryResp, _err := client.QuerySendDetails(queryReq)
		if _err != nil {
			return _err
		}

		dtos := queryResp.Body.SmsSendDetailDTOs.SmsSendDetailDTO
		// 打印结果
		for _, dto := range dtos {
			if tea.BoolValue(util.EqualString(tea.String(tea.ToString(tea.Int64Value(dto.SendStatus))), tea.String("3"))) {
				console.Log(tea.String(tea.StringValue(dto.PhoneNum) + " 发送成功，接收时间: " + tea.StringValue(dto.ReceiveDate)))
			} else if tea.BoolValue(util.EqualString(tea.String(tea.ToString(tea.Int64Value(dto.SendStatus))), tea.String("2"))) {
				console.Log(tea.String(tea.StringValue(dto.PhoneNum) + " 发送失败"))
			} else {
				console.Log(tea.String(tea.StringValue(dto.PhoneNum) + " 正在发送中..."))
			}

		}
	}
	return _err
}

func TestAliyunSMS(t *testing.T) {
	args := []string{
		"15802246228",
		"普米智图Prismer",
		"SMS_468850709",
		"{\"code\":\"1234\"}",
	}
	err := SendSMSCode(tea.StringSlice(args))
	if err != nil {
		panic(err)
	}
}
