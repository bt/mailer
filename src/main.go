package main

import (
	"log"
	"sync"

	"mailer/src/emailing"
	"mailer/src/emailing/sendgrid"
	"mailer/src/internal"
)

const BeanstalkHost = "127.0.0.1"

func main() {
	var wg sync.WaitGroup

	// Initialise email providers
	initialiseEmailProviders()

	// Initialise Beanstalk queue
	bQueue, err := internal.NewBeanstalkQueue(BeanstalkHost)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Create a new worker
	worker := internal.NewWorker(bQueue)

	// Wait for 1 job to complete
	wg.Add(1)

	// Start the worker
	go func() {
		defer wg.Done()
		worker.Start()
	}()

	// Wait indefinitely
	wg.Wait()
}

func initialiseEmailProviders() {
	emailing.RegisterProvider(sendgrid.NewSendgridProvider("@TODO"))
}
