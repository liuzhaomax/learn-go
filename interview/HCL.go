package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(timeMinutes("101:01:20"))
}

// timeMinutes converts a time string in format "hh:mm:ss" to minutes. //assume all data validation has passed for the input "timeStr".
func timeMinutes(timeStr string) float64 {
	arr := make([]string, 0, 0)
	arr = strings.Split(timeStr, ":")
	var result float64
	h, _ := strconv.Atoi(arr[0])
	m, _ := strconv.Atoi(arr[1])
	s, _ := strconv.Atoi(arr[2])
	result += float64(h) * 60
	result += float64(m)
	result += math.Round(float64(s)/60*100) / 100

	return result
}
