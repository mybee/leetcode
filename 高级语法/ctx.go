package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 12*time.Second)
	go watch(ctx, "【监控1】")

	chCtx, cl := context.WithCancel(ctx)
	go watch(chCtx, "【监控2】")

	vCtx := context.WithValue(ctx, "name", "mafeng")
	go watch(vCtx, "【监控3】")

	time.Sleep(time.Second)
	cl()

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func wa(ctx context.Context, name string) {
	fmt.Println(name, ctx.Value("name"))
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控退出，停止了...")
			return
		default:
			fmt.Println(name, "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}

func watch(ctx context.Context, name string) {
	go wa(ctx, name+"[子 ctx]")
	fmt.Println(name, ctx.Value("name"))
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控退出，停止了...")
			return
		default:
			fmt.Println(name, "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}
