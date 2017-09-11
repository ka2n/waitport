package waitport

import (
	"context"
	"net"
	"time"
)

type Watcher struct {
	Interval time.Duration
}

// Do block until listen is success or context expired.
func (w Watcher) Do(ctx context.Context, listen string) error {
	var dialer net.Dialer
	for {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		if conn, err := dialer.DialContext(ctx, "tcp", listen); err != nil {
			time.Sleep(w.Interval)
			if ctx.Err() != nil {
				return err
			}
			continue
		} else {
			conn.Close()
			return nil
		}
	}
}
