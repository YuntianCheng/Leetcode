package main

import (
	"context"
	"fmt"
	"time"
)

func searchText(query string) string {
	time.Sleep(time.Millisecond * 2)
	return string("text " + query)
}

func searchImage(query string) string {
	time.Sleep(time.Millisecond * 8)
	return string("image " + query)
}

func searchVideo(query string) string {
	time.Sleep(time.Millisecond * 20)
	return string("video " + query)
}

func run(ctx context.Context, ch chan<- string, f func(string2 string) string, arg string) {
	go func() {
		select {
		case ch <- f(arg):
		case <-ctx.Done():
		}
	}()
}

func Search(ctx context.Context, query string) (ret []string) {
	timeC, cancelFunc := context.WithTimeout(ctx, time.Millisecond*10)
	defer cancelFunc()
	ch := make(chan string)
	run(timeC, ch, searchImage, query)
	run(timeC, ch, searchVideo, query)
	run(timeC, ch, searchText, query)
	for {
		select {
		case d := <-ch:
			ret = append(ret, d)
		case <-timeC.Done():
			return
		}

	}
}

func main() {
	fmt.Println(Search(context.Background(), "123"))
}
