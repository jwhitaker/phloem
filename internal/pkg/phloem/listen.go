package phloem

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// ListenAndRoute will startup the listening process
func ListenAndRoute(consumer Consumer, eventRouter EventRouter) {
	log.Printf("Starting up event service")

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	consumer.Subscribe(eventRouter.Events())

	run := true

	for run == true {
		select {
		case sig := <-sigchan:
			log.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := consumer.Poll()

			if ev == nil {
				continue
			}

			handler, ok := eventRouter.GetHandler(ev.Event)

			if !ok {
				log.Printf("Could not find handler for %s.  Ignoring", ev.Event)
				continue
			}

			handler(ev)
		}
	}

	consumer.Close()
}
