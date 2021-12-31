package main

import "github.com/Leonardo-Antonio/goemail"

func main() {
	//"secw rgzf vbdl ddij"
	//"smtp.gmail.com:587"

	mail := goemail.New(
		&goemail.Config{
			UserName: "leo2001.nl08@gmail.com",
			Password: "secw rgzf vbdl ddij",
			Host:     "smtp.gmail.com",
			Port:     "587",
		})

	mail.Send(goemail.Email{
		From:    "leo2001.nl08@gmail.com",
		To:      []string{"leonardo@prodequa.com", "ale2002.nn13@gmail.com"},
		Subject: "Mensaje de tu amorcito",
		Mime:    goemail.HTML,
		Body:    "<h1>TE AMO <3</h1>",
	})

}
