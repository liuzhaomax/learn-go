/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/1/31 15:14
 * @version     v1.0
 * @filename    SyncHash.go
 * @description
 ***************************************************************************/
package main

import (
	"errors"
	"time"
)

type Result struct {
	Hash   string
	Number int32
}

var timeoutErr = errors.New("Timeout error")
var maxDelay = 100 * time.Millisecond

func SyncHash(steps int, resultChan0 chan Result, resultChan1 chan Result) (map[string]int32, error) {
	var res = make(map[string]int32)
	var finish = make(chan bool)
	go func() {
		for i := 0; i < steps; i++ {
			result0 := <-resultChan0
			res[result0.Hash] += result0.Number
			result1 := <-resultChan1
			res[result1.Hash] += result1.Number
		}
		finish <- true
	}()
	for {
		select {
		case <-time.After(maxDelay):
			return nil, timeoutErr
		case <-finish:
			return res, nil
		}
	}
}

//func main() {
//	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
//
//	stdout, err := os.Create("./basic/OUTPUT_PATH")
//	checkError(err)
//
//	defer stdout.Close()
//
//	writer := bufio.NewWriterSize(stdout, 16*1024*1024)
//
//	hashesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//	checkError(err)
//
//	var hashes []string
//
//	for i := 0; i < int(hashesCount); i++ {
//		hashesItem := readLine(reader)
//		hashes = append(hashes, hashesItem)
//	}
//
//	numbers0Count, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//	checkError(err)
//
//	var numbers0 []int32
//
//	for i := 0; i < int(numbers0Count); i++ {
//		numbers0ItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//		checkError(err)
//		numbers0Item := int32(numbers0ItemTemp)
//		numbers0 = append(numbers0, numbers0Item)
//	}
//
//	numbers1Count, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//	checkError(err)
//
//	var numbers1 []int32
//
//	for i := 0; i < int(numbers1Count); i++ {
//		numbers1ItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//		checkError(err)
//		numbers1Item := int32(numbers1ItemTemp)
//		numbers1 = append(numbers1, numbers1Item)
//	}
//
//	thresholdTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//	checkError(err)
//	threshold := int32(thresholdTemp)
//
//	server0, server1 := makeServer(), makeServer()
//	resultChan0 := server0(hashes, numbers0, threshold)
//	resultChan1 := server1(hashes, numbers1, threshold)
//	result, err := SyncHash(len(hashes), resultChan0, resultChan1)
//	if err == nil {
//		for _, hash := range hashes {
//			fmt.Fprintf(writer, hash)
//			fmt.Fprintf(writer, " ")
//			fmt.Fprintf(writer, fmt.Sprintf("%d", result[hash]))
//			fmt.Fprintf(writer, "\n")
//		}
//	} else {
//		fmt.Fprintf(writer, err.Error())
//	}
//
//	writer.Flush()
//}
//
//
//func makeServer() func([]string, []int32, int32) chan Result {
//	return func(hashes []string, numbers []int32, threshold int32) chan Result {
//		resultChan := make(chan Result)
//		go func() {
//			indexes := []int{}
//			for i := 0; i < len(hashes); i++ {
//				indexes = append(indexes, i)
//			}
//			rand.Seed(time.Now().UnixNano())
//			rand.Shuffle(len(indexes), func(i, j int) { indexes[i], indexes[j] = indexes[j], indexes[i] })
//			for _, i := range indexes {
//				if numbers[i] > threshold {
//					time.Sleep(125 * time.Millisecond)
//				} else {
//					time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
//				}
//				resultChan <- Result{
//					Hash:   hashes[i],
//					Number: numbers[i],
//				}
//			}
//		}()
//		return resultChan
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
