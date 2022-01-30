/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/1/29 17:55
 * @version     v1.0
 * @filename    Fib.go
 * @description
 ***************************************************************************/
package main

import (
	"math"
	"time"
)

func ModuloFibonacciSequence(requestChan chan bool, resultChan chan int) {
	x, y := 1, 1
	for {
		<-requestChan
		time.Sleep(time.Millisecond * 10)
		x, y = y, x+y
		if x > int(math.Pow(10, 9)) {
			x = x % int(math.Pow(10, 9))
		}
		resultChan <- x
	}
}

//func main() {
//	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
//
//	skipTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//	checkError(err)
//	skip := int32(skipTemp)
//
//	totalTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//	checkError(err)
//	total := int32(totalTemp)
//
//	resultChan := make(chan int)
//	requestChan := make(chan bool)
//	go ModuloFibonacciSequence(requestChan, resultChan)
//	for i := int32(0); i < skip+total; i++ {
//		start := time.Now().UnixNano()
//		requestChan <- true
//		news := <-resultChan
//		if i < skip {
//			continue
//		}
//		end := time.Now().UnixNano()
//		timeDiff := (end - start) / 1000000
//		if timeDiff < 3 {
//			fmt.Println("Rate is too high")
//			break
//		}
//		fmt.Println(news)
//	}
//}
//
//func readLine(reader *bufio.Reader) string {
//	str, _, err := reader.ReadLine()
//	if err == io.EOF {
//		return ""
//	}
//
//	return strings.TrimRight(string(str), "\r\n")
//}
//
//func checkError(err error) {
//	if err != nil {
//		panic(err)
//	}
//}
