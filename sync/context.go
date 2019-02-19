package main

import (
	"context"
	"fmt"
	"time"
)

type myString string

var key myString = "name"

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	//附加值
	valueCtx := context.WithValue(ctx, key, "【监控1】")
	go watch(valueCtx)
	valueCtx2 := context.WithValue(ctx, key, "【监控2】")
	go watch(valueCtx2)

	time.Sleep(time.Millisecond * 500)
	fmt.Println("可以了，通知监控停止")
	cancelFunc()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(time.Millisecond * 200)
}

func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			//取出值
			fmt.Println(ctx.Value(key), "监控退出，停止了...")
			return
		default:
			//取出值
			fmt.Println(ctx.Value(key), "goroutine监控中...")
			time.Sleep(time.Millisecond * 200)
		}
	}
}
