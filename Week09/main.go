package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"week09/server"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	svr := server.NewTCPServer("127.0.0.1", 9999)

	// tcp server
	g.Go(func() error {
		// fmt.Println("tcp")
		go func() {
			<-ctx.Done()
			fmt.Println("tcp ctx done")
			svr.Shutdown(context.TODO())
		}()
		return svr.Start(context.Background())
	})

	// signal
	g.Go(func() error {
		exitSignals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT} // SIGTERM is POSIX specific
		sig := make(chan os.Signal, len(exitSignals))
		signal.Notify(sig, exitSignals...)
		for {
			// fmt.Println("signal")
			select {
			case <-ctx.Done():
				fmt.Println("signal ctx done")
				return ctx.Err()
			case <-sig:
				// do something
				svr.Shutdown(ctx)
				return nil
			}
		}
	})

	// inject error
	// g.Go(func() error {
	// 	fmt.Println("inject")
	// 	time.Sleep(time.Second)
	// 	fmt.Println("inject finish")
	// 	return errors.New("inject error")
	// })

	err := g.Wait() // first error return
	fmt.Println(err)
}
