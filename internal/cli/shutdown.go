package cli

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

var totalDuration time.Duration = 5 // sec

func CreateContextWithTimeout(d time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), d)
}

func SetupSignalHandler(w io.Writer, cancelFunc context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM) // SIGINT(Ctrl+C) , SIGTERM(Program ShutDown)
	go func() {
		s := <-c
		fmt.Fprintf(w, "Got signal: %v\n", s)
		cancelFunc()
	}()
}

func ExecuteCommand(ctx context.Context, command string, arg string) error {
	return exec.CommandContext(ctx, command, arg).Run()
}
