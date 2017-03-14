package main

import (
	"fmt"
	"log"

	"github.com/nutrun/lentil"
)

const QueueHost = "127.0.0.1"

func main() {
	// Example email
	email := `{"from": "b@bertramtruong.com", "to": "mnt@mailinator.com", "subject": "Test Email", "body": "This is a test email."}`

	// Establish connection to Beanstalkd
	conn, err := lentil.Dial(fmt.Sprintf("%s:6379", QueueHost))
	if err != nil {
		log.Fatal(err)
		return
	}

	// Push work onto queue
	jobId, err := conn.Put(0, 0, 60, []byte(email))

	// Feedback
	fmt.Println(fmt.Sprintf("✓ Email '%s' sent to queue!", jobId))
}
