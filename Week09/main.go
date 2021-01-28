package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	// tcp server
	g.Go(func() error {
		return startSvr(ctx, "127.0.0.1:9999")
	})

	// signal
	g.Go(func() error {
		return recvSignal(ctx)
	})

	err := g.Wait() // first error return
	fmt.Println(err)

	time.Sleep(time.Second)
}

func startSvr(ctx context.Context, addr string) error {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return errors.WithMessage(err, "start tcp server listen failure")
	}

	go func() {
		select {
		case <-ctx.Done():
			listen.Close()
		}
	}()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept error:%v\n", err)
			return err
		}

		// 如何回收
		go handlerConn(ctx, conn)
	}
}

func handlerConn(ctx context.Context, conn net.Conn) {
	defer conn.Close()

	type User struct {
		Addr           string
		MessageChannel chan string
	}

	user := &User{
		Addr:           conn.RemoteAddr().String(),
		MessageChannel: make(chan string, 8),
	}

	log.Println(user.Addr + ": enter ")

	// 如何回收
	go sendMessage(ctx, conn, user.MessageChannel)

	input := bufio.NewScanner(conn)

	for input.Scan() {
		fmt.Println(user.Addr + ": --> " + input.Text())
		fmt.Println(user.Addr + ": <-- " + input.Text())
		user.MessageChannel <- input.Text()
	}

	if err := input.Err(); err != nil {
		log.Println("read data err:", err)
	}

	log.Println(user.Addr + ": leave ")
}

func sendMessage(ctx context.Context, conn net.Conn, ch <-chan string) {
	select {
	case <-ctx.Done():
		return
	case msg := <-ch:
		fmt.Fprintln(conn, msg)
	}
}

// 处理信息
func recvSignal(ctx context.Context) error {
	exitSignals := []os.Signal{os.Interrupt, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT} // SIGTERM is POSIX specific
	sig := make(chan os.Signal, len(exitSignals))
	signal.Notify(sig, exitSignals...)

	select {
	case <-sig:
		return errors.New("exit -> close signal")
	case <-ctx.Done():
		return errors.New("exit -> ctx done")
	}
}
