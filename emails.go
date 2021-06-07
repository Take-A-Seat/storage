package storage

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
)

// To use the SendEmail function you have to provide an email subject, an object that contains information about
// the receiver of the email (his name and email, you have to create an ToInfo object), two strings (textContent and htmlContent).
// The difference between these two strings is that the second one (htmlContent) can contain html tags.
// example of function call:
// var toInfo storage.ToInfo
// toInfo.Name = "andrei"
// toInfo.Email = "calinciucandrei98@gmail.com"
// subject := "Test send email"
// textContent := "Abracadabra abracada proadmin"
// htmlContent := "<b>Abracadabra</b> abracada <br/> proadmin"
// storage.SendEmail(subject, toInfo, textContent, htmlContent)

type ToInfo struct {
	Name  string
	Email string
}

func SendEmail(subject string, toInfo ToInfo, textContent string, htmlContent string) {
	from := mail.NewEmail("Take a seat", "info@takeaseat.site")
	to := mail.NewEmail(toInfo.Name, toInfo.Email)
	message := mail.NewSingleEmail(from, subject, to, textContent, htmlContent)
	client := sendgrid.NewSendClient("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")

	response, err := client.Send(message)

	if err != nil {
		log.Println("Send email error", err)
	}

	log.Println("Email sent", response.StatusCode, response.Body, response.Headers)
}