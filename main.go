package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
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
		dialer net.Dialer
		ctx    context.Context
		cancel context.CancelFunc
	)

	if timeout > 0 {
		ctx, cancel = context.WithDeadline(context.Background(), time.Now().Add(timeout))
		defer cancel()
	} else {
		ctx = context.Background()
	}

CHECK:
	for {
		if conn, err := dialer.DialContext(ctx, "tcp", listen); err != nil {
			time.Sleep(time.Second * 1)
			if ctx.Err() != nil {
				return 1
			}
			continue
		} else {
			conn.Close()
		}
		break CHECK
	}
	return 0
}
