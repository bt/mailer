package emailing

type Email struct {
	From    string
	To      []string
	Subject string
	Body    string
}
