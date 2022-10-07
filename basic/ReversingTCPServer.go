package main

import (
	"fmt"
	"net"
)

func TCPServer(ready chan bool) {
	listener, err := net.Listen("tcp", "127.0.0.1:3333")
	if err != nil {
		fmt.Println("net error", err)
		return
	}
	ready <- true
	handle(listener)
}
func handle(listener net.Listener) {
	reversed := ""
	for {
		fmt.Println("listening...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("net error", err)
			return
		}
		fmt.Println("connection established! ")

		buf := make([]byte, 1024)
		fmt.Println("ready to receive")

		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("net error", err)
			return
		}
		msg := string(buf[:cnt])
		reversed = reverse(msg)
		sendData := []byte(reversed)
		cnt, err = conn.Write(sendData)
		if err != nil {
			fmt.Println("net error", err)
			return
		}
		conn.Close()
	}
}

func reverse(msg string) string {
	runeMsg := []rune(msg)
	for i, j := 0, len(runeMsg)-1; i < j; i, j = i+1, j-1 {
		runeMsg[i], runeMsg[j] = runeMsg[j], runeMsg[i]
	}
	return string(runeMsg)
}

const maxBufferSize = 1024
const address = "127.0.0.1:3333"

//func main() {
//	stdout, err := os.Create("./basic/OUTPUT_PATH")
//	checkError(err)
//
//	defer stdout.Close()
//
//	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
//	writer := bufio.NewWriterSize(stdout, 16*1024*1024)
//
//	messagesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//	checkError(err)
//
//	var messages []string
//
//	for i := 0; i < int(messagesCount); i++ {
//		messagesItem := readLine(reader)
//		messages = append(messages, messagesItem)
//	}
//
//	ready := make(chan bool)
//	go TCPServer(ready)
//	<-ready
//	reversed, err := tcpClient(messages)
//	if err != nil {
//		panic(err)
//	}
//	for _, msg := range reversed {
//		fmt.Fprintf(writer, "%s\n", msg)
//	}
//	writer.Flush()
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
//
//func tcpClient(messages []string) ([]string, error) {
//	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
//	if err != nil {
//		return []string{}, err
//	}
//
//	reversed := []string{}
//
//	for _, msg := range messages {
//
//		conn, err := net.DialTCP("tcp", nil, tcpAddr)
//		if err != nil {
//			return []string{}, err
//		}
//		_, err = conn.Write([]byte(msg))
//		if err != nil {
//			return []string{}, err
//		}
//
//		reply := make([]byte, maxBufferSize)
//
//		n, err := conn.Read(reply)
//		if err != nil {
//			return []string{}, err
//		}
//
//		reversed = append(reversed, string(reply[:n]))
//		conn.Close()
//	}
//
//	return reversed, nil
//}
