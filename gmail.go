package main

import (
	"fmt"
	"github.com/SlyMarbo/gmail"
	"go.code.as/writeas.v2"
)

func main() {
// Grab the Write.as post using its ID
	c := writeas.NewClient()
	p, err := c.GetPost("n6p1133qsl5f5hhe")
	if err != nil {
		fmt.Printf("%v", err)
	} else {
	// All you need is the title as the subject and content as the body of the email
		subject := p.Title
		body := p.Content

	// Now we use the gmail library
  // Sign in with your Gmail account, compose the email, and add recepients to it
  	email.From = "blahblah@gmail.com"
		email.Password = "blahblah"
    email := gmail.Compose(subject, body)
		email.AddRecipient("blahblah@gmail.com")

		// Send that email!
		err := email.Send()
		if err != nil {
			fmt.Printf("%v", err)
		} else {
			fmt.Println("Check your mail!")
		}
	}

}
