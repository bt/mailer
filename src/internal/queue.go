package internal

import (
	"encoding/json"
	"fmt"
	"log"

	"mailer/src/emailing"

	"github.com/nutrun/lentil"
)

type Queue interface {
	// Will start receiving emails and pass it to the specified channel
	StartReceiving(chan *emailing.Email)
}

// A queue using Beanstalkd.
type BeanstalkQueue struct {
	conn *lentil.Beanstalkd
}

func (q *BeanstalkQueue) StartReceiving(ch chan *emailing.Email) {
	for {
		// Receive an email job.
		email, err := q.receiveJob()

		// An error occurred, break the receiving.
		if err != nil {
			log.Fatal(err)
			break
		}

		// Send it off for processing.
		ch <- email
	}
}

func (q *BeanstalkQueue) receiveJob() (*emailing.Email, error) {
	// Reserve the next job.
	job, err := q.conn.Reserve()
	if err != nil {
		return nil, err
	}

	// We'll get a JSON back, unmarshal it into an email object
	var email emailing.Email
	if err := json.Unmarshal(job.Body, &email); err != nil {
		return nil, err
	}

	return &email, nil
}

func NewBeanstalkQueue(host string) (*BeanstalkQueue, error) {
	// Establish connection to Beanstalkd
	conn, err := lentil.Dial(fmt.Sprintf("%s:11300", host))
	if err != nil {
		return nil, err
	}

	return &BeanstalkQueue{
		conn: conn,
	}, nil
}
