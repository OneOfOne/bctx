package bctx

import (
	"runtime"
	"time"
)

func Select(ctxs ...Context) Context {
	for {
		for _, ctx := range ctxs {
			select {
			case <-ctx.Done():
				return ctx
			default:
			}
		}
		runtime.Gosched()
		time.Sleep(time.Millisecond)
	}
}
