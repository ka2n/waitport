package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/ka2n/waitport"
)

func main() {
	os.Exit(mainCLI())
}

func mainCLI() int {
	var (
		timeout time.Duration
		listen  string
		logout  = os.Stdout
	)

	flag.DurationVar(&timeout, "timeout", 0, "timout, '0' waits forever")
	flag.StringVar(&listen, "listen", "", "where to listen. e.g. :8080, /tmp/unix.sock")
	flag.Parse()

	if listen == "" {
		fmt.Fprintf(logout, "-listen required")
		return 1
	}

	fmt.Fprintf(logout, "Waiting %v (timeout: %v)\n", listen, timeout)

	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	if timeout > 0 {
		ctx, cancel = context.WithDeadline(context.Background(), time.Now().Add(timeout))
		defer cancel()
	} else {
		ctx = context.Background()
	}

	w := waitport.Watcher{Interval: time.Second}
	if err := w.Do(ctx, listen); err != nil {
		fmt.Fprintf(logout, "Error: %s", err)
		return 1
	}
	return 0
}
