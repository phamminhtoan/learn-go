package main

import (
	"context"
	"fmt"
	"time"
)

func main()  {
	ctx, cancel := context.WithCancel(context.Background())

	timeoutCtx, _ := context.WithTimeout(ctx, time.Second*10)
	time.AfterFunc(time.Second, func() {
		cancel()
	})

	//select {
	//case <- ctx.Done():
	//	fmt.Print("Done")
	//}

	select {
	case <- timeoutCtx.Done():
		fmt.Print("Done timeout ctx")
	}
}