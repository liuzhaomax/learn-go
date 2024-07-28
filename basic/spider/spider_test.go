package spider

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
	"testing"
)

// fetchHTML 获取网页的HTML内容
func fetchHTML(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

// extractText 从HTML节点中提取所有文本内容
func extractText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	var sb strings.Builder
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		sb.WriteString(extractText(c))
	}
	return sb.String()
}

func TestSpider(t *testing.T) {
	url := "https://q2chemistry.net/resumeChinese.html"
	// url := "https://fan.princeton.edu/sites/g/files/toruqf5476/files/documents/ChineseBiography1.pdf"
	doc, err := fetchHTML(url)
	if err != nil {
		fmt.Printf("获取网页失败: %v\n", err)
		return
	}

	text := extractText(doc)
	fmt.Println("网页文本内容:")
	fmt.Println(text)
}
