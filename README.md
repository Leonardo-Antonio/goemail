# SENDING EMAIL - GoEmail

### Install librery

```shell
go get github.com/Leonardo-Antonio/goemail
```

### Use

***Configuración y envío del mensaje (html template)***
```go
func main () {
    mail := goemail.New(
	&goemail.Config{
		UserName: "leo2001.nl08@gmail.com",
		Password: "password", // se recomienda crear una contraseña de aplicación
		Host:     "smtp.gmail.com", // example -> smtp.gmail.com
		Port:     "587", // port the host smtp 
	},)

    mail.Send(goemail.Email{
	From:    "master@example.com",
	To:      []string{"mail2@example.com", "mail1@example.com"},
	Subject: "Subject test",
	Mime:    goemail.HTML,
	Body:    "<h1>Hola Mundo</h1>",
}
```
