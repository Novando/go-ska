package common

import (
	"context"
	"fmt"
	"github.com/novando/go-ska/pkg/logger"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Operation func(ctx context.Context) error

// GracefulShutdown waits for termination syscalls and doing clean up operations after received it.
func GracefulShutdown(
	ctx context.Context,
	timeout time.Duration,
	ops map[string]Operation,
	l ...*logger.Logger,
) <-chan struct{} {
	log := logger.Call()
	if len(l) > 0 {
		log = l[0]
	}
	wait := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)

		// add any other syscalls that you want to be notified with
		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-s

		log.Infof("shutting down")

		// set timeout for the ops to be done to prevent system hang
		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Errorf(fmt.Sprintf("timeout %d ms has been elapsed, force exit", timeout.Milliseconds()))
			os.Exit(0)
		})

		defer timeoutFunc.Stop()

		var wg sync.WaitGroup

		// Do the operations asynchronously to save time
		for key, op := range ops {
			wg.Add(1)
			innerOp := op
			innerKey := key
			go func() {
				defer wg.Done()

				log.Infof(fmt.Sprintf("cleaning up: %s", innerKey))
				if err := innerOp(ctx); err != nil {
					log.Errorf(fmt.Sprintf("%s: clean up failed: %s", innerKey, err.Error()))
					return
				}

				log.Infof(fmt.Sprintf("%s was shutdown gracefully", innerKey))
			}()
		}

		wg.Wait()

		close(wait)
	}()

	return wait
}
