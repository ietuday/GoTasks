package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func printOdd(ctx context.Context, n int, oddChan, evenChan chan struct{}, wg *sync.WaitGroup, logger *slog.Logger) {
	defer wg.Done()
	for i := 1; i <= n; i += 2 {
		select {
		case <-ctx.Done():
			logger.Warn("Odd printer exiting: context canceled")
			return
		case <-oddChan:
			logger.Info("Odd", "value", i)
			fmt.Println(i)
			if i+1 <= n {
				evenChan <- struct{}{}
			}
		}
	}
}

func printEven(ctx context.Context, n int, oddChan, evenChan chan struct{}, wg *sync.WaitGroup, logger *slog.Logger) {
	defer wg.Done()
	for i := 2; i <= n; i += 2 {
		select {
		case <-ctx.Done():
			logger.Warn("Even printer exiting: context canceled")
			return
		case <-evenChan:
			logger.Info("Even", "value", i)
			fmt.Println(i)
			if i+1 <= n {
				oddChan <- struct{}{}
			}
		}
	}
}

func startPrinting(n int) {
	// Structured logger setup
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Context with cancellation on OS signal
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Buffered channels prevent blocking under edge conditions
	oddChan := make(chan struct{}, 1)
	evenChan := make(chan struct{}, 1)

	var wg sync.WaitGroup
	wg.Add(2)

	// Launch goroutines
	go printOdd(ctx, n, oddChan, evenChan, &wg, logger)
	go printEven(ctx, n, oddChan, evenChan, &wg, logger)

	// Start the sequence
	oddChan <- struct{}{}

	// Wait for both goroutines or shutdown
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		logger.Info("All numbers printed successfully.")
	case <-ctx.Done():
		logger.Error("Forced shutdown due to OS signal or timeout.")
	}
}

func main() {
	startPrinting(10)
}
