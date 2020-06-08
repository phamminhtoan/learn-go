package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*1)
	//
	//time.AfterFunc(time.Second, func() {
	//	cancel()
	//})
	//
	//select {
	//case <- ctx.Done():
	//	fmt.Println(ctx.Err())
	//}

	doSomething(ctx)
	fmt.Println("end main")
}

func doSomething(ctx context.Context){
	cancelChannel := make(chan bool)
	go func() {
		select {
		case <- ctx.Done():
			fmt.Println(ctx.Err())
			cancelChannel <- true
			return
		}
	}()

	isCancelChannel := <- cancelChannel
	if isCancelChannel {
		close(cancelChannel)
		return
	}

	time.Sleep(time.Second*10)
	fmt.Println("end")


}