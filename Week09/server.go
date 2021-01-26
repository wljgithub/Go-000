package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"sync/atomic"
)

var port = flag.String("port", "6666", "tcp port")
var uid int64 = 1

type User struct {
	Id        int
	Addr      string
	MessageCh chan string
}

func main() {

	l, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	fmt.Println("tcp server running on:", *port)
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	//io.Copy(conn,conn)
	ch := make(chan string, 16)

	user := User{
		Id:        genUid(),
		Addr:      conn.RemoteAddr().String(),
		MessageCh: ch,
	}
	go sendBack(conn, ch)
	for {
		receive, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("[uid: %v] [addr: %v]: %v", user.Id, user.Addr, receive)
		ch <- receive
	}
}

func sendBack(conn net.Conn, ch chan string) {
	for m := range ch {
		conn.Write([]byte(m))
	}
}
func genUid() int {
	atomic.AddInt64(&uid, 1)
	return int(uid)
}
