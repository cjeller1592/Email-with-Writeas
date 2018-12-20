package main

import (
	"fmt"
	"log"
	"net/mail"
	"net/smtp"

	"github.com/scorredoira/email"
	"go.code.as/writeas.v2"
)

func main() {
	// Grab the Write.as post using its ID
	c := writeas.NewClient()
	p, err := c.GetPost("zf2j0t1rbex12bg1")
	if err != nil {
		fmt.Printf("%v", err)
	} else {
		// All you need is the title as the subject and content as the body of the email
		subject := p.Title
		body := p.Content

		// Now we use the email library
		// Compose the email by feeding in Write.as post and adding the To/From Addresses
		m := email.NewMessage(subject, body)
		m.From = mail.Address{Name: "Something", Address: "something@gmail.com"}
		m.To = []string{"cjeller1592@gmail.com"}

		// Authorize and send!
		auth := smtp.PlainAuth("", "something@gmail.com", "password", "smtp.gmail.com")
		// Send that email!
		if err := email.Send("smtp.gmail.com:587", auth, m); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Check your mail!")
	}

}
