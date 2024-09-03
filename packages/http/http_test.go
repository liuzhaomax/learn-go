package http

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHeadRequest(t *testing.T) {
	url := "https://www.researchgate.net/profile/Alexander-Hernandez-13/publication/324692924/figure/fig1/AS:618476335017986@1524467657868/Mobile-based-Gradebook-with-Student-Outcomes-Analytics-Architecture_Q320.jpg"
	resp, _ := http.Head(url)
	fmt.Println(resp.StatusCode)
}
