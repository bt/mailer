package internal

import (
	"log"
	"mailer/src/emailing"
)

type Worker struct {
	queue Queue
	ch    chan *emailing.Email
}

func (w *Worker) Start() {
	// Start a thread to receive from the queue
	go w.queue.StartReceiving(w.ch)

	// Wait for emails to arrive in that channel
	for {
		select {
		case email := <-w.ch:
			if err := w.processEmail(email); err != nil {
				// Fatal error, break
				log.Fatal(err)
				break
			}
		}
	}
}

// Process the email.
func (w *Worker) processEmail(e *emailing.Email) error {
	provider, err := emailing.AvailableProvider()
	if err != nil {
		return err
	}

	// Failed to send email
	if failed := provider.SendEmail(e); failed {
		provider.Deactivate()

		// Start a new thread to reinsert email into queue.
		// Without this thread there will be a deadlock.
		go func() {
			w.ch <- e
		}()
	}

	return nil
}

// Returns a new worker.
func NewWorker(queue Queue) *Worker {
	return &Worker{
		queue: queue,
		ch:    make(chan *emailing.Email),
	}
}
