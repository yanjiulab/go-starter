package main

import (
	"context"
	"fmt"
	"time"
)

// job respects cancellation and timeout from context.
func job(ctx context.Context) error {
	select {
	case <-time.After(150 * time.Millisecond):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	// WithValue carries request-scoped metadata.
	base := context.WithValue(context.Background(), "traceID", "demo-ctx-1")
	// WithTimeout auto-cancels after deadline.
	ctx, cancel := context.WithTimeout(base, 100*time.Millisecond)
	defer cancel()

	fmt.Println("traceID:", ctx.Value("traceID"))
	if err := job(ctx); err != nil {
		fmt.Println("job canceled:", err)
	}

	// WithCancel allows manual cancellation.
	ctx2, cancel2 := context.WithCancel(context.Background())
	go func() {
		time.Sleep(20 * time.Millisecond)
		cancel2()
	}()
	<-ctx2.Done()
	fmt.Println("manual cancel:", ctx2.Err())
}
