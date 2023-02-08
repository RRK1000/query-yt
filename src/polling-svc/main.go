package main

import (
	"context"
	"time"

	"github.com/query-yt/src/polling-svc/internal"
)

func main() {
	svc := internal.InitPollingSvc()
	ctx := context.Background()

	ticker := time.NewTicker(time.Second * 1).C
	go func() {
		for {
			select {
			case <-ticker:
				svc.PollYTApi(ctx)
			}
		}

	}()

	time.Sleep(time.Second * 10)

}
