package helpers

import (
	"context"
	"fmt"
	"time"
)

func Loading(ctx context.Context, prefix string) {
	go func() {
		select {
		case <-ctx.Done():
			return
		default:
			for i := 0; ; i++ {
				time.Sleep(75 * time.Millisecond)
				fmt.Printf("\r \r")
				fmt.Printf("[%s] ", prefix)
				switch i % 4 {
				case 0:
					fmt.Printf("-")
				case 1:
					fmt.Printf("\\")
				case 2:
					fmt.Printf("|")
				case 3:
					fmt.Printf("/")
				}
				if i == 4 {
					i = 0
				}
			}
		}
	}()
}
