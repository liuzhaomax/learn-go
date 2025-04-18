package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/chromedp/chromedp"
)

// 正常的
// var url = "https://ts4.cn.mm.bing.net/th?id=OIP-C.A6DEhzqPChDhvcBrrofibwHaL2&pid=Api"

// 报错跨域  返回403
// var url = "https://www.researchgate.net/profile/Alexander-Hernandez-13/publication/324692924/figure/fig1/AS:618476335017986@1524467657868/Mobile-based-Gradebook-with-Student-Outcomes-Analytics-Architecture_Q320.jpg"
var url = "https://www.pnas.org/cms/10.1073/pnas.1619316114/asset/3f2b6fbb-a602-421f-af31-4b722d453da0/assets/graphic/pnas.1619316114fig04.jpeg"

// 报错(failed)net::ERR_BLOCKED_BY_ORB   返回200，但是network报错403
// var url = "https://img-blog.csdnimg.cn/img_convert/052dfa02ef0d15a1d2f8924c782f752a.png"

func TestHeadRequest(t *testing.T) {
	// resp, _ := http.Head(url)
	// contentType := resp.Header.Get("Content-Type")
	// fmt.Println(contentType)
	req, err := http.NewRequest("OPTIONS", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Access-Control-Allow-Origin"))
}

func TestChromedp(t *testing.T) {
	// 创建一个上下文
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	// 变量用于存储结果
	var res string
	// 执行Chromedp任务
	err := chromedp.Run(ctx,
		// 启动浏览器并访问URL
		chromedp.Navigate(url),
		// 等待页面加载完成
		chromedp.WaitVisible("body", chromedp.ByQuery),
		// 获取页面的某些内容或状态
		chromedp.OuterHTML("html", &res),
	)
	if err != nil {
		log.Fatal(err)
	}
	// 打印结果
	fmt.Println("页面加载成功，返回的HTML：", res)
}
