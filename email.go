package goemail

type Email struct {
	From    string
	To      []string
	Subject string
	Mime    string
	Body    string
}
