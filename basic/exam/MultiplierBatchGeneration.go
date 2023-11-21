package exam

import (
	"time"
)

func BurstyRateLimiter(requestChan chan bool, resultChan chan int, batchSize int, toAdd int) {
	res := 0
	for {
		<-requestChan
		time.Sleep(time.Millisecond * 10)
		for i := 0; i < batchSize; i++ {
			resultChan <- res
			res = res + toAdd
		}
	}
}

//func main() {
//	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
//
//	skipTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//	checkError(err)
//	skipBatches := int(skipTemp)
//
//	totalTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//	checkError(err)
//	printBatches := int(totalTemp)
//
//	batchSizeTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//	checkError(err)
//	batchSize := int(batchSizeTemp)
//
//	toAddTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//	checkError(err)
//	toAdd := int(toAddTemp)
//
//	resultChan := make(chan int)
//	requestChan := make(chan bool)
//	go BurstyRateLimiter(requestChan, resultChan, batchSize, toAdd)
//	for i := 0; i < skipBatches+printBatches; i++ {
//		start := time.Now().UnixNano()
//		requestChan <- true
//		for j := 0; j < batchSize; j++ {
//			news := <-resultChan
//			if i < skipBatches {
//				continue
//			}
//			fmt.Println(news)
//		}
//		if i >= skipBatches && i != skipBatches+printBatches-1 {
//			fmt.Println("-----")
//		}
//		end := time.Now().UnixNano()
//		timeDiff := (end - start) / 1000000
//		if timeDiff < 3 {
//			fmt.Println("Rate is too high")
//			break
//		}
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
