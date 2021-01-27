/*
 * @Author: your name
 * @Date: 2021-01-27 20:02:09
 * @LastEditTime: 2021-01-27 20:56:18
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /Week09/server/server.go
 */
package server

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type TCPServer struct {
	Ip     string
	Port   int
	Listen net.Listener
}

func NewTCPServer(ip string, port int) *TCPServer {
	return &TCPServer{
		Ip:   ip,
		Port: port,
	}
}

func (svr *TCPServer) Start(ctx context.Context) (err error) {
	svr.Listen, err = net.Listen("tcp", fmt.Sprintf("%s:%d", svr.Ip, svr.Port))
	if err != nil {
		return errors.WithMessage(err, "start tcp server listen failure")
	}

	for {
		conn, err := svr.Listen.Accept()
		if err != nil {
			log.Printf("accept error:%v\n", err)
			// continue
			return err
		}

		go handlerConn(ctx, conn)

	}
}

func handlerConn(ctx context.Context, conn net.Conn) {
	defer conn.Close()

	user := &User{
		ID:             GenUserID(),
		Addr:           conn.RemoteAddr().String(),
		EnterAt:        time.Now(),
		MessageChannel: make(chan string, 8),
	}
	go sendMessage(conn, user.MessageChannel)

	input := bufio.NewScanner(conn)
	for input.Scan() {
		fmt.Println(strconv.Itoa(user.ID) + ":" + input.Text())
		// echo
		user.MessageChannel <- input.Text()
	}

	if err := input.Err(); err != nil {
		log.Println("read data err:", err)
	}

}

func sendMessage(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func (svr *TCPServer) Shutdown(ctx context.Context) {
	svr.Listen.Close()
}
