package goemail

import (
	"fmt"
	"net/smtp"
)

type goEmail struct {
	auth   smtp.Auth
	Config Config
}

func New(config *Config) *goEmail {
	return &goEmail{smtp.PlainAuth("", config.UserName, config.Password, config.Host), *config}
}

func (c *goEmail) Send(data Email) {
	data.Subject = fmt.Sprintf("Subject: %s \n", data.Subject)

	var to string
	for _, v := range data.To {
		to += fmt.Sprintf("To: %s \n", v)
	}

	config := c.Config
	smtp.SendMail(
		fmt.Sprintf("%s:%s", config.Host, config.Port),
		c.auth,
		data.From,
		data.To,
		[]byte(
			data.Subject+
				to+
				data.Mime+
				data.Body,
		),
	)
}
