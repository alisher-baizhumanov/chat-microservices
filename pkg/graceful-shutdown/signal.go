package graceful_shutdown

import (
	"os"
	"os/signal"
	"syscall"
)

// WaitSignal waits for a termination signal (SIGTERM or SIGINT) and returns it.
func WaitSignal() os.Signal {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	return <-stop
}
