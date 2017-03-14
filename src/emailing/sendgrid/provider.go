package sendgrid

import (
	"mailer/src/emailing"
)

type SendGridProvider struct {
	apiKey string
	active bool
}

func NewSendgridProvider(apiKey string) *SendGridProvider {
	return &SendGridProvider{
		apiKey: apiKey,
		active: true,
	}
}

func (p *SendGridProvider) SendEmail(e *emailing.Email) bool {
	// @TODO
	return false
}

func (p *SendGridProvider) Active() bool {
	return p.active
}

func (p *SendGridProvider) Deactivate() {
	p.active = false
}
