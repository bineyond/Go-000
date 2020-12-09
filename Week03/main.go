/*
 * @Author: your name
 * @Date: 2020-12-09 21:00:25
 * @LastEditTime: 2020-12-09 22:25:45
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /Week03/main.go
 */
package main

import (
	"context"
	"errors"
	"fmt"
	"html"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {

	g, ctx := errgroup.WithContext(context.Background())

	c := make(chan os.Signal)
	svrErr := make(chan error)
	stop := make(chan struct{})

	g.Go(func() error {

		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		})

		s := &http.Server{
			Addr:    "localhost:8080",
			Handler: mux,
		}

		go func() {
			svrErr <- s.ListenAndServe()
		}()

		select {
		case <-stop:
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			s.Shutdown(ctx)
			fmt.Println("http server shutdown by signal")
			return nil
		case <-svrErr:
			return errors.New("http server shutdown")
		case <-ctx.Done():
			return nil
		}

	})

	g.Go(func() error {
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, os.Interrupt)

		select {
		case <-c:
			// 退出信号
			stop <- struct{}{}
			return nil
		case <-ctx.Done():
			return nil
		}
	})

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}

}
